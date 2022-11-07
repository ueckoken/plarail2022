// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var block_pb = require('./block_pb.js');

function serialize_BlockStateList(arg) {
  if (!(arg instanceof block_pb.BlockStateList)) {
    throw new Error('Expected argument of type BlockStateList');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_BlockStateList(buffer_arg) {
  return block_pb.BlockStateList.deserializeBinary(new Uint8Array(buffer_arg));
}


var BlockStateSyncService = exports.BlockStateSyncService = {
  notifyState: {
    path: '/BlockStateSync/NotifyState',
    requestStream: false,
    responseStream: false,
    requestType: block_pb.BlockStateList,
    responseType: block_pb.BlockStateList,
    requestSerialize: serialize_BlockStateList,
    requestDeserialize: deserialize_BlockStateList,
    responseSerialize: serialize_BlockStateList,
    responseDeserialize: deserialize_BlockStateList,
  },
};

exports.BlockStateSyncClient = grpc.makeGenericClientConstructor(BlockStateSyncService);
