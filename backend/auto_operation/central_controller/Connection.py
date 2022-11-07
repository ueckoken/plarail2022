# externalとproxyとの繋ぎ込みをするためのクラス

from concurrent import futures
import grpc
import spec.ats_pb2 as ats_pb2
import spec.ats_pb2_grpc as ats_pb2_grpc
import spec.block_pb2 as block_pb2
import spec.block_pb2_grpc as block_pb2_grpc

# Proxyと通信するためのサーバー
class Ats(ats_pb2_grpc.AtsServicer):
  def SendStatus(self, request, context):
    print("Reveived: " + str(request.sensor))
    return ats_pb2.SendStatusResponse(response=1)

def test():
  with grpc.insecure_channel('localhost:6543') as channel:
    stub = ats_pb2_grpc.AtsStub(channel)
    response = stub.SendStatus(ats_pb2.SendStatusRequest(sensor=1))
  print("Received: " + str(response.response))

# gRPCサーバーを起動する
def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    ats_pb2_grpc.add_AtsServicer_to_server(Ats(), server)
    server.add_insecure_port('[::]:6543')
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    serve()
    test()
