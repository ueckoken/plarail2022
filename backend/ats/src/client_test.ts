import * as grpc from '@grpc/grpc-js';
import { SendStatusRequest, SendStatusResponse } from '../spec/ats_pb';
import { AtsClient } from '../spec/ats_grpc_pb';

function ats(): Promise<SendStatusResponse> {
  const client = new AtsClient(
    `localhost:6543`,
    grpc.credentials.createInsecure()
  )

  const request = new SendStatusRequest();
  request.setSensor(24);
  return new Promise((resolve, reject) => {
    client.sendStatus(request, (err, response) => {
      if (err) {
        reject(err);
      }
      if (response === undefined) return reject(new Error('No data'));
      resolve(response);
    });
  });
}

(async () => {
  console.log("Client Start");
  console.log((await ats()).getResponse());
  console.log("Client End");
})();
