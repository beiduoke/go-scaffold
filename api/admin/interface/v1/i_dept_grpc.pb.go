// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: admin/interface/v1/i_dept.proto

package v1

import (
	context "context"
	pagination "github.com/beiduoke/go-scaffold/api/common/pagination"
	v1 "github.com/beiduoke/go-scaffold/api/core/service/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	DeptService_ListDept_FullMethodName   = "/admin.interface.v1.DeptService/ListDept"
	DeptService_GetDept_FullMethodName    = "/admin.interface.v1.DeptService/GetDept"
	DeptService_CreateDept_FullMethodName = "/admin.interface.v1.DeptService/CreateDept"
	DeptService_UpdateDept_FullMethodName = "/admin.interface.v1.DeptService/UpdateDept"
	DeptService_DeleteDept_FullMethodName = "/admin.interface.v1.DeptService/DeleteDept"
)

// DeptServiceClient is the client API for DeptService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 部门管理服务
type DeptServiceClient interface {
	// 获取部门列表
	ListDept(ctx context.Context, in *pagination.PagingRequest, opts ...grpc.CallOption) (*v1.ListDeptResponse, error)
	// 获取部门数据
	GetDept(ctx context.Context, in *v1.GetDeptRequest, opts ...grpc.CallOption) (*v1.Dept, error)
	// 创建部门
	CreateDept(ctx context.Context, in *v1.CreateDeptRequest, opts ...grpc.CallOption) (*v1.CreateDeptResponse, error)
	// 更新部门
	UpdateDept(ctx context.Context, in *v1.UpdateDeptRequest, opts ...grpc.CallOption) (*v1.UpdateDeptResponse, error)
	// 删除部门
	DeleteDept(ctx context.Context, in *v1.DeleteDeptRequest, opts ...grpc.CallOption) (*v1.DeleteDeptResponse, error)
}

type deptServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeptServiceClient(cc grpc.ClientConnInterface) DeptServiceClient {
	return &deptServiceClient{cc}
}

func (c *deptServiceClient) ListDept(ctx context.Context, in *pagination.PagingRequest, opts ...grpc.CallOption) (*v1.ListDeptResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.ListDeptResponse)
	err := c.cc.Invoke(ctx, DeptService_ListDept_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deptServiceClient) GetDept(ctx context.Context, in *v1.GetDeptRequest, opts ...grpc.CallOption) (*v1.Dept, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.Dept)
	err := c.cc.Invoke(ctx, DeptService_GetDept_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deptServiceClient) CreateDept(ctx context.Context, in *v1.CreateDeptRequest, opts ...grpc.CallOption) (*v1.CreateDeptResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.CreateDeptResponse)
	err := c.cc.Invoke(ctx, DeptService_CreateDept_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deptServiceClient) UpdateDept(ctx context.Context, in *v1.UpdateDeptRequest, opts ...grpc.CallOption) (*v1.UpdateDeptResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.UpdateDeptResponse)
	err := c.cc.Invoke(ctx, DeptService_UpdateDept_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deptServiceClient) DeleteDept(ctx context.Context, in *v1.DeleteDeptRequest, opts ...grpc.CallOption) (*v1.DeleteDeptResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.DeleteDeptResponse)
	err := c.cc.Invoke(ctx, DeptService_DeleteDept_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeptServiceServer is the server API for DeptService service.
// All implementations must embed UnimplementedDeptServiceServer
// for forward compatibility.
//
// 部门管理服务
type DeptServiceServer interface {
	// 获取部门列表
	ListDept(context.Context, *pagination.PagingRequest) (*v1.ListDeptResponse, error)
	// 获取部门数据
	GetDept(context.Context, *v1.GetDeptRequest) (*v1.Dept, error)
	// 创建部门
	CreateDept(context.Context, *v1.CreateDeptRequest) (*v1.CreateDeptResponse, error)
	// 更新部门
	UpdateDept(context.Context, *v1.UpdateDeptRequest) (*v1.UpdateDeptResponse, error)
	// 删除部门
	DeleteDept(context.Context, *v1.DeleteDeptRequest) (*v1.DeleteDeptResponse, error)
	mustEmbedUnimplementedDeptServiceServer()
}

// UnimplementedDeptServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDeptServiceServer struct{}

func (UnimplementedDeptServiceServer) ListDept(context.Context, *pagination.PagingRequest) (*v1.ListDeptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDept not implemented")
}
func (UnimplementedDeptServiceServer) GetDept(context.Context, *v1.GetDeptRequest) (*v1.Dept, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDept not implemented")
}
func (UnimplementedDeptServiceServer) CreateDept(context.Context, *v1.CreateDeptRequest) (*v1.CreateDeptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDept not implemented")
}
func (UnimplementedDeptServiceServer) UpdateDept(context.Context, *v1.UpdateDeptRequest) (*v1.UpdateDeptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDept not implemented")
}
func (UnimplementedDeptServiceServer) DeleteDept(context.Context, *v1.DeleteDeptRequest) (*v1.DeleteDeptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDept not implemented")
}
func (UnimplementedDeptServiceServer) mustEmbedUnimplementedDeptServiceServer() {}
func (UnimplementedDeptServiceServer) testEmbeddedByValue()                     {}

// UnsafeDeptServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeptServiceServer will
// result in compilation errors.
type UnsafeDeptServiceServer interface {
	mustEmbedUnimplementedDeptServiceServer()
}

func RegisterDeptServiceServer(s grpc.ServiceRegistrar, srv DeptServiceServer) {
	// If the following call pancis, it indicates UnimplementedDeptServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DeptService_ServiceDesc, srv)
}

func _DeptService_ListDept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pagination.PagingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeptServiceServer).ListDept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeptService_ListDept_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeptServiceServer).ListDept(ctx, req.(*pagination.PagingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeptService_GetDept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetDeptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeptServiceServer).GetDept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeptService_GetDept_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeptServiceServer).GetDept(ctx, req.(*v1.GetDeptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeptService_CreateDept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.CreateDeptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeptServiceServer).CreateDept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeptService_CreateDept_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeptServiceServer).CreateDept(ctx, req.(*v1.CreateDeptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeptService_UpdateDept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.UpdateDeptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeptServiceServer).UpdateDept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeptService_UpdateDept_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeptServiceServer).UpdateDept(ctx, req.(*v1.UpdateDeptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeptService_DeleteDept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.DeleteDeptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeptServiceServer).DeleteDept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeptService_DeleteDept_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeptServiceServer).DeleteDept(ctx, req.(*v1.DeleteDeptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeptService_ServiceDesc is the grpc.ServiceDesc for DeptService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeptService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.interface.v1.DeptService",
	HandlerType: (*DeptServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDept",
			Handler:    _DeptService_ListDept_Handler,
		},
		{
			MethodName: "GetDept",
			Handler:    _DeptService_GetDept_Handler,
		},
		{
			MethodName: "CreateDept",
			Handler:    _DeptService_CreateDept_Handler,
		},
		{
			MethodName: "UpdateDept",
			Handler:    _DeptService_UpdateDept_Handler,
		},
		{
			MethodName: "DeleteDept",
			Handler:    _DeptService_DeleteDept_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/interface/v1/i_dept.proto",
}
