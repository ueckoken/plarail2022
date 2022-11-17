from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

BLOCKSTATE_CLOSE: BlockStateEnum
BLOCKSTATE_OPEN: BlockStateEnum
BLOCKSTATE_UNKNOWN: BlockStateEnum
DESCRIPTOR: _descriptor.FileDescriptor
chofu_b1: BlockId
chofu_b2: BlockId
chofu_b3: BlockId
chofu_b4: BlockId
chofu_b5: BlockId
hachioji_b1: BlockId
hashimoto_b1: BlockId
hashimoto_b2: BlockId
hashioji_b2: BlockId
sakurajosui_b1: BlockId
sakurajosui_b2: BlockId
sakurajosui_b3: BlockId
sakurajosui_b4: BlockId
sakurajosui_b5: BlockId
sakurajosui_b6: BlockId
shinjuku_b1: BlockId
shinjuku_b2: BlockId
unknown: BlockId

class BlockAndState(_message.Message):
    __slots__ = ["blockId", "state"]
    BLOCKID_FIELD_NUMBER: _ClassVar[int]
    STATE_FIELD_NUMBER: _ClassVar[int]
    blockId: BlockId
    state: BlockStateEnum
    def __init__(self, blockId: _Optional[_Union[BlockId, str]] = ..., state: _Optional[_Union[BlockStateEnum, str]] = ...) -> None: ...

class NotifyBlockStateRequest(_message.Message):
    __slots__ = ["state"]
    STATE_FIELD_NUMBER: _ClassVar[int]
    state: BlockAndState
    def __init__(self, state: _Optional[_Union[BlockAndState, _Mapping]] = ...) -> None: ...

class NotifyBlockStateResponse(_message.Message):
    __slots__ = []
    def __init__(self) -> None: ...

class UpdateBlockStateRequest(_message.Message):
    __slots__ = ["state"]
    STATE_FIELD_NUMBER: _ClassVar[int]
    state: BlockAndState
    def __init__(self, state: _Optional[_Union[BlockAndState, _Mapping]] = ...) -> None: ...

class UpdateBlockStateResponse(_message.Message):
    __slots__ = []
    def __init__(self) -> None: ...

class BlockStateEnum(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []

class BlockId(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
