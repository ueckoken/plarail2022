// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var block_pb = require('./block_pb.js');

function serialize_NotifyBlockStateRequest(arg) {
  if (!(arg instanceof block_pb.NotifyBlockStateRequest)) {
    throw new Error('Expected argument of type NotifyBlockStateRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_NotifyBlockStateRequest(buffer_arg) {
  return block_pb.NotifyBlockStateRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_NotifyBlockStateResponse(arg) {
  if (!(arg instanceof block_pb.NotifyBlockStateResponse)) {
    throw new Error('Expected argument of type NotifyBlockStateResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_NotifyBlockStateResponse(buffer_arg) {
  return block_pb.NotifyBlockStateResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var BlockStateSyncService = exports.BlockStateSyncService = {
  notifyState: {
    path: '/BlockStateSync/NotifyState',
    requestStream: false,
    responseStream: false,
    requestType: block_pb.NotifyBlockStateRequest,
    responseType: block_pb.NotifyBlockStateResponse,
    requestSerialize: serialize_NotifyBlockStateRequest,
    requestDeserialize: deserialize_NotifyBlockStateRequest,
    responseSerialize: serialize_NotifyBlockStateResponse,
    responseDeserialize: deserialize_NotifyBlockStateResponse,
  },
};

exports.BlockStateSyncClient = grpc.makeGenericClientConstructor(BlockStateSyncService);
