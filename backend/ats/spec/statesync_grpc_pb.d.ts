// GENERATED CODE -- DO NOT EDIT!

// package: 
// file: statesync.proto

import * as statesync_pb from "./statesync_pb";
import * as grpc from "@grpc/grpc-js";

interface IControlService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  command2Internal: grpc.MethodDefinition<statesync_pb.RequestSync, statesync_pb.ResponseSync>;
}

export const ControlService: IControlService;

export interface IControlServer extends grpc.UntypedServiceImplementation {
  command2Internal: grpc.handleUnaryCall<statesync_pb.RequestSync, statesync_pb.ResponseSync>;
}

export class ControlClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  command2Internal(argument: statesync_pb.RequestSync, callback: grpc.requestCallback<statesync_pb.ResponseSync>): grpc.ClientUnaryCall;
  command2Internal(argument: statesync_pb.RequestSync, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<statesync_pb.ResponseSync>): grpc.ClientUnaryCall;
  command2Internal(argument: statesync_pb.RequestSync, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<statesync_pb.ResponseSync>): grpc.ClientUnaryCall;
}
