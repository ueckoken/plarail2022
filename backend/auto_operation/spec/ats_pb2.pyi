from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SendStatusRequest(_message.Message):
    __slots__ = ["sensor"]

    class SensorName(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = []
    SENSOR_FIELD_NUMBER: _ClassVar[int]
    chofu_d1: SendStatusRequest.SensorName
    chofu_d2: SendStatusRequest.SensorName
    chofu_d3: SendStatusRequest.SensorName
    chofu_d4: SendStatusRequest.SensorName
    chofu_d5: SendStatusRequest.SensorName
    chofu_d6: SendStatusRequest.SensorName
    hachioji_d1: SendStatusRequest.SensorName
    hachioji_d2: SendStatusRequest.SensorName
    hashimoto_d1: SendStatusRequest.SensorName
    hashimoto_d2: SendStatusRequest.SensorName
    sakurajosui_d1: SendStatusRequest.SensorName
    sakurajosui_d2: SendStatusRequest.SensorName
    sakurajosui_d3: SendStatusRequest.SensorName
    sakurajosui_d4: SendStatusRequest.SensorName
    sensor: SendStatusRequest.SensorName
    shinjuku_d1: SendStatusRequest.SensorName
    shinjuku_d2: SendStatusRequest.SensorName
    unknown: SendStatusRequest.SensorName
    wakabadai_d1: SendStatusRequest.SensorName
    wakabadai_d2: SendStatusRequest.SensorName
    def __init__(
        self, sensor: _Optional[_Union[SendStatusRequest.SensorName, str]] = ...
    ) -> None: ...

class SendStatusResponse(_message.Message):
    __slots__ = ["response"]

    class Response(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = []
    FAILED: SendStatusResponse.Response
    RESPONSE_FIELD_NUMBER: _ClassVar[int]
    SUCCESS: SendStatusResponse.Response
    UNKNOWN: SendStatusResponse.Response
    response: SendStatusResponse.Response
    def __init__(
        self, response: _Optional[_Union[SendStatusResponse.Response, str]] = ...
    ) -> None: ...
