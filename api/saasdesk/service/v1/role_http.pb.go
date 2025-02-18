// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.3
// - protoc             (unknown)
// source: saasdesk/service/v1/role.proto

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

const OperationRoleServiceCreateRole = "/saasdesk.service.v1.RoleService/CreateRole"

type RoleServiceHTTPServer interface {
	CreateRole(context.Context, *CreateRoleRequest) (*CreateRoleResponse, error)
}

func RegisterRoleServiceHTTPServer(s *http.Server, srv RoleServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/role", _RoleService_CreateRole0_HTTP_Handler(srv))
}

func _RoleService_CreateRole0_HTTP_Handler(srv RoleServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateRoleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRoleServiceCreateRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateRole(ctx, req.(*CreateRoleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateRoleResponse)
		return ctx.Result(200, reply)
	}
}

type RoleServiceHTTPClient interface {
	CreateRole(ctx context.Context, req *CreateRoleRequest, opts ...http.CallOption) (rsp *CreateRoleResponse, err error)
}

type RoleServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewRoleServiceHTTPClient(client *http.Client) RoleServiceHTTPClient {
	return &RoleServiceHTTPClientImpl{client}
}

func (c *RoleServiceHTTPClientImpl) CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...http.CallOption) (*CreateRoleResponse, error) {
	var out CreateRoleResponse
	pattern := "/v1/role"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRoleServiceCreateRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
