# externalとproxyとの繋ぎ込みをするためのクラス

"""
auto_operation <-> external (gRPC)
(statesync) 現在のハードウェアの状態を更新・管理する
(block) 閉塞の状態を管理する
auto_operation <-> proxy (gRPC)
通知をする
"""


import queue
import threading
import time
from concurrent import futures
from dataclasses import dataclass
from typing import Optional

import grpc

import spec.ats_pb2 as ats_pb2
import spec.ats_pb2_grpc as ats_pb2_grpc
import spec.block_pb2 as block_pb2
import spec.block_pb2_grpc as block_pb2_grpc
import spec.statesync_pb2 as statesync_pb2
import spec.statesync_pb2_grpc as statesync_pb2_grpc


class Connection:
    serverAddress: str
    externalServerAddress: str
    atsServicer: "AtsServicer"
    controlServicer: "ControlServicer"
    serverThread: Optional[threading.Thread]

    def __init__(self, serverAddress: str, externalServerAddress: str) -> None:
        self.serverAddress = serverAddress
        self.externalServerAddress = externalServerAddress
        self.atsServicer = AtsServicer()
        self.controlServicer = ControlServicer()
        self.serverThread = None

    @staticmethod
    def serveAndWait(
        serverAddress: str, atsServicer: "AtsServicer", controlServicer: "ControlServicer"
    ) -> None:
        server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        ats_pb2_grpc.add_AtsServicer_to_server(atsServicer, server)
        statesync_pb2_grpc.add_ControlServicer_to_server(controlServicer, server)
        server.add_insecure_port(serverAddress)
        server.start()
        server.wait_for_termination()

    def startServerThread(self) -> threading.Thread:
        if self.serverThread is not None:
            return self.serverThread

        self.serverThread = threading.Thread(
            target=Connection.serveAndWait,
            kwargs={
                "serverAddress": self.serverAddress,
                "atsServicer": self.atsServicer,
                "controlServicer": self.controlServicer,
            },
            daemon=True,
        )
        self.serverThread.start()
        return self.serverThread

    def sendStop(self, stationId: str, state: str) -> None:
        try:
            with grpc.insecure_channel(self.externalServerAddress) as channel:
                stub = statesync_pb2_grpc.ControlStub(channel)
                response = stub.Command2Internal(
                    statesync_pb2.RequestSync(
                        state=state, station=statesync_pb2.Stations(stationId=stationId)
                    )
                )
            print(f"Send Stop: {statesync_pb2.ResponseSync.Response.Name(response.response)}")
        except grpc._channel._InactiveRpcError as e:
            print(e)

    def sendBlock(self, blockId: str, state: str) -> None:
        try:
            with grpc.insecure_channel(self.externalServerAddress) as channel:
                stub = block_pb2_grpc.BlockStateSyncStub(channel)
                response = stub.NotifyState(
                    block_pb2.NotifyStateRequest(
                        state=state, block=block_pb2.Blocks(blockId=blockId)
                    )
                )
            print(f"Send Block: {block_pb2.NotifyStateResponse.Response.Name(response.response)}")
        except grpc._channel._InactiveRpcError as e:
            print(e)


# Proxyと通信するためのサーバー(センサーの検知結果が飛んでくる)
class AtsServicer(ats_pb2_grpc.AtsServicer):
    @dataclass
    class SensorData:
        sensorId: str

    sensorQueue: queue.Queue[SensorData]

    def __init__(self) -> None:
        super().__init__()
        self.sensorQueue = queue.Queue()

    def SendStatus(self, request: ats_pb2.SendStatusRequest, context) -> ats_pb2.SendStatusResponse:
        sensorId = ats_pb2.SendStatusRequest.SensorName.Name(request.sensor)
        data = AtsServicer.SensorData(sensorId=sensorId)
        print(f"Received: {data}")
        self.sensorQueue.put(data)
        return ats_pb2.SendStatusResponse(response=ats_pb2.SendStatusResponse.SUCCESS)


# externalと通信するためのサーバー(新しい状態の更新を受けとる)
class ControlServicer(statesync_pb2_grpc.ControlServicer):
    @dataclass
    class PointData:
        stationId: str
        state: str

    pointQueue: queue.Queue[PointData]

    def __init__(self) -> None:
        super().__init__()
        self.pointQueue = queue.Queue()

    def Command2Internal(
        self, request: statesync_pb2.Command2InternalRequest, context
    ) -> statesync_pb2.Command2InternalResponse:
        stationId = statesync_pb2.Stations.StationId.Name(request.station.stationId)
        state = statesync_pb2.Command2InternalRequest.State.Name(request.state)
        data = ControlServicer.PointData(stationId=stationId, state=state)
        print(f"Received: {data}")
        self.pointQueue.put(data)
        return statesync_pb2.Command2InternalResponse(
            response=statesync_pb2.Command2InternalResponse.SUCCESS
        )


# 使い方の例
def main() -> None:
    connection = Connection(
        serverAddress="[::]:6543",
        externalServerAddress="localhost:6543",
    )
    connection.startServerThread()

    while True:
        # センサー情報を受信
        while connection.atsServicer.sensorQueue.qsize():
            connection.atsServicer.sensorQueue.get()

        # ポイント情報を受信
        while connection.controlServicer.pointQueue.qsize():
            connection.controlServicer.pointQueue.get()

        # ストップ情報を送信
        connection.sendStop(stationId="shinjuku_s1", state="ON")

        # 閉塞情報を送信
        connection.sendBlock(blockId="shinjuku_b1", state="OPEN")

        time.sleep(1)


if __name__ == "__main__":
    main()
