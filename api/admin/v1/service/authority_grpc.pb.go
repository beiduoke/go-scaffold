// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: admin/v1/service/authority.proto

package service

import (
	context "context"
	protobuf "github.com/beiduoke/go-scaffold/api/protobuf"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuthorityServiceClient is the client API for AuthorityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthorityServiceClient interface {
	// 列表角色
	ListRole(ctx context.Context, in *protobuf.PagingReq, opts ...grpc.CallOption) (*protobuf.PagingReply, error)
	// 创建角色
	CreateRole(ctx context.Context, in *CreateRoleReq, opts ...grpc.CallOption) (*CreateRoleReply, error)
	// 获取角色
	GetRole(ctx context.Context, in *GetRoleReq, opts ...grpc.CallOption) (*Role, error)
	// 修改角色
	UpdateRole(ctx context.Context, in *UpdateRoleReq, opts ...grpc.CallOption) (*UpdateRoleReply, error)
	// 删除角色
	DeleteRole(ctx context.Context, in *DeleteRoleReq, opts ...grpc.CallOption) (*DeleteRoleReply, error)
	// 权限模块
	// 列表权限
	ListAuthority(ctx context.Context, in *protobuf.PagingReq, opts ...grpc.CallOption) (*protobuf.PagingReply, error)
	// 创建权限
	CreateAuthority(ctx context.Context, in *CreateAuthorityReq, opts ...grpc.CallOption) (*CreateAuthorityReply, error)
	// 获取权限
	GetAuthority(ctx context.Context, in *GetAuthorityReq, opts ...grpc.CallOption) (*Authority, error)
	// 修改权限
	UpdateAuthority(ctx context.Context, in *UpdateAuthorityReq, opts ...grpc.CallOption) (*UpdateAuthorityReply, error)
	// 删除权限
	DeleteAuthority(ctx context.Context, in *DeleteAuthorityReq, opts ...grpc.CallOption) (*DeleteAuthorityReply, error)
}

type authorityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorityServiceClient(cc grpc.ClientConnInterface) AuthorityServiceClient {
	return &authorityServiceClient{cc}
}

func (c *authorityServiceClient) ListRole(ctx context.Context, in *protobuf.PagingReq, opts ...grpc.CallOption) (*protobuf.PagingReply, error) {
	out := new(protobuf.PagingReply)
	err := c.cc.Invoke(ctx, "/api.admin.v1.service.AuthorityService/ListRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorityServiceClient) CreateRole(ctx context.Context, in *CreateRoleReq, opts ...grpc.CallOption) (*CreateRoleReply, error) {
	out := new(CreateRoleReply)
	err := c.cc.Invoke(ctx, "/api.admin.v1.service.AuthorityService/CreateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorityServiceClient) GetRole(ctx context.Context, in *GetRoleReq, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := c.cc.Invoke(ctx, "/api.admin.v1.service.AuthorityService/GetRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorityServiceClient) UpdateRole(ctx context.Context, in *UpdateRoleReq, opts ...grpc.CallOption) (*UpdateRoleReply, error) {
	out := new(UpdateRoleReply)
	err := c.cc.Invoke(ctx, "/api.admin.v1.service.AuthorityService/UpdateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorityServiceClient) DeleteRole(ctx context.Context, in *DeleteRoleReq, opts ...grpc.CallOption) (*DeleteRoleReply, error) {
	out := new(DeleteRoleReply)
	err := c.cc.Invoke(ctx, "/api.admin.v1.service.AuthorityService/DeleteRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorityServiceClient) ListAuthority(ctx context.Context, in *protobuf.PagingReq, opts ...grpc.CallOption) (*protobuf.PagingReply, error) {
	out := new(protobuf.PagingReply)
	err := c.cc.Invoke(ctx, "/api.admin.v1.service.AuthorityService/ListAuthority", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorityServiceClient) CreateAuthority(ctx context.Context, in *CreateAuthorityReq, opts ...grpc.CallOption) (*CreateAuthorityReply, error) {
	out := new(CreateAuthorityReply)
	err := c.cc.Invoke(ctx, "/api.admin.v1.service.AuthorityService/CreateAuthority", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorityServiceClient) GetAuthority(ctx context.Context, in *GetAuthorityReq, opts ...grpc.CallOption) (*Authority, error) {
	out := new(Authority)
	err := c.cc.Invoke(ctx, "/api.admin.v1.service.AuthorityService/GetAuthority", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorityServiceClient) UpdateAuthority(ctx context.Context, in *UpdateAuthorityReq, opts ...grpc.CallOption) (*UpdateAuthorityReply, error) {
	out := new(UpdateAuthorityReply)
	err := c.cc.Invoke(ctx, "/api.admin.v1.service.AuthorityService/UpdateAuthority", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorityServiceClient) DeleteAuthority(ctx context.Context, in *DeleteAuthorityReq, opts ...grpc.CallOption) (*DeleteAuthorityReply, error) {
	out := new(DeleteAuthorityReply)
	err := c.cc.Invoke(ctx, "/api.admin.v1.service.AuthorityService/DeleteAuthority", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorityServiceServer is the server API for AuthorityService service.
// All implementations must embed UnimplementedAuthorityServiceServer
// for forward compatibility
type AuthorityServiceServer interface {
	// 列表角色
	ListRole(context.Context, *protobuf.PagingReq) (*protobuf.PagingReply, error)
	// 创建角色
	CreateRole(context.Context, *CreateRoleReq) (*CreateRoleReply, error)
	// 获取角色
	GetRole(context.Context, *GetRoleReq) (*Role, error)
	// 修改角色
	UpdateRole(context.Context, *UpdateRoleReq) (*UpdateRoleReply, error)
	// 删除角色
	DeleteRole(context.Context, *DeleteRoleReq) (*DeleteRoleReply, error)
	// 权限模块
	// 列表权限
	ListAuthority(context.Context, *protobuf.PagingReq) (*protobuf.PagingReply, error)
	// 创建权限
	CreateAuthority(context.Context, *CreateAuthorityReq) (*CreateAuthorityReply, error)
	// 获取权限
	GetAuthority(context.Context, *GetAuthorityReq) (*Authority, error)
	// 修改权限
	UpdateAuthority(context.Context, *UpdateAuthorityReq) (*UpdateAuthorityReply, error)
	// 删除权限
	DeleteAuthority(context.Context, *DeleteAuthorityReq) (*DeleteAuthorityReply, error)
	mustEmbedUnimplementedAuthorityServiceServer()
}

// UnimplementedAuthorityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthorityServiceServer struct {
}

func (UnimplementedAuthorityServiceServer) ListRole(context.Context, *protobuf.PagingReq) (*protobuf.PagingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRole not implemented")
}
func (UnimplementedAuthorityServiceServer) CreateRole(context.Context, *CreateRoleReq) (*CreateRoleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}
func (UnimplementedAuthorityServiceServer) GetRole(context.Context, *GetRoleReq) (*Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRole not implemented")
}
func (UnimplementedAuthorityServiceServer) UpdateRole(context.Context, *UpdateRoleReq) (*UpdateRoleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRole not implemented")
}
func (UnimplementedAuthorityServiceServer) DeleteRole(context.Context, *DeleteRoleReq) (*DeleteRoleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRole not implemented")
}
func (UnimplementedAuthorityServiceServer) ListAuthority(context.Context, *protobuf.PagingReq) (*protobuf.PagingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAuthority not implemented")
}
func (UnimplementedAuthorityServiceServer) CreateAuthority(context.Context, *CreateAuthorityReq) (*CreateAuthorityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAuthority not implemented")
}
func (UnimplementedAuthorityServiceServer) GetAuthority(context.Context, *GetAuthorityReq) (*Authority, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthority not implemented")
}
func (UnimplementedAuthorityServiceServer) UpdateAuthority(context.Context, *UpdateAuthorityReq) (*UpdateAuthorityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAuthority not implemented")
}
func (UnimplementedAuthorityServiceServer) DeleteAuthority(context.Context, *DeleteAuthorityReq) (*DeleteAuthorityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAuthority not implemented")
}
func (UnimplementedAuthorityServiceServer) mustEmbedUnimplementedAuthorityServiceServer() {}

// UnsafeAuthorityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthorityServiceServer will
// result in compilation errors.
type UnsafeAuthorityServiceServer interface {
	mustEmbedUnimplementedAuthorityServiceServer()
}

func RegisterAuthorityServiceServer(s grpc.ServiceRegistrar, srv AuthorityServiceServer) {
	s.RegisterService(&AuthorityService_ServiceDesc, srv)
}

func _AuthorityService_ListRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protobuf.PagingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityServiceServer).ListRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.admin.v1.service.AuthorityService/ListRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityServiceServer).ListRole(ctx, req.(*protobuf.PagingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorityService_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityServiceServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.admin.v1.service.AuthorityService/CreateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityServiceServer).CreateRole(ctx, req.(*CreateRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorityService_GetRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityServiceServer).GetRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.admin.v1.service.AuthorityService/GetRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityServiceServer).GetRole(ctx, req.(*GetRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorityService_UpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityServiceServer).UpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.admin.v1.service.AuthorityService/UpdateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityServiceServer).UpdateRole(ctx, req.(*UpdateRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorityService_DeleteRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityServiceServer).DeleteRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.admin.v1.service.AuthorityService/DeleteRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityServiceServer).DeleteRole(ctx, req.(*DeleteRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorityService_ListAuthority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protobuf.PagingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityServiceServer).ListAuthority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.admin.v1.service.AuthorityService/ListAuthority",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityServiceServer).ListAuthority(ctx, req.(*protobuf.PagingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorityService_CreateAuthority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAuthorityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityServiceServer).CreateAuthority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.admin.v1.service.AuthorityService/CreateAuthority",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityServiceServer).CreateAuthority(ctx, req.(*CreateAuthorityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorityService_GetAuthority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthorityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityServiceServer).GetAuthority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.admin.v1.service.AuthorityService/GetAuthority",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityServiceServer).GetAuthority(ctx, req.(*GetAuthorityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorityService_UpdateAuthority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAuthorityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityServiceServer).UpdateAuthority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.admin.v1.service.AuthorityService/UpdateAuthority",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityServiceServer).UpdateAuthority(ctx, req.(*UpdateAuthorityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorityService_DeleteAuthority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAuthorityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityServiceServer).DeleteAuthority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.admin.v1.service.AuthorityService/DeleteAuthority",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityServiceServer).DeleteAuthority(ctx, req.(*DeleteAuthorityReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthorityService_ServiceDesc is the grpc.ServiceDesc for AuthorityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthorityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.admin.v1.service.AuthorityService",
	HandlerType: (*AuthorityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRole",
			Handler:    _AuthorityService_ListRole_Handler,
		},
		{
			MethodName: "CreateRole",
			Handler:    _AuthorityService_CreateRole_Handler,
		},
		{
			MethodName: "GetRole",
			Handler:    _AuthorityService_GetRole_Handler,
		},
		{
			MethodName: "UpdateRole",
			Handler:    _AuthorityService_UpdateRole_Handler,
		},
		{
			MethodName: "DeleteRole",
			Handler:    _AuthorityService_DeleteRole_Handler,
		},
		{
			MethodName: "ListAuthority",
			Handler:    _AuthorityService_ListAuthority_Handler,
		},
		{
			MethodName: "CreateAuthority",
			Handler:    _AuthorityService_CreateAuthority_Handler,
		},
		{
			MethodName: "GetAuthority",
			Handler:    _AuthorityService_GetAuthority_Handler,
		},
		{
			MethodName: "UpdateAuthority",
			Handler:    _AuthorityService_UpdateAuthority_Handler,
		},
		{
			MethodName: "DeleteAuthority",
			Handler:    _AuthorityService_DeleteAuthority_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/v1/service/authority.proto",
}
