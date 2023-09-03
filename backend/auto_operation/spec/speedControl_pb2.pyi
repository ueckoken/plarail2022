from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SendSpeed(_message.Message):
    __slots__ = ["Speed", "train"]
    class Train(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = []
        UNKNOWN: _ClassVar[SendSpeed.Train]
        TAKAO: _ClassVar[SendSpeed.Train]
        CHICHIBU: _ClassVar[SendSpeed.Train]
        HAKONE: _ClassVar[SendSpeed.Train]
        OKUTAMA: _ClassVar[SendSpeed.Train]
        NIKKO: _ClassVar[SendSpeed.Train]
        ENOSHIMA: _ClassVar[SendSpeed.Train]
        KAMAKURA: _ClassVar[SendSpeed.Train]
        YOKOSUKA: _ClassVar[SendSpeed.Train]
    UNKNOWN: SendSpeed.Train
    TAKAO: SendSpeed.Train
    CHICHIBU: SendSpeed.Train
    HAKONE: SendSpeed.Train
    OKUTAMA: SendSpeed.Train
    NIKKO: SendSpeed.Train
    ENOSHIMA: SendSpeed.Train
    KAMAKURA: SendSpeed.Train
    YOKOSUKA: SendSpeed.Train
    SPEED_FIELD_NUMBER: _ClassVar[int]
    TRAIN_FIELD_NUMBER: _ClassVar[int]
    Speed: int
    train: SendSpeed.Train
    def __init__(self, Speed: _Optional[int] = ..., train: _Optional[_Union[SendSpeed.Train, str]] = ...) -> None: ...

class StatusCode(_message.Message):
    __slots__ = ["code"]
    class Code(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = []
        UNKNOWN: _ClassVar[StatusCode.Code]
        SUCCESS: _ClassVar[StatusCode.Code]
        FAILED: _ClassVar[StatusCode.Code]
    UNKNOWN: StatusCode.Code
    SUCCESS: StatusCode.Code
    FAILED: StatusCode.Code
    CODE_FIELD_NUMBER: _ClassVar[int]
    code: StatusCode.Code
    def __init__(self, code: _Optional[_Union[StatusCode.Code, str]] = ...) -> None: ...
