// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             (unknown)
// source: admin/service/v1/user.proto

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

const OperationUserServiceCreateUser = "/admin.service.v1.UserService/CreateUser"
const OperationUserServiceDeleteUser = "/admin.service.v1.UserService/DeleteUser"
const OperationUserServiceExistUserName = "/admin.service.v1.UserService/ExistUserName"
const OperationUserServiceGetUser = "/admin.service.v1.UserService/GetUser"
const OperationUserServiceGetUserInfo = "/admin.service.v1.UserService/GetUserInfo"
const OperationUserServiceGetUserProfile = "/admin.service.v1.UserService/GetUserProfile"
const OperationUserServiceListUser = "/admin.service.v1.UserService/ListUser"
const OperationUserServiceListUserRole = "/admin.service.v1.UserService/ListUserRole"
const OperationUserServiceListUserRoleMenuRouterTree = "/admin.service.v1.UserService/ListUserRoleMenuRouterTree"
const OperationUserServiceListUserRoleMenuTree = "/admin.service.v1.UserService/ListUserRoleMenuTree"
const OperationUserServiceListUserRolePermission = "/admin.service.v1.UserService/ListUserRolePermission"
const OperationUserServiceUpdateUser = "/admin.service.v1.UserService/UpdateUser"

type UserServiceHTTPServer interface {
	// CreateUser 创建用户
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	// DeleteUser 删除用户
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	// ExistUserName 验证用户名是否存在
	ExistUserName(context.Context, *ExistUserNameRequest) (*ExistUserNameResponse, error)
	// GetUser 获取用户
	GetUser(context.Context, *GetUserRequest) (*User, error)
	// GetUserInfo User 用户模块
	// 当前登录用户概述
	GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoResponse, error)
	// GetUserProfile 当前登录用户概述
	GetUserProfile(context.Context, *GetUserProfileRequest) (*GetUserProfileResponse, error)
	// ListUser 列表用户
	ListUser(context.Context, *ListUserRequest) (*ListUserResponse, error)
	// ListUserRole 当前登录用户拥有角色
	ListUserRole(context.Context, *ListUserRoleRequest) (*ListUserRoleResponse, error)
	// ListUserRoleMenuRouterTree 获取角色菜单路由树形列表
	ListUserRoleMenuRouterTree(context.Context, *ListUserRoleMenuRouterTreeRequest) (*ListUserRoleMenuRouterTreeResponse, error)
	// ListUserRoleMenuTree 获取角色菜单路由树形列表
	ListUserRoleMenuTree(context.Context, *ListUserRoleMenuTreeRequest) (*ListUserRoleMenuTreeResponse, error)
	// ListUserRolePermission 获取角色权限列表
	ListUserRolePermission(context.Context, *ListUserRolePermissionRequest) (*ListUserRolePermissionResponse, error)
	// UpdateUser 修改用户
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
}

func RegisterUserServiceHTTPServer(s *http.Server, srv UserServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/users/info", _UserService_GetUserInfo0_HTTP_Handler(srv))
	r.GET("/v1/users/profiles", _UserService_GetUserProfile0_HTTP_Handler(srv))
	r.GET("/v1/users/roles", _UserService_ListUserRole0_HTTP_Handler(srv))
	r.GET("/v1/users/roles/{role_id}/routers/trees", _UserService_ListUserRoleMenuRouterTree0_HTTP_Handler(srv))
	r.GET("/v1/users/routers/trees", _UserService_ListUserRoleMenuRouterTree1_HTTP_Handler(srv))
	r.GET("/v1/users/roles/{role_id}/menus/trees", _UserService_ListUserRoleMenuTree0_HTTP_Handler(srv))
	r.GET("/v1/users/menus/trees", _UserService_ListUserRoleMenuTree1_HTTP_Handler(srv))
	r.GET("/v1/users/roles/{role_id}/permissions", _UserService_ListUserRolePermission0_HTTP_Handler(srv))
	r.GET("/v1/users/permissions", _UserService_ListUserRolePermission1_HTTP_Handler(srv))
	r.GET("/v1/users", _UserService_ListUser0_HTTP_Handler(srv))
	r.POST("/v1/users", _UserService_CreateUser0_HTTP_Handler(srv))
	r.GET("/v1/users/{id}", _UserService_GetUser0_HTTP_Handler(srv))
	r.PUT("/v1/users/{id}", _UserService_UpdateUser0_HTTP_Handler(srv))
	r.DELETE("/v1/users/{id}", _UserService_DeleteUser0_HTTP_Handler(srv))
	r.POST("/v1/users/existName", _UserService_ExistUserName0_HTTP_Handler(srv))
}

