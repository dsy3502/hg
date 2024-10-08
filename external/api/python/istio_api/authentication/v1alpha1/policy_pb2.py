# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: authentication/v1alpha1/policy.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import field_behavior_pb2 as google_dot_api_dot_field__behavior__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='authentication/v1alpha1/policy.proto',
  package='istio.authentication.v1alpha1',
  syntax='proto3',
  serialized_options=b'Z$istio.io/api/authentication/v1alpha1',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n$authentication/v1alpha1/policy.proto\x12\x1distio.authentication.v1alpha1\x1a\x1fgoogle/api/field_behavior.proto\"\x7f\n\x0bStringMatch\x12\x16\n\x05\x65xact\x18\x01 \x01(\tH\x00R\x05\x65xact\x12\x18\n\x06prefix\x18\x02 \x01(\tH\x00R\x06prefix\x12\x18\n\x06suffix\x18\x03 \x01(\tH\x00R\x06suffix\x12\x16\n\x05regex\x18\x04 \x01(\tH\x00R\x05regexB\x0c\n\nmatch_type\"\x93\x01\n\tMutualTls\x12\x1f\n\tallow_tls\x18\x01 \x01(\x08\x42\x02\x18\x01R\x08\x61llowTls\x12\x41\n\x04mode\x18\x02 \x01(\x0e\x32-.istio.authentication.v1alpha1.MutualTls.ModeR\x04mode\"\"\n\x04Mode\x12\n\n\x06STRICT\x10\x00\x12\x0e\n\nPERMISSIVE\x10\x01\"\xb5\x03\n\x03Jwt\x12\x16\n\x06issuer\x18\x01 \x01(\tR\x06issuer\x12\x1c\n\taudiences\x18\x02 \x03(\tR\taudiences\x12\x19\n\x08jwks_uri\x18\x03 \x01(\tR\x07jwksUri\x12\x12\n\x04jwks\x18\n \x01(\tR\x04jwks\x12\x1f\n\x0bjwt_headers\x18\x06 \x03(\tR\njwtHeaders\x12\x1d\n\njwt_params\x18\x07 \x03(\tR\tjwtParams\x12S\n\rtrigger_rules\x18\t \x03(\x0b\x32..istio.authentication.v1alpha1.Jwt.TriggerRuleR\x0ctriggerRules\x1a\xb3\x01\n\x0bTriggerRule\x12Q\n\x0e\x65xcluded_paths\x18\x01 \x03(\x0b\x32*.istio.authentication.v1alpha1.StringMatchR\rexcludedPaths\x12Q\n\x0eincluded_paths\x18\x02 \x03(\x0b\x32*.istio.authentication.v1alpha1.StringMatchR\rincludedPaths\"\xa0\x01\n\x18PeerAuthenticationMethod\x12>\n\x04mtls\x18\x01 \x01(\x0b\x32(.istio.authentication.v1alpha1.MutualTlsH\x00R\x04mtls\x12:\n\x03jwt\x18\x02 \x01(\x0b\x32\".istio.authentication.v1alpha1.JwtB\x02\x18\x01H\x00R\x03jwtB\x08\n\x06params\"R\n\x1aOriginAuthenticationMethod\x12\x34\n\x03jwt\x18\x01 \x01(\x0b\x32\".istio.authentication.v1alpha1.JwtR\x03jwt\"\xbf\x03\n\x06Policy\x12K\n\x07targets\x18\x01 \x03(\x0b\x32-.istio.authentication.v1alpha1.TargetSelectorB\x02\x18\x01R\x07targets\x12M\n\x05peers\x18\x02 \x03(\x0b\x32\x37.istio.authentication.v1alpha1.PeerAuthenticationMethodR\x05peers\x12,\n\x10peer_is_optional\x18\x03 \x01(\x08\x42\x02\x18\x01R\x0epeerIsOptional\x12W\n\x07origins\x18\x04 \x03(\x0b\x32\x39.istio.authentication.v1alpha1.OriginAuthenticationMethodB\x02\x18\x01R\x07origins\x12\x30\n\x12origin_is_optional\x18\x05 \x01(\x08\x42\x02\x18\x01R\x10originIsOptional\x12`\n\x11principal_binding\x18\x06 \x01(\x0e\x32/.istio.authentication.v1alpha1.PrincipalBindingB\x02\x18\x01R\x10principalBinding\"{\n\x0eTargetSelector\x12\x18\n\x04name\x18\x01 \x01(\tB\x04\xe2\x41\x01\x02R\x04name\x12\x41\n\x05ports\x18\x02 \x03(\x0b\x32+.istio.authentication.v1alpha1.PortSelectorR\x05portsJ\x04\x08\x03\x10\x04R\x06labels\"F\n\x0cPortSelector\x12\x18\n\x06number\x18\x01 \x01(\rH\x00R\x06number\x12\x14\n\x04name\x18\x02 \x01(\tH\x00R\x04nameB\x06\n\x04port*0\n\x10PrincipalBinding\x12\x0c\n\x08USE_PEER\x10\x00\x12\x0e\n\nUSE_ORIGIN\x10\x01\x42&Z$istio.io/api/authentication/v1alpha1b\x06proto3'
  ,
  dependencies=[google_dot_api_dot_field__behavior__pb2.DESCRIPTOR,])

