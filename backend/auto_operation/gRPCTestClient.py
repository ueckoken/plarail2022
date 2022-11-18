# externalとproxyとの繋ぎ込みをするためのクラス


import certifi
import grpc

import spec.ats_pb2 as ats_pb2
import spec.ats_pb2_grpc as ats_pb2_grpc

cred = grpc.ssl_channel_credentials(root_certificates=bytes(certifi.contents(), "utf-8"))


def testATS():
    with grpc.secure_channel("operate.chofufes2022.ueckoken.club", cred) as channel:
        stub = ats_pb2_grpc.AtsStub(channel)
        response = stub.SendStatus(ats_pb2.SendStatusRequest(sensor=1))
        print("Received: " + str(response))


# def testBlock():
#     with grpc.insecure_channel("localhost:6543") as channel:
#         stub = block_pb2_grpc.BlockStateSyncStub(channel)
#         response = stub.NotifyState(
#             block_pb2.NotifyStateRequest(state=2, block=block_pb2.Blocks(blockId=10))
#         )
#     print("Received: " + str(response.response))


# def testStateSync():
#     with grpc.insecure_channel("localhost:6543") as channel:
#         stub = statesync_pb2_grpc.ControlStub(channel)
#         response = stub.Command2Internal(
#             statesync_pb2.RequestSync(state=2, station=statesync_pb2.Stations(stationId=1))
#         )
#     print("Received: " + str(response.response))


if __name__ == "__main__":
    # testBlock()
    # testStateSync()
    testATS()
