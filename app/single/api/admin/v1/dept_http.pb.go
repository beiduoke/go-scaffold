// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             (unknown)
// source: admin/v1/dept.proto

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

const OperationDeptServiceCreateDept = "/admin.v1.DeptService/CreateDept"
const OperationDeptServiceDeleteDept = "/admin.v1.DeptService/DeleteDept"
const OperationDeptServiceGetDept = "/admin.v1.DeptService/GetDept"
const OperationDeptServiceListDept = "/admin.v1.DeptService/ListDept"
const OperationDeptServiceListDeptTree = "/admin.v1.DeptService/ListDeptTree"
const OperationDeptServiceUpdateDept = "/admin.v1.DeptService/UpdateDept"
const OperationDeptServiceUpdateDeptState = "/admin.v1.DeptService/UpdateDeptState"

type DeptServiceHTTPServer interface {
	// CreateDept 创建部门
	CreateDept(context.Context, *CreateDeptRequest) (*CreateDeptResponse, error)
	// DeleteDept 删除部门
	DeleteDept(context.Context, *DeleteDeptRequest) (*DeleteDeptResponse, error)
	// GetDept 获取部门
	GetDept(context.Context, *GetDeptRequest) (*Dept, error)
	// ListDept 列表部门
	ListDept(context.Context, *ListDeptRequest) (*ListDeptResponse, error)
	// ListDeptTree 获取全部部门树形
	ListDeptTree(context.Context, *ListDeptTreeRequest) (*ListDeptTreeResponse, error)
	// UpdateDept 修改部门
	UpdateDept(context.Context, *UpdateDeptRequest) (*UpdateDeptResponse, error)
	// UpdateDeptState 更新指定ID角色状态
	UpdateDeptState(context.Context, *UpdateDeptStateRequest) (*UpdateDeptStateResponse, error)
}

func RegisterDeptServiceHTTPServer(s *http.Server, srv DeptServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/depts", _DeptService_ListDept0_HTTP_Handler(srv))
	r.POST("/v1/depts", _DeptService_CreateDept0_HTTP_Handler(srv))
	r.GET("/v1/depts/{id}", _DeptService_GetDept0_HTTP_Handler(srv))
	r.PUT("/v1/depts/{id}", _DeptService_UpdateDept0_HTTP_Handler(srv))
	r.PUT("/v1/depts/{id}/state", _DeptService_UpdateDeptState0_HTTP_Handler(srv))
	r.DELETE("/v1/depts/{id}", _DeptService_DeleteDept0_HTTP_Handler(srv))
	r.GET("/v1/depts/{id}/trees", _DeptService_ListDeptTree0_HTTP_Handler(srv))
}

func _DeptService_ListDept0_HTTP_Handler(srv DeptServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListDeptRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDeptServiceListDept)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListDept(ctx, req.(*ListDeptRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListDeptResponse)
		return ctx.Result(200, reply)
	}
}

func _DeptService_CreateDept0_HTTP_Handler(srv DeptServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateDeptRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDeptServiceCreateDept)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateDept(ctx, req.(*CreateDeptRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateDeptResponse)
		return ctx.Result(200, reply)
	}
}

func _DeptService_GetDept0_HTTP_Handler(srv DeptServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetDeptRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDeptServiceGetDept)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetDept(ctx, req.(*GetDeptRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Dept)
		return ctx.Result(200, reply)
	}
}

func _DeptService_UpdateDept0_HTTP_Handler(srv DeptServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateDeptRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDeptServiceUpdateDept)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateDept(ctx, req.(*UpdateDeptRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateDeptResponse)
		return ctx.Result(200, reply)
	}
}

func _DeptService_UpdateDeptState0_HTTP_Handler(srv DeptServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateDeptStateRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDeptServiceUpdateDeptState)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateDeptState(ctx, req.(*UpdateDeptStateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateDeptStateResponse)
		return ctx.Result(200, reply)
	}
}

func _DeptService_DeleteDept0_HTTP_Handler(srv DeptServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteDeptRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDeptServiceDeleteDept)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteDept(ctx, req.(*DeleteDeptRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteDeptResponse)
		return ctx.Result(200, reply)
	}
}

func _DeptService_ListDeptTree0_HTTP_Handler(srv DeptServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListDeptTreeRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDeptServiceListDeptTree)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListDeptTree(ctx, req.(*ListDeptTreeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListDeptTreeResponse)
		return ctx.Result(200, reply)
	}
}

type DeptServiceHTTPClient interface {
	CreateDept(ctx context.Context, req *CreateDeptRequest, opts ...http.CallOption) (rsp *CreateDeptResponse, err error)
	DeleteDept(ctx context.Context, req *DeleteDeptRequest, opts ...http.CallOption) (rsp *DeleteDeptResponse, err error)
	GetDept(ctx context.Context, req *GetDeptRequest, opts ...http.CallOption) (rsp *Dept, err error)
	ListDept(ctx context.Context, req *ListDeptRequest, opts ...http.CallOption) (rsp *ListDeptResponse, err error)
	ListDeptTree(ctx context.Context, req *ListDeptTreeRequest, opts ...http.CallOption) (rsp *ListDeptTreeResponse, err error)
	UpdateDept(ctx context.Context, req *UpdateDeptRequest, opts ...http.CallOption) (rsp *UpdateDeptResponse, err error)
	UpdateDeptState(ctx context.Context, req *UpdateDeptStateRequest, opts ...http.CallOption) (rsp *UpdateDeptStateResponse, err error)
}

type DeptServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewDeptServiceHTTPClient(client *http.Client) DeptServiceHTTPClient {
	return &DeptServiceHTTPClientImpl{client}
}

func (c *DeptServiceHTTPClientImpl) CreateDept(ctx context.Context, in *CreateDeptRequest, opts ...http.CallOption) (*CreateDeptResponse, error) {
	var out CreateDeptResponse
	pattern := "/v1/depts"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDeptServiceCreateDept))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DeptServiceHTTPClientImpl) DeleteDept(ctx context.Context, in *DeleteDeptRequest, opts ...http.CallOption) (*DeleteDeptResponse, error) {
	var out DeleteDeptResponse
	pattern := "/v1/depts/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDeptServiceDeleteDept))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DeptServiceHTTPClientImpl) GetDept(ctx context.Context, in *GetDeptRequest, opts ...http.CallOption) (*Dept, error) {
	var out Dept
	pattern := "/v1/depts/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDeptServiceGetDept))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DeptServiceHTTPClientImpl) ListDept(ctx context.Context, in *ListDeptRequest, opts ...http.CallOption) (*ListDeptResponse, error) {
	var out ListDeptResponse
	pattern := "/v1/depts"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDeptServiceListDept))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DeptServiceHTTPClientImpl) ListDeptTree(ctx context.Context, in *ListDeptTreeRequest, opts ...http.CallOption) (*ListDeptTreeResponse, error) {
	var out ListDeptTreeResponse
	pattern := "/v1/depts/{id}/trees"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDeptServiceListDeptTree))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DeptServiceHTTPClientImpl) UpdateDept(ctx context.Context, in *UpdateDeptRequest, opts ...http.CallOption) (*UpdateDeptResponse, error) {
	var out UpdateDeptResponse
	pattern := "/v1/depts/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDeptServiceUpdateDept))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DeptServiceHTTPClientImpl) UpdateDeptState(ctx context.Context, in *UpdateDeptStateRequest, opts ...http.CallOption) (*UpdateDeptStateResponse, error) {
	var out UpdateDeptStateResponse
	pattern := "/v1/depts/{id}/state"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDeptServiceUpdateDeptState))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
