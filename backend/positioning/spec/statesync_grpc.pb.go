// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: proto/statesync.proto

package spec

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StateManagerClient is the client API for StateManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StateManagerClient interface {
	// UpdatePointStateはexternalへPointState更新要求を送る。
	UpdatePointState(ctx context.Context, in *UpdatePointStateRequest, opts ...grpc.CallOption) (*UpdatePointStateResponse, error)
}

type stateManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewStateManagerClient(cc grpc.ClientConnInterface) StateManagerClient {
	return &stateManagerClient{cc}
}

func (c *stateManagerClient) UpdatePointState(ctx context.Context, in *UpdatePointStateRequest, opts ...grpc.CallOption) (*UpdatePointStateResponse, error) {
	out := new(UpdatePointStateResponse)
	err := c.cc.Invoke(ctx, "/StateManager/UpdatePointState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StateManagerServer is the server API for StateManager service.
// All implementations must embed UnimplementedStateManagerServer
// for forward compatibility
type StateManagerServer interface {
	// UpdatePointStateはexternalへPointState更新要求を送る。
	UpdatePointState(context.Context, *UpdatePointStateRequest) (*UpdatePointStateResponse, error)
	mustEmbedUnimplementedStateManagerServer()
}

// UnimplementedStateManagerServer must be embedded to have forward compatible implementations.
type UnimplementedStateManagerServer struct {
}

func (UnimplementedStateManagerServer) UpdatePointState(context.Context, *UpdatePointStateRequest) (*UpdatePointStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePointState not implemented")
}
func (UnimplementedStateManagerServer) mustEmbedUnimplementedStateManagerServer() {}

// UnsafeStateManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StateManagerServer will
// result in compilation errors.
type UnsafeStateManagerServer interface {
	mustEmbedUnimplementedStateManagerServer()
}

func RegisterStateManagerServer(s grpc.ServiceRegistrar, srv StateManagerServer) {
	s.RegisterService(&StateManager_ServiceDesc, srv)
}

func _StateManager_UpdatePointState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePointStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StateManagerServer).UpdatePointState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/StateManager/UpdatePointState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StateManagerServer).UpdatePointState(ctx, req.(*UpdatePointStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StateManager_ServiceDesc is the grpc.ServiceDesc for StateManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StateManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "StateManager",
	HandlerType: (*StateManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdatePointState",
			Handler:    _StateManager_UpdatePointState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/statesync.proto",
}

// PointStateNotificationClient is the client API for PointStateNotification service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PointStateNotificationClient interface {
	// NotifyPointStateはexternalからauto-operationやinternalへPointStateの更新情報を伝える。
	NotifyPointState(ctx context.Context, in *NotifyPointStateRequest, opts ...grpc.CallOption) (*NotifyPointStateResponse, error)
}

type pointStateNotificationClient struct {
	cc grpc.ClientConnInterface
}

func NewPointStateNotificationClient(cc grpc.ClientConnInterface) PointStateNotificationClient {
	return &pointStateNotificationClient{cc}
}

func (c *pointStateNotificationClient) NotifyPointState(ctx context.Context, in *NotifyPointStateRequest, opts ...grpc.CallOption) (*NotifyPointStateResponse, error) {
	out := new(NotifyPointStateResponse)
	err := c.cc.Invoke(ctx, "/PointStateNotification/NotifyPointState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PointStateNotificationServer is the server API for PointStateNotification service.
// All implementations must embed UnimplementedPointStateNotificationServer
// for forward compatibility
type PointStateNotificationServer interface {
	// NotifyPointStateはexternalからauto-operationやinternalへPointStateの更新情報を伝える。
	NotifyPointState(context.Context, *NotifyPointStateRequest) (*NotifyPointStateResponse, error)
	mustEmbedUnimplementedPointStateNotificationServer()
}

// UnimplementedPointStateNotificationServer must be embedded to have forward compatible implementations.
type UnimplementedPointStateNotificationServer struct {
}

func (UnimplementedPointStateNotificationServer) NotifyPointState(context.Context, *NotifyPointStateRequest) (*NotifyPointStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyPointState not implemented")
}
func (UnimplementedPointStateNotificationServer) mustEmbedUnimplementedPointStateNotificationServer() {
}

// UnsafePointStateNotificationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PointStateNotificationServer will
// result in compilation errors.
type UnsafePointStateNotificationServer interface {
	mustEmbedUnimplementedPointStateNotificationServer()
}

func RegisterPointStateNotificationServer(s grpc.ServiceRegistrar, srv PointStateNotificationServer) {
	s.RegisterService(&PointStateNotification_ServiceDesc, srv)
}

func _PointStateNotification_NotifyPointState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyPointStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointStateNotificationServer).NotifyPointState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PointStateNotification/NotifyPointState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointStateNotificationServer).NotifyPointState(ctx, req.(*NotifyPointStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PointStateNotification_ServiceDesc is the grpc.ServiceDesc for PointStateNotification service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PointStateNotification_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PointStateNotification",
	HandlerType: (*PointStateNotificationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NotifyPointState",
			Handler:    _PointStateNotification_NotifyPointState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/statesync.proto",
}
