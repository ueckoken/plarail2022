# externalとproxyとの繋ぎ込みをするためのクラス

"""
auto_operation <-> external (WebSocket)
現在のハードウェアの状態を管理する
auto_operation <-> 
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
import inspect

# Proxyと通信するためのサーバー
class Ats(ats_pb2_grpc.AtsServicer):
  def SendStatus(self, request, context):
    sensorId = request.sensor
    print("Reveived: " + sensorId2sensorName(sensorId))
    return ats_pb2.SendStatusResponse(response=ats_pb2.SendStatusResponse.Response.Value('SUCCESS'))

# externalと通信するためのサーバー
class StateSync(statesync_pb2_grpc.ControlServicer):
  def Command2Internal(self, request, context):
    print("Received: StationId: "+ str(request.station).strip() + ", State:"+str(request.state))
    return statesync_pb2.Command2InternalResponse(response=1)

def test():
  with grpc.insecure_channel('localhost:6543') as channel:
    stub = ats_pb2_grpc.AtsStub(channel)
    response = stub.SendStatus(ats_pb2.SendStatusRequest(sensor=1))
  print("Received: " + str(response.response))
  
def sensorId2sensorName(sensorId):
  return ats_pb2.SendStatusRequest.SensorName.Name(sensorId)

# gRPCサーバーを起動する
def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    ats_pb2_grpc.add_AtsServicer_to_server(Ats(), server)
    server.add_insecure_port('[::]:6543')
    server.start()
    server.wait_for_termination()

def testStateSync():
  address = os.environ['EXTERNALSERVER_ADDR']
  port = os.environ['EXTERNALSERVER_PORT']
  with grpc.insecure_channel(f'{address}:{port}') as channel:
    stub = statesync_pb2_grpc.ControlStub(channel)
    response = stub.Command2Internal(statesync_pb2.RequestSync(state=2,station=statesync_pb2.Stations(stationId=1)))
  print("Received: " + str(response))

class Connection:
  def __init__(self):
    self.channel = grpc.insecure_channel('localhost:6543')
    self.stub = ats_pb2_grpc.AtsStub(self.channel)

  def send_status(self, sensor):
    response = self.stub.SendStatus(ats_pb2.SendStatusRequest(sensor=sensor))
    return response.response

if __name__ == '__main__':
  serve()
