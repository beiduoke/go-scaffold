// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: admin/v1/dept.proto

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
	DeptService_ListDept_FullMethodName        = "/admin.v1.DeptService/ListDept"
	DeptService_CreateDept_FullMethodName      = "/admin.v1.DeptService/CreateDept"
	DeptService_GetDept_FullMethodName         = "/admin.v1.DeptService/GetDept"
	DeptService_UpdateDept_FullMethodName      = "/admin.v1.DeptService/UpdateDept"
	DeptService_UpdateDeptState_FullMethodName = "/admin.v1.DeptService/UpdateDeptState"
	DeptService_DeleteDept_FullMethodName      = "/admin.v1.DeptService/DeleteDept"
	DeptService_ListDeptTree_FullMethodName    = "/admin.v1.DeptService/ListDeptTree"
)

// DeptServiceClient is the client API for DeptService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeptServiceClient interface {
	// 列表部门
	ListDept(ctx context.Context, in *ListDeptRequest, opts ...grpc.CallOption) (*ListDeptResponse, error)
	// 创建部门
	CreateDept(ctx context.Context, in *CreateDeptRequest, opts ...grpc.CallOption) (*CreateDeptResponse, error)
	// 获取部门
	GetDept(ctx context.Context, in *GetDeptRequest, opts ...grpc.CallOption) (*Dept, error)
	// 修改部门
	UpdateDept(ctx context.Context, in *UpdateDeptRequest, opts ...grpc.CallOption) (*UpdateDeptResponse, error)
	// 更新指定ID角色状态
	UpdateDeptState(ctx context.Context, in *UpdateDeptStateRequest, opts ...grpc.CallOption) (*UpdateDeptStateResponse, error)
	// 删除部门
	DeleteDept(ctx context.Context, in *DeleteDeptRequest, opts ...grpc.CallOption) (*DeleteDeptResponse, error)
	// 获取全部部门树形
	ListDeptTree(ctx context.Context, in *ListDeptTreeRequest, opts ...grpc.CallOption) (*ListDeptTreeResponse, error)
}

type deptServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeptServiceClient(cc grpc.ClientConnInterface) DeptServiceClient {
	return &deptServiceClient{cc}
}