func _UserService_GetUserInfo0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserInfoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceGetUserInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserInfo(ctx, req.(*GetUserInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserInfoResponse)
		return ctx.Result(200, reply)
	}
}

func _UserService_GetUserProfile0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserProfileRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceGetUserProfile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserProfile(ctx, req.(*GetUserProfileRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserProfileResponse)
		return ctx.Result(200, reply)
	}
}

func _UserService_ListUserRole0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListUserRoleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceListUserRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUserRole(ctx, req.(*ListUserRoleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListUserRoleResponse)
		return ctx.Result(200, reply)
	}
}

func _UserService_ListUserRoleMenuRouterTree0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListUserRoleMenuRouterTreeRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceListUserRoleMenuRouterTree)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUserRoleMenuRouterTree(ctx, req.(*ListUserRoleMenuRouterTreeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListUserRoleMenuRouterTreeResponse)
		return ctx.Result(200, reply)
	}
}

func _UserService_ListUserRoleMenuRouterTree1_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListUserRoleMenuRouterTreeRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceListUserRoleMenuRouterTree)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUserRoleMenuRouterTree(ctx, req.(*ListUserRoleMenuRouterTreeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListUserRoleMenuRouterTreeResponse)
		return ctx.Result(200, reply)
	}
}

func _UserService_ListUserRoleMenuTree0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListUserRoleMenuTreeRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceListUserRoleMenuTree)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUserRoleMenuTree(ctx, req.(*ListUserRoleMenuTreeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListUserRoleMenuTreeResponse)
		return ctx.Result(200, reply)
	}
}

func _UserService_ListUserRoleMenuTree1_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListUserRoleMenuTreeRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceListUserRoleMenuTree)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUserRoleMenuTree(ctx, req.(*ListUserRoleMenuTreeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListUserRoleMenuTreeResponse)
		return ctx.Result(200, reply)
	}
}

func _UserService_ListUserRolePermission0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListUserRolePermissionRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceListUserRolePermission)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUserRolePermission(ctx, req.(*ListUserRolePermissionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListUserRolePermissionResponse)
		return ctx.Result(200, reply)
	}
}

func _UserService_ListUserRolePermission1_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListUserRolePermissionRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceListUserRolePermission)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUserRolePermission(ctx, req.(*ListUserRolePermissionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListUserRolePermissionResponse)
		return ctx.Result(200, reply)
	}
}

func _UserService_ListUser0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceListUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUser(ctx, req.(*ListUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListUserResponse)
		return ctx.Result(200, reply)
	}
}

func _UserService_CreateUser0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceCreateUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateUser(ctx, req.(*CreateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateUserResponse)
		return ctx.Result(200, reply)
	}
}

func _UserService_GetUser0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceGetUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUser(ctx, req.(*GetUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*User)
		return ctx.Result(200, reply)
	}
}

func _UserService_UpdateUser0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceUpdateUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUser(ctx, req.(*UpdateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateUserResponse)
		return ctx.Result(200, reply)
	}
}

func _UserService_DeleteUser0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceDeleteUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteUser(ctx, req.(*DeleteUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteUserResponse)
		return ctx.Result(200, reply)
	}
}

func _UserService_ExistUserName0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ExistUserNameRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceExistUserName)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ExistUserName(ctx, req.(*ExistUserNameRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ExistUserNameResponse)
		return ctx.Result(200, reply)
	}
}

