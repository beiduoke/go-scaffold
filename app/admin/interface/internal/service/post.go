package service

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/admin/interface/v1"
	"github.com/beiduoke/go-scaffold/api/common/pagination"
	coreV1 "github.com/beiduoke/go-scaffold/api/core/service/v1"
	"github.com/beiduoke/go-scaffold/pkg/middleware/authn"
	"github.com/go-kratos/kratos/v2/log"
)

type PostService struct {
	v1.PostServiceHTTPServer

	uc  coreV1.PostServiceClient
	log *log.Helper
}

func NewPostService(logger log.Logger, uc coreV1.PostServiceClient) *PostService {
	l := log.NewHelper(log.With(logger, "module", "post/service/admin-service"))
	return &PostService{
		log: l,
		uc:  uc,
	}
}

func (s *PostService) ListPost(ctx context.Context, req *pagination.PagingRequest) (*coreV1.ListPostResponse, error) {
	return s.uc.ListPost(ctx, req)
}

func (s *PostService) GetPost(ctx context.Context, req *coreV1.GetPostRequest) (*coreV1.Post, error) {
	return s.uc.GetPost(ctx, req)
}

func (s *PostService) CreatePost(ctx context.Context, req *coreV1.CreatePostRequest) (*coreV1.CreatePostResponse, error) {
	// return s.uc.CreatePost(ctx, req)
	authInfo, err := authn.AuthFromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, v1.ErrorAccessForbidden("用户认证失败")
	}

	if req.Post == nil {
		return nil, v1.ErrorBadRequest("错误的参数")
	}

	req.OperatorId = authInfo.UserId
	return s.uc.CreatePost(ctx, req)
}

func (s *PostService) UpdatePost(ctx context.Context, req *coreV1.UpdatePostRequest) (*coreV1.UpdatePostResponse, error) {
	authInfo, err := authn.AuthFromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, v1.ErrorAccessForbidden("用户认证失败")
	}

	if req.Post == nil {
		return nil, v1.ErrorBadRequest("错误的参数")
	}

	req.OperatorId = authInfo.UserId

	return s.uc.UpdatePost(ctx, req)
}

func (s *PostService) DeletePost(ctx context.Context, req *coreV1.DeletePostRequest) (*coreV1.DeletePostResponse, error) {
	authInfo, err := authn.AuthFromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, v1.ErrorAccessForbidden("用户认证失败")
	}

	req.OperatorId = authInfo.UserId

	return s.uc.DeletePost(ctx, req)
}
