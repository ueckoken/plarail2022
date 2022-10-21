// GENERATED CODE -- DO NOT EDIT!

// package: 
// file: ats.proto

import * as ats_pb from "./ats_pb";
import * as grpc from "@grpc/grpc-js";

interface IAtsService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  sendStatus: grpc.MethodDefinition<ats_pb.SendStatusRequest, ats_pb.SendStatusResponse>;
}

export const AtsService: IAtsService;

export interface IAtsServer extends grpc.UntypedServiceImplementation {
  sendStatus: grpc.handleUnaryCall<ats_pb.SendStatusRequest, ats_pb.SendStatusResponse>;
}

export class AtsClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  sendStatus(argument: ats_pb.SendStatusRequest, callback: grpc.requestCallback<ats_pb.SendStatusResponse>): grpc.ClientUnaryCall;
  sendStatus(argument: ats_pb.SendStatusRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<ats_pb.SendStatusResponse>): grpc.ClientUnaryCall;
  sendStatus(argument: ats_pb.SendStatusRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<ats_pb.SendStatusResponse>): grpc.ClientUnaryCall;
}
