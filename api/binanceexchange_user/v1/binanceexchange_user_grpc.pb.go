// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.7
// source: api/binanceexchange_user/v1/binanceexchange_user.proto

package v1

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

const (
	BinanceUser_SetUser_FullMethodName        = "/BinanceUser/SetUser"
	BinanceUser_PullUserStatus_FullMethodName = "/BinanceUser/PullUserStatus"
)

// BinanceUserClient is the client API for BinanceUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BinanceUserClient interface {
	SetUser(ctx context.Context, in *SetUserRequest, opts ...grpc.CallOption) (*SetUserReply, error)
	PullUserStatus(ctx context.Context, in *PullUserStatusRequest, opts ...grpc.CallOption) (*PullUserStatusReply, error)
}

type binanceUserClient struct {
	cc grpc.ClientConnInterface
}

func NewBinanceUserClient(cc grpc.ClientConnInterface) BinanceUserClient {
	return &binanceUserClient{cc}
}

func (c *binanceUserClient) SetUser(ctx context.Context, in *SetUserRequest, opts ...grpc.CallOption) (*SetUserReply, error) {
	out := new(SetUserReply)
	err := c.cc.Invoke(ctx, BinanceUser_SetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *binanceUserClient) PullUserStatus(ctx context.Context, in *PullUserStatusRequest, opts ...grpc.CallOption) (*PullUserStatusReply, error) {
	out := new(PullUserStatusReply)
	err := c.cc.Invoke(ctx, BinanceUser_PullUserStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BinanceUserServer is the server API for BinanceUser service.
// All implementations must embed UnimplementedBinanceUserServer
// for forward compatibility
type BinanceUserServer interface {
	SetUser(context.Context, *SetUserRequest) (*SetUserReply, error)
	PullUserStatus(context.Context, *PullUserStatusRequest) (*PullUserStatusReply, error)
	mustEmbedUnimplementedBinanceUserServer()
}

// UnimplementedBinanceUserServer must be embedded to have forward compatible implementations.
type UnimplementedBinanceUserServer struct {
}

func (UnimplementedBinanceUserServer) SetUser(context.Context, *SetUserRequest) (*SetUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUser not implemented")
}
func (UnimplementedBinanceUserServer) PullUserStatus(context.Context, *PullUserStatusRequest) (*PullUserStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PullUserStatus not implemented")
}
func (UnimplementedBinanceUserServer) mustEmbedUnimplementedBinanceUserServer() {}

// UnsafeBinanceUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BinanceUserServer will
// result in compilation errors.
type UnsafeBinanceUserServer interface {
	mustEmbedUnimplementedBinanceUserServer()
}

func RegisterBinanceUserServer(s grpc.ServiceRegistrar, srv BinanceUserServer) {
	s.RegisterService(&BinanceUser_ServiceDesc, srv)
}

func _BinanceUser_SetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BinanceUserServer).SetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BinanceUser_SetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BinanceUserServer).SetUser(ctx, req.(*SetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BinanceUser_PullUserStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullUserStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BinanceUserServer).PullUserStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BinanceUser_PullUserStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BinanceUserServer).PullUserStatus(ctx, req.(*PullUserStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BinanceUser_ServiceDesc is the grpc.ServiceDesc for BinanceUser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BinanceUser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BinanceUser",
	HandlerType: (*BinanceUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetUser",
			Handler:    _BinanceUser_SetUser_Handler,
		},
		{
			MethodName: "PullUserStatus",
			Handler:    _BinanceUser_PullUserStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/binanceexchange_user/v1/binanceexchange_user.proto",
}
