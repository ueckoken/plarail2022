// GENERATED CODE -- DO NOT EDIT!

// package: 
// file: block.proto

import * as block_pb from "./block_pb";
import * as grpc from "@grpc/grpc-js";

interface IBlockStateSyncService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  notifyState: grpc.MethodDefinition<block_pb.NotifyBlockStateRequest, block_pb.NotifyBlockStateResponse>;
}

export const BlockStateSyncService: IBlockStateSyncService;

export interface IBlockStateSyncServer extends grpc.UntypedServiceImplementation {
  notifyState: grpc.handleUnaryCall<block_pb.NotifyBlockStateRequest, block_pb.NotifyBlockStateResponse>;
}

export class BlockStateSyncClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  notifyState(argument: block_pb.NotifyBlockStateRequest, callback: grpc.requestCallback<block_pb.NotifyBlockStateResponse>): grpc.ClientUnaryCall;
  notifyState(argument: block_pb.NotifyBlockStateRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<block_pb.NotifyBlockStateResponse>): grpc.ClientUnaryCall;
  notifyState(argument: block_pb.NotifyBlockStateRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<block_pb.NotifyBlockStateResponse>): grpc.ClientUnaryCall;
}
