// GENERATED CODE -- DO NOT EDIT!

// package: 
// file: block.proto

import * as block_pb from "./block_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as grpc from "@grpc/grpc-js";

interface IBlockStateSyncService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  notifyState: grpc.MethodDefinition<google_protobuf_empty_pb.Empty, block_pb.BlockStateList>;
}

export const BlockStateSyncService: IBlockStateSyncService;

export interface IBlockStateSyncServer extends grpc.UntypedServiceImplementation {
  notifyState: grpc.handleUnaryCall<google_protobuf_empty_pb.Empty, block_pb.BlockStateList>;
}

export class BlockStateSyncClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  notifyState(argument: google_protobuf_empty_pb.Empty, callback: grpc.requestCallback<block_pb.BlockStateList>): grpc.ClientUnaryCall;
  notifyState(argument: google_protobuf_empty_pb.Empty, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<block_pb.BlockStateList>): grpc.ClientUnaryCall;
  notifyState(argument: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<block_pb.BlockStateList>): grpc.ClientUnaryCall;
}
