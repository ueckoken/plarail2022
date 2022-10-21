// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var speedControl_pb = require('./speedControl_pb.js');

function serialize_SendSpeed(arg) {
  if (!(arg instanceof speedControl_pb.SendSpeed)) {
    throw new Error('Expected argument of type SendSpeed');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_SendSpeed(buffer_arg) {
  return speedControl_pb.SendSpeed.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_StatusCode(arg) {
  if (!(arg instanceof speedControl_pb.StatusCode)) {
    throw new Error('Expected argument of type StatusCode');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_StatusCode(buffer_arg) {
  return speedControl_pb.StatusCode.deserializeBinary(new Uint8Array(buffer_arg));
}


var SpeedService = exports.SpeedService = {
  controlSpeed: {
    path: '/Speed/ControlSpeed',
    requestStream: false,
    responseStream: false,
    requestType: speedControl_pb.SendSpeed,
    responseType: speedControl_pb.StatusCode,
    requestSerialize: serialize_SendSpeed,
    requestDeserialize: deserialize_SendSpeed,
    responseSerialize: serialize_StatusCode,
    responseDeserialize: deserialize_StatusCode,
  },
};

exports.SpeedClient = grpc.makeGenericClientConstructor(SpeedService);
