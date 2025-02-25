from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor
POINTSTATE_OFF: PointStateEnum
POINTSTATE_ON: PointStateEnum
POINTSTATE_UNKNOWN: PointStateEnum
chofu_p1: StationId
chofu_s0: StationId
chofu_s1: StationId
chofu_s2: StationId
chofu_s3: StationId
chofu_s4: StationId
hachioji_s1: StationId
hachioji_s2: StationId
hashimoto_s1: StationId
hashimoto_s2: StationId
sakurajosui_p1: StationId
sakurajosui_p2: StationId
sakurajosui_s0: StationId
sakurajosui_s1: StationId
sakurajosui_s2: StationId
sakurajosui_s3: StationId
sakurajosui_s4: StationId
sakurajosui_s5: StationId
shinjuku_s1: StationId
shinjuku_s2: StationId
stationid_unknown: StationId

class NotifyPointStateRequest(_message.Message):
    __slots__ = ["state"]
    STATE_FIELD_NUMBER: _ClassVar[int]
    state: PointAndState
    def __init__(self, state: _Optional[_Union[PointAndState, _Mapping]] = ...) -> None: ...

class NotifyPointStateResponse(_message.Message):
    __slots__ = []
    def __init__(self) -> None: ...

class PointAndState(_message.Message):
    __slots__ = ["state", "station"]
    STATE_FIELD_NUMBER: _ClassVar[int]
    STATION_FIELD_NUMBER: _ClassVar[int]
    state: PointStateEnum
    station: Station
    def __init__(self, station: _Optional[_Union[Station, _Mapping]] = ..., state: _Optional[_Union[PointStateEnum, str]] = ...) -> None: ...

class Station(_message.Message):
    __slots__ = ["stationId"]
    STATIONID_FIELD_NUMBER: _ClassVar[int]
    stationId: StationId
    def __init__(self, stationId: _Optional[_Union[StationId, str]] = ...) -> None: ...

class UpdatePointStateRequest(_message.Message):
    __slots__ = ["state"]
    STATE_FIELD_NUMBER: _ClassVar[int]
    state: PointAndState
    def __init__(self, state: _Optional[_Union[PointAndState, _Mapping]] = ...) -> None: ...

class UpdatePointStateResponse(_message.Message):
    __slots__ = []
    def __init__(self) -> None: ...

class StationId(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []

class PointStateEnum(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
