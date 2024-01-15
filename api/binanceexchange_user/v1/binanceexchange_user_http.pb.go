// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             v3.21.7
// source: api/binanceexchange_user/v1/binanceexchange_user.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationBinanceUserGetUser = "/BinanceUser/GetUser"
const OperationBinanceUserPullUserCredentialsBsc = "/BinanceUser/PullUserCredentialsBsc"
const OperationBinanceUserPullUserStatus = "/BinanceUser/PullUserStatus"
const OperationBinanceUserSetUser = "/BinanceUser/SetUser"

type BinanceUserHTTPServer interface {
	GetUser(context.Context, *GetUserRequest) (*GetUserReply, error)
	PullUserCredentialsBsc(context.Context, *PullUserCredentialsBscRequest) (*PullUserCredentialsBscReply, error)
	PullUserStatus(context.Context, *PullUserStatusRequest) (*PullUserStatusReply, error)
	SetUser(context.Context, *SetUserRequest) (*SetUserReply, error)
}

func RegisterBinanceUserHTTPServer(s *http.Server, srv BinanceUserHTTPServer) {
	r := s.Route("/")
	r.POST("/api/binanceexchange_user/set_user", _BinanceUser_SetUser0_HTTP_Handler(srv))
	r.GET("/api/binanceexchange_user/get_user", _BinanceUser_GetUser0_HTTP_Handler(srv))
	r.GET("/api/binanceexchange_user/pull_user_status", _BinanceUser_PullUserStatus0_HTTP_Handler(srv))
	r.GET("/api/binanceexchange_user/pull_user_credentials_bsc", _BinanceUser_PullUserCredentialsBsc0_HTTP_Handler(srv))
}

func _BinanceUser_SetUser0_HTTP_Handler(srv BinanceUserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SetUserRequest
		if err := ctx.Bind(&in.SendBody); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBinanceUserSetUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SetUser(ctx, req.(*SetUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SetUserReply)
		return ctx.Result(200, reply)
	}
}

func _BinanceUser_GetUser0_HTTP_Handler(srv BinanceUserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBinanceUserGetUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUser(ctx, req.(*GetUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserReply)
		return ctx.Result(200, reply)
	}
}

func _BinanceUser_PullUserStatus0_HTTP_Handler(srv BinanceUserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PullUserStatusRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBinanceUserPullUserStatus)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PullUserStatus(ctx, req.(*PullUserStatusRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PullUserStatusReply)
		return ctx.Result(200, reply)
	}
}

func _BinanceUser_PullUserCredentialsBsc0_HTTP_Handler(srv BinanceUserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PullUserCredentialsBscRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBinanceUserPullUserCredentialsBsc)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PullUserCredentialsBsc(ctx, req.(*PullUserCredentialsBscRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PullUserCredentialsBscReply)
		return ctx.Result(200, reply)
	}
}

type BinanceUserHTTPClient interface {
	GetUser(ctx context.Context, req *GetUserRequest, opts ...http.CallOption) (rsp *GetUserReply, err error)
	PullUserCredentialsBsc(ctx context.Context, req *PullUserCredentialsBscRequest, opts ...http.CallOption) (rsp *PullUserCredentialsBscReply, err error)
	PullUserStatus(ctx context.Context, req *PullUserStatusRequest, opts ...http.CallOption) (rsp *PullUserStatusReply, err error)
	SetUser(ctx context.Context, req *SetUserRequest, opts ...http.CallOption) (rsp *SetUserReply, err error)
}

type BinanceUserHTTPClientImpl struct {
	cc *http.Client
}

func NewBinanceUserHTTPClient(client *http.Client) BinanceUserHTTPClient {
	return &BinanceUserHTTPClientImpl{client}
}

func (c *BinanceUserHTTPClientImpl) GetUser(ctx context.Context, in *GetUserRequest, opts ...http.CallOption) (*GetUserReply, error) {
	var out GetUserReply
	pattern := "/api/binanceexchange_user/get_user"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBinanceUserGetUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BinanceUserHTTPClientImpl) PullUserCredentialsBsc(ctx context.Context, in *PullUserCredentialsBscRequest, opts ...http.CallOption) (*PullUserCredentialsBscReply, error) {
	var out PullUserCredentialsBscReply
	pattern := "/api/binanceexchange_user/pull_user_credentials_bsc"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBinanceUserPullUserCredentialsBsc))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BinanceUserHTTPClientImpl) PullUserStatus(ctx context.Context, in *PullUserStatusRequest, opts ...http.CallOption) (*PullUserStatusReply, error) {
	var out PullUserStatusReply
	pattern := "/api/binanceexchange_user/pull_user_status"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBinanceUserPullUserStatus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BinanceUserHTTPClientImpl) SetUser(ctx context.Context, in *SetUserRequest, opts ...http.CallOption) (*SetUserReply, error) {
	var out SetUserReply
	pattern := "/api/binanceexchange_user/set_user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBinanceUserSetUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.SendBody, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
