// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var block_pb = require('./block_pb.js');

function serialize_NotifyStateRequest(arg) {
  if (!(arg instanceof block_pb.NotifyStateRequest)) {
    throw new Error('Expected argument of type NotifyStateRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_NotifyStateRequest(buffer_arg) {
  return block_pb.NotifyStateRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_NotifyStateResponse(arg) {
  if (!(arg instanceof block_pb.NotifyStateResponse)) {
    throw new Error('Expected argument of type NotifyStateResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_NotifyStateResponse(buffer_arg) {
  return block_pb.NotifyStateResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var BlockStateSyncService = exports.BlockStateSyncService = {
  notifyState: {
    path: '/BlockStateSync/NotifyState',
    requestStream: false,
    responseStream: false,
    requestType: block_pb.NotifyStateRequest,
    responseType: block_pb.NotifyStateResponse,
    requestSerialize: serialize_NotifyStateRequest,
    requestDeserialize: deserialize_NotifyStateRequest,
    responseSerialize: serialize_NotifyStateResponse,
    responseDeserialize: deserialize_NotifyStateResponse,
  },
};

exports.BlockStateSyncClient = grpc.makeGenericClientConstructor(BlockStateSyncService);
