# externalとproxyとの繋ぎ込みをするためのクラス

from concurrent import futures
import grpc
import spec.ats_pb2 as ats_pb2
import spec.ats_pb2_grpc as ats_pb2_grpc
import spec.block_pb2 as block_pb2
import spec.block_pb2_grpc as block_pb2_grpc

def testATS():
  with grpc.insecure_channel('localhost:6543') as channel:
    stub = ats_pb2_grpc.AtsStub(channel)
    response = stub.SendStatus(ats_pb2.SendStatusRequest(sensor=13))
  print("Received: " + str(response.response))

def testBlock():
  with grpc.insecure_channel('localhost:6543') as channel:
    stub = block_pb2_grpc.BlockStateSyncStub(channel)
    response = stub.NotifyState(block_pb2.NotifyStateRequest(state=2,blockId=2))
  print("Received: " + str(response.response))

if __name__ == '__main__':
  testBlock()