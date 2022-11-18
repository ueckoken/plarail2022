# externalとproxyとの繋ぎ込みをするためのクラス

import grpc

import spec.ats_pb2 as ats_pb2
import spec.ats_pb2_grpc as ats_pb2_grpc
import spec.block_pb2 as block_pb2
import spec.block_pb2_grpc as block_pb2_grpc
import spec.statesync_pb2 as statesync_pb2

import spec.statesync_pb2_grpc as statesync_pb2_grpc


def testATS():
    with grpc.insecure_channel("operate.chofufes2022.ueckoken.club") as channel:
        stub = ats_pb2_grpc.AtsStub(channel)
        response = stub.SendStatus(ats_pb2.SendStatusRequest(sensor=13))
    print("Received: " + str(response))


def testBlock():
    with grpc.insecure_channel("operate.chofufes2022.ueckoken.club") as channel:
        stub = block_pb2_grpc.BlockStateSyncStub(channel)
        response = stub.NotifyState(
            block_pb2.NotifyStateRequest(state=2, block=block_pb2.Blocks(blockId=10))
        )
    print("Received: " + str(response.response))


def testStateSync():
    with grpc.insecure_channel("operate.chofufes2022.ueckoken.club") as channel:
        stub = statesync_pb2_grpc.ControlStub(channel)
        response = stub.Command2Internal(
            statesync_pb2.RequestSync(state=2, station=statesync_pb2.Stations(stationId=1))
        )
    print("Received: " + str(response.response))


if __name__ == "__main__":
    # testBlock()
    # testStateSync()
    testATS()
