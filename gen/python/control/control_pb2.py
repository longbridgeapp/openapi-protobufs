# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: control/control.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x15\x63ontrol/control.proto\x12\x18longbridgeapp.control.v1\"\xe4\x01\n\x05\x43lose\x12\x38\n\x04\x63ode\x18\x01 \x01(\x0e\x32$.longbridgeapp.control.v1.Close.CodeR\x04\x63ode\x12\x16\n\x06reason\x18\x02 \x01(\tR\x06reason\"\x88\x01\n\x04\x43ode\x12\x14\n\x10HeartbeatTimeout\x10\x00\x12\x0f\n\x0bServerError\x10\x01\x12\x12\n\x0eServerShutdown\x10\x02\x12\x0f\n\x0bUnpackError\x10\x03\x12\r\n\tAuthError\x10\x04\x12\x0f\n\x0bSessExpired\x10\x05\x12\x14\n\x10\x43onnectDuplicate\x10\x06\"b\n\tHeartbeat\x12\x1c\n\ttimestamp\x18\x01 \x01(\x03R\ttimestamp\x12&\n\x0cheartbeat_id\x18\x02 \x01(\x05H\x00R\x0bheartbeatId\x88\x01\x01\x42\x0f\n\r_heartbeat_id\"#\n\x0b\x41uthRequest\x12\x14\n\x05token\x18\x01 \x01(\tR\x05token\"G\n\x0c\x41uthResponse\x12\x1d\n\nsession_id\x18\x01 \x01(\tR\tsessionId\x12\x18\n\x07\x65xpires\x18\x02 \x01(\x03R\x07\x65xpires\"1\n\x10ReconnectRequest\x12\x1d\n\nsession_id\x18\x01 \x01(\tR\tsessionId\"L\n\x11ReconnectResponse\x12\x1d\n\nsession_id\x18\x01 \x01(\tR\tsessionId\x12\x18\n\x07\x65xpires\x18\x02 \x01(\x03R\x07\x65xpires*L\n\x07\x43ommand\x12\r\n\tCMD_CLOSE\x10\x00\x12\x11\n\rCMD_HEARTBEAT\x10\x01\x12\x0c\n\x08\x43MD_AUTH\x10\x02\x12\x11\n\rCMD_RECONNECT\x10\x03\x42\xf3\x01\n\x1c\x63om.longbridgeapp.control.v1B\x0c\x43ontrolProtoP\x01ZCgithub.com/longbridgeapp/openapi-protobufs/gen/go/control;controlv1\xa2\x02\x03LCX\xaa\x02\x18Longbridgeapp.Control.V1\xca\x02\x18Longbridgeapp\\Control\\V1\xe2\x02$Longbridgeapp\\Control\\V1\\GPBMetadata\xea\x02\x1aLongbridgeapp::Control::V1b\x06proto3')

_COMMAND = DESCRIPTOR.enum_types_by_name['Command']
Command = enum_type_wrapper.EnumTypeWrapper(_COMMAND)
CMD_CLOSE = 0
CMD_HEARTBEAT = 1
CMD_AUTH = 2
CMD_RECONNECT = 3


_CLOSE = DESCRIPTOR.message_types_by_name['Close']
_HEARTBEAT = DESCRIPTOR.message_types_by_name['Heartbeat']
_AUTHREQUEST = DESCRIPTOR.message_types_by_name['AuthRequest']
_AUTHRESPONSE = DESCRIPTOR.message_types_by_name['AuthResponse']
_RECONNECTREQUEST = DESCRIPTOR.message_types_by_name['ReconnectRequest']
_RECONNECTRESPONSE = DESCRIPTOR.message_types_by_name['ReconnectResponse']
_CLOSE_CODE = _CLOSE.enum_types_by_name['Code']
Close = _reflection.GeneratedProtocolMessageType('Close', (_message.Message,), {
  'DESCRIPTOR' : _CLOSE,
  '__module__' : 'control.control_pb2'
  # @@protoc_insertion_point(class_scope:longbridgeapp.control.v1.Close)
  })
_sym_db.RegisterMessage(Close)

Heartbeat = _reflection.GeneratedProtocolMessageType('Heartbeat', (_message.Message,), {
  'DESCRIPTOR' : _HEARTBEAT,
  '__module__' : 'control.control_pb2'
  # @@protoc_insertion_point(class_scope:longbridgeapp.control.v1.Heartbeat)
  })
_sym_db.RegisterMessage(Heartbeat)

AuthRequest = _reflection.GeneratedProtocolMessageType('AuthRequest', (_message.Message,), {
  'DESCRIPTOR' : _AUTHREQUEST,
  '__module__' : 'control.control_pb2'
  # @@protoc_insertion_point(class_scope:longbridgeapp.control.v1.AuthRequest)
  })
_sym_db.RegisterMessage(AuthRequest)

AuthResponse = _reflection.GeneratedProtocolMessageType('AuthResponse', (_message.Message,), {
  'DESCRIPTOR' : _AUTHRESPONSE,
  '__module__' : 'control.control_pb2'
  # @@protoc_insertion_point(class_scope:longbridgeapp.control.v1.AuthResponse)
  })
_sym_db.RegisterMessage(AuthResponse)

ReconnectRequest = _reflection.GeneratedProtocolMessageType('ReconnectRequest', (_message.Message,), {
  'DESCRIPTOR' : _RECONNECTREQUEST,
  '__module__' : 'control.control_pb2'
  # @@protoc_insertion_point(class_scope:longbridgeapp.control.v1.ReconnectRequest)
  })
_sym_db.RegisterMessage(ReconnectRequest)

ReconnectResponse = _reflection.GeneratedProtocolMessageType('ReconnectResponse', (_message.Message,), {
  'DESCRIPTOR' : _RECONNECTRESPONSE,
  '__module__' : 'control.control_pb2'
  # @@protoc_insertion_point(class_scope:longbridgeapp.control.v1.ReconnectResponse)
  })
_sym_db.RegisterMessage(ReconnectResponse)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\034com.longbridgeapp.control.v1B\014ControlProtoP\001ZCgithub.com/longbridgeapp/openapi-protobufs/gen/go/control;controlv1\242\002\003LCX\252\002\030Longbridgeapp.Control.V1\312\002\030Longbridgeapp\\Control\\V1\342\002$Longbridgeapp\\Control\\V1\\GPBMetadata\352\002\032Longbridgeapp::Control::V1'
  _COMMAND._serialized_start=621
  _COMMAND._serialized_end=697
  _CLOSE._serialized_start=52
  _CLOSE._serialized_end=280
  _CLOSE_CODE._serialized_start=144
  _CLOSE_CODE._serialized_end=280
  _HEARTBEAT._serialized_start=282
  _HEARTBEAT._serialized_end=380
  _AUTHREQUEST._serialized_start=382
  _AUTHREQUEST._serialized_end=417
  _AUTHRESPONSE._serialized_start=419
  _AUTHRESPONSE._serialized_end=490
  _RECONNECTREQUEST._serialized_start=492
  _RECONNECTREQUEST._serialized_end=541
  _RECONNECTRESPONSE._serialized_start=543
  _RECONNECTRESPONSE._serialized_end=619
# @@protoc_insertion_point(module_scope)
