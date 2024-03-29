// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             (unknown)
// source: admin/v1/dict.proto

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

const OperationDictServiceCreateDict = "/admin.v1.DictService/CreateDict"
const OperationDictServiceCreateDictData = "/admin.v1.DictService/CreateDictData"
const OperationDictServiceDeleteDict = "/admin.v1.DictService/DeleteDict"
const OperationDictServiceDeleteDictData = "/admin.v1.DictService/DeleteDictData"
const OperationDictServiceGetDict = "/admin.v1.DictService/GetDict"
const OperationDictServiceGetDictData = "/admin.v1.DictService/GetDictData"
const OperationDictServiceListDict = "/admin.v1.DictService/ListDict"
const OperationDictServiceListDictData = "/admin.v1.DictService/ListDictData"
const OperationDictServiceUpdateDict = "/admin.v1.DictService/UpdateDict"
const OperationDictServiceUpdateDictData = "/admin.v1.DictService/UpdateDictData"
const OperationDictServiceUpdateDictDataState = "/admin.v1.DictService/UpdateDictDataState"
const OperationDictServiceUpdateDictState = "/admin.v1.DictService/UpdateDictState"

type DictServiceHTTPServer interface {
	// CreateDict 创建字典
	CreateDict(context.Context, *CreateDictRequest) (*CreateDictResponse, error)
	// CreateDictData 创建字典数据
	CreateDictData(context.Context, *CreateDictDataRequest) (*CreateDictDataResponse, error)
	// DeleteDict 删除字典
	DeleteDict(context.Context, *DeleteDictRequest) (*DeleteDictResponse, error)
	// DeleteDictData 删除字典数据
	DeleteDictData(context.Context, *DeleteDictDataRequest) (*DeleteDictDataResponse, error)
	// GetDict 获取字典
	GetDict(context.Context, *GetDictRequest) (*Dict, error)
	// GetDictData 获取字典数据
	GetDictData(context.Context, *GetDictDataRequest) (*DictData, error)
	// ListDict 列表字典
	ListDict(context.Context, *ListDictRequest) (*ListDictResponse, error)
	// ListDictData 列表字典数据
	ListDictData(context.Context, *ListDictDataRequest) (*ListDictDataResponse, error)
	// UpdateDict 修改字典
	UpdateDict(context.Context, *UpdateDictRequest) (*UpdateDictResponse, error)
	// UpdateDictData 修改字典数据
	UpdateDictData(context.Context, *UpdateDictDataRequest) (*UpdateDictDataResponse, error)
	// UpdateDictDataState 设置字典数据状态
	UpdateDictDataState(context.Context, *UpdateDictDataStateRequest) (*UpdateDictDataStateResponse, error)
	// UpdateDictState 设置字典状态
	UpdateDictState(context.Context, *UpdateDictStateRequest) (*UpdateDictStateResponse, error)
}

func RegisterDictServiceHTTPServer(s *http.Server, srv DictServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/dicts", _DictService_ListDict0_HTTP_Handler(srv))
	r.POST("/v1/dicts", _DictService_CreateDict0_HTTP_Handler(srv))
	r.GET("/v1/dicts/{id}", _DictService_GetDict0_HTTP_Handler(srv))
	r.PUT("/v1/dicts/{id}", _DictService_UpdateDict0_HTTP_Handler(srv))
	r.DELETE("/v1/dicts/{id}", _DictService_DeleteDict0_HTTP_Handler(srv))
	r.PUT("/v1/dicts/{id}/state", _DictService_UpdateDictState0_HTTP_Handler(srv))
	r.GET("/v1/dictData", _DictService_ListDictData0_HTTP_Handler(srv))
	r.POST("/v1/dictData", _DictService_CreateDictData0_HTTP_Handler(srv))
	r.GET("/v1/dictData/{id}", _DictService_GetDictData0_HTTP_Handler(srv))
	r.PUT("/v1/dictData/{id}", _DictService_UpdateDictData0_HTTP_Handler(srv))
	r.DELETE("/v1/dictData/{id}", _DictService_DeleteDictData0_HTTP_Handler(srv))
	r.PUT("/v1/dictData/{id}/state", _DictService_UpdateDictDataState0_HTTP_Handler(srv))
}

