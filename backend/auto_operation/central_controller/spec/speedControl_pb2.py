# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: speedControl.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x12speedControl.proto\"\xb8\x01\n\tSendSpeed\x12\r\n\x05Speed\x18\x01 \x01(\x05\x12\x1f\n\x05train\x18\x02 \x01(\x0e\x32\x10.SendSpeed.Train\"{\n\x05Train\x12\x0b\n\x07UNKNOWN\x10\x00\x12\t\n\x05TAKAO\x10\x01\x12\x0c\n\x08\x43HICHIBU\x10\x02\x12\n\n\x06HAKONE\x10\x03\x12\x0b\n\x07OKUTAMA\x10\x04\x12\t\n\x05NIKKO\x10\x05\x12\x0c\n\x08\x45NOSHIMA\x10\x06\x12\x0c\n\x08KAMAKURA\x10\x07\x12\x0c\n\x08YOKOSUKA\x10\x08\"Z\n\nStatusCode\x12\x1e\n\x04\x63ode\x18\x01 \x01(\x0e\x32\x10.StatusCode.Code\",\n\x04\x43ode\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x0b\n\x07SUCCESS\x10\x01\x12\n\n\x06\x46\x41ILED\x10\x02\x32\x32\n\x05Speed\x12)\n\x0c\x43ontrolSpeed\x12\n.SendSpeed\x1a\x0b.StatusCode\"\x00\x42\x08Z\x06./specb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'speedControl_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\006./spec'
  _SENDSPEED._serialized_start=23
  _SENDSPEED._serialized_end=207
  _SENDSPEED_TRAIN._serialized_start=84
  _SENDSPEED_TRAIN._serialized_end=207
  _STATUSCODE._serialized_start=209
  _STATUSCODE._serialized_end=299
  _STATUSCODE_CODE._serialized_start=255
  _STATUSCODE_CODE._serialized_end=299
  _SPEED._serialized_start=301
  _SPEED._serialized_end=351
# @@protoc_insertion_point(module_scope)