_PRINCIPALBINDING = _descriptor.EnumDescriptor(
  name='PrincipalBinding',
  full_name='istio.authentication.v1alpha1.PrincipalBinding',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='USE_PEER', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='USE_ORIGIN', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1717,
  serialized_end=1765,
)
_sym_db.RegisterEnumDescriptor(_PRINCIPALBINDING)

PrincipalBinding = enum_type_wrapper.EnumTypeWrapper(_PRINCIPALBINDING)
USE_PEER = 0
USE_ORIGIN = 1


_MUTUALTLS_MODE = _descriptor.EnumDescriptor(
  name='Mode',
  full_name='istio.authentication.v1alpha1.MutualTls.Mode',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='STRICT', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='PERMISSIVE', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=347,
  serialized_end=381,
)
_sym_db.RegisterEnumDescriptor(_MUTUALTLS_MODE)


_STRINGMATCH = _descriptor.Descriptor(
  name='StringMatch',
  full_name='istio.authentication.v1alpha1.StringMatch',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='exact', full_name='istio.authentication.v1alpha1.StringMatch.exact', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='exact', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='prefix', full_name='istio.authentication.v1alpha1.StringMatch.prefix', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='prefix', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='suffix', full_name='istio.authentication.v1alpha1.StringMatch.suffix', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='suffix', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='regex', full_name='istio.authentication.v1alpha1.StringMatch.regex', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='regex', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='match_type', full_name='istio.authentication.v1alpha1.StringMatch.match_type',
      index=0, containing_type=None,
      create_key=_descriptor._internal_create_key,
    fields=[]),
  ],
  serialized_start=104,
  serialized_end=231,
)


_MUTUALTLS = _descriptor.Descriptor(
  name='MutualTls',
  full_name='istio.authentication.v1alpha1.MutualTls',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='allow_tls', full_name='istio.authentication.v1alpha1.MutualTls.allow_tls', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\030\001', json_name='allowTls', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='mode', full_name='istio.authentication.v1alpha1.MutualTls.mode', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='mode', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _MUTUALTLS_MODE,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=234,
  serialized_end=381,
)


_JWT_TRIGGERRULE = _descriptor.Descriptor(
  name='TriggerRule',
  full_name='istio.authentication.v1alpha1.Jwt.TriggerRule',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='excluded_paths', full_name='istio.authentication.v1alpha1.Jwt.TriggerRule.excluded_paths', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='excludedPaths', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='included_paths', full_name='istio.authentication.v1alpha1.Jwt.TriggerRule.included_paths', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='includedPaths', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=642,
  serialized_end=821,
)

