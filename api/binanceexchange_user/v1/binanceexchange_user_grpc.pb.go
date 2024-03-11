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
	BinanceUser_GetUser_FullMethodName                  = "/BinanceUser/GetUser"
	BinanceUser_PullUserDeposit_FullMethodName          = "/BinanceUser/PullUserDeposit"
	BinanceUser_PullUserDeposit2_FullMethodName         = "/BinanceUser/PullUserDeposit2"
	BinanceUser_PullUserCredentialsBsc_FullMethodName   = "/BinanceUser/PullUserCredentialsBsc"
	BinanceUser_BindTrader_FullMethodName               = "/BinanceUser/BindTrader"
	BinanceUser_ListenTraderAndUserOrder_FullMethodName = "/BinanceUser/ListenTraderAndUserOrder"
	BinanceUser_OrderHandle_FullMethodName              = "/BinanceUser/OrderHandle"
	BinanceUser_OrderHandleTwo_FullMethodName           = "/BinanceUser/OrderHandleTwo"
	BinanceUser_Analyze_FullMethodName                  = "/BinanceUser/Analyze"
	BinanceUser_CloseOrderAfterBind_FullMethodName      = "/BinanceUser/CloseOrderAfterBind"
	BinanceUser_CloseOrderAfterBindTwo_FullMethodName   = "/BinanceUser/CloseOrderAfterBindTwo"
	BinanceUser_InitOrderAfterBind_FullMethodName       = "/BinanceUser/InitOrderAfterBind"
	BinanceUser_InitOrderAfterBindTwo_FullMethodName    = "/BinanceUser/InitOrderAfterBindTwo"
	BinanceUser_OverOrderAfterBind_FullMethodName       = "/BinanceUser/OverOrderAfterBind"
	BinanceUser_OverOrderAfterBindTwo_FullMethodName    = "/BinanceUser/OverOrderAfterBindTwo"
)

// BinanceUserClient is the client API for BinanceUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BinanceUserClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserReply, error)
	PullUserDeposit(ctx context.Context, in *PullUserDepositRequest, opts ...grpc.CallOption) (*PullUserDepositReply, error)
	PullUserDeposit2(ctx context.Context, in *PullUserDepositRequest, opts ...grpc.CallOption) (*PullUserDepositReply, error)
	PullUserCredentialsBsc(ctx context.Context, in *PullUserCredentialsBscRequest, opts ...grpc.CallOption) (*PullUserCredentialsBscReply, error)
	BindTrader(ctx context.Context, in *BindTraderRequest, opts ...grpc.CallOption) (*BindTraderReply, error)
	ListenTraderAndUserOrder(ctx context.Context, in *ListenTraderAndUserOrderRequest, opts ...grpc.CallOption) (*ListenTraderAndUserOrderReply, error)
	OrderHandle(ctx context.Context, in *OrderHandleRequest, opts ...grpc.CallOption) (*OrderHandleReply, error)
	OrderHandleTwo(ctx context.Context, in *OrderHandleRequest, opts ...grpc.CallOption) (*OrderHandleReply, error)
	Analyze(ctx context.Context, in *AnalyzeRequest, opts ...grpc.CallOption) (*AnalyzeReply, error)
	CloseOrderAfterBind(ctx context.Context, in *CloseOrderAfterBindRequest, opts ...grpc.CallOption) (*CloseOrderAfterBindReply, error)
	CloseOrderAfterBindTwo(ctx context.Context, in *CloseOrderAfterBindRequest, opts ...grpc.CallOption) (*CloseOrderAfterBindReply, error)
	InitOrderAfterBind(ctx context.Context, in *InitOrderAfterBindRequest, opts ...grpc.CallOption) (*InitOrderAfterBindReply, error)
	InitOrderAfterBindTwo(ctx context.Context, in *InitOrderAfterBindRequest, opts ...grpc.CallOption) (*InitOrderAfterBindReply, error)
	OverOrderAfterBind(ctx context.Context, in *OverOrderAfterBindRequest, opts ...grpc.CallOption) (*OverOrderAfterBindReply, error)
	OverOrderAfterBindTwo(ctx context.Context, in *OverOrderAfterBindRequest, opts ...grpc.CallOption) (*OverOrderAfterBindReply, error)
}

