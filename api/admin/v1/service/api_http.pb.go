// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.4.0
// - protoc             v3.21.5
// source: admin/v1/service/api.proto

package service

import (
	context "context"
	protobuf "github.com/beiduoke/go-scaffold/api/protobuf"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationApiServiceCreateApi = "/api.admin.v1.service.ApiService/CreateApi"
const OperationApiServiceDeleteApi = "/api.admin.v1.service.ApiService/DeleteApi"
const OperationApiServiceGetApi = "/api.admin.v1.service.ApiService/GetApi"
const OperationApiServiceListApi = "/api.admin.v1.service.ApiService/ListApi"
const OperationApiServiceUpdateApi = "/api.admin.v1.service.ApiService/UpdateApi"

type ApiServiceHTTPServer interface {
	CreateApi(context.Context, *CreateApiReq) (*CreateApiReply, error)
	DeleteApi(context.Context, *DeleteApiReq) (*DeleteApiReply, error)
	GetApi(context.Context, *GetApiReq) (*Api, error)
	ListApi(context.Context, *protobuf.PagingReq) (*protobuf.PagingReply, error)
	UpdateApi(context.Context, *UpdateApiReq) (*UpdateApiReply, error)
}

func RegisterApiServiceHTTPServer(s *http.Server, srv ApiServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/admin/v1/apis", _ApiService_ListApi0_HTTP_Handler(srv))
	r.POST("/admin/v1/apis", _ApiService_CreateApi0_HTTP_Handler(srv))
	r.GET("/admin/v1/apis/{id}", _ApiService_GetApi0_HTTP_Handler(srv))
	r.PUT("/admin/v1/apis/{id}", _ApiService_UpdateApi0_HTTP_Handler(srv))
	r.DELETE("/admin/v1/apis/{id}", _ApiService_DeleteApi0_HTTP_Handler(srv))
}

func _ApiService_ListApi0_HTTP_Handler(srv ApiServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in protobuf.PagingReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationApiServiceListApi)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListApi(ctx, req.(*protobuf.PagingReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*protobuf.PagingReply)
		return ctx.Result(200, reply)
	}
}

func _ApiService_CreateApi0_HTTP_Handler(srv ApiServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateApiReq
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationApiServiceCreateApi)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateApi(ctx, req.(*CreateApiReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateApiReply)
		return ctx.Result(200, reply)
	}
}

func _ApiService_GetApi0_HTTP_Handler(srv ApiServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetApiReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationApiServiceGetApi)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetApi(ctx, req.(*GetApiReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Api)
		return ctx.Result(200, reply)
	}
}

func _ApiService_UpdateApi0_HTTP_Handler(srv ApiServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateApiReq
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationApiServiceUpdateApi)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateApi(ctx, req.(*UpdateApiReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateApiReply)
		return ctx.Result(200, reply)
	}
}

func _ApiService_DeleteApi0_HTTP_Handler(srv ApiServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteApiReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationApiServiceDeleteApi)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteApi(ctx, req.(*DeleteApiReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteApiReply)
		return ctx.Result(200, reply)
	}
}

type ApiServiceHTTPClient interface {
	CreateApi(ctx context.Context, req *CreateApiReq, opts ...http.CallOption) (rsp *CreateApiReply, err error)
	DeleteApi(ctx context.Context, req *DeleteApiReq, opts ...http.CallOption) (rsp *DeleteApiReply, err error)
	GetApi(ctx context.Context, req *GetApiReq, opts ...http.CallOption) (rsp *Api, err error)
	ListApi(ctx context.Context, req *protobuf.PagingReq, opts ...http.CallOption) (rsp *protobuf.PagingReply, err error)
	UpdateApi(ctx context.Context, req *UpdateApiReq, opts ...http.CallOption) (rsp *UpdateApiReply, err error)
}

type ApiServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewApiServiceHTTPClient(client *http.Client) ApiServiceHTTPClient {
	return &ApiServiceHTTPClientImpl{client}
}

func (c *ApiServiceHTTPClientImpl) CreateApi(ctx context.Context, in *CreateApiReq, opts ...http.CallOption) (*CreateApiReply, error) {
	var out CreateApiReply
	pattern := "/admin/v1/apis"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationApiServiceCreateApi))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ApiServiceHTTPClientImpl) DeleteApi(ctx context.Context, in *DeleteApiReq, opts ...http.CallOption) (*DeleteApiReply, error) {
	var out DeleteApiReply
	pattern := "/admin/v1/apis/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationApiServiceDeleteApi))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ApiServiceHTTPClientImpl) GetApi(ctx context.Context, in *GetApiReq, opts ...http.CallOption) (*Api, error) {
	var out Api
	pattern := "/admin/v1/apis/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationApiServiceGetApi))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ApiServiceHTTPClientImpl) ListApi(ctx context.Context, in *protobuf.PagingReq, opts ...http.CallOption) (*protobuf.PagingReply, error) {
	var out protobuf.PagingReply
	pattern := "/admin/v1/apis"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationApiServiceListApi))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ApiServiceHTTPClientImpl) UpdateApi(ctx context.Context, in *UpdateApiReq, opts ...http.CallOption) (*UpdateApiReply, error) {
	var out UpdateApiReply
	pattern := "/admin/v1/apis/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationApiServiceUpdateApi))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
