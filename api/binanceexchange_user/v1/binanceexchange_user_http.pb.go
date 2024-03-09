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

const OperationBinanceUserAnalyze = "/BinanceUser/Analyze"
const OperationBinanceUserBindTrader = "/BinanceUser/BindTrader"
const OperationBinanceUserCloseOrderAfterBind = "/BinanceUser/CloseOrderAfterBind"
const OperationBinanceUserCloseOrderAfterBindTwo = "/BinanceUser/CloseOrderAfterBindTwo"
const OperationBinanceUserGetUser = "/BinanceUser/GetUser"
const OperationBinanceUserInitOrderAfterBind = "/BinanceUser/InitOrderAfterBind"
const OperationBinanceUserInitOrderAfterBindTwo = "/BinanceUser/InitOrderAfterBindTwo"
const OperationBinanceUserListenTraderAndUserOrder = "/BinanceUser/ListenTraderAndUserOrder"
const OperationBinanceUserOrderHandle = "/BinanceUser/OrderHandle"
const OperationBinanceUserOrderHandleTwo = "/BinanceUser/OrderHandleTwo"
const OperationBinanceUserPullUserCredentialsBsc = "/BinanceUser/PullUserCredentialsBsc"
const OperationBinanceUserPullUserDeposit = "/BinanceUser/PullUserDeposit"
const OperationBinanceUserPullUserDeposit2 = "/BinanceUser/PullUserDeposit2"

type BinanceUserHTTPServer interface {
	Analyze(context.Context, *AnalyzeRequest) (*AnalyzeReply, error)
	BindTrader(context.Context, *BindTraderRequest) (*BindTraderReply, error)
	CloseOrderAfterBind(context.Context, *CloseOrderAfterBindRequest) (*CloseOrderAfterBindReply, error)
	CloseOrderAfterBindTwo(context.Context, *CloseOrderAfterBindRequest) (*CloseOrderAfterBindReply, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserReply, error)
	InitOrderAfterBind(context.Context, *InitOrderAfterBindRequest) (*InitOrderAfterBindReply, error)
	InitOrderAfterBindTwo(context.Context, *InitOrderAfterBindRequest) (*InitOrderAfterBindReply, error)
	ListenTraderAndUserOrder(context.Context, *ListenTraderAndUserOrderRequest) (*ListenTraderAndUserOrderReply, error)
	OrderHandle(context.Context, *OrderHandleRequest) (*OrderHandleReply, error)
	OrderHandleTwo(context.Context, *OrderHandleRequest) (*OrderHandleReply, error)
	PullUserCredentialsBsc(context.Context, *PullUserCredentialsBscRequest) (*PullUserCredentialsBscReply, error)
	PullUserDeposit(context.Context, *PullUserDepositRequest) (*PullUserDepositReply, error)
	PullUserDeposit2(context.Context, *PullUserDepositRequest) (*PullUserDepositReply, error)
}

func RegisterBinanceUserHTTPServer(s *http.Server, srv BinanceUserHTTPServer) {
	r := s.Route("/")
	r.GET("/api/binanceexchange_user/get_user", _BinanceUser_GetUser0_HTTP_Handler(srv))
	r.GET("/api/binanceexchange_user/pull_user_deposit", _BinanceUser_PullUserDeposit0_HTTP_Handler(srv))
	r.GET("/api/binanceexchange_user/pull_user_deposit_2", _BinanceUser_PullUserDeposit20_HTTP_Handler(srv))
	r.GET("/api/binanceexchange_user/pull_user_credentials_bsc", _BinanceUser_PullUserCredentialsBsc0_HTTP_Handler(srv))
	r.GET("/api/binanceexchange_user/bind_trader", _BinanceUser_BindTrader0_HTTP_Handler(srv))
	r.POST("/api/binanceexchange_user/listen_trader_and_user_order", _BinanceUser_ListenTraderAndUserOrder0_HTTP_Handler(srv))
	r.GET("/api/binanceexchange_user/order_handle", _BinanceUser_OrderHandle0_HTTP_Handler(srv))
	r.GET("/api/binanceexchange_user/order_handle_two", _BinanceUser_OrderHandleTwo0_HTTP_Handler(srv))
	r.GET("/api/binanceexchange_user/analyze", _BinanceUser_Analyze0_HTTP_Handler(srv))
	r.GET("/api/binanceexchange_user/close_order_after_bind", _BinanceUser_CloseOrderAfterBind0_HTTP_Handler(srv))
	r.GET("/api/binanceexchange_user/close_order_after_bind_tfi", _BinanceUser_CloseOrderAfterBindTwo0_HTTP_Handler(srv))
	r.GET("/api/binanceexchange_user/init_order_after_bind", _BinanceUser_InitOrderAfterBind0_HTTP_Handler(srv))
	r.GET("/api/binanceexchange_user/init_order_after_bind_tfi", _BinanceUser_InitOrderAfterBindTwo0_HTTP_Handler(srv))
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

func _BinanceUser_PullUserDeposit0_HTTP_Handler(srv BinanceUserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PullUserDepositRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBinanceUserPullUserDeposit)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PullUserDeposit(ctx, req.(*PullUserDepositRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PullUserDepositReply)
		return ctx.Result(200, reply)
	}
}

