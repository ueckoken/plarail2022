// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var statesync_pb = require('./statesync_pb.js');

function serialize_NotifyPointStateRequest(arg) {
  if (!(arg instanceof statesync_pb.NotifyPointStateRequest)) {
    throw new Error('Expected argument of type NotifyPointStateRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_NotifyPointStateRequest(buffer_arg) {
  return statesync_pb.NotifyPointStateRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_NotifyPointStateResponse(arg) {
  if (!(arg instanceof statesync_pb.NotifyPointStateResponse)) {
    throw new Error('Expected argument of type NotifyPointStateResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_NotifyPointStateResponse(buffer_arg) {
  return statesync_pb.NotifyPointStateResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_UpdatePointStateRequest(arg) {
  if (!(arg instanceof statesync_pb.UpdatePointStateRequest)) {
    throw new Error('Expected argument of type UpdatePointStateRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_UpdatePointStateRequest(buffer_arg) {
  return statesync_pb.UpdatePointStateRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_UpdatePointStateResponse(arg) {
  if (!(arg instanceof statesync_pb.UpdatePointStateResponse)) {
    throw new Error('Expected argument of type UpdatePointStateResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_UpdatePointStateResponse(buffer_arg) {
  return statesync_pb.UpdatePointStateResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// StateManagerはexternalにて立ち上げます。ポイント情報などを管理します。
var StateManagerService = exports.StateManagerService = {
  // UpdatePointStateはexternalへPointState更新要求を送る。
updatePointState: {
    path: '/StateManager/UpdatePointState',
    requestStream: false,
    responseStream: false,
    requestType: statesync_pb.UpdatePointStateRequest,
    responseType: statesync_pb.UpdatePointStateResponse,
    requestSerialize: serialize_UpdatePointStateRequest,
    requestDeserialize: deserialize_UpdatePointStateRequest,
    responseSerialize: serialize_UpdatePointStateResponse,
    responseDeserialize: deserialize_UpdatePointStateResponse,
  },
};

exports.StateManagerClient = grpc.makeGenericClientConstructor(StateManagerService);
// Notificationはauto_operationなどで立ち上げます。StateManagerからポイント情報などの変更を受けとります。
var PointStateNotificationService = exports.PointStateNotificationService = {
  // NotifyPointStateはexternalからauto-operationやinternalへPointStateの更新情報を伝える。
notifyPointState: {
    path: '/PointStateNotification/NotifyPointState',
    requestStream: false,
    responseStream: false,
    requestType: statesync_pb.NotifyPointStateRequest,
    responseType: statesync_pb.NotifyPointStateResponse,
    requestSerialize: serialize_NotifyPointStateRequest,
    requestDeserialize: deserialize_NotifyPointStateRequest,
    responseSerialize: serialize_NotifyPointStateResponse,
    responseDeserialize: deserialize_NotifyPointStateResponse,
  },
};

exports.PointStateNotificationClient = grpc.makeGenericClientConstructor(PointStateNotificationService);