func _DictService_ListDict0_HTTP_Handler(srv DictServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListDictRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictServiceListDict)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListDict(ctx, req.(*ListDictRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListDictResponse)
		return ctx.Result(200, reply)
	}
}

func _DictService_CreateDict0_HTTP_Handler(srv DictServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateDictRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictServiceCreateDict)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateDict(ctx, req.(*CreateDictRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateDictResponse)
		return ctx.Result(200, reply)
	}
}

func _DictService_GetDict0_HTTP_Handler(srv DictServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetDictRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictServiceGetDict)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetDict(ctx, req.(*GetDictRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Dict)
		return ctx.Result(200, reply)
	}
}

func _DictService_UpdateDict0_HTTP_Handler(srv DictServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateDictRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictServiceUpdateDict)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateDict(ctx, req.(*UpdateDictRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateDictResponse)
		return ctx.Result(200, reply)
	}
}

func _DictService_DeleteDict0_HTTP_Handler(srv DictServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteDictRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictServiceDeleteDict)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteDict(ctx, req.(*DeleteDictRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteDictResponse)
		return ctx.Result(200, reply)
	}
}

func _DictService_UpdateDictState0_HTTP_Handler(srv DictServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateDictStateRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictServiceUpdateDictState)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateDictState(ctx, req.(*UpdateDictStateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateDictStateResponse)
		return ctx.Result(200, reply)
	}
}

func _DictService_ListDictData0_HTTP_Handler(srv DictServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListDictDataRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictServiceListDictData)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListDictData(ctx, req.(*ListDictDataRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListDictDataResponse)
		return ctx.Result(200, reply)
	}
}

func _DictService_CreateDictData0_HTTP_Handler(srv DictServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateDictDataRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictServiceCreateDictData)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateDictData(ctx, req.(*CreateDictDataRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateDictDataResponse)
		return ctx.Result(200, reply)
	}
}

func _DictService_GetDictData0_HTTP_Handler(srv DictServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetDictDataRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictServiceGetDictData)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetDictData(ctx, req.(*GetDictDataRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DictData)
		return ctx.Result(200, reply)
	}
}

func _DictService_UpdateDictData0_HTTP_Handler(srv DictServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateDictDataRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictServiceUpdateDictData)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateDictData(ctx, req.(*UpdateDictDataRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateDictDataResponse)
		return ctx.Result(200, reply)
	}
}

func _DictService_DeleteDictData0_HTTP_Handler(srv DictServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteDictDataRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictServiceDeleteDictData)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteDictData(ctx, req.(*DeleteDictDataRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteDictDataResponse)
		return ctx.Result(200, reply)
	}
}

func _DictService_UpdateDictDataState0_HTTP_Handler(srv DictServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateDictDataStateRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictServiceUpdateDictDataState)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateDictDataState(ctx, req.(*UpdateDictDataStateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateDictDataStateResponse)
		return ctx.Result(200, reply)
	}
}

type DictServiceHTTPClient interface {
	CreateDict(ctx context.Context, req *CreateDictRequest, opts ...http.CallOption) (rsp *CreateDictResponse, err error)
	CreateDictData(ctx context.Context, req *CreateDictDataRequest, opts ...http.CallOption) (rsp *CreateDictDataResponse, err error)
	DeleteDict(ctx context.Context, req *DeleteDictRequest, opts ...http.CallOption) (rsp *DeleteDictResponse, err error)
	DeleteDictData(ctx context.Context, req *DeleteDictDataRequest, opts ...http.CallOption) (rsp *DeleteDictDataResponse, err error)
	GetDict(ctx context.Context, req *GetDictRequest, opts ...http.CallOption) (rsp *Dict, err error)
	GetDictData(ctx context.Context, req *GetDictDataRequest, opts ...http.CallOption) (rsp *DictData, err error)
	ListDict(ctx context.Context, req *ListDictRequest, opts ...http.CallOption) (rsp *ListDictResponse, err error)
	ListDictData(ctx context.Context, req *ListDictDataRequest, opts ...http.CallOption) (rsp *ListDictDataResponse, err error)
	UpdateDict(ctx context.Context, req *UpdateDictRequest, opts ...http.CallOption) (rsp *UpdateDictResponse, err error)
	UpdateDictData(ctx context.Context, req *UpdateDictDataRequest, opts ...http.CallOption) (rsp *UpdateDictDataResponse, err error)
	UpdateDictDataState(ctx context.Context, req *UpdateDictDataStateRequest, opts ...http.CallOption) (rsp *UpdateDictDataStateResponse, err error)
	UpdateDictState(ctx context.Context, req *UpdateDictStateRequest, opts ...http.CallOption) (rsp *UpdateDictStateResponse, err error)
}

type DictServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewDictServiceHTTPClient(client *http.Client) DictServiceHTTPClient {
	return &DictServiceHTTPClientImpl{client}
}

func (c *DictServiceHTTPClientImpl) CreateDict(ctx context.Context, in *CreateDictRequest, opts ...http.CallOption) (*CreateDictResponse, error) {
	var out CreateDictResponse
	pattern := "/v1/dicts"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDictServiceCreateDict))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DictServiceHTTPClientImpl) CreateDictData(ctx context.Context, in *CreateDictDataRequest, opts ...http.CallOption) (*CreateDictDataResponse, error) {
	var out CreateDictDataResponse
	pattern := "/v1/dictData"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDictServiceCreateDictData))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DictServiceHTTPClientImpl) DeleteDict(ctx context.Context, in *DeleteDictRequest, opts ...http.CallOption) (*DeleteDictResponse, error) {
	var out DeleteDictResponse
	pattern := "/v1/dicts/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDictServiceDeleteDict))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DictServiceHTTPClientImpl) DeleteDictData(ctx context.Context, in *DeleteDictDataRequest, opts ...http.CallOption) (*DeleteDictDataResponse, error) {
	var out DeleteDictDataResponse
	pattern := "/v1/dictData/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDictServiceDeleteDictData))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DictServiceHTTPClientImpl) GetDict(ctx context.Context, in *GetDictRequest, opts ...http.CallOption) (*Dict, error) {
	var out Dict
	pattern := "/v1/dicts/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDictServiceGetDict))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DictServiceHTTPClientImpl) GetDictData(ctx context.Context, in *GetDictDataRequest, opts ...http.CallOption) (*DictData, error) {
	var out DictData
	pattern := "/v1/dictData/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDictServiceGetDictData))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DictServiceHTTPClientImpl) ListDict(ctx context.Context, in *ListDictRequest, opts ...http.CallOption) (*ListDictResponse, error) {
	var out ListDictResponse
	pattern := "/v1/dicts"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDictServiceListDict))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DictServiceHTTPClientImpl) ListDictData(ctx context.Context, in *ListDictDataRequest, opts ...http.CallOption) (*ListDictDataResponse, error) {
	var out ListDictDataResponse
	pattern := "/v1/dictData"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDictServiceListDictData))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DictServiceHTTPClientImpl) UpdateDict(ctx context.Context, in *UpdateDictRequest, opts ...http.CallOption) (*UpdateDictResponse, error) {
	var out UpdateDictResponse
	pattern := "/v1/dicts/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDictServiceUpdateDict))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DictServiceHTTPClientImpl) UpdateDictData(ctx context.Context, in *UpdateDictDataRequest, opts ...http.CallOption) (*UpdateDictDataResponse, error) {
	var out UpdateDictDataResponse
	pattern := "/v1/dictData/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDictServiceUpdateDictData))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DictServiceHTTPClientImpl) UpdateDictDataState(ctx context.Context, in *UpdateDictDataStateRequest, opts ...http.CallOption) (*UpdateDictDataStateResponse, error) {
	var out UpdateDictDataStateResponse
	pattern := "/v1/dictData/{id}/state"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDictServiceUpdateDictDataState))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DictServiceHTTPClientImpl) UpdateDictState(ctx context.Context, in *UpdateDictStateRequest, opts ...http.CallOption) (*UpdateDictStateResponse, error) {
	var out UpdateDictStateResponse
	pattern := "/v1/dicts/{id}/state"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDictServiceUpdateDictState))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