type UserServiceHTTPClient interface {
	CreateUser(ctx context.Context, req *CreateUserRequest, opts ...http.CallOption) (rsp *CreateUserResponse, err error)
	DeleteUser(ctx context.Context, req *DeleteUserRequest, opts ...http.CallOption) (rsp *DeleteUserResponse, err error)
	ExistUserName(ctx context.Context, req *ExistUserNameRequest, opts ...http.CallOption) (rsp *ExistUserNameResponse, err error)
	GetUser(ctx context.Context, req *GetUserRequest, opts ...http.CallOption) (rsp *User, err error)
	GetUserInfo(ctx context.Context, req *GetUserInfoRequest, opts ...http.CallOption) (rsp *GetUserInfoResponse, err error)
	GetUserProfile(ctx context.Context, req *GetUserProfileRequest, opts ...http.CallOption) (rsp *GetUserProfileResponse, err error)
	ListUser(ctx context.Context, req *ListUserRequest, opts ...http.CallOption) (rsp *ListUserResponse, err error)
	ListUserRole(ctx context.Context, req *ListUserRoleRequest, opts ...http.CallOption) (rsp *ListUserRoleResponse, err error)
	ListUserRoleMenuRouterTree(ctx context.Context, req *ListUserRoleMenuRouterTreeRequest, opts ...http.CallOption) (rsp *ListUserRoleMenuRouterTreeResponse, err error)
	ListUserRoleMenuTree(ctx context.Context, req *ListUserRoleMenuTreeRequest, opts ...http.CallOption) (rsp *ListUserRoleMenuTreeResponse, err error)
	ListUserRolePermission(ctx context.Context, req *ListUserRolePermissionRequest, opts ...http.CallOption) (rsp *ListUserRolePermissionResponse, err error)
	UpdateUser(ctx context.Context, req *UpdateUserRequest, opts ...http.CallOption) (rsp *UpdateUserResponse, err error)
}

type UserServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewUserServiceHTTPClient(client *http.Client) UserServiceHTTPClient {
	return &UserServiceHTTPClientImpl{client}
}

func (c *UserServiceHTTPClientImpl) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...http.CallOption) (*CreateUserResponse, error) {
	var out CreateUserResponse
	pattern := "/v1/users"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserServiceCreateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserServiceHTTPClientImpl) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...http.CallOption) (*DeleteUserResponse, error) {
	var out DeleteUserResponse
	pattern := "/v1/users/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserServiceDeleteUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserServiceHTTPClientImpl) ExistUserName(ctx context.Context, in *ExistUserNameRequest, opts ...http.CallOption) (*ExistUserNameResponse, error) {
	var out ExistUserNameResponse
	pattern := "/v1/users/existName"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserServiceExistUserName))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserServiceHTTPClientImpl) GetUser(ctx context.Context, in *GetUserRequest, opts ...http.CallOption) (*User, error) {
	var out User
	pattern := "/v1/users/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserServiceGetUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserServiceHTTPClientImpl) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...http.CallOption) (*GetUserInfoResponse, error) {
	var out GetUserInfoResponse
	pattern := "/v1/users/info"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserServiceGetUserInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserServiceHTTPClientImpl) GetUserProfile(ctx context.Context, in *GetUserProfileRequest, opts ...http.CallOption) (*GetUserProfileResponse, error) {
	var out GetUserProfileResponse
	pattern := "/v1/users/profiles"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserServiceGetUserProfile))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserServiceHTTPClientImpl) ListUser(ctx context.Context, in *ListUserRequest, opts ...http.CallOption) (*ListUserResponse, error) {
	var out ListUserResponse
	pattern := "/v1/users"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserServiceListUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserServiceHTTPClientImpl) ListUserRole(ctx context.Context, in *ListUserRoleRequest, opts ...http.CallOption) (*ListUserRoleResponse, error) {
	var out ListUserRoleResponse
	pattern := "/v1/users/roles"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserServiceListUserRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserServiceHTTPClientImpl) ListUserRoleMenuRouterTree(ctx context.Context, in *ListUserRoleMenuRouterTreeRequest, opts ...http.CallOption) (*ListUserRoleMenuRouterTreeResponse, error) {
	var out ListUserRoleMenuRouterTreeResponse
	pattern := "/v1/users/routers/trees"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserServiceListUserRoleMenuRouterTree))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserServiceHTTPClientImpl) ListUserRoleMenuTree(ctx context.Context, in *ListUserRoleMenuTreeRequest, opts ...http.CallOption) (*ListUserRoleMenuTreeResponse, error) {
	var out ListUserRoleMenuTreeResponse
	pattern := "/v1/users/menus/trees"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserServiceListUserRoleMenuTree))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserServiceHTTPClientImpl) ListUserRolePermission(ctx context.Context, in *ListUserRolePermissionRequest, opts ...http.CallOption) (*ListUserRolePermissionResponse, error) {
	var out ListUserRolePermissionResponse
	pattern := "/v1/users/permissions"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserServiceListUserRolePermission))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserServiceHTTPClientImpl) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...http.CallOption) (*UpdateUserResponse, error) {
	var out UpdateUserResponse
	pattern := "/v1/users/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserServiceUpdateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
