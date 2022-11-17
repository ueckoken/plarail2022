// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: proto/block.proto

package spec

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BlockStateEnum int32

const (
	BlockStateEnum_BLOCKSTATE_UNKNOWN BlockStateEnum = 0
	BlockStateEnum_BLOCKSTATE_OPEN    BlockStateEnum = 1
	BlockStateEnum_BLOCKSTATE_CLOSE   BlockStateEnum = 2
)

// Enum value maps for BlockStateEnum.
var (
	BlockStateEnum_name = map[int32]string{
		0: "BLOCKSTATE_UNKNOWN",
		1: "BLOCKSTATE_OPEN",
		2: "BLOCKSTATE_CLOSE",
	}
	BlockStateEnum_value = map[string]int32{
		"BLOCKSTATE_UNKNOWN": 0,
		"BLOCKSTATE_OPEN":    1,
		"BLOCKSTATE_CLOSE":   2,
	}
)

func (x BlockStateEnum) Enum() *BlockStateEnum {
	p := new(BlockStateEnum)
	*p = x
	return p
}

func (x BlockStateEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BlockStateEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_block_proto_enumTypes[0].Descriptor()
}

func (BlockStateEnum) Type() protoreflect.EnumType {
	return &file_proto_block_proto_enumTypes[0]
}

func (x BlockStateEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BlockStateEnum.Descriptor instead.
func (BlockStateEnum) EnumDescriptor() ([]byte, []int) {
	return file_proto_block_proto_rawDescGZIP(), []int{0}
}

type BlockId int32

const (
	BlockId_unknown        BlockId = 0
	BlockId_shinjuku_b1    BlockId = 1
	BlockId_shinjuku_b2    BlockId = 2
	BlockId_sakurajosui_b1 BlockId = 11
	BlockId_sakurajosui_b2 BlockId = 12
	BlockId_sakurajosui_b3 BlockId = 13
	BlockId_sakurajosui_b4 BlockId = 14
	BlockId_sakurajosui_b5 BlockId = 15
	BlockId_sakurajosui_b6 BlockId = 16
	BlockId_chofu_b1       BlockId = 21
	BlockId_chofu_b2       BlockId = 22
	BlockId_chofu_b3       BlockId = 23
	BlockId_chofu_b4       BlockId = 24
	BlockId_chofu_b5       BlockId = 25
	BlockId_hashimoto_b1   BlockId = 31
	BlockId_hashimoto_b2   BlockId = 32
	BlockId_hachioji_b1    BlockId = 41
	BlockId_hachioji_b2    BlockId = 42
)

// Enum value maps for BlockId.
var (
	BlockId_name = map[int32]string{
		0:  "unknown",
		1:  "shinjuku_b1",
		2:  "shinjuku_b2",
		11: "sakurajosui_b1",
		12: "sakurajosui_b2",
		13: "sakurajosui_b3",
		14: "sakurajosui_b4",
		15: "sakurajosui_b5",
		16: "sakurajosui_b6",
		21: "chofu_b1",
		22: "chofu_b2",
		23: "chofu_b3",
		24: "chofu_b4",
		25: "chofu_b5",
		31: "hashimoto_b1",
		32: "hashimoto_b2",
		41: "hachioji_b1",
		42: "hachioji_b2",
	}
	BlockId_value = map[string]int32{
		"unknown":        0,
		"shinjuku_b1":    1,
		"shinjuku_b2":    2,
		"sakurajosui_b1": 11,
		"sakurajosui_b2": 12,
		"sakurajosui_b3": 13,
		"sakurajosui_b4": 14,
		"sakurajosui_b5": 15,
		"sakurajosui_b6": 16,
		"chofu_b1":       21,
		"chofu_b2":       22,
		"chofu_b3":       23,
		"chofu_b4":       24,
		"chofu_b5":       25,
		"hashimoto_b1":   31,
		"hashimoto_b2":   32,
		"hachioji_b1":    41,
		"hachioji_b2":    42,
	}
)

func (x BlockId) Enum() *BlockId {
	p := new(BlockId)
	*p = x
	return p
}

func (x BlockId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BlockId) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_block_proto_enumTypes[1].Descriptor()
}

func (BlockId) Type() protoreflect.EnumType {
	return &file_proto_block_proto_enumTypes[1]
}

func (x BlockId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BlockId.Descriptor instead.
func (BlockId) EnumDescriptor() ([]byte, []int) {
	return file_proto_block_proto_rawDescGZIP(), []int{1}
}

type BlockAndState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockId BlockId        `protobuf:"varint,3,opt,name=blockId,proto3,enum=BlockId" json:"blockId,omitempty"`
	State   BlockStateEnum `protobuf:"varint,2,opt,name=state,proto3,enum=BlockStateEnum" json:"state,omitempty"`
}

func (x *BlockAndState) Reset() {
	*x = BlockAndState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockAndState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockAndState) ProtoMessage() {}

