from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Command2InternalRequest(_message.Message):
    __slots__ = ["state", "station"]
    class State(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = []
    OFF: Command2InternalRequest.State
    ON: Command2InternalRequest.State
    STATE_FIELD_NUMBER: _ClassVar[int]
    STATION_FIELD_NUMBER: _ClassVar[int]
    UNKNOWN: Command2InternalRequest.State
    state: Command2InternalRequest.State
    station: Stations
    def __init__(self, station: _Optional[_Union[Stations, _Mapping]] = ..., state: _Optional[_Union[Command2InternalRequest.State, str]] = ...) -> None: ...

class Command2InternalResponse(_message.Message):
    __slots__ = ["response"]
    class Response(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = []
    FAILED: Command2InternalResponse.Response
    RESPONSE_FIELD_NUMBER: _ClassVar[int]
    SUCCESS: Command2InternalResponse.Response
    UNKNOWN: Command2InternalResponse.Response
    response: Command2InternalResponse.Response
    def __init__(self, response: _Optional[_Union[Command2InternalResponse.Response, str]] = ...) -> None: ...

class RequestSync(_message.Message):
    __slots__ = ["state", "station"]
    class State(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = []
    OFF: RequestSync.State
    ON: RequestSync.State
    STATE_FIELD_NUMBER: _ClassVar[int]
    STATION_FIELD_NUMBER: _ClassVar[int]
    UNKNOWN: RequestSync.State
    state: RequestSync.State
    station: Stations
    def __init__(self, station: _Optional[_Union[Stations, _Mapping]] = ..., state: _Optional[_Union[RequestSync.State, str]] = ...) -> None: ...

class ResponseSync(_message.Message):
    __slots__ = ["response"]
    class Response(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = []
    FAILED: ResponseSync.Response
    RESPONSE_FIELD_NUMBER: _ClassVar[int]
    SUCCESS: ResponseSync.Response
    UNKNOWN: ResponseSync.Response
    response: ResponseSync.Response
    def __init__(self, response: _Optional[_Union[ResponseSync.Response, str]] = ...) -> None: ...

class Stations(_message.Message):
    __slots__ = ["stationId"]
    class StationId(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = []
    STATIONID_FIELD_NUMBER: _ClassVar[int]
    chofu_p1: Stations.StationId
    chofu_s0: Stations.StationId
    chofu_s1: Stations.StationId
    chofu_s2: Stations.StationId
    chofu_s3: Stations.StationId
    chofu_s4: Stations.StationId
    hachioji_s1: Stations.StationId
    hachioji_s2: Stations.StationId
    hashimoto_s1: Stations.StationId
    hashimoto_s2: Stations.StationId
    sakurajosui_p1: Stations.StationId
    sakurajosui_p2: Stations.StationId
    sakurajosui_s1: Stations.StationId
    sakurajosui_s2: Stations.StationId
    sakurajosui_s3: Stations.StationId
    sakurajosui_s4: Stations.StationId
    sakurajosui_s5: Stations.StationId
    sakurajosui_s6: Stations.StationId
    shinjuku_s1: Stations.StationId
    shinjuku_s2: Stations.StationId
    stationId: Stations.StationId
    unknown: Stations.StationId
    def __init__(self, stationId: _Optional[_Union[Stations.StationId, str]] = ...) -> None: ...
