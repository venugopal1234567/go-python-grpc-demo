# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: score.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0bscore.proto\x12\tlivescore\"%\n\x12ListMatchesRequest\x12\x0f\n\x07\x63ountry\x18\x01 \x01(\t\"1\n\x12MatchScoreResponse\x12\r\n\x05score\x18\x01 \x01(\t\x12\x0c\n\x04live\x18\x02 \x01(\x08\"D\n\x13ListMatchesResponse\x12-\n\x06scores\x18\x01 \x03(\x0b\x32\x1d.livescore.MatchScoreResponse2\\\n\x0cScoreService\x12L\n\x0bListMatches\x12\x1d.livescore.ListMatchesRequest\x1a\x1e.livescore.ListMatchesResponseB\x13Z\x11./proto/livescoreb\x06proto3')



_LISTMATCHESREQUEST = DESCRIPTOR.message_types_by_name['ListMatchesRequest']
_MATCHSCORERESPONSE = DESCRIPTOR.message_types_by_name['MatchScoreResponse']
_LISTMATCHESRESPONSE = DESCRIPTOR.message_types_by_name['ListMatchesResponse']
ListMatchesRequest = _reflection.GeneratedProtocolMessageType('ListMatchesRequest', (_message.Message,), {
  'DESCRIPTOR' : _LISTMATCHESREQUEST,
  '__module__' : 'score_pb2'
  # @@protoc_insertion_point(class_scope:livescore.ListMatchesRequest)
  })
_sym_db.RegisterMessage(ListMatchesRequest)

MatchScoreResponse = _reflection.GeneratedProtocolMessageType('MatchScoreResponse', (_message.Message,), {
  'DESCRIPTOR' : _MATCHSCORERESPONSE,
  '__module__' : 'score_pb2'
  # @@protoc_insertion_point(class_scope:livescore.MatchScoreResponse)
  })
_sym_db.RegisterMessage(MatchScoreResponse)

ListMatchesResponse = _reflection.GeneratedProtocolMessageType('ListMatchesResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTMATCHESRESPONSE,
  '__module__' : 'score_pb2'
  # @@protoc_insertion_point(class_scope:livescore.ListMatchesResponse)
  })
_sym_db.RegisterMessage(ListMatchesResponse)

_SCORESERVICE = DESCRIPTOR.services_by_name['ScoreService']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\021./proto/livescore'
  _LISTMATCHESREQUEST._serialized_start=26
  _LISTMATCHESREQUEST._serialized_end=63
  _MATCHSCORERESPONSE._serialized_start=65
  _MATCHSCORERESPONSE._serialized_end=114
  _LISTMATCHESRESPONSE._serialized_start=116
  _LISTMATCHESRESPONSE._serialized_end=184
  _SCORESERVICE._serialized_start=186
  _SCORESERVICE._serialized_end=278
# @@protoc_insertion_point(module_scope)