func (x *BlockAndState) ProtoReflect() protoreflect.Message {
	mi := &file_proto_block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockAndState.ProtoReflect.Descriptor instead.
func (*BlockAndState) Descriptor() ([]byte, []int) {
	return file_proto_block_proto_rawDescGZIP(), []int{0}
}

func (x *BlockAndState) GetBlockId() BlockId {
	if x != nil {
		return x.BlockId
	}
	return BlockId_unknown
}

func (x *BlockAndState) GetState() BlockStateEnum {
	if x != nil {
		return x.State
	}
	return BlockStateEnum_BLOCKSTATE_UNKNOWN
}

type UpdateBlockStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State *BlockAndState `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *UpdateBlockStateRequest) Reset() {
	*x = UpdateBlockStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_block_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBlockStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBlockStateRequest) ProtoMessage() {}

func (x *UpdateBlockStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_block_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBlockStateRequest.ProtoReflect.Descriptor instead.
func (*UpdateBlockStateRequest) Descriptor() ([]byte, []int) {
	return file_proto_block_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateBlockStateRequest) GetState() *BlockAndState {
	if x != nil {
		return x.State
	}
	return nil
}

type UpdateBlockStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateBlockStateResponse) Reset() {
	*x = UpdateBlockStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_block_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBlockStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBlockStateResponse) ProtoMessage() {}

func (x *UpdateBlockStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_block_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBlockStateResponse.ProtoReflect.Descriptor instead.
func (*UpdateBlockStateResponse) Descriptor() ([]byte, []int) {
	return file_proto_block_proto_rawDescGZIP(), []int{2}
}

type NotifyBlockStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State *BlockAndState `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *NotifyBlockStateRequest) Reset() {
	*x = NotifyBlockStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_block_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyBlockStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyBlockStateRequest) ProtoMessage() {}

func (x *NotifyBlockStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_block_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyBlockStateRequest.ProtoReflect.Descriptor instead.
func (*NotifyBlockStateRequest) Descriptor() ([]byte, []int) {
	return file_proto_block_proto_rawDescGZIP(), []int{3}
}

func (x *NotifyBlockStateRequest) GetState() *BlockAndState {
	if x != nil {
		return x.State
	}
	return nil
}

type NotifyBlockStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NotifyBlockStateResponse) Reset() {
	*x = NotifyBlockStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_block_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyBlockStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyBlockStateResponse) ProtoMessage() {}