type binanceUserClient struct {
	cc grpc.ClientConnInterface
}

func NewBinanceUserClient(cc grpc.ClientConnInterface) BinanceUserClient {
	return &binanceUserClient{cc}
}

func (c *binanceUserClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserReply, error) {
	out := new(GetUserReply)
	err := c.cc.Invoke(ctx, BinanceUser_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *binanceUserClient) PullUserDeposit(ctx context.Context, in *PullUserDepositRequest, opts ...grpc.CallOption) (*PullUserDepositReply, error) {
	out := new(PullUserDepositReply)
	err := c.cc.Invoke(ctx, BinanceUser_PullUserDeposit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *binanceUserClient) PullUserDeposit2(ctx context.Context, in *PullUserDepositRequest, opts ...grpc.CallOption) (*PullUserDepositReply, error) {
	out := new(PullUserDepositReply)
	err := c.cc.Invoke(ctx, BinanceUser_PullUserDeposit2_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *binanceUserClient) PullUserCredentialsBsc(ctx context.Context, in *PullUserCredentialsBscRequest, opts ...grpc.CallOption) (*PullUserCredentialsBscReply, error) {
	out := new(PullUserCredentialsBscReply)
	err := c.cc.Invoke(ctx, BinanceUser_PullUserCredentialsBsc_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *binanceUserClient) BindTrader(ctx context.Context, in *BindTraderRequest, opts ...grpc.CallOption) (*BindTraderReply, error) {
	out := new(BindTraderReply)
	err := c.cc.Invoke(ctx, BinanceUser_BindTrader_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *binanceUserClient) ListenTraderAndUserOrder(ctx context.Context, in *ListenTraderAndUserOrderRequest, opts ...grpc.CallOption) (*ListenTraderAndUserOrderReply, error) {
	out := new(ListenTraderAndUserOrderReply)
	err := c.cc.Invoke(ctx, BinanceUser_ListenTraderAndUserOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *binanceUserClient) OrderHandle(ctx context.Context, in *OrderHandleRequest, opts ...grpc.CallOption) (*OrderHandleReply, error) {
	out := new(OrderHandleReply)
	err := c.cc.Invoke(ctx, BinanceUser_OrderHandle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *binanceUserClient) OrderHandleTwo(ctx context.Context, in *OrderHandleRequest, opts ...grpc.CallOption) (*OrderHandleReply, error) {
	out := new(OrderHandleReply)
	err := c.cc.Invoke(ctx, BinanceUser_OrderHandleTwo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *binanceUserClient) Analyze(ctx context.Context, in *AnalyzeRequest, opts ...grpc.CallOption) (*AnalyzeReply, error) {
	out := new(AnalyzeReply)
	err := c.cc.Invoke(ctx, BinanceUser_Analyze_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *binanceUserClient) CloseOrderAfterBind(ctx context.Context, in *CloseOrderAfterBindRequest, opts ...grpc.CallOption) (*CloseOrderAfterBindReply, error) {
	out := new(CloseOrderAfterBindReply)
	err := c.cc.Invoke(ctx, BinanceUser_CloseOrderAfterBind_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *binanceUserClient) CloseOrderAfterBindTwo(ctx context.Context, in *CloseOrderAfterBindRequest, opts ...grpc.CallOption) (*CloseOrderAfterBindReply, error) {
	out := new(CloseOrderAfterBindReply)
	err := c.cc.Invoke(ctx, BinanceUser_CloseOrderAfterBindTwo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *binanceUserClient) InitOrderAfterBind(ctx context.Context, in *InitOrderAfterBindRequest, opts ...grpc.CallOption) (*InitOrderAfterBindReply, error) {
	out := new(InitOrderAfterBindReply)
	err := c.cc.Invoke(ctx, BinanceUser_InitOrderAfterBind_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *binanceUserClient) InitOrderAfterBindTwo(ctx context.Context, in *InitOrderAfterBindRequest, opts ...grpc.CallOption) (*InitOrderAfterBindReply, error) {
	out := new(InitOrderAfterBindReply)
	err := c.cc.Invoke(ctx, BinanceUser_InitOrderAfterBindTwo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *binanceUserClient) OverOrderAfterBind(ctx context.Context, in *OverOrderAfterBindRequest, opts ...grpc.CallOption) (*OverOrderAfterBindReply, error) {
	out := new(OverOrderAfterBindReply)
	err := c.cc.Invoke(ctx, BinanceUser_OverOrderAfterBind_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *binanceUserClient) OverOrderAfterBindTwo(ctx context.Context, in *OverOrderAfterBindRequest, opts ...grpc.CallOption) (*OverOrderAfterBindReply, error) {
	out := new(OverOrderAfterBindReply)
	err := c.cc.Invoke(ctx, BinanceUser_OverOrderAfterBindTwo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BinanceUserServer is the server API for BinanceUser service.
// All implementations must embed UnimplementedBinanceUserServer
// for forward compatibility
type BinanceUserServer interface {
	GetUser(context.Context, *GetUserRequest) (*GetUserReply, error)
	PullUserDeposit(context.Context, *PullUserDepositRequest) (*PullUserDepositReply, error)
	PullUserDeposit2(context.Context, *PullUserDepositRequest) (*PullUserDepositReply, error)
	PullUserCredentialsBsc(context.Context, *PullUserCredentialsBscRequest) (*PullUserCredentialsBscReply, error)
	BindTrader(context.Context, *BindTraderRequest) (*BindTraderReply, error)
	ListenTraderAndUserOrder(context.Context, *ListenTraderAndUserOrderRequest) (*ListenTraderAndUserOrderReply, error)
	OrderHandle(context.Context, *OrderHandleRequest) (*OrderHandleReply, error)
	OrderHandleTwo(context.Context, *OrderHandleRequest) (*OrderHandleReply, error)
	Analyze(context.Context, *AnalyzeRequest) (*AnalyzeReply, error)
	CloseOrderAfterBind(context.Context, *CloseOrderAfterBindRequest) (*CloseOrderAfterBindReply, error)
	CloseOrderAfterBindTwo(context.Context, *CloseOrderAfterBindRequest) (*CloseOrderAfterBindReply, error)
	InitOrderAfterBind(context.Context, *InitOrderAfterBindRequest) (*InitOrderAfterBindReply, error)
	InitOrderAfterBindTwo(context.Context, *InitOrderAfterBindRequest) (*InitOrderAfterBindReply, error)
	OverOrderAfterBind(context.Context, *OverOrderAfterBindRequest) (*OverOrderAfterBindReply, error)
	OverOrderAfterBindTwo(context.Context, *OverOrderAfterBindRequest) (*OverOrderAfterBindReply, error)
	mustEmbedUnimplementedBinanceUserServer()
}

// UnimplementedBinanceUserServer must be embedded to have forward compatible implementations.
type UnimplementedBinanceUserServer struct {
}

func (UnimplementedBinanceUserServer) GetUser(context.Context, *GetUserRequest) (*GetUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedBinanceUserServer) PullUserDeposit(context.Context, *PullUserDepositRequest) (*PullUserDepositReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PullUserDeposit not implemented")
}
func (UnimplementedBinanceUserServer) PullUserDeposit2(context.Context, *PullUserDepositRequest) (*PullUserDepositReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PullUserDeposit2 not implemented")
}
func (UnimplementedBinanceUserServer) PullUserCredentialsBsc(context.Context, *PullUserCredentialsBscRequest) (*PullUserCredentialsBscReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PullUserCredentialsBsc not implemented")
}
func (UnimplementedBinanceUserServer) BindTrader(context.Context, *BindTraderRequest) (*BindTraderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BindTrader not implemented")
}
func (UnimplementedBinanceUserServer) ListenTraderAndUserOrder(context.Context, *ListenTraderAndUserOrderRequest) (*ListenTraderAndUserOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListenTraderAndUserOrder not implemented")
}
func (UnimplementedBinanceUserServer) OrderHandle(context.Context, *OrderHandleRequest) (*OrderHandleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderHandle not implemented")
}
func (UnimplementedBinanceUserServer) OrderHandleTwo(context.Context, *OrderHandleRequest) (*OrderHandleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderHandleTwo not implemented")
}
func (UnimplementedBinanceUserServer) Analyze(context.Context, *AnalyzeRequest) (*AnalyzeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Analyze not implemented")
}
func (UnimplementedBinanceUserServer) CloseOrderAfterBind(context.Context, *CloseOrderAfterBindRequest) (*CloseOrderAfterBindReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseOrderAfterBind not implemented")
}
func (UnimplementedBinanceUserServer) CloseOrderAfterBindTwo(context.Context, *CloseOrderAfterBindRequest) (*CloseOrderAfterBindReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseOrderAfterBindTwo not implemented")
}
func (UnimplementedBinanceUserServer) InitOrderAfterBind(context.Context, *InitOrderAfterBindRequest) (*InitOrderAfterBindReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitOrderAfterBind not implemented")
}
func (UnimplementedBinanceUserServer) InitOrderAfterBindTwo(context.Context, *InitOrderAfterBindRequest) (*InitOrderAfterBindReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitOrderAfterBindTwo not implemented")
}
func (UnimplementedBinanceUserServer) OverOrderAfterBind(context.Context, *OverOrderAfterBindRequest) (*OverOrderAfterBindReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OverOrderAfterBind not implemented")
}
func (UnimplementedBinanceUserServer) OverOrderAfterBindTwo(context.Context, *OverOrderAfterBindRequest) (*OverOrderAfterBindReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OverOrderAfterBindTwo not implemented")
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

func _BinanceUser_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BinanceUserServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BinanceUser_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BinanceUserServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BinanceUser_PullUserDeposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullUserDepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BinanceUserServer).PullUserDeposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BinanceUser_PullUserDeposit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BinanceUserServer).PullUserDeposit(ctx, req.(*PullUserDepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BinanceUser_PullUserDeposit2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullUserDepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BinanceUserServer).PullUserDeposit2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BinanceUser_PullUserDeposit2_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BinanceUserServer).PullUserDeposit2(ctx, req.(*PullUserDepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BinanceUser_PullUserCredentialsBsc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullUserCredentialsBscRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BinanceUserServer).PullUserCredentialsBsc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BinanceUser_PullUserCredentialsBsc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BinanceUserServer).PullUserCredentialsBsc(ctx, req.(*PullUserCredentialsBscRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BinanceUser_BindTrader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindTraderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BinanceUserServer).BindTrader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BinanceUser_BindTrader_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BinanceUserServer).BindTrader(ctx, req.(*BindTraderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BinanceUser_ListenTraderAndUserOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListenTraderAndUserOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BinanceUserServer).ListenTraderAndUserOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BinanceUser_ListenTraderAndUserOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BinanceUserServer).ListenTraderAndUserOrder(ctx, req.(*ListenTraderAndUserOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BinanceUser_OrderHandle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderHandleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BinanceUserServer).OrderHandle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BinanceUser_OrderHandle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BinanceUserServer).OrderHandle(ctx, req.(*OrderHandleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BinanceUser_OrderHandleTwo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderHandleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BinanceUserServer).OrderHandleTwo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BinanceUser_OrderHandleTwo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BinanceUserServer).OrderHandleTwo(ctx, req.(*OrderHandleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BinanceUser_Analyze_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnalyzeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BinanceUserServer).Analyze(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BinanceUser_Analyze_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BinanceUserServer).Analyze(ctx, req.(*AnalyzeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BinanceUser_CloseOrderAfterBind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseOrderAfterBindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BinanceUserServer).CloseOrderAfterBind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BinanceUser_CloseOrderAfterBind_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BinanceUserServer).CloseOrderAfterBind(ctx, req.(*CloseOrderAfterBindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BinanceUser_CloseOrderAfterBindTwo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseOrderAfterBindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BinanceUserServer).CloseOrderAfterBindTwo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BinanceUser_CloseOrderAfterBindTwo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BinanceUserServer).CloseOrderAfterBindTwo(ctx, req.(*CloseOrderAfterBindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BinanceUser_InitOrderAfterBind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitOrderAfterBindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BinanceUserServer).InitOrderAfterBind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BinanceUser_InitOrderAfterBind_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BinanceUserServer).InitOrderAfterBind(ctx, req.(*InitOrderAfterBindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BinanceUser_InitOrderAfterBindTwo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitOrderAfterBindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BinanceUserServer).InitOrderAfterBindTwo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BinanceUser_InitOrderAfterBindTwo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BinanceUserServer).InitOrderAfterBindTwo(ctx, req.(*InitOrderAfterBindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BinanceUser_OverOrderAfterBind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OverOrderAfterBindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BinanceUserServer).OverOrderAfterBind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BinanceUser_OverOrderAfterBind_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BinanceUserServer).OverOrderAfterBind(ctx, req.(*OverOrderAfterBindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BinanceUser_OverOrderAfterBindTwo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OverOrderAfterBindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BinanceUserServer).OverOrderAfterBindTwo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BinanceUser_OverOrderAfterBindTwo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BinanceUserServer).OverOrderAfterBindTwo(ctx, req.(*OverOrderAfterBindRequest))
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
			MethodName: "GetUser",
			Handler:    _BinanceUser_GetUser_Handler,
		},
		{
			MethodName: "PullUserDeposit",
			Handler:    _BinanceUser_PullUserDeposit_Handler,
		},
		{
			MethodName: "PullUserDeposit2",
			Handler:    _BinanceUser_PullUserDeposit2_Handler,
		},
		{
			MethodName: "PullUserCredentialsBsc",
			Handler:    _BinanceUser_PullUserCredentialsBsc_Handler,
		},
		{
			MethodName: "BindTrader",
			Handler:    _BinanceUser_BindTrader_Handler,
		},
		{
			MethodName: "ListenTraderAndUserOrder",
			Handler:    _BinanceUser_ListenTraderAndUserOrder_Handler,
		},
		{
			MethodName: "OrderHandle",
			Handler:    _BinanceUser_OrderHandle_Handler,
		},
		{
			MethodName: "OrderHandleTwo",
			Handler:    _BinanceUser_OrderHandleTwo_Handler,
		},
		{
			MethodName: "Analyze",
			Handler:    _BinanceUser_Analyze_Handler,
		},
		{
			MethodName: "CloseOrderAfterBind",
			Handler:    _BinanceUser_CloseOrderAfterBind_Handler,
		},
		{
			MethodName: "CloseOrderAfterBindTwo",
			Handler:    _BinanceUser_CloseOrderAfterBindTwo_Handler,
		},
		{
			MethodName: "InitOrderAfterBind",
			Handler:    _BinanceUser_InitOrderAfterBind_Handler,
		},
		{
			MethodName: "InitOrderAfterBindTwo",
			Handler:    _BinanceUser_InitOrderAfterBindTwo_Handler,
		},
		{
			MethodName: "OverOrderAfterBind",
			Handler:    _BinanceUser_OverOrderAfterBind_Handler,
		},
		{
			MethodName: "OverOrderAfterBindTwo",
			Handler:    _BinanceUser_OverOrderAfterBindTwo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/binanceexchange_user/v1/binanceexchange_user.proto",
}
