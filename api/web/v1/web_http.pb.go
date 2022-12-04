// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.4.0
// - protoc             v3.21.5
// source: web/v1/web.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationWebListUser = "/api.web.v1.Web/ListUser"
const OperationWebLogin = "/api.web.v1.Web/Login"
const OperationWebLogout = "/api.web.v1.Web/Logout"

type WebHTTPServer interface {
	ListUser(context.Context, *emptypb.Empty) (*ListUserReply, error)
	Login(context.Context, *LoginReq) (*User, error)
	Logout(context.Context, *LogoutReq) (*LogoutReply, error)
}

func RegisterWebHTTPServer(s *http.Server, srv WebHTTPServer) {
	r := s.Route("/")
	r.POST("/web/v1/login", _Web_Login0_HTTP_Handler(srv))
	r.POST("/web/v1/logout", _Web_Logout0_HTTP_Handler(srv))
	r.GET("/web/v1/users", _Web_ListUser0_HTTP_Handler(srv))
}

func _Web_Login0_HTTP_Handler(srv WebHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWebLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*User)
		return ctx.Result(200, reply)
	}
}

func _Web_Logout0_HTTP_Handler(srv WebHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LogoutReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWebLogout)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Logout(ctx, req.(*LogoutReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LogoutReply)
		return ctx.Result(200, reply)
	}
}

func _Web_ListUser0_HTTP_Handler(srv WebHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWebListUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUser(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListUserReply)
		return ctx.Result(200, reply)
	}
}

type WebHTTPClient interface {
	ListUser(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *ListUserReply, err error)
	Login(ctx context.Context, req *LoginReq, opts ...http.CallOption) (rsp *User, err error)
	Logout(ctx context.Context, req *LogoutReq, opts ...http.CallOption) (rsp *LogoutReply, err error)
}

type WebHTTPClientImpl struct {
	cc *http.Client
}

func NewWebHTTPClient(client *http.Client) WebHTTPClient {
	return &WebHTTPClientImpl{client}
}

func (c *WebHTTPClientImpl) ListUser(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*ListUserReply, error) {
	var out ListUserReply
	pattern := "/web/v1/users"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationWebListUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WebHTTPClientImpl) Login(ctx context.Context, in *LoginReq, opts ...http.CallOption) (*User, error) {
	var out User
	pattern := "/web/v1/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationWebLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WebHTTPClientImpl) Logout(ctx context.Context, in *LogoutReq, opts ...http.CallOption) (*LogoutReply, error) {
	var out LogoutReply
	pattern := "/web/v1/logout"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationWebLogout))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
