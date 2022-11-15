# externalとproxyとの繋ぎ込みをするためのクラス

"""
auto_operation <-> external (gRPC)
(statesync) 現在のハードウェアの状態を更新・管理する
(block) 閉塞の状態を管理する
auto_operation <-> proxy (gRPC)
通知をする
"""


from concurrent import futures
import grpc
import spec.ats_pb2 as ats_pb2
import spec.ats_pb2_grpc as ats_pb2_grpc
import spec.block_pb2 as block_pb2
import spec.block_pb2_grpc as block_pb2_grpc
import spec.statesync_pb2 as statesync_pb2
import spec.statesync_pb2_grpc as statesync_pb2_grpc
import os
import datetime
import inspect


class Connection:
    def __init__(self):
        self.externalServer = "localhost:6543"
        self.sensorBuffer = []

    def updateStatus(self, StationId, State):
        with grpc.insecure_channel(self.externalServer) as channel:
            stub = statesync_pb2_grpc.ControlStub(channel)
            response = stub.Command2Internal(
                statesync_pb2.RequestSync(
                    state=State, station=statesync_pb2.Stations(stationId=StationId)
                )
            )
        print(
            "Update Status: "
            + str(statesync_pb2.ResponseSync.Response.Name(response.response))
        )

    def updateBlock(self, BlockId, State):
        with grpc.insecure_channel("localhost:6543") as channel:
            stub = block_pb2_grpc.BlockStateSyncStub(channel)
            response = stub.NotifyState(
                block_pb2.NotifyStateRequest(
                    state=State, block=block_pb2.Blocks(blockId=BlockId)
                )
            )
        print(
            "Update Block: "
            + str(block_pb2.NotifyStateResponse.Response.Name(response.response))
        )


# sensorIdからsensorNameを取得する
def sensorId2sensorName(sensorId):
    return ats_pb2.SendStatusRequest.SensorName.Name(sensorId)


# Proxyと通信するためのサーバー(センサーの検知結果が飛んでくる)
class Ats(ats_pb2_grpc.AtsServicer):
    def SendStatus(self, request, context):
        sensorId = request.sensor
        print(
            f"Time: {datetime.datetime.timestamp( datetime.datetime.now())}, Reveived Sensor:  {sensorId2sensorName(sensorId)}"
        )
        return ats_pb2.SendStatusResponse(
            response=ats_pb2.SendStatusResponse.Response.Value("SUCCESS")
        )


# externalと通信するためのサーバー(新しい状態の更新を受けとる)
class StateSync(statesync_pb2_grpc.ControlServicer):
    def Command2Internal(self, request, context):
        print(
            "Received: StationId: "
            + str(request.station).strip()
            + "\t\tState:"
            + str(request.state)
        )
        return statesync_pb2.Command2InternalResponse(response=1)


# gRPCサーバーを起動する
def serve(con: Connection):
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    ats_pb2_grpc.add_AtsServicer_to_server(Ats(), server)
    server.add_insecure_port("[::]:6543")
    server.start()
    server.wait_for_termination()


if __name__ == "__main__":
    con = Connection()
    con.updateStatus(1, 2)
    serve(con)
    # con.updateBlock(1,2)