func _BinanceUser_PullUserDeposit20_HTTP_Handler(srv BinanceUserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PullUserDepositRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBinanceUserPullUserDeposit2)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PullUserDeposit2(ctx, req.(*PullUserDepositRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PullUserDepositReply)
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

func _BinanceUser_BindTrader0_HTTP_Handler(srv BinanceUserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BindTraderRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBinanceUserBindTrader)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BindTrader(ctx, req.(*BindTraderRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BindTraderReply)
		return ctx.Result(200, reply)
	}
}

func _BinanceUser_ListenTraderAndUserOrder0_HTTP_Handler(srv BinanceUserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListenTraderAndUserOrderRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBinanceUserListenTraderAndUserOrder)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListenTraderAndUserOrder(ctx, req.(*ListenTraderAndUserOrderRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListenTraderAndUserOrderReply)
		return ctx.Result(200, reply)
	}
}

func _BinanceUser_OrderHandle0_HTTP_Handler(srv BinanceUserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in OrderHandleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBinanceUserOrderHandle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.OrderHandle(ctx, req.(*OrderHandleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*OrderHandleReply)
		return ctx.Result(200, reply)
	}
}

func _BinanceUser_OrderHandleTwo0_HTTP_Handler(srv BinanceUserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in OrderHandleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBinanceUserOrderHandleTwo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.OrderHandleTwo(ctx, req.(*OrderHandleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*OrderHandleReply)
		return ctx.Result(200, reply)
	}
}

func _BinanceUser_Analyze0_HTTP_Handler(srv BinanceUserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AnalyzeRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBinanceUserAnalyze)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Analyze(ctx, req.(*AnalyzeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AnalyzeReply)
		return ctx.Result(200, reply)
	}
}

func _BinanceUser_CloseOrderAfterBind0_HTTP_Handler(srv BinanceUserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CloseOrderAfterBindRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBinanceUserCloseOrderAfterBind)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CloseOrderAfterBind(ctx, req.(*CloseOrderAfterBindRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CloseOrderAfterBindReply)
		return ctx.Result(200, reply)
	}
}

func _BinanceUser_CloseOrderAfterBindTwo0_HTTP_Handler(srv BinanceUserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CloseOrderAfterBindRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBinanceUserCloseOrderAfterBindTwo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CloseOrderAfterBindTwo(ctx, req.(*CloseOrderAfterBindRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CloseOrderAfterBindReply)
		return ctx.Result(200, reply)
	}
}

func _BinanceUser_InitOrderAfterBind0_HTTP_Handler(srv BinanceUserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in InitOrderAfterBindRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBinanceUserInitOrderAfterBind)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.InitOrderAfterBind(ctx, req.(*InitOrderAfterBindRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*InitOrderAfterBindReply)
		return ctx.Result(200, reply)
	}
}

func _BinanceUser_InitOrderAfterBindTwo0_HTTP_Handler(srv BinanceUserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in InitOrderAfterBindRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBinanceUserInitOrderAfterBindTwo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.InitOrderAfterBindTwo(ctx, req.(*InitOrderAfterBindRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*InitOrderAfterBindReply)
		return ctx.Result(200, reply)
	}
}

type BinanceUserHTTPClient interface {
	Analyze(ctx context.Context, req *AnalyzeRequest, opts ...http.CallOption) (rsp *AnalyzeReply, err error)
	BindTrader(ctx context.Context, req *BindTraderRequest, opts ...http.CallOption) (rsp *BindTraderReply, err error)
	CloseOrderAfterBind(ctx context.Context, req *CloseOrderAfterBindRequest, opts ...http.CallOption) (rsp *CloseOrderAfterBindReply, err error)
	CloseOrderAfterBindTwo(ctx context.Context, req *CloseOrderAfterBindRequest, opts ...http.CallOption) (rsp *CloseOrderAfterBindReply, err error)
	GetUser(ctx context.Context, req *GetUserRequest, opts ...http.CallOption) (rsp *GetUserReply, err error)
	InitOrderAfterBind(ctx context.Context, req *InitOrderAfterBindRequest, opts ...http.CallOption) (rsp *InitOrderAfterBindReply, err error)
	InitOrderAfterBindTwo(ctx context.Context, req *InitOrderAfterBindRequest, opts ...http.CallOption) (rsp *InitOrderAfterBindReply, err error)
	ListenTraderAndUserOrder(ctx context.Context, req *ListenTraderAndUserOrderRequest, opts ...http.CallOption) (rsp *ListenTraderAndUserOrderReply, err error)
	OrderHandle(ctx context.Context, req *OrderHandleRequest, opts ...http.CallOption) (rsp *OrderHandleReply, err error)
	OrderHandleTwo(ctx context.Context, req *OrderHandleRequest, opts ...http.CallOption) (rsp *OrderHandleReply, err error)
	PullUserCredentialsBsc(ctx context.Context, req *PullUserCredentialsBscRequest, opts ...http.CallOption) (rsp *PullUserCredentialsBscReply, err error)
	PullUserDeposit(ctx context.Context, req *PullUserDepositRequest, opts ...http.CallOption) (rsp *PullUserDepositReply, err error)
	PullUserDeposit2(ctx context.Context, req *PullUserDepositRequest, opts ...http.CallOption) (rsp *PullUserDepositReply, err error)
}

