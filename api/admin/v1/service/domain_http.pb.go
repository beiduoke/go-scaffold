// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.4.0
// - protoc             v3.21.5
// source: admin/v1/service/domain.proto

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

const OperationDomainServiceCreateDomain = "/api.admin.v1.service.DomainService/CreateDomain"
const OperationDomainServiceDeleteDomain = "/api.admin.v1.service.DomainService/DeleteDomain"
const OperationDomainServiceGetDomain = "/api.admin.v1.service.DomainService/GetDomain"
const OperationDomainServiceListDomain = "/api.admin.v1.service.DomainService/ListDomain"
const OperationDomainServiceUpdateDomain = "/api.admin.v1.service.DomainService/UpdateDomain"

type DomainServiceHTTPServer interface {
	CreateDomain(context.Context, *CreateDomainReq) (*CreateDomainReply, error)
	DeleteDomain(context.Context, *DeleteDomainReq) (*DeleteDomainReply, error)
	GetDomain(context.Context, *GetDomainReq) (*Domain, error)
	ListDomain(context.Context, *protobuf.PagingReq) (*protobuf.PagingReply, error)
	UpdateDomain(context.Context, *UpdateDomainReq) (*UpdateDomainReply, error)
}

func RegisterDomainServiceHTTPServer(s *http.Server, srv DomainServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/admin/v1/domains", _DomainService_ListDomain0_HTTP_Handler(srv))
	r.POST("/admin/v1/domains", _DomainService_CreateDomain0_HTTP_Handler(srv))
	r.GET("/admin/v1/domains/{id}", _DomainService_GetDomain0_HTTP_Handler(srv))
	r.PUT("/admin/v1/domains/{id}", _DomainService_UpdateDomain0_HTTP_Handler(srv))
	r.DELETE("/admin/v1/domains/{id}", _DomainService_DeleteDomain0_HTTP_Handler(srv))
}

func _DomainService_ListDomain0_HTTP_Handler(srv DomainServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in protobuf.PagingReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDomainServiceListDomain)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListDomain(ctx, req.(*protobuf.PagingReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*protobuf.PagingReply)
		return ctx.Result(200, reply)
	}
}

func _DomainService_CreateDomain0_HTTP_Handler(srv DomainServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateDomainReq
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDomainServiceCreateDomain)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateDomain(ctx, req.(*CreateDomainReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateDomainReply)
		return ctx.Result(200, reply)
	}
}

func _DomainService_GetDomain0_HTTP_Handler(srv DomainServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetDomainReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDomainServiceGetDomain)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetDomain(ctx, req.(*GetDomainReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Domain)
		return ctx.Result(200, reply)
	}
}

func _DomainService_UpdateDomain0_HTTP_Handler(srv DomainServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateDomainReq
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDomainServiceUpdateDomain)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateDomain(ctx, req.(*UpdateDomainReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateDomainReply)
		return ctx.Result(200, reply)
	}
}

func _DomainService_DeleteDomain0_HTTP_Handler(srv DomainServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteDomainReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDomainServiceDeleteDomain)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteDomain(ctx, req.(*DeleteDomainReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteDomainReply)
		return ctx.Result(200, reply)
	}
}

type DomainServiceHTTPClient interface {
	CreateDomain(ctx context.Context, req *CreateDomainReq, opts ...http.CallOption) (rsp *CreateDomainReply, err error)
	DeleteDomain(ctx context.Context, req *DeleteDomainReq, opts ...http.CallOption) (rsp *DeleteDomainReply, err error)
	GetDomain(ctx context.Context, req *GetDomainReq, opts ...http.CallOption) (rsp *Domain, err error)
	ListDomain(ctx context.Context, req *protobuf.PagingReq, opts ...http.CallOption) (rsp *protobuf.PagingReply, err error)
	UpdateDomain(ctx context.Context, req *UpdateDomainReq, opts ...http.CallOption) (rsp *UpdateDomainReply, err error)
}

type DomainServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewDomainServiceHTTPClient(client *http.Client) DomainServiceHTTPClient {
	return &DomainServiceHTTPClientImpl{client}
}

func (c *DomainServiceHTTPClientImpl) CreateDomain(ctx context.Context, in *CreateDomainReq, opts ...http.CallOption) (*CreateDomainReply, error) {
	var out CreateDomainReply
	pattern := "/admin/v1/domains"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDomainServiceCreateDomain))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DomainServiceHTTPClientImpl) DeleteDomain(ctx context.Context, in *DeleteDomainReq, opts ...http.CallOption) (*DeleteDomainReply, error) {
	var out DeleteDomainReply
	pattern := "/admin/v1/domains/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDomainServiceDeleteDomain))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DomainServiceHTTPClientImpl) GetDomain(ctx context.Context, in *GetDomainReq, opts ...http.CallOption) (*Domain, error) {
	var out Domain
	pattern := "/admin/v1/domains/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDomainServiceGetDomain))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DomainServiceHTTPClientImpl) ListDomain(ctx context.Context, in *protobuf.PagingReq, opts ...http.CallOption) (*protobuf.PagingReply, error) {
	var out protobuf.PagingReply
	pattern := "/admin/v1/domains"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDomainServiceListDomain))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DomainServiceHTTPClientImpl) UpdateDomain(ctx context.Context, in *UpdateDomainReq, opts ...http.CallOption) (*UpdateDomainReply, error) {
	var out UpdateDomainReply
	pattern := "/admin/v1/domains/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDomainServiceUpdateDomain))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
