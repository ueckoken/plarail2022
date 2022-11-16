// GENERATED CODE -- DO NOT EDIT!

// package: 
// file: statesync.proto

import * as statesync_pb from "./statesync_pb";
import * as grpc from "@grpc/grpc-js";

interface IControlService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  updatePointState: grpc.MethodDefinition<statesync_pb.UpdatePointStateRequest, statesync_pb.UpdatePointStateResponse>;
  notifyPointState: grpc.MethodDefinition<statesync_pb.NotifyPointStateRequest, statesync_pb.NotifyPointStateResponse>;
}

export const ControlService: IControlService;

export interface IControlServer extends grpc.UntypedServiceImplementation {
  updatePointState: grpc.handleUnaryCall<statesync_pb.UpdatePointStateRequest, statesync_pb.UpdatePointStateResponse>;
  notifyPointState: grpc.handleUnaryCall<statesync_pb.NotifyPointStateRequest, statesync_pb.NotifyPointStateResponse>;
}

export class ControlClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  updatePointState(argument: statesync_pb.UpdatePointStateRequest, callback: grpc.requestCallback<statesync_pb.UpdatePointStateResponse>): grpc.ClientUnaryCall;
  updatePointState(argument: statesync_pb.UpdatePointStateRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<statesync_pb.UpdatePointStateResponse>): grpc.ClientUnaryCall;
  updatePointState(argument: statesync_pb.UpdatePointStateRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<statesync_pb.UpdatePointStateResponse>): grpc.ClientUnaryCall;
  notifyPointState(argument: statesync_pb.NotifyPointStateRequest, callback: grpc.requestCallback<statesync_pb.NotifyPointStateResponse>): grpc.ClientUnaryCall;
  notifyPointState(argument: statesync_pb.NotifyPointStateRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<statesync_pb.NotifyPointStateResponse>): grpc.ClientUnaryCall;
  notifyPointState(argument: statesync_pb.NotifyPointStateRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<statesync_pb.NotifyPointStateResponse>): grpc.ClientUnaryCall;
}
