# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: ats.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\tats.proto\"\x84\x03\n\x11SendStatusRequest\x12-\n\x06sensor\x18\x01 \x01(\x0e\x32\x1d.SendStatusRequest.SensorName\"\xbf\x02\n\nSensorName\x12\x0b\n\x07unknown\x10\x00\x12\x0f\n\x0bshinjuku_d1\x10\x01\x12\x0f\n\x0bshinjuku_d2\x10\x02\x12\x12\n\x0esakurajosui_d1\x10\x0b\x12\x12\n\x0esakurajosui_d2\x10\x0c\x12\x12\n\x0esakurajosui_d3\x10\r\x12\x12\n\x0esakurajosui_d4\x10\x0e\x12\x12\n\x0esakurajosui_d5\x10\x0f\x12\x12\n\x0esakurajosui_d6\x10\x10\x12\x0c\n\x08\x63hofu_d1\x10\x15\x12\x0c\n\x08\x63hofu_d2\x10\x16\x12\x0c\n\x08\x63hofu_d3\x10\x17\x12\x0c\n\x08\x63hofu_d4\x10\x18\x12\x0c\n\x08\x63hofu_d5\x10\x19\x12\x10\n\x0chashimoto_d1\x10\x1f\x12\x10\n\x0chashimoto_d2\x10 \x12\x0f\n\x0bhachioji_d1\x10)\x12\x0f\n\x0bhachioji_d2\x10*\"v\n\x12SendStatusResponse\x12.\n\x08response\x18\x01 \x01(\x0e\x32\x1c.SendStatusResponse.Response\"0\n\x08Response\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x0b\n\x07SUCCESS\x10\x01\x12\n\n\x06\x46\x41ILED\x10\x02\x32>\n\x03\x41ts\x12\x37\n\nSendStatus\x12\x12.SendStatusRequest\x1a\x13.SendStatusResponse\"\x00\x42\x08Z\x06./specb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'ats_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\006./spec'
  _globals['_SENDSTATUSREQUEST']._serialized_start=14
  _globals['_SENDSTATUSREQUEST']._serialized_end=402
  _globals['_SENDSTATUSREQUEST_SENSORNAME']._serialized_start=83
  _globals['_SENDSTATUSREQUEST_SENSORNAME']._serialized_end=402
  _globals['_SENDSTATUSRESPONSE']._serialized_start=404
  _globals['_SENDSTATUSRESPONSE']._serialized_end=522
  _globals['_SENDSTATUSRESPONSE_RESPONSE']._serialized_start=474
  _globals['_SENDSTATUSRESPONSE_RESPONSE']._serialized_end=522
  _globals['_ATS']._serialized_start=524
  _globals['_ATS']._serialized_end=586
# @@protoc_insertion_point(module_scope)
