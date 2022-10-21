// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var ats_pb = require('./ats_pb.js');

function serialize_SendStatusRequest(arg) {
  if (!(arg instanceof ats_pb.SendStatusRequest)) {
    throw new Error('Expected argument of type SendStatusRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_SendStatusRequest(buffer_arg) {
  return ats_pb.SendStatusRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_SendStatusResponse(arg) {
  if (!(arg instanceof ats_pb.SendStatusResponse)) {
    throw new Error('Expected argument of type SendStatusResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_SendStatusResponse(buffer_arg) {
  return ats_pb.SendStatusResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var AtsService = exports.AtsService = {
  sendStatus: {
    path: '/Ats/SendStatus',
    requestStream: false,
    responseStream: false,
    requestType: ats_pb.SendStatusRequest,
    responseType: ats_pb.SendStatusResponse,
    requestSerialize: serialize_SendStatusRequest,
    requestDeserialize: deserialize_SendStatusRequest,
    responseSerialize: serialize_SendStatusResponse,
    responseDeserialize: deserialize_SendStatusResponse,
  },
};

exports.AtsClient = grpc.makeGenericClientConstructor(AtsService);
