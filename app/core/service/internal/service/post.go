package service

import (
	"context"

	"github.com/beiduoke/go-scaffold/api/common/pagination"
	v1 "github.com/beiduoke/go-scaffold/api/core/service/v1"
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

type PostService struct {
	v1.UnimplementedPostServiceServer
	log *log.Helper
	ac  *data.PostRepo
}

func NewPostService(logger log.Logger, ac *data.PostRepo) *PostService {
	l := log.NewHelper(log.With(logger, "module", "post/service"))
	return &PostService{
		log: l,
		ac:  ac,
	}
}

func (s *PostService) CreatePost(ctx context.Context, req *v1.CreatePostRequest) (*v1.CreatePostResponse, error) {
	return s.ac.CreatePost(ctx, req)
}

func (s *PostService) UpdatePost(ctx context.Context, req *v1.UpdatePostRequest) (*v1.UpdatePostResponse, error) {
	return s.ac.UpdatePost(ctx, req)
}

func (s *PostService) DeletePost(ctx context.Context, req *v1.DeletePostRequest) (*v1.DeletePostResponse, error) {
	return s.ac.DeletePost(ctx, req)
}

func (s *PostService) GetPost(ctx context.Context, req *v1.GetPostRequest) (*v1.Post, error) {
	return s.ac.GetPost(ctx, req)
}

func (s *PostService) ListPost(ctx context.Context, req *pagination.PagingRequest) (*v1.ListPostResponse, error) {
	return s.ac.ListPost(ctx, req)
}
