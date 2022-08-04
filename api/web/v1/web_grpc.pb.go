// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.0
// source: web/v1/web.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WebClient is the client API for Web service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WebClient interface {
	// 登陆
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*User, error)
	// 登出
	Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*LogoutReply, error)
	// 用户列表
	ListUser(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListUserReply, error)
}

type webClient struct {
	cc grpc.ClientConnInterface
}

func NewWebClient(cc grpc.ClientConnInterface) WebClient {
	return &webClient{cc}
}

func (c *webClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/api.web.v1.Web/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webClient) Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*LogoutReply, error) {
	out := new(LogoutReply)
	err := c.cc.Invoke(ctx, "/api.web.v1.Web/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webClient) ListUser(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListUserReply, error) {
	out := new(ListUserReply)
	err := c.cc.Invoke(ctx, "/api.web.v1.Web/ListUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebServer is the server API for Web service.
// All implementations must embed UnimplementedWebServer
// for forward compatibility
type WebServer interface {
	// 登陆
	Login(context.Context, *LoginReq) (*User, error)
	// 登出
	Logout(context.Context, *LogoutReq) (*LogoutReply, error)
	// 用户列表
	ListUser(context.Context, *emptypb.Empty) (*ListUserReply, error)
	mustEmbedUnimplementedWebServer()
}

// UnimplementedWebServer must be embedded to have forward compatible implementations.
type UnimplementedWebServer struct {
}

func (UnimplementedWebServer) Login(context.Context, *LoginReq) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedWebServer) Logout(context.Context, *LogoutReq) (*LogoutReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedWebServer) ListUser(context.Context, *emptypb.Empty) (*ListUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUser not implemented")
}
func (UnimplementedWebServer) mustEmbedUnimplementedWebServer() {}

// UnsafeWebServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WebServer will
// result in compilation errors.
type UnsafeWebServer interface {
	mustEmbedUnimplementedWebServer()
}

func RegisterWebServer(s grpc.ServiceRegistrar, srv WebServer) {
	s.RegisterService(&Web_ServiceDesc, srv)
}

func _Web_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.web.v1.Web/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Web_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.web.v1.Web/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebServer).Logout(ctx, req.(*LogoutReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Web_ListUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebServer).ListUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.web.v1.Web/ListUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebServer).ListUser(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Web_ServiceDesc is the grpc.ServiceDesc for Web service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Web_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.web.v1.Web",
	HandlerType: (*WebServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Web_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Web_Logout_Handler,
		},
		{
			MethodName: "ListUser",
			Handler:    _Web_ListUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "web/v1/web.proto",
}
