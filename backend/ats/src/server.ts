import { SendStatusRequest, SendStatusResponse } from '../spec/ats_pb';
import { AtsService } from '../spec/ats_grpc_pb';
import * as grpc from '@grpc/grpc-js'
import { sendUnaryData } from '@grpc/grpc-js/build/src/server-call';

const AtsServer = {
  sendStatus: (call: grpc.ServerUnaryCall<SendStatusRequest, SendStatusResponse>, callback: sendUnaryData<SendStatusResponse>) => {
    const request = call.request;
    const response = new SendStatusResponse();
    response.setResponse(1);
    console.log(`Train passed: ${request.getSensor()}`);
    callback(null, response);
  }
}

function serve(): void {
  const server = new grpc.Server();
  server.addService(AtsService, AtsServer);
  server.bindAsync('localhost:6543', grpc.ServerCredentials.createInsecure(), (err, port) => {
    if (err) {
      throw err;
    }
    console.log(`Listening on ${port}`);
    server.start();
  });
}

serve();