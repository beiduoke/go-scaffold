// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: admin/service/v1/domain.proto

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
	DomainService_ListDomain_FullMethodName               = "/admin.service.v1.DomainService/ListDomain"
	DomainService_ListDomainTree_FullMethodName           = "/admin.service.v1.DomainService/ListDomainTree"
	DomainService_CreateDomain_FullMethodName             = "/admin.service.v1.DomainService/CreateDomain"
	DomainService_GetDomain_FullMethodName                = "/admin.service.v1.DomainService/GetDomain"
	DomainService_GetDomainCode_FullMethodName            = "/admin.service.v1.DomainService/GetDomainCode"
	DomainService_GetDomainName_FullMethodName            = "/admin.service.v1.DomainService/GetDomainName"
	DomainService_UpdateDomain_FullMethodName             = "/admin.service.v1.DomainService/UpdateDomain"
	DomainService_DeleteDomain_FullMethodName             = "/admin.service.v1.DomainService/DeleteDomain"
	DomainService_UpdateDomainState_FullMethodName        = "/admin.service.v1.DomainService/UpdateDomainState"
	DomainService_ListDomainMenu_FullMethodName           = "/admin.service.v1.DomainService/ListDomainMenu"
	DomainService_HandleDomainMenu_FullMethodName         = "/admin.service.v1.DomainService/HandleDomainMenu"
	DomainService_ListDomainPackage_FullMethodName        = "/admin.service.v1.DomainService/ListDomainPackage"
	DomainService_CreateDomainPackage_FullMethodName      = "/admin.service.v1.DomainService/CreateDomainPackage"
	DomainService_GetDomainPackage_FullMethodName         = "/admin.service.v1.DomainService/GetDomainPackage"
	DomainService_UpdateDomainPackage_FullMethodName      = "/admin.service.v1.DomainService/UpdateDomainPackage"
	DomainService_UpdateDomainPackageState_FullMethodName = "/admin.service.v1.DomainService/UpdateDomainPackageState"
	DomainService_DeleteDomainPackage_FullMethodName      = "/admin.service.v1.DomainService/DeleteDomainPackage"
)

// DomainServiceClient is the client API for DomainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DomainServiceClient interface {
	// 列表租户
	ListDomain(ctx context.Context, in *ListDomainRequest, opts ...grpc.CallOption) (*ListDomainResponse, error)
	// 获取租户树形列表
	ListDomainTree(ctx context.Context, in *ListDomainTreeRequest, opts ...grpc.CallOption) (*ListDomainTreeResponse, error)
	// 创建租户
	CreateDomain(ctx context.Context, in *CreateDomainRequest, opts ...grpc.CallOption) (*CreateDomainResponse, error)
	// 获取租户
	GetDomain(ctx context.Context, in *GetDomainRequest, opts ...grpc.CallOption) (*Domain, error)
	// 获取租户
	GetDomainCode(ctx context.Context, in *GetDomainCodeRequest, opts ...grpc.CallOption) (*Domain, error)
	// 获取租户
	GetDomainName(ctx context.Context, in *GetDomainNameRequest, opts ...grpc.CallOption) (*Domain, error)
	// 修改租户
	UpdateDomain(ctx context.Context, in *UpdateDomainRequest, opts ...grpc.CallOption) (*UpdateDomainResponse, error)
	// 删除租户
	DeleteDomain(ctx context.Context, in *DeleteDomainRequest, opts ...grpc.CallOption) (*DeleteDomainResponse, error)
	// 设置租户状态
	UpdateDomainState(ctx context.Context, in *UpdateDomainStateRequest, opts ...grpc.CallOption) (*UpdateDomainStateResponse, error)
	// 获取租户菜单
	ListDomainMenu(ctx context.Context, in *ListDomainMenuRequest, opts ...grpc.CallOption) (*ListDomainMenuResponse, error)
	// 处理租户菜单
	HandleDomainMenu(ctx context.Context, in *HandleDomainMenuRequest, opts ...grpc.CallOption) (*HandleDomainMenuResponse, error)
	// 列表租户套餐
	ListDomainPackage(ctx context.Context, in *ListDomainPackageRequest, opts ...grpc.CallOption) (*ListDomainPackageResponse, error)
	// 创建租户套餐
	CreateDomainPackage(ctx context.Context, in *CreateDomainPackageRequest, opts ...grpc.CallOption) (*CreateDomainPackageResponse, error)
	// 获取租户套餐
	GetDomainPackage(ctx context.Context, in *GetDomainPackageRequest, opts ...grpc.CallOption) (*DomainPackage, error)
	// 修改租户套餐
	UpdateDomainPackage(ctx context.Context, in *UpdateDomainPackageRequest, opts ...grpc.CallOption) (*UpdateDomainPackageResponse, error)
	// 更新指定ID套餐状态
	UpdateDomainPackageState(ctx context.Context, in *UpdateDomainPackageStateRequest, opts ...grpc.CallOption) (*UpdateDomainPackageStateResponse, error)
	// 删除租户套餐
	DeleteDomainPackage(ctx context.Context, in *DeleteDomainPackageRequest, opts ...grpc.CallOption) (*DeleteDomainPackageResponse, error)
}

type domainServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDomainServiceClient(cc grpc.ClientConnInterface) DomainServiceClient {
	return &domainServiceClient{cc}
}

func (c *domainServiceClient) ListDomain(ctx context.Context, in *ListDomainRequest, opts ...grpc.CallOption) (*ListDomainResponse, error) {
	out := new(ListDomainResponse)
	err := c.cc.Invoke(ctx, DomainService_ListDomain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainServiceClient) ListDomainTree(ctx context.Context, in *ListDomainTreeRequest, opts ...grpc.CallOption) (*ListDomainTreeResponse, error) {
	out := new(ListDomainTreeResponse)
	err := c.cc.Invoke(ctx, DomainService_ListDomainTree_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainServiceClient) CreateDomain(ctx context.Context, in *CreateDomainRequest, opts ...grpc.CallOption) (*CreateDomainResponse, error) {
	out := new(CreateDomainResponse)
	err := c.cc.Invoke(ctx, DomainService_CreateDomain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainServiceClient) GetDomain(ctx context.Context, in *GetDomainRequest, opts ...grpc.CallOption) (*Domain, error) {
	out := new(Domain)
	err := c.cc.Invoke(ctx, DomainService_GetDomain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainServiceClient) GetDomainCode(ctx context.Context, in *GetDomainCodeRequest, opts ...grpc.CallOption) (*Domain, error) {
	out := new(Domain)
	err := c.cc.Invoke(ctx, DomainService_GetDomainCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainServiceClient) GetDomainName(ctx context.Context, in *GetDomainNameRequest, opts ...grpc.CallOption) (*Domain, error) {
	out := new(Domain)
	err := c.cc.Invoke(ctx, DomainService_GetDomainName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainServiceClient) UpdateDomain(ctx context.Context, in *UpdateDomainRequest, opts ...grpc.CallOption) (*UpdateDomainResponse, error) {
	out := new(UpdateDomainResponse)
	err := c.cc.Invoke(ctx, DomainService_UpdateDomain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainServiceClient) DeleteDomain(ctx context.Context, in *DeleteDomainRequest, opts ...grpc.CallOption) (*DeleteDomainResponse, error) {
	out := new(DeleteDomainResponse)
	err := c.cc.Invoke(ctx, DomainService_DeleteDomain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainServiceClient) UpdateDomainState(ctx context.Context, in *UpdateDomainStateRequest, opts ...grpc.CallOption) (*UpdateDomainStateResponse, error) {
	out := new(UpdateDomainStateResponse)
	err := c.cc.Invoke(ctx, DomainService_UpdateDomainState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainServiceClient) ListDomainMenu(ctx context.Context, in *ListDomainMenuRequest, opts ...grpc.CallOption) (*ListDomainMenuResponse, error) {
	out := new(ListDomainMenuResponse)
	err := c.cc.Invoke(ctx, DomainService_ListDomainMenu_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainServiceClient) HandleDomainMenu(ctx context.Context, in *HandleDomainMenuRequest, opts ...grpc.CallOption) (*HandleDomainMenuResponse, error) {
	out := new(HandleDomainMenuResponse)
	err := c.cc.Invoke(ctx, DomainService_HandleDomainMenu_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainServiceClient) ListDomainPackage(ctx context.Context, in *ListDomainPackageRequest, opts ...grpc.CallOption) (*ListDomainPackageResponse, error) {
	out := new(ListDomainPackageResponse)
	err := c.cc.Invoke(ctx, DomainService_ListDomainPackage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainServiceClient) CreateDomainPackage(ctx context.Context, in *CreateDomainPackageRequest, opts ...grpc.CallOption) (*CreateDomainPackageResponse, error) {
	out := new(CreateDomainPackageResponse)
	err := c.cc.Invoke(ctx, DomainService_CreateDomainPackage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainServiceClient) GetDomainPackage(ctx context.Context, in *GetDomainPackageRequest, opts ...grpc.CallOption) (*DomainPackage, error) {
	out := new(DomainPackage)
	err := c.cc.Invoke(ctx, DomainService_GetDomainPackage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainServiceClient) UpdateDomainPackage(ctx context.Context, in *UpdateDomainPackageRequest, opts ...grpc.CallOption) (*UpdateDomainPackageResponse, error) {
	out := new(UpdateDomainPackageResponse)
	err := c.cc.Invoke(ctx, DomainService_UpdateDomainPackage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainServiceClient) UpdateDomainPackageState(ctx context.Context, in *UpdateDomainPackageStateRequest, opts ...grpc.CallOption) (*UpdateDomainPackageStateResponse, error) {
	out := new(UpdateDomainPackageStateResponse)
	err := c.cc.Invoke(ctx, DomainService_UpdateDomainPackageState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainServiceClient) DeleteDomainPackage(ctx context.Context, in *DeleteDomainPackageRequest, opts ...grpc.CallOption) (*DeleteDomainPackageResponse, error) {
	out := new(DeleteDomainPackageResponse)
	err := c.cc.Invoke(ctx, DomainService_DeleteDomainPackage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DomainServiceServer is the server API for DomainService service.
// All implementations must embed UnimplementedDomainServiceServer
// for forward compatibility
type DomainServiceServer interface {
	// 列表租户
	ListDomain(context.Context, *ListDomainRequest) (*ListDomainResponse, error)
	// 获取租户树形列表
	ListDomainTree(context.Context, *ListDomainTreeRequest) (*ListDomainTreeResponse, error)
	// 创建租户
	CreateDomain(context.Context, *CreateDomainRequest) (*CreateDomainResponse, error)
	// 获取租户
	GetDomain(context.Context, *GetDomainRequest) (*Domain, error)
	// 获取租户
	GetDomainCode(context.Context, *GetDomainCodeRequest) (*Domain, error)
	// 获取租户
	GetDomainName(context.Context, *GetDomainNameRequest) (*Domain, error)
	// 修改租户
	UpdateDomain(context.Context, *UpdateDomainRequest) (*UpdateDomainResponse, error)
	// 删除租户
	DeleteDomain(context.Context, *DeleteDomainRequest) (*DeleteDomainResponse, error)
	// 设置租户状态
	UpdateDomainState(context.Context, *UpdateDomainStateRequest) (*UpdateDomainStateResponse, error)
	// 获取租户菜单
	ListDomainMenu(context.Context, *ListDomainMenuRequest) (*ListDomainMenuResponse, error)
	// 处理租户菜单
	HandleDomainMenu(context.Context, *HandleDomainMenuRequest) (*HandleDomainMenuResponse, error)
	// 列表租户套餐
	ListDomainPackage(context.Context, *ListDomainPackageRequest) (*ListDomainPackageResponse, error)
	// 创建租户套餐
	CreateDomainPackage(context.Context, *CreateDomainPackageRequest) (*CreateDomainPackageResponse, error)
	// 获取租户套餐
	GetDomainPackage(context.Context, *GetDomainPackageRequest) (*DomainPackage, error)
	// 修改租户套餐
	UpdateDomainPackage(context.Context, *UpdateDomainPackageRequest) (*UpdateDomainPackageResponse, error)
	// 更新指定ID套餐状态
	UpdateDomainPackageState(context.Context, *UpdateDomainPackageStateRequest) (*UpdateDomainPackageStateResponse, error)
	// 删除租户套餐
	DeleteDomainPackage(context.Context, *DeleteDomainPackageRequest) (*DeleteDomainPackageResponse, error)
	mustEmbedUnimplementedDomainServiceServer()
}

// UnimplementedDomainServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDomainServiceServer struct {
}

func (UnimplementedDomainServiceServer) ListDomain(context.Context, *ListDomainRequest) (*ListDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDomain not implemented")
}
func (UnimplementedDomainServiceServer) ListDomainTree(context.Context, *ListDomainTreeRequest) (*ListDomainTreeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDomainTree not implemented")
}
func (UnimplementedDomainServiceServer) CreateDomain(context.Context, *CreateDomainRequest) (*CreateDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDomain not implemented")
}
func (UnimplementedDomainServiceServer) GetDomain(context.Context, *GetDomainRequest) (*Domain, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDomain not implemented")
}
func (UnimplementedDomainServiceServer) GetDomainCode(context.Context, *GetDomainCodeRequest) (*Domain, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDomainCode not implemented")
}
func (UnimplementedDomainServiceServer) GetDomainName(context.Context, *GetDomainNameRequest) (*Domain, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDomainName not implemented")
}
func (UnimplementedDomainServiceServer) UpdateDomain(context.Context, *UpdateDomainRequest) (*UpdateDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDomain not implemented")
}
func (UnimplementedDomainServiceServer) DeleteDomain(context.Context, *DeleteDomainRequest) (*DeleteDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDomain not implemented")
}
func (UnimplementedDomainServiceServer) UpdateDomainState(context.Context, *UpdateDomainStateRequest) (*UpdateDomainStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDomainState not implemented")
}
func (UnimplementedDomainServiceServer) ListDomainMenu(context.Context, *ListDomainMenuRequest) (*ListDomainMenuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDomainMenu not implemented")
}
func (UnimplementedDomainServiceServer) HandleDomainMenu(context.Context, *HandleDomainMenuRequest) (*HandleDomainMenuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleDomainMenu not implemented")
}
func (UnimplementedDomainServiceServer) ListDomainPackage(context.Context, *ListDomainPackageRequest) (*ListDomainPackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDomainPackage not implemented")
}
func (UnimplementedDomainServiceServer) CreateDomainPackage(context.Context, *CreateDomainPackageRequest) (*CreateDomainPackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDomainPackage not implemented")
}
func (UnimplementedDomainServiceServer) GetDomainPackage(context.Context, *GetDomainPackageRequest) (*DomainPackage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDomainPackage not implemented")
}
func (UnimplementedDomainServiceServer) UpdateDomainPackage(context.Context, *UpdateDomainPackageRequest) (*UpdateDomainPackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDomainPackage not implemented")
}
func (UnimplementedDomainServiceServer) UpdateDomainPackageState(context.Context, *UpdateDomainPackageStateRequest) (*UpdateDomainPackageStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDomainPackageState not implemented")
}
func (UnimplementedDomainServiceServer) DeleteDomainPackage(context.Context, *DeleteDomainPackageRequest) (*DeleteDomainPackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDomainPackage not implemented")
}
func (UnimplementedDomainServiceServer) mustEmbedUnimplementedDomainServiceServer() {}

// UnsafeDomainServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DomainServiceServer will
// result in compilation errors.
type UnsafeDomainServiceServer interface {
	mustEmbedUnimplementedDomainServiceServer()
}

func RegisterDomainServiceServer(s grpc.ServiceRegistrar, srv DomainServiceServer) {
	s.RegisterService(&DomainService_ServiceDesc, srv)
}

func _DomainService_ListDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServiceServer).ListDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainService_ListDomain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServiceServer).ListDomain(ctx, req.(*ListDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainService_ListDomainTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDomainTreeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServiceServer).ListDomainTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainService_ListDomainTree_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServiceServer).ListDomainTree(ctx, req.(*ListDomainTreeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainService_CreateDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServiceServer).CreateDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainService_CreateDomain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServiceServer).CreateDomain(ctx, req.(*CreateDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainService_GetDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServiceServer).GetDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainService_GetDomain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServiceServer).GetDomain(ctx, req.(*GetDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainService_GetDomainCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDomainCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServiceServer).GetDomainCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainService_GetDomainCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServiceServer).GetDomainCode(ctx, req.(*GetDomainCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainService_GetDomainName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDomainNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServiceServer).GetDomainName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainService_GetDomainName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServiceServer).GetDomainName(ctx, req.(*GetDomainNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainService_UpdateDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServiceServer).UpdateDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainService_UpdateDomain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServiceServer).UpdateDomain(ctx, req.(*UpdateDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainService_DeleteDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServiceServer).DeleteDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainService_DeleteDomain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServiceServer).DeleteDomain(ctx, req.(*DeleteDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainService_UpdateDomainState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDomainStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServiceServer).UpdateDomainState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainService_UpdateDomainState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServiceServer).UpdateDomainState(ctx, req.(*UpdateDomainStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainService_ListDomainMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDomainMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServiceServer).ListDomainMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainService_ListDomainMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServiceServer).ListDomainMenu(ctx, req.(*ListDomainMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainService_HandleDomainMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleDomainMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServiceServer).HandleDomainMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainService_HandleDomainMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServiceServer).HandleDomainMenu(ctx, req.(*HandleDomainMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainService_ListDomainPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDomainPackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServiceServer).ListDomainPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainService_ListDomainPackage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServiceServer).ListDomainPackage(ctx, req.(*ListDomainPackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainService_CreateDomainPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDomainPackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServiceServer).CreateDomainPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainService_CreateDomainPackage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServiceServer).CreateDomainPackage(ctx, req.(*CreateDomainPackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainService_GetDomainPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDomainPackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServiceServer).GetDomainPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainService_GetDomainPackage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServiceServer).GetDomainPackage(ctx, req.(*GetDomainPackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainService_UpdateDomainPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDomainPackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServiceServer).UpdateDomainPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainService_UpdateDomainPackage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServiceServer).UpdateDomainPackage(ctx, req.(*UpdateDomainPackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainService_UpdateDomainPackageState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDomainPackageStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServiceServer).UpdateDomainPackageState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainService_UpdateDomainPackageState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServiceServer).UpdateDomainPackageState(ctx, req.(*UpdateDomainPackageStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainService_DeleteDomainPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDomainPackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServiceServer).DeleteDomainPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainService_DeleteDomainPackage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServiceServer).DeleteDomainPackage(ctx, req.(*DeleteDomainPackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DomainService_ServiceDesc is the grpc.ServiceDesc for DomainService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DomainService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.service.v1.DomainService",
	HandlerType: (*DomainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDomain",
			Handler:    _DomainService_ListDomain_Handler,
		},
		{
			MethodName: "ListDomainTree",
			Handler:    _DomainService_ListDomainTree_Handler,
		},
		{
			MethodName: "CreateDomain",
			Handler:    _DomainService_CreateDomain_Handler,
		},
		{
			MethodName: "GetDomain",
			Handler:    _DomainService_GetDomain_Handler,
		},
		{
			MethodName: "GetDomainCode",
			Handler:    _DomainService_GetDomainCode_Handler,
		},
		{
			MethodName: "GetDomainName",
			Handler:    _DomainService_GetDomainName_Handler,
		},
		{
			MethodName: "UpdateDomain",
			Handler:    _DomainService_UpdateDomain_Handler,
		},
		{
			MethodName: "DeleteDomain",
			Handler:    _DomainService_DeleteDomain_Handler,
		},
		{
			MethodName: "UpdateDomainState",
			Handler:    _DomainService_UpdateDomainState_Handler,
		},
		{
			MethodName: "ListDomainMenu",
			Handler:    _DomainService_ListDomainMenu_Handler,
		},
		{
			MethodName: "HandleDomainMenu",
			Handler:    _DomainService_HandleDomainMenu_Handler,
		},
		{
			MethodName: "ListDomainPackage",
			Handler:    _DomainService_ListDomainPackage_Handler,
		},
		{
			MethodName: "CreateDomainPackage",
			Handler:    _DomainService_CreateDomainPackage_Handler,
		},
		{
			MethodName: "GetDomainPackage",
			Handler:    _DomainService_GetDomainPackage_Handler,
		},
		{
			MethodName: "UpdateDomainPackage",
			Handler:    _DomainService_UpdateDomainPackage_Handler,
		},
		{
			MethodName: "UpdateDomainPackageState",
			Handler:    _DomainService_UpdateDomainPackageState_Handler,
		},
		{
			MethodName: "DeleteDomainPackage",
			Handler:    _DomainService_DeleteDomainPackage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/service/v1/domain.proto",
}
