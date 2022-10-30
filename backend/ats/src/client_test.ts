import * as grpc from '@grpc/grpc-js';
import { SendStatusRequest, SendStatusResponse } from '../spec/ats_pb';
import { AtsClient } from '../spec/ats_grpc_pb';

function ats(): Promise<SendStatusResponse> {
  const client = new AtsClient(
    `localhost:6543`,
    grpc.credentials.createInsecure()
  )

  const request = new SendStatusRequest();
  request.setSensor(0);
  return new Promise<SendStatusResponse>((resolve, reject) => {
    console.log("Send Message");
    client.sendStatus(request, (err, response) => {
      if (err) {
        reject(err);
      }
      console.log("Receive Message");
      if(response==null){
        throw new Error("response is null");
      }
      return resolve(response);
    }
    );
  });
}

(async () => {
  console.log("Client Start");
  const result = await ats();
  console.log("Client End");
})();
