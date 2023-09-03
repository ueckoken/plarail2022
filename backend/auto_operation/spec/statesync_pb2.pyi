from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class StationId(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    stationid_unknown: _ClassVar[StationId]
    shinjuku_s1: _ClassVar[StationId]
    shinjuku_s2: _ClassVar[StationId]
    sakurajosui_p1: _ClassVar[StationId]
    sakurajosui_p2: _ClassVar[StationId]
    sakurajosui_s0: _ClassVar[StationId]
    sakurajosui_s1: _ClassVar[StationId]
    sakurajosui_s2: _ClassVar[StationId]
    sakurajosui_s3: _ClassVar[StationId]
    sakurajosui_s4: _ClassVar[StationId]
    sakurajosui_s5: _ClassVar[StationId]
    chofu_p1: _ClassVar[StationId]
    chofu_s0: _ClassVar[StationId]
    chofu_s1: _ClassVar[StationId]
    chofu_s2: _ClassVar[StationId]
    chofu_s3: _ClassVar[StationId]
    chofu_s4: _ClassVar[StationId]
    hashimoto_s1: _ClassVar[StationId]
    hashimoto_s2: _ClassVar[StationId]
    hachioji_s1: _ClassVar[StationId]
    hachioji_s2: _ClassVar[StationId]

class PointStateEnum(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    POINTSTATE_UNKNOWN: _ClassVar[PointStateEnum]
    POINTSTATE_ON: _ClassVar[PointStateEnum]
    POINTSTATE_OFF: _ClassVar[PointStateEnum]
stationid_unknown: StationId
shinjuku_s1: StationId
shinjuku_s2: StationId
sakurajosui_p1: StationId
sakurajosui_p2: StationId
sakurajosui_s0: StationId
sakurajosui_s1: StationId
sakurajosui_s2: StationId
sakurajosui_s3: StationId
sakurajosui_s4: StationId
sakurajosui_s5: StationId
chofu_p1: StationId
chofu_s0: StationId
chofu_s1: StationId
chofu_s2: StationId
chofu_s3: StationId
chofu_s4: StationId
hashimoto_s1: StationId
hashimoto_s2: StationId
hachioji_s1: StationId
hachioji_s2: StationId
POINTSTATE_UNKNOWN: PointStateEnum
POINTSTATE_ON: PointStateEnum
POINTSTATE_OFF: PointStateEnum

class Station(_message.Message):
    __slots__ = ["stationId"]
    STATIONID_FIELD_NUMBER: _ClassVar[int]
    stationId: StationId
    def __init__(self, stationId: _Optional[_Union[StationId, str]] = ...) -> None: ...

class PointAndState(_message.Message):
    __slots__ = ["station", "state"]
    STATION_FIELD_NUMBER: _ClassVar[int]
    STATE_FIELD_NUMBER: _ClassVar[int]
    station: Station
    state: PointStateEnum
    def __init__(self, station: _Optional[_Union[Station, _Mapping]] = ..., state: _Optional[_Union[PointStateEnum, str]] = ...) -> None: ...

class UpdatePointStateRequest(_message.Message):
    __slots__ = ["state"]
    STATE_FIELD_NUMBER: _ClassVar[int]
    state: PointAndState
    def __init__(self, state: _Optional[_Union[PointAndState, _Mapping]] = ...) -> None: ...

class UpdatePointStateResponse(_message.Message):
    __slots__ = []
    def __init__(self) -> None: ...

class NotifyPointStateRequest(_message.Message):
    __slots__ = ["state"]
    STATE_FIELD_NUMBER: _ClassVar[int]
    state: PointAndState
    def __init__(self, state: _Optional[_Union[PointAndState, _Mapping]] = ...) -> None: ...

class NotifyPointStateResponse(_message.Message):
    __slots__ = []
    def __init__(self) -> None: ...
