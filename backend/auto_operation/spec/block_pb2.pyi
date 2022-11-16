from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Blocks(_message.Message):
    __slots__ = ["blockId"]
    class BlockId(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = []
    BLOCKID_FIELD_NUMBER: _ClassVar[int]
    blockId: Blocks.BlockId
    chofu_b1: Blocks.BlockId
    chofu_b2: Blocks.BlockId
    chofu_b3: Blocks.BlockId
    chofu_b4: Blocks.BlockId
    chofu_b5: Blocks.BlockId
    hachioji_b1: Blocks.BlockId
    hashimoto_b1: Blocks.BlockId
    hashimoto_b2: Blocks.BlockId
    hashioji_b2: Blocks.BlockId
    sakurajosui_b1: Blocks.BlockId
    sakurajosui_b2: Blocks.BlockId
    sakurajosui_b3: Blocks.BlockId
    sakurajosui_b4: Blocks.BlockId
    sakurajosui_b5: Blocks.BlockId
    sakurajosui_b6: Blocks.BlockId
    shinjuku_b1: Blocks.BlockId
    shinjuku_b2: Blocks.BlockId
    unknown: Blocks.BlockId
    def __init__(self, blockId: _Optional[_Union[Blocks.BlockId, str]] = ...) -> None: ...

class NotifyBlockStateRequest(_message.Message):
    __slots__ = ["block", "state"]
    class State(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = []
    BLOCK_FIELD_NUMBER: _ClassVar[int]
    CLOSE: NotifyBlockStateRequest.State
    OPEN: NotifyBlockStateRequest.State
    STATE_FIELD_NUMBER: _ClassVar[int]
    UNKNOWN: NotifyBlockStateRequest.State
    block: Blocks
    state: NotifyBlockStateRequest.State
    def __init__(self, state: _Optional[_Union[NotifyBlockStateRequest.State, str]] = ..., block: _Optional[_Union[Blocks, _Mapping]] = ...) -> None: ...

class NotifyBlockStateResponse(_message.Message):
    __slots__ = ["response"]
    class Response(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = []
    FAILED: NotifyBlockStateResponse.Response
    RESPONSE_FIELD_NUMBER: _ClassVar[int]
    SUCCESS: NotifyBlockStateResponse.Response
    UNKNOWN: NotifyBlockStateResponse.Response
    response: NotifyBlockStateResponse.Response
    def __init__(self, response: _Optional[_Union[NotifyBlockStateResponse.Response, str]] = ...) -> None: ...
