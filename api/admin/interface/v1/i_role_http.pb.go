// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             (unknown)
// source: admin/interface/v1/i_role.proto

package v1

import (
	context "context"
	pagination "github.com/beiduoke/go-scaffold/api/common/pagination"
	v1 "github.com/beiduoke/go-scaffold/api/core/service/v1"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationRoleServiceCreateRole = "/admin.interface.v1.RoleService/CreateRole"
const OperationRoleServiceDeleteRole = "/admin.interface.v1.RoleService/DeleteRole"
const OperationRoleServiceGetRole = "/admin.interface.v1.RoleService/GetRole"
const OperationRoleServiceListRole = "/admin.interface.v1.RoleService/ListRole"
const OperationRoleServiceUpdateRole = "/admin.interface.v1.RoleService/UpdateRole"

type RoleServiceHTTPServer interface {
	// CreateRole 创建角色
	CreateRole(context.Context, *v1.CreateRoleRequest) (*v1.CreateRoleResponse, error)
	// DeleteRole 删除角色
	DeleteRole(context.Context, *v1.DeleteRoleRequest) (*v1.DeleteRoleResponse, error)
	// GetRole 获取角色数据
	GetRole(context.Context, *v1.GetRoleRequest) (*v1.Role, error)
	// ListRole 获取角色列表
	ListRole(context.Context, *pagination.PagingRequest) (*v1.ListRoleResponse, error)
	// UpdateRole 更新角色
	UpdateRole(context.Context, *v1.UpdateRoleRequest) (*v1.UpdateRoleResponse, error)
}

func RegisterRoleServiceHTTPServer(s *http.Server, srv RoleServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/admin/v1/roles", _RoleService_ListRole0_HTTP_Handler(srv))
	r.GET("/admin/v1/roles/{id}", _RoleService_GetRole0_HTTP_Handler(srv))
	r.POST("/admin/v1/roles", _RoleService_CreateRole0_HTTP_Handler(srv))
	r.PUT("/admin/v1/roles/{id}", _RoleService_UpdateRole0_HTTP_Handler(srv))
	r.DELETE("/admin/v1/roles/{id}", _RoleService_DeleteRole0_HTTP_Handler(srv))
}

func _RoleService_ListRole0_HTTP_Handler(srv RoleServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in pagination.PagingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRoleServiceListRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListRole(ctx, req.(*pagination.PagingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.ListRoleResponse)
		return ctx.Result(200, reply)
	}
}

func _RoleService_GetRole0_HTTP_Handler(srv RoleServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.GetRoleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRoleServiceGetRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetRole(ctx, req.(*v1.GetRoleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.Role)
		return ctx.Result(200, reply)
	}
}

func _RoleService_CreateRole0_HTTP_Handler(srv RoleServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.CreateRoleRequest
		if err := ctx.Bind(&in.User); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRoleServiceCreateRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateRole(ctx, req.(*v1.CreateRoleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.CreateRoleResponse)
		return ctx.Result(200, reply)
	}
}

func _RoleService_UpdateRole0_HTTP_Handler(srv RoleServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.UpdateRoleRequest
		if err := ctx.Bind(&in.User); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRoleServiceUpdateRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateRole(ctx, req.(*v1.UpdateRoleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.UpdateRoleResponse)
		return ctx.Result(200, reply)
	}
}

func _RoleService_DeleteRole0_HTTP_Handler(srv RoleServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.DeleteRoleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRoleServiceDeleteRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteRole(ctx, req.(*v1.DeleteRoleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.DeleteRoleResponse)
		return ctx.Result(200, reply)
	}
}

type RoleServiceHTTPClient interface {
	CreateRole(ctx context.Context, req *v1.CreateRoleRequest, opts ...http.CallOption) (rsp *v1.CreateRoleResponse, err error)
	DeleteRole(ctx context.Context, req *v1.DeleteRoleRequest, opts ...http.CallOption) (rsp *v1.DeleteRoleResponse, err error)
	GetRole(ctx context.Context, req *v1.GetRoleRequest, opts ...http.CallOption) (rsp *v1.Role, err error)
	ListRole(ctx context.Context, req *pagination.PagingRequest, opts ...http.CallOption) (rsp *v1.ListRoleResponse, err error)
	UpdateRole(ctx context.Context, req *v1.UpdateRoleRequest, opts ...http.CallOption) (rsp *v1.UpdateRoleResponse, err error)
}

type RoleServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewRoleServiceHTTPClient(client *http.Client) RoleServiceHTTPClient {
	return &RoleServiceHTTPClientImpl{client}
}

func (c *RoleServiceHTTPClientImpl) CreateRole(ctx context.Context, in *v1.CreateRoleRequest, opts ...http.CallOption) (*v1.CreateRoleResponse, error) {
	var out v1.CreateRoleResponse
	pattern := "/admin/v1/roles"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRoleServiceCreateRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.User, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RoleServiceHTTPClientImpl) DeleteRole(ctx context.Context, in *v1.DeleteRoleRequest, opts ...http.CallOption) (*v1.DeleteRoleResponse, error) {
	var out v1.DeleteRoleResponse
	pattern := "/admin/v1/roles/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRoleServiceDeleteRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RoleServiceHTTPClientImpl) GetRole(ctx context.Context, in *v1.GetRoleRequest, opts ...http.CallOption) (*v1.Role, error) {
	var out v1.Role
	pattern := "/admin/v1/roles/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRoleServiceGetRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RoleServiceHTTPClientImpl) ListRole(ctx context.Context, in *pagination.PagingRequest, opts ...http.CallOption) (*v1.ListRoleResponse, error) {
	var out v1.ListRoleResponse
	pattern := "/admin/v1/roles"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRoleServiceListRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RoleServiceHTTPClientImpl) UpdateRole(ctx context.Context, in *v1.UpdateRoleRequest, opts ...http.CallOption) (*v1.UpdateRoleResponse, error) {
	var out v1.UpdateRoleResponse
	pattern := "/admin/v1/roles/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRoleServiceUpdateRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.User, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
