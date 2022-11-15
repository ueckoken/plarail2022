from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SendSpeed(_message.Message):
    __slots__ = ["Speed", "train"]

    class Train(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = []
    CHICHIBU: SendSpeed.Train
    ENOSHIMA: SendSpeed.Train
    HAKONE: SendSpeed.Train
    KAMAKURA: SendSpeed.Train
    NIKKO: SendSpeed.Train
    OKUTAMA: SendSpeed.Train
    SPEED_FIELD_NUMBER: _ClassVar[int]
    Speed: int
    TAKAO: SendSpeed.Train
    TRAIN_FIELD_NUMBER: _ClassVar[int]
    UNKNOWN: SendSpeed.Train
    YOKOSUKA: SendSpeed.Train
    train: SendSpeed.Train
    def __init__(
        self,
        Speed: _Optional[int] = ...,
        train: _Optional[_Union[SendSpeed.Train, str]] = ...,
    ) -> None: ...

class StatusCode(_message.Message):
    __slots__ = ["code"]

    class Code(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = []
    CODE_FIELD_NUMBER: _ClassVar[int]
    FAILED: StatusCode.Code
    SUCCESS: StatusCode.Code
    UNKNOWN: StatusCode.Code
    code: StatusCode.Code
    def __init__(self, code: _Optional[_Union[StatusCode.Code, str]] = ...) -> None: ...