_JWT = _descriptor.Descriptor(
  name='Jwt',
  full_name='istio.authentication.v1alpha1.Jwt',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='issuer', full_name='istio.authentication.v1alpha1.Jwt.issuer', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='issuer', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='audiences', full_name='istio.authentication.v1alpha1.Jwt.audiences', index=1,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='audiences', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='jwks_uri', full_name='istio.authentication.v1alpha1.Jwt.jwks_uri', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='jwksUri', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='jwks', full_name='istio.authentication.v1alpha1.Jwt.jwks', index=3,
      number=10, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='jwks', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='jwt_headers', full_name='istio.authentication.v1alpha1.Jwt.jwt_headers', index=4,
      number=6, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='jwtHeaders', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='jwt_params', full_name='istio.authentication.v1alpha1.Jwt.jwt_params', index=5,
      number=7, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='jwtParams', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='trigger_rules', full_name='istio.authentication.v1alpha1.Jwt.trigger_rules', index=6,
      number=9, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='triggerRules', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[_JWT_TRIGGERRULE, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=384,
  serialized_end=821,
)


_PEERAUTHENTICATIONMETHOD = _descriptor.Descriptor(
  name='PeerAuthenticationMethod',
  full_name='istio.authentication.v1alpha1.PeerAuthenticationMethod',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='mtls', full_name='istio.authentication.v1alpha1.PeerAuthenticationMethod.mtls', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='mtls', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='jwt', full_name='istio.authentication.v1alpha1.PeerAuthenticationMethod.jwt', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\030\001', json_name='jwt', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='params', full_name='istio.authentication.v1alpha1.PeerAuthenticationMethod.params',
      index=0, containing_type=None,
      create_key=_descriptor._internal_create_key,
    fields=[]),
  ],
  serialized_start=824,
  serialized_end=984,
)


_ORIGINAUTHENTICATIONMETHOD = _descriptor.Descriptor(
  name='OriginAuthenticationMethod',
  full_name='istio.authentication.v1alpha1.OriginAuthenticationMethod',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='jwt', full_name='istio.authentication.v1alpha1.OriginAuthenticationMethod.jwt', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='jwt', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=986,
  serialized_end=1068,
)


_POLICY = _descriptor.Descriptor(
  name='Policy',
  full_name='istio.authentication.v1alpha1.Policy',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='targets', full_name='istio.authentication.v1alpha1.Policy.targets', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\030\001', json_name='targets', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='peers', full_name='istio.authentication.v1alpha1.Policy.peers', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='peers', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='peer_is_optional', full_name='istio.authentication.v1alpha1.Policy.peer_is_optional', index=2,
      number=3, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\030\001', json_name='peerIsOptional', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='origins', full_name='istio.authentication.v1alpha1.Policy.origins', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\030\001', json_name='origins', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='origin_is_optional', full_name='istio.authentication.v1alpha1.Policy.origin_is_optional', index=4,
      number=5, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\030\001', json_name='originIsOptional', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='principal_binding', full_name='istio.authentication.v1alpha1.Policy.principal_binding', index=5,
      number=6, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\030\001', json_name='principalBinding', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1071,
  serialized_end=1518,
)


_TARGETSELECTOR = _descriptor.Descriptor(
  name='TargetSelector',
  full_name='istio.authentication.v1alpha1.TargetSelector',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='istio.authentication.v1alpha1.TargetSelector.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342A\001\002', json_name='name', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='ports', full_name='istio.authentication.v1alpha1.TargetSelector.ports', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='ports', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1520,
  serialized_end=1643,
)


_PORTSELECTOR = _descriptor.Descriptor(
  name='PortSelector',
  full_name='istio.authentication.v1alpha1.PortSelector',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='number', full_name='istio.authentication.v1alpha1.PortSelector.number', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='number', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='istio.authentication.v1alpha1.PortSelector.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='name', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='port', full_name='istio.authentication.v1alpha1.PortSelector.port',
      index=0, containing_type=None,
      create_key=_descriptor._internal_create_key,
    fields=[]),
  ],
  serialized_start=1645,
  serialized_end=1715,
)

_STRINGMATCH.oneofs_by_name['match_type'].fields.append(
  _STRINGMATCH.fields_by_name['exact'])
_STRINGMATCH.fields_by_name['exact'].containing_oneof = _STRINGMATCH.oneofs_by_name['match_type']
_STRINGMATCH.oneofs_by_name['match_type'].fields.append(
  _STRINGMATCH.fields_by_name['prefix'])
_STRINGMATCH.fields_by_name['prefix'].containing_oneof = _STRINGMATCH.oneofs_by_name['match_type']
_STRINGMATCH.oneofs_by_name['match_type'].fields.append(
  _STRINGMATCH.fields_by_name['suffix'])
