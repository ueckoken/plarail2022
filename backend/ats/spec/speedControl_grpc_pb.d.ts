// GENERATED CODE -- DO NOT EDIT!

// package: 
// file: speedControl.proto

import * as speedControl_pb from "./speedControl_pb";
import * as grpc from "@grpc/grpc-js";

interface ISpeedService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  controlSpeed: grpc.MethodDefinition<speedControl_pb.SendSpeed, speedControl_pb.StatusCode>;
}

export const SpeedService: ISpeedService;

export interface ISpeedServer extends grpc.UntypedServiceImplementation {
  controlSpeed: grpc.handleUnaryCall<speedControl_pb.SendSpeed, speedControl_pb.StatusCode>;
}

export class SpeedClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  controlSpeed(argument: speedControl_pb.SendSpeed, callback: grpc.requestCallback<speedControl_pb.StatusCode>): grpc.ClientUnaryCall;
  controlSpeed(argument: speedControl_pb.SendSpeed, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<speedControl_pb.StatusCode>): grpc.ClientUnaryCall;
  controlSpeed(argument: speedControl_pb.SendSpeed, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<speedControl_pb.StatusCode>): grpc.ClientUnaryCall;
}
