// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var block_pb = require('./block_pb.js');
var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');

function serialize_BlockStateList(arg) {
  if (!(arg instanceof block_pb.BlockStateList)) {
    throw new Error('Expected argument of type BlockStateList');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_BlockStateList(buffer_arg) {
  return block_pb.BlockStateList.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_google_protobuf_Empty(arg) {
  if (!(arg instanceof google_protobuf_empty_pb.Empty)) {
    throw new Error('Expected argument of type google.protobuf.Empty');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_google_protobuf_Empty(buffer_arg) {
  return google_protobuf_empty_pb.Empty.deserializeBinary(new Uint8Array(buffer_arg));
}


var BlockStateSyncService = exports.BlockStateSyncService = {
  notifyState: {
    path: '/BlockStateSync/NotifyState',
    requestStream: false,
    responseStream: false,
    requestType: google_protobuf_empty_pb.Empty,
    responseType: block_pb.BlockStateList,
    requestSerialize: serialize_google_protobuf_Empty,
    requestDeserialize: deserialize_google_protobuf_Empty,
    responseSerialize: serialize_BlockStateList,
    responseDeserialize: deserialize_BlockStateList,
  },
};

exports.BlockStateSyncClient = grpc.makeGenericClientConstructor(BlockStateSyncService);