_STRINGMATCH.fields_by_name['suffix'].containing_oneof = _STRINGMATCH.oneofs_by_name['match_type']
_STRINGMATCH.oneofs_by_name['match_type'].fields.append(
  _STRINGMATCH.fields_by_name['regex'])
_STRINGMATCH.fields_by_name['regex'].containing_oneof = _STRINGMATCH.oneofs_by_name['match_type']
_MUTUALTLS.fields_by_name['mode'].enum_type = _MUTUALTLS_MODE
_MUTUALTLS_MODE.containing_type = _MUTUALTLS
_JWT_TRIGGERRULE.fields_by_name['excluded_paths'].message_type = _STRINGMATCH
_JWT_TRIGGERRULE.fields_by_name['included_paths'].message_type = _STRINGMATCH
_JWT_TRIGGERRULE.containing_type = _JWT
_JWT.fields_by_name['trigger_rules'].message_type = _JWT_TRIGGERRULE
_PEERAUTHENTICATIONMETHOD.fields_by_name['mtls'].message_type = _MUTUALTLS
_PEERAUTHENTICATIONMETHOD.fields_by_name['jwt'].message_type = _JWT
_PEERAUTHENTICATIONMETHOD.oneofs_by_name['params'].fields.append(
  _PEERAUTHENTICATIONMETHOD.fields_by_name['mtls'])
_PEERAUTHENTICATIONMETHOD.fields_by_name['mtls'].containing_oneof = _PEERAUTHENTICATIONMETHOD.oneofs_by_name['params']
_PEERAUTHENTICATIONMETHOD.oneofs_by_name['params'].fields.append(
  _PEERAUTHENTICATIONMETHOD.fields_by_name['jwt'])
_PEERAUTHENTICATIONMETHOD.fields_by_name['jwt'].containing_oneof = _PEERAUTHENTICATIONMETHOD.oneofs_by_name['params']
_ORIGINAUTHENTICATIONMETHOD.fields_by_name['jwt'].message_type = _JWT
_POLICY.fields_by_name['targets'].message_type = _TARGETSELECTOR
_POLICY.fields_by_name['peers'].message_type = _PEERAUTHENTICATIONMETHOD
_POLICY.fields_by_name['origins'].message_type = _ORIGINAUTHENTICATIONMETHOD
_POLICY.fields_by_name['principal_binding'].enum_type = _PRINCIPALBINDING
_TARGETSELECTOR.fields_by_name['ports'].message_type = _PORTSELECTOR
_PORTSELECTOR.oneofs_by_name['port'].fields.append(
  _PORTSELECTOR.fields_by_name['number'])
_PORTSELECTOR.fields_by_name['number'].containing_oneof = _PORTSELECTOR.oneofs_by_name['port']
_PORTSELECTOR.oneofs_by_name['port'].fields.append(
  _PORTSELECTOR.fields_by_name['name'])
_PORTSELECTOR.fields_by_name['name'].containing_oneof = _PORTSELECTOR.oneofs_by_name['port']
DESCRIPTOR.message_types_by_name['StringMatch'] = _STRINGMATCH
DESCRIPTOR.message_types_by_name['MutualTls'] = _MUTUALTLS
DESCRIPTOR.message_types_by_name['Jwt'] = _JWT
DESCRIPTOR.message_types_by_name['PeerAuthenticationMethod'] = _PEERAUTHENTICATIONMETHOD
DESCRIPTOR.message_types_by_name['OriginAuthenticationMethod'] = _ORIGINAUTHENTICATIONMETHOD
DESCRIPTOR.message_types_by_name['Policy'] = _POLICY
DESCRIPTOR.message_types_by_name['TargetSelector'] = _TARGETSELECTOR
DESCRIPTOR.message_types_by_name['PortSelector'] = _PORTSELECTOR
DESCRIPTOR.enum_types_by_name['PrincipalBinding'] = _PRINCIPALBINDING
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

StringMatch = _reflection.GeneratedProtocolMessageType('StringMatch', (_message.Message,), {
  'DESCRIPTOR' : _STRINGMATCH,
  '__module__' : 'authentication.v1alpha1.policy_pb2'
  # @@protoc_insertion_point(class_scope:istio.authentication.v1alpha1.StringMatch)
  })
