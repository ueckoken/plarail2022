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

function serialize_UpdateBlockStateRequest(arg) {
  if (!(arg instanceof block_pb.UpdateBlockStateRequest)) {
    throw new Error('Expected argument of type UpdateBlockStateRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_UpdateBlockStateRequest(buffer_arg) {
  return block_pb.UpdateBlockStateRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_UpdateBlockStateResponse(arg) {
  if (!(arg instanceof block_pb.UpdateBlockStateResponse)) {
    throw new Error('Expected argument of type UpdateBlockStateResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_UpdateBlockStateResponse(buffer_arg) {
  return block_pb.UpdateBlockStateResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// BlockStateManagerはblock状態を管理します。externalで動作します。
var BlockStateManagerService = exports.BlockStateManagerService = {
  // UpdateBlockStateはサーバのBlockStateを変更します。
updateBlockState: {
    path: '/BlockStateManager/UpdateBlockState',
    requestStream: false,
    responseStream: false,
    requestType: block_pb.UpdateBlockStateRequest,
    responseType: block_pb.UpdateBlockStateResponse,
    requestSerialize: serialize_UpdateBlockStateRequest,
    requestDeserialize: deserialize_UpdateBlockStateRequest,
    responseSerialize: serialize_UpdateBlockStateResponse,
    responseDeserialize: deserialize_UpdateBlockStateResponse,
  },
};

exports.BlockStateManagerClient = grpc.makeGenericClientConstructor(BlockStateManagerService);
// BlockStateNotificationはblock状態の通知を受けます。auto_operationで動作します。
var BlockStateNotificationService = exports.BlockStateNotificationService = {
  notifyBlockState: {
    path: '/BlockStateNotification/NotifyBlockState',
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

exports.BlockStateNotificationClient = grpc.makeGenericClientConstructor(BlockStateNotificationService);
