// GENERATED CODE -- DO NOT EDIT!

// package: 
// file: block.proto

import * as block_pb from "./block_pb";
import * as grpc from "@grpc/grpc-js";

interface IBlockStateManagerService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  updateBlockState: grpc.MethodDefinition<block_pb.UpdateBlockStateRequest, block_pb.UpdateBlockStateResponse>;
}

export const BlockStateManagerService: IBlockStateManagerService;

export interface IBlockStateManagerServer extends grpc.UntypedServiceImplementation {
  updateBlockState: grpc.handleUnaryCall<block_pb.UpdateBlockStateRequest, block_pb.UpdateBlockStateResponse>;
}

export class BlockStateManagerClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  updateBlockState(argument: block_pb.UpdateBlockStateRequest, callback: grpc.requestCallback<block_pb.UpdateBlockStateResponse>): grpc.ClientUnaryCall;
  updateBlockState(argument: block_pb.UpdateBlockStateRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<block_pb.UpdateBlockStateResponse>): grpc.ClientUnaryCall;
  updateBlockState(argument: block_pb.UpdateBlockStateRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<block_pb.UpdateBlockStateResponse>): grpc.ClientUnaryCall;
}

interface IBlockStateNotificationService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  notifyBlockState: grpc.MethodDefinition<block_pb.NotifyBlockStateRequest, block_pb.NotifyBlockStateResponse>;
}

export const BlockStateNotificationService: IBlockStateNotificationService;

export interface IBlockStateNotificationServer extends grpc.UntypedServiceImplementation {
  notifyBlockState: grpc.handleUnaryCall<block_pb.NotifyBlockStateRequest, block_pb.NotifyBlockStateResponse>;
}

export class BlockStateNotificationClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  notifyBlockState(argument: block_pb.NotifyBlockStateRequest, callback: grpc.requestCallback<block_pb.NotifyBlockStateResponse>): grpc.ClientUnaryCall;
  notifyBlockState(argument: block_pb.NotifyBlockStateRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<block_pb.NotifyBlockStateResponse>): grpc.ClientUnaryCall;
  notifyBlockState(argument: block_pb.NotifyBlockStateRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<block_pb.NotifyBlockStateResponse>): grpc.ClientUnaryCall;
}