_sym_db.RegisterMessage(StringMatch)

MutualTls = _reflection.GeneratedProtocolMessageType('MutualTls', (_message.Message,), {
  'DESCRIPTOR' : _MUTUALTLS,
  '__module__' : 'authentication.v1alpha1.policy_pb2'
  # @@protoc_insertion_point(class_scope:istio.authentication.v1alpha1.MutualTls)
  })
_sym_db.RegisterMessage(MutualTls)

Jwt = _reflection.GeneratedProtocolMessageType('Jwt', (_message.Message,), {

  'TriggerRule' : _reflection.GeneratedProtocolMessageType('TriggerRule', (_message.Message,), {
    'DESCRIPTOR' : _JWT_TRIGGERRULE,
    '__module__' : 'authentication.v1alpha1.policy_pb2'
    # @@protoc_insertion_point(class_scope:istio.authentication.v1alpha1.Jwt.TriggerRule)
    })
  ,
  'DESCRIPTOR' : _JWT,
  '__module__' : 'authentication.v1alpha1.policy_pb2'
  # @@protoc_insertion_point(class_scope:istio.authentication.v1alpha1.Jwt)
  })
_sym_db.RegisterMessage(Jwt)
_sym_db.RegisterMessage(Jwt.TriggerRule)

PeerAuthenticationMethod = _reflection.GeneratedProtocolMessageType('PeerAuthenticationMethod', (_message.Message,), {
  'DESCRIPTOR' : _PEERAUTHENTICATIONMETHOD,
  '__module__' : 'authentication.v1alpha1.policy_pb2'
  # @@protoc_insertion_point(class_scope:istio.authentication.v1alpha1.PeerAuthenticationMethod)
  })
_sym_db.RegisterMessage(PeerAuthenticationMethod)

OriginAuthenticationMethod = _reflection.GeneratedProtocolMessageType('OriginAuthenticationMethod', (_message.Message,), {
  'DESCRIPTOR' : _ORIGINAUTHENTICATIONMETHOD,
  '__module__' : 'authentication.v1alpha1.policy_pb2'
  # @@protoc_insertion_point(class_scope:istio.authentication.v1alpha1.OriginAuthenticationMethod)
  })
_sym_db.RegisterMessage(OriginAuthenticationMethod)

Policy = _reflection.GeneratedProtocolMessageType('Policy', (_message.Message,), {
  'DESCRIPTOR' : _POLICY,
  '__module__' : 'authentication.v1alpha1.policy_pb2'
  # @@protoc_insertion_point(class_scope:istio.authentication.v1alpha1.Policy)
  })
_sym_db.RegisterMessage(Policy)

TargetSelector = _reflection.GeneratedProtocolMessageType('TargetSelector', (_message.Message,), {
  'DESCRIPTOR' : _TARGETSELECTOR,
  '__module__' : 'authentication.v1alpha1.policy_pb2'
  # @@protoc_insertion_point(class_scope:istio.authentication.v1alpha1.TargetSelector)
  })
_sym_db.RegisterMessage(TargetSelector)

PortSelector = _reflection.GeneratedProtocolMessageType('PortSelector', (_message.Message,), {
  'DESCRIPTOR' : _PORTSELECTOR,
  '__module__' : 'authentication.v1alpha1.policy_pb2'
  # @@protoc_insertion_point(class_scope:istio.authentication.v1alpha1.PortSelector)
  })
_sym_db.RegisterMessage(PortSelector)


DESCRIPTOR._options = None
_MUTUALTLS.fields_by_name['allow_tls']._options = None
_PEERAUTHENTICATIONMETHOD.fields_by_name['jwt']._options = None
_POLICY.fields_by_name['targets']._options = None
_POLICY.fields_by_name['peer_is_optional']._options = None
_POLICY.fields_by_name['origins']._options = None
_POLICY.fields_by_name['origin_is_optional']._options = None
_POLICY.fields_by_name['principal_binding']._options = None
_TARGETSELECTOR.fields_by_name['name']._options = None
# @@protoc_insertion_point(module_scope)