func (x *NotifyBlockStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_block_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyBlockStateResponse.ProtoReflect.Descriptor instead.
func (*NotifyBlockStateResponse) Descriptor() ([]byte, []int) {
	return file_proto_block_proto_rawDescGZIP(), []int{4}
}

var File_proto_block_proto protoreflect.FileDescriptor

var file_proto_block_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x0d, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x41, 0x6e, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x08, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x52,
	0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22,
	0x3f, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x41, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x22, 0x1a, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3f, 0x0a, 0x17,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x41, 0x6e,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x1a, 0x0a,
	0x18, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x53, 0x0a, 0x0e, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x16, 0x0a, 0x12, 0x42,
	0x4c, 0x4f, 0x43, 0x4b, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x42, 0x4c, 0x4f, 0x43,
	0x4b, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x10, 0x02, 0x2a, 0xbc,
	0x02, 0x0a, 0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x0b, 0x0a, 0x07, 0x75, 0x6e,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x73, 0x68, 0x69, 0x6e, 0x6a,
	0x75, 0x6b, 0x75, 0x5f, 0x62, 0x31, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x73, 0x68, 0x69, 0x6e,
	0x6a, 0x75, 0x6b, 0x75, 0x5f, 0x62, 0x32, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x73, 0x61, 0x6b,
	0x75, 0x72, 0x61, 0x6a, 0x6f, 0x73, 0x75, 0x69, 0x5f, 0x62, 0x31, 0x10, 0x0b, 0x12, 0x12, 0x0a,
	0x0e, 0x73, 0x61, 0x6b, 0x75, 0x72, 0x61, 0x6a, 0x6f, 0x73, 0x75, 0x69, 0x5f, 0x62, 0x32, 0x10,
	0x0c, 0x12, 0x12, 0x0a, 0x0e, 0x73, 0x61, 0x6b, 0x75, 0x72, 0x61, 0x6a, 0x6f, 0x73, 0x75, 0x69,
	0x5f, 0x62, 0x33, 0x10, 0x0d, 0x12, 0x12, 0x0a, 0x0e, 0x73, 0x61, 0x6b, 0x75, 0x72, 0x61, 0x6a,
	0x6f, 0x73, 0x75, 0x69, 0x5f, 0x62, 0x34, 0x10, 0x0e, 0x12, 0x12, 0x0a, 0x0e, 0x73, 0x61, 0x6b,
	0x75, 0x72, 0x61, 0x6a, 0x6f, 0x73, 0x75, 0x69, 0x5f, 0x62, 0x35, 0x10, 0x0f, 0x12, 0x12, 0x0a,
	0x0e, 0x73, 0x61, 0x6b, 0x75, 0x72, 0x61, 0x6a, 0x6f, 0x73, 0x75, 0x69, 0x5f, 0x62, 0x36, 0x10,
	0x10, 0x12, 0x0c, 0x0a, 0x08, 0x63, 0x68, 0x6f, 0x66, 0x75, 0x5f, 0x62, 0x31, 0x10, 0x15, 0x12,
	0x0c, 0x0a, 0x08, 0x63, 0x68, 0x6f, 0x66, 0x75, 0x5f, 0x62, 0x32, 0x10, 0x16, 0x12, 0x0c, 0x0a,
	0x08, 0x63, 0x68, 0x6f, 0x66, 0x75, 0x5f, 0x62, 0x33, 0x10, 0x17, 0x12, 0x0c, 0x0a, 0x08, 0x63,
	0x68, 0x6f, 0x66, 0x75, 0x5f, 0x62, 0x34, 0x10, 0x18, 0x12, 0x0c, 0x0a, 0x08, 0x63, 0x68, 0x6f,
	0x66, 0x75, 0x5f, 0x62, 0x35, 0x10, 0x19, 0x12, 0x10, 0x0a, 0x0c, 0x68, 0x61, 0x73, 0x68, 0x69,
	0x6d, 0x6f, 0x74, 0x6f, 0x5f, 0x62, 0x31, 0x10, 0x1f, 0x12, 0x10, 0x0a, 0x0c, 0x68, 0x61, 0x73,
	0x68, 0x69, 0x6d, 0x6f, 0x74, 0x6f, 0x5f, 0x62, 0x32, 0x10, 0x20, 0x12, 0x0f, 0x0a, 0x0b, 0x68,
	0x61, 0x63, 0x68, 0x69, 0x6f, 0x6a, 0x69, 0x5f, 0x62, 0x31, 0x10, 0x29, 0x12, 0x0f, 0x0a, 0x0b,
	0x68, 0x61, 0x63, 0x68, 0x69, 0x6f, 0x6a, 0x69, 0x5f, 0x62, 0x32, 0x10, 0x2a, 0x32, 0x5e, 0x0a,
	0x11, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x12, 0x49, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x63, 0x0a,
	0x16, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x10, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_block_proto_rawDescOnce sync.Once
	file_proto_block_proto_rawDescData = file_proto_block_proto_rawDesc
)

func file_proto_block_proto_rawDescGZIP() []byte {
	file_proto_block_proto_rawDescOnce.Do(func() {
		file_proto_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_block_proto_rawDescData)
	})
	return file_proto_block_proto_rawDescData
}

var file_proto_block_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_block_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_block_proto_goTypes = []interface{}{
	(BlockStateEnum)(0),              // 0: BlockStateEnum
	(BlockId)(0),                     // 1: BlockId
	(*BlockAndState)(nil),            // 2: BlockAndState
	(*UpdateBlockStateRequest)(nil),  // 3: UpdateBlockStateRequest
	(*UpdateBlockStateResponse)(nil), // 4: UpdateBlockStateResponse
	(*NotifyBlockStateRequest)(nil),  // 5: NotifyBlockStateRequest
	(*NotifyBlockStateResponse)(nil), // 6: NotifyBlockStateResponse
}
var file_proto_block_proto_depIdxs = []int32{
	1, // 0: BlockAndState.blockId:type_name -> BlockId
	0, // 1: BlockAndState.state:type_name -> BlockStateEnum
	2, // 2: UpdateBlockStateRequest.state:type_name -> BlockAndState
	2, // 3: NotifyBlockStateRequest.state:type_name -> BlockAndState
	3, // 4: BlockStateManager.UpdateBlockState:input_type -> UpdateBlockStateRequest
	5, // 5: BlockStateNotification.NotifyBlockState:input_type -> NotifyBlockStateRequest
	4, // 6: BlockStateManager.UpdateBlockState:output_type -> UpdateBlockStateResponse
	6, // 7: BlockStateNotification.NotifyBlockState:output_type -> NotifyBlockStateResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_block_proto_init() }
func file_proto_block_proto_init() {
	if File_proto_block_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_block_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockAndState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_block_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBlockStateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_block_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBlockStateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_block_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyBlockStateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_block_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyBlockStateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_block_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_block_proto_goTypes,
		DependencyIndexes: file_proto_block_proto_depIdxs,
		EnumInfos:         file_proto_block_proto_enumTypes,
		MessageInfos:      file_proto_block_proto_msgTypes,
	}.Build()
	File_proto_block_proto = out.File
	file_proto_block_proto_rawDesc = nil
	file_proto_block_proto_goTypes = nil
	file_proto_block_proto_depIdxs = nil
}
