// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var statesync_pb = require('./statesync_pb.js');

function serialize_RequestSync(arg) {
  if (!(arg instanceof statesync_pb.RequestSync)) {
    throw new Error('Expected argument of type RequestSync');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_RequestSync(buffer_arg) {
  return statesync_pb.RequestSync.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ResponseSync(arg) {
  if (!(arg instanceof statesync_pb.ResponseSync)) {
    throw new Error('Expected argument of type ResponseSync');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ResponseSync(buffer_arg) {
  return statesync_pb.ResponseSync.deserializeBinary(new Uint8Array(buffer_arg));
}


var ControlService = exports.ControlService = {
  command2Internal: {
    path: '/Control/Command2Internal',
    requestStream: false,
    responseStream: false,
    requestType: statesync_pb.RequestSync,
    responseType: statesync_pb.ResponseSync,
    requestSerialize: serialize_RequestSync,
    requestDeserialize: deserialize_RequestSync,
    responseSerialize: serialize_ResponseSync,
    responseDeserialize: deserialize_ResponseSync,
  },
};

exports.ControlClient = grpc.makeGenericClientConstructor(ControlService);