func (c *deptServiceClient) ListDept(ctx context.Context, in *ListDeptRequest, opts ...grpc.CallOption) (*ListDeptResponse, error) {
	out := new(ListDeptResponse)
	err := c.cc.Invoke(ctx, DeptService_ListDept_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deptServiceClient) CreateDept(ctx context.Context, in *CreateDeptRequest, opts ...grpc.CallOption) (*CreateDeptResponse, error) {
	out := new(CreateDeptResponse)
	err := c.cc.Invoke(ctx, DeptService_CreateDept_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deptServiceClient) GetDept(ctx context.Context, in *GetDeptRequest, opts ...grpc.CallOption) (*Dept, error) {
	out := new(Dept)
	err := c.cc.Invoke(ctx, DeptService_GetDept_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deptServiceClient) UpdateDept(ctx context.Context, in *UpdateDeptRequest, opts ...grpc.CallOption) (*UpdateDeptResponse, error) {
	out := new(UpdateDeptResponse)
	err := c.cc.Invoke(ctx, DeptService_UpdateDept_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deptServiceClient) UpdateDeptState(ctx context.Context, in *UpdateDeptStateRequest, opts ...grpc.CallOption) (*UpdateDeptStateResponse, error) {
	out := new(UpdateDeptStateResponse)
	err := c.cc.Invoke(ctx, DeptService_UpdateDeptState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deptServiceClient) DeleteDept(ctx context.Context, in *DeleteDeptRequest, opts ...grpc.CallOption) (*DeleteDeptResponse, error) {
	out := new(DeleteDeptResponse)
	err := c.cc.Invoke(ctx, DeptService_DeleteDept_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deptServiceClient) ListDeptTree(ctx context.Context, in *ListDeptTreeRequest, opts ...grpc.CallOption) (*ListDeptTreeResponse, error) {
	out := new(ListDeptTreeResponse)
	err := c.cc.Invoke(ctx, DeptService_ListDeptTree_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeptServiceServer is the server API for DeptService service.
// All implementations must embed UnimplementedDeptServiceServer
// for forward compatibility
type DeptServiceServer interface {
	// 列表部门
	ListDept(context.Context, *ListDeptRequest) (*ListDeptResponse, error)
	// 创建部门
	CreateDept(context.Context, *CreateDeptRequest) (*CreateDeptResponse, error)
	// 获取部门
	GetDept(context.Context, *GetDeptRequest) (*Dept, error)
	// 修改部门
	UpdateDept(context.Context, *UpdateDeptRequest) (*UpdateDeptResponse, error)
	// 更新指定ID角色状态
	UpdateDeptState(context.Context, *UpdateDeptStateRequest) (*UpdateDeptStateResponse, error)
	// 删除部门
	DeleteDept(context.Context, *DeleteDeptRequest) (*DeleteDeptResponse, error)
	// 获取全部部门树形
	ListDeptTree(context.Context, *ListDeptTreeRequest) (*ListDeptTreeResponse, error)
	mustEmbedUnimplementedDeptServiceServer()
}

// UnimplementedDeptServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDeptServiceServer struct {
}

func (UnimplementedDeptServiceServer) ListDept(context.Context, *ListDeptRequest) (*ListDeptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDept not implemented")
}
func (UnimplementedDeptServiceServer) CreateDept(context.Context, *CreateDeptRequest) (*CreateDeptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDept not implemented")
}
func (UnimplementedDeptServiceServer) GetDept(context.Context, *GetDeptRequest) (*Dept, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDept not implemented")
}
func (UnimplementedDeptServiceServer) UpdateDept(context.Context, *UpdateDeptRequest) (*UpdateDeptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDept not implemented")
}
func (UnimplementedDeptServiceServer) UpdateDeptState(context.Context, *UpdateDeptStateRequest) (*UpdateDeptStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDeptState not implemented")
}
func (UnimplementedDeptServiceServer) DeleteDept(context.Context, *DeleteDeptRequest) (*DeleteDeptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDept not implemented")
}
func (UnimplementedDeptServiceServer) ListDeptTree(context.Context, *ListDeptTreeRequest) (*ListDeptTreeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDeptTree not implemented")
}
func (UnimplementedDeptServiceServer) mustEmbedUnimplementedDeptServiceServer() {}

// UnsafeDeptServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeptServiceServer will
// result in compilation errors.
type UnsafeDeptServiceServer interface {
	mustEmbedUnimplementedDeptServiceServer()
}

func RegisterDeptServiceServer(s grpc.ServiceRegistrar, srv DeptServiceServer) {
	s.RegisterService(&DeptService_ServiceDesc, srv)
}

func _DeptService_ListDept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDeptRequest)
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
		return srv.(DeptServiceServer).ListDept(ctx, req.(*ListDeptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeptService_CreateDept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeptRequest)
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
		return srv.(DeptServiceServer).CreateDept(ctx, req.(*CreateDeptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeptService_GetDept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeptRequest)
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
		return srv.(DeptServiceServer).GetDept(ctx, req.(*GetDeptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeptService_UpdateDept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeptRequest)
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
		return srv.(DeptServiceServer).UpdateDept(ctx, req.(*UpdateDeptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeptService_UpdateDeptState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeptStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeptServiceServer).UpdateDeptState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeptService_UpdateDeptState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeptServiceServer).UpdateDeptState(ctx, req.(*UpdateDeptStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeptService_DeleteDept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeptRequest)
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
		return srv.(DeptServiceServer).DeleteDept(ctx, req.(*DeleteDeptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeptService_ListDeptTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDeptTreeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeptServiceServer).ListDeptTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeptService_ListDeptTree_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeptServiceServer).ListDeptTree(ctx, req.(*ListDeptTreeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeptService_ServiceDesc is the grpc.ServiceDesc for DeptService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeptService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.v1.DeptService",
	HandlerType: (*DeptServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDept",
			Handler:    _DeptService_ListDept_Handler,
		},
		{
			MethodName: "CreateDept",
			Handler:    _DeptService_CreateDept_Handler,
		},
		{
			MethodName: "GetDept",
			Handler:    _DeptService_GetDept_Handler,
		},
		{
			MethodName: "UpdateDept",
			Handler:    _DeptService_UpdateDept_Handler,
		},
		{
			MethodName: "UpdateDeptState",
			Handler:    _DeptService_UpdateDeptState_Handler,
		},
		{
			MethodName: "DeleteDept",
			Handler:    _DeptService_DeleteDept_Handler,
		},
		{
			MethodName: "ListDeptTree",
			Handler:    _DeptService_ListDeptTree_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/v1/dept.proto",
}