type BinanceUserHTTPClientImpl struct {
	cc *http.Client
}

func NewBinanceUserHTTPClient(client *http.Client) BinanceUserHTTPClient {
	return &BinanceUserHTTPClientImpl{client}
}

func (c *BinanceUserHTTPClientImpl) Analyze(ctx context.Context, in *AnalyzeRequest, opts ...http.CallOption) (*AnalyzeReply, error) {
	var out AnalyzeReply
	pattern := "/api/binanceexchange_user/analyze"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBinanceUserAnalyze))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BinanceUserHTTPClientImpl) BindTrader(ctx context.Context, in *BindTraderRequest, opts ...http.CallOption) (*BindTraderReply, error) {
	var out BindTraderReply
	pattern := "/api/binanceexchange_user/bind_trader"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBinanceUserBindTrader))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BinanceUserHTTPClientImpl) CloseOrderAfterBind(ctx context.Context, in *CloseOrderAfterBindRequest, opts ...http.CallOption) (*CloseOrderAfterBindReply, error) {
	var out CloseOrderAfterBindReply
	pattern := "/api/binanceexchange_user/close_order_after_bind"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBinanceUserCloseOrderAfterBind))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BinanceUserHTTPClientImpl) CloseOrderAfterBindTwo(ctx context.Context, in *CloseOrderAfterBindRequest, opts ...http.CallOption) (*CloseOrderAfterBindReply, error) {
	var out CloseOrderAfterBindReply
	pattern := "/api/binanceexchange_user/close_order_after_bind_tfi"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBinanceUserCloseOrderAfterBindTwo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
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

func (c *BinanceUserHTTPClientImpl) InitOrderAfterBind(ctx context.Context, in *InitOrderAfterBindRequest, opts ...http.CallOption) (*InitOrderAfterBindReply, error) {
	var out InitOrderAfterBindReply
	pattern := "/api/binanceexchange_user/init_order_after_bind"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBinanceUserInitOrderAfterBind))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BinanceUserHTTPClientImpl) InitOrderAfterBindTwo(ctx context.Context, in *InitOrderAfterBindRequest, opts ...http.CallOption) (*InitOrderAfterBindReply, error) {
	var out InitOrderAfterBindReply
	pattern := "/api/binanceexchange_user/init_order_after_bind_tfi"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBinanceUserInitOrderAfterBindTwo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BinanceUserHTTPClientImpl) ListenTraderAndUserOrder(ctx context.Context, in *ListenTraderAndUserOrderRequest, opts ...http.CallOption) (*ListenTraderAndUserOrderReply, error) {
	var out ListenTraderAndUserOrderReply
	pattern := "/api/binanceexchange_user/listen_trader_and_user_order"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBinanceUserListenTraderAndUserOrder))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BinanceUserHTTPClientImpl) OrderHandle(ctx context.Context, in *OrderHandleRequest, opts ...http.CallOption) (*OrderHandleReply, error) {
	var out OrderHandleReply
	pattern := "/api/binanceexchange_user/order_handle"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBinanceUserOrderHandle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BinanceUserHTTPClientImpl) OrderHandleTwo(ctx context.Context, in *OrderHandleRequest, opts ...http.CallOption) (*OrderHandleReply, error) {
	var out OrderHandleReply
	pattern := "/api/binanceexchange_user/order_handle_two"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBinanceUserOrderHandleTwo))
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

func (c *BinanceUserHTTPClientImpl) PullUserDeposit(ctx context.Context, in *PullUserDepositRequest, opts ...http.CallOption) (*PullUserDepositReply, error) {
	var out PullUserDepositReply
	pattern := "/api/binanceexchange_user/pull_user_deposit"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBinanceUserPullUserDeposit))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BinanceUserHTTPClientImpl) PullUserDeposit2(ctx context.Context, in *PullUserDepositRequest, opts ...http.CallOption) (*PullUserDepositReply, error) {
	var out PullUserDepositReply
	pattern := "/api/binanceexchange_user/pull_user_deposit_2"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBinanceUserPullUserDeposit2))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
