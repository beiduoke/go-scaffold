// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             (unknown)
// source: admin/interface/v1/i_post.proto

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

const OperationPostServiceCreatePost = "/admin.interface.v1.PostService/CreatePost"
const OperationPostServiceDeletePost = "/admin.interface.v1.PostService/DeletePost"
const OperationPostServiceGetPost = "/admin.interface.v1.PostService/GetPost"
const OperationPostServiceListPost = "/admin.interface.v1.PostService/ListPost"
const OperationPostServiceUpdatePost = "/admin.interface.v1.PostService/UpdatePost"

type PostServiceHTTPServer interface {
	// CreatePost 创建岗位
	CreatePost(context.Context, *v1.CreatePostRequest) (*v1.CreatePostResponse, error)
	// DeletePost 删除岗位
	DeletePost(context.Context, *v1.DeletePostRequest) (*v1.DeletePostResponse, error)
	// GetPost 获取岗位数据
	GetPost(context.Context, *v1.GetPostRequest) (*v1.Post, error)
	// ListPost 获取岗位列表
	ListPost(context.Context, *pagination.PagingRequest) (*v1.ListPostResponse, error)
	// UpdatePost 更新岗位
	UpdatePost(context.Context, *v1.UpdatePostRequest) (*v1.UpdatePostResponse, error)
}

func RegisterPostServiceHTTPServer(s *http.Server, srv PostServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/admin/v1/posts", _PostService_ListPost0_HTTP_Handler(srv))
	r.GET("/admin/v1/posts/{id}", _PostService_GetPost0_HTTP_Handler(srv))
	r.POST("/admin/v1/posts", _PostService_CreatePost0_HTTP_Handler(srv))
	r.PUT("/admin/v1/posts/{id}", _PostService_UpdatePost0_HTTP_Handler(srv))
	r.DELETE("/admin/v1/posts/{id}", _PostService_DeletePost0_HTTP_Handler(srv))
}

func _PostService_ListPost0_HTTP_Handler(srv PostServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in pagination.PagingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPostServiceListPost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListPost(ctx, req.(*pagination.PagingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.ListPostResponse)
		return ctx.Result(200, reply)
	}
}

func _PostService_GetPost0_HTTP_Handler(srv PostServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.GetPostRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPostServiceGetPost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPost(ctx, req.(*v1.GetPostRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.Post)
		return ctx.Result(200, reply)
	}
}

func _PostService_CreatePost0_HTTP_Handler(srv PostServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.CreatePostRequest
		if err := ctx.Bind(&in.User); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPostServiceCreatePost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreatePost(ctx, req.(*v1.CreatePostRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.CreatePostResponse)
		return ctx.Result(200, reply)
	}
}

func _PostService_UpdatePost0_HTTP_Handler(srv PostServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.UpdatePostRequest
		if err := ctx.Bind(&in.User); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPostServiceUpdatePost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdatePost(ctx, req.(*v1.UpdatePostRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.UpdatePostResponse)
		return ctx.Result(200, reply)
	}
}

func _PostService_DeletePost0_HTTP_Handler(srv PostServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.DeletePostRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPostServiceDeletePost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeletePost(ctx, req.(*v1.DeletePostRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.DeletePostResponse)
		return ctx.Result(200, reply)
	}
}

type PostServiceHTTPClient interface {
	CreatePost(ctx context.Context, req *v1.CreatePostRequest, opts ...http.CallOption) (rsp *v1.CreatePostResponse, err error)
	DeletePost(ctx context.Context, req *v1.DeletePostRequest, opts ...http.CallOption) (rsp *v1.DeletePostResponse, err error)
	GetPost(ctx context.Context, req *v1.GetPostRequest, opts ...http.CallOption) (rsp *v1.Post, err error)
	ListPost(ctx context.Context, req *pagination.PagingRequest, opts ...http.CallOption) (rsp *v1.ListPostResponse, err error)
	UpdatePost(ctx context.Context, req *v1.UpdatePostRequest, opts ...http.CallOption) (rsp *v1.UpdatePostResponse, err error)
}

type PostServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewPostServiceHTTPClient(client *http.Client) PostServiceHTTPClient {
	return &PostServiceHTTPClientImpl{client}
}

func (c *PostServiceHTTPClientImpl) CreatePost(ctx context.Context, in *v1.CreatePostRequest, opts ...http.CallOption) (*v1.CreatePostResponse, error) {
	var out v1.CreatePostResponse
	pattern := "/admin/v1/posts"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPostServiceCreatePost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.User, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PostServiceHTTPClientImpl) DeletePost(ctx context.Context, in *v1.DeletePostRequest, opts ...http.CallOption) (*v1.DeletePostResponse, error) {
	var out v1.DeletePostResponse
	pattern := "/admin/v1/posts/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPostServiceDeletePost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PostServiceHTTPClientImpl) GetPost(ctx context.Context, in *v1.GetPostRequest, opts ...http.CallOption) (*v1.Post, error) {
	var out v1.Post
	pattern := "/admin/v1/posts/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPostServiceGetPost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PostServiceHTTPClientImpl) ListPost(ctx context.Context, in *pagination.PagingRequest, opts ...http.CallOption) (*v1.ListPostResponse, error) {
	var out v1.ListPostResponse
	pattern := "/admin/v1/posts"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPostServiceListPost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PostServiceHTTPClientImpl) UpdatePost(ctx context.Context, in *v1.UpdatePostRequest, opts ...http.CallOption) (*v1.UpdatePostResponse, error) {
	var out v1.UpdatePostResponse
	pattern := "/admin/v1/posts/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPostServiceUpdatePost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.User, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
