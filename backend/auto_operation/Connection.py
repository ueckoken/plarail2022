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
    autoOperationServerAddress: str
    externalServerAddress: str
    atsServicer: "AtsServicer"
    pointStateNotificationServicer: "PointStateNotificationServicer"
    serverThread: Optional[threading.Thread]

    def __init__(self, autoOperationServerAddress: str, externalServerAddress: str) -> None:
        self.autoOperationServerAddress = autoOperationServerAddress
        self.externalServerAddress = externalServerAddress
        self.atsServicer = AtsServicer()
        self.pointStateNotificationServicer = PointStateNotificationServicer()
        self.serverThread = None

    @staticmethod
    def serveAndWait(
        autoOperationServerAddress: str,
        atsServicer: "AtsServicer",
        pointStateNotificationServicer: "PointStateNotificationServicer",
    ) -> None:
        server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        ats_pb2_grpc.add_AtsServicer_to_server(atsServicer, server)
        statesync_pb2_grpc.add_PointStateNotificationServicer_to_server(
            pointStateNotificationServicer, server
        )
        server.add_insecure_port(autoOperationServerAddress)
        server.start()
        server.wait_for_termination()

    def startServerThread(self) -> threading.Thread:
        if self.serverThread is not None:
            return self.serverThread

        self.serverThread = threading.Thread(
            target=Connection.serveAndWait,
            kwargs={
                "autoOperationServerAddress": self.autoOperationServerAddress,
                "atsServicer": self.atsServicer,
                "pointStateNotificationServicer": self.pointStateNotificationServicer,
            },
            daemon=True,
        )
        self.serverThread.start()
        return self.serverThread

    def sendStop(self, stationId: str, state: str) -> None:
        try:
            with grpc.insecure_channel(self.externalServerAddress) as channel:
                stub = statesync_pb2_grpc.StateManagerStub(channel)
                response = stub.UpdatePointState(
                    statesync_pb2.UpdatePointStateRequest(
                        state=statesync_pb2.PointAndState(
                            station=statesync_pb2.Station(stationId=stationId),
                            state=state,
                        )
                    )
                )
            print(f"Send Stop: {response}")
        except grpc._channel._InactiveRpcError as e:
            print(e)

    def sendPoint(self, pointId: str, state: str) -> None:
        try:
            with grpc.insecure_channel(self.externalServerAddress) as channel:
                stub = statesync_pb2_grpc.StateManagerStub(channel)
                response = stub.UpdatePointState(
                    statesync_pb2.UpdatePointStateRequest(
                        state=statesync_pb2.PointAndState(
                            station=statesync_pb2.Station(stationId=pointId), state=state
                        )
                    )
                )
            print(f"Send Point: {response}")
        except grpc._channel._InactiveRpcError as e:
            print(e)

    def sendBlock(self, blockId: str, state: str) -> None:
        try:
            with grpc.insecure_channel(self.externalServerAddress) as channel:
                stub = block_pb2_grpc.BlockStateManagerStub(channel)
                response = stub.UpdateBlockState(
                    block_pb2.UpdateBlockStateRequest(
                        state=block_pb2.BlockAndState(
                            blockId=blockId,
                            state=state,
                        )
                    )
                )
            print(f"Send Block: {response}")
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
class PointStateNotificationServicer(statesync_pb2_grpc.PointStateNotificationServicer):
    @dataclass
    class PointData:
        stationId: str
        state: str

    pointQueue: queue.Queue[PointData]

    def __init__(self) -> None:
        super().__init__()
        self.pointQueue = queue.Queue()

    def NotifyPointState(
        self, request: statesync_pb2.NotifyPointStateRequest, context
    ) -> statesync_pb2.NotifyPointStateResponse:
        stationId = statesync_pb2.StationId.Name(request.state.station.stationId)
        state = statesync_pb2.PointStateEnum.Name(request.state.state)
        data = PointStateNotificationServicer.PointData(stationId=stationId, state=state)
        print(f"Received: {data}")
        self.pointQueue.put(data)
        return statesync_pb2.NotifyPointStateResponse()


# 使い方の例
def test_connection() -> None:
    connection = Connection(
        autoOperationServerAddress="[::]:6543",
        externalServerAddress="localhost:6543",
    )
    connection.startServerThread()

    while True:
        # センサー情報を受信
        while connection.atsServicer.sensorQueue.qsize():
            connection.atsServicer.sensorQueue.get()

        # ポイント情報を受信
        while connection.pointStateNotificationServicer.pointQueue.qsize():
            connection.pointStateNotificationServicer.pointQueue.get()

        # ストップ情報を送信
        connection.sendStop(stationId="shinjuku_s1", state="POINTSTATE_ON")

        # 閉塞情報を送信
        connection.sendBlock(blockId="shinjuku_b1", state="BLOCKSTATE_OPEN")

        time.sleep(1)


if __name__ == "__main__":
    test_connection()
