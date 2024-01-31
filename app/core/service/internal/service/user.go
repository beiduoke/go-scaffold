package service

import (
	"context"

	"github.com/beiduoke/go-scaffold/api/common/pagination"
	v1 "github.com/beiduoke/go-scaffold/api/core/service/v1"
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
	v1.UnimplementedUserServiceServer
	uc  *data.UserRepo
	log *log.Helper
}

func NewUserService(logger log.Logger, uc *data.UserRepo) *UserService {
	l := log.NewHelper(log.With(logger, "module", "user/service"))
	return &UserService{
		log: l,
		uc:  uc,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	return s.uc.CreateUser(ctx, req)
}
func (s *UserService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UpdateUserResponse, error) {
	return s.uc.UpdateUser(ctx, req)
}
func (s *UserService) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*v1.DeleteUserResponse, error) {
	return s.uc.DeleteUser(ctx, req)
}
func (s *UserService) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.User, error) {
	return s.uc.GetUser(ctx, req)
}
func (s *UserService) GetUserByUserName(ctx context.Context, req *v1.GetUserByUserNameRequest) (*v1.User, error) {
	return s.uc.GetUserByUserName(ctx, req)
}
func (s *UserService) ListUser(ctx context.Context, req *pagination.PagingRequest) (*v1.ListUserResponse, error) {
	return s.uc.ListUser(ctx, req)
}
func (s *UserService) VerifyPassword(ctx context.Context, req *v1.VerifyPasswordRequest) (*v1.VerifyPasswordResponse, error) {
	return s.uc.VerifyPassword(ctx, req)
}
func (s *UserService) UserExists(ctx context.Context, req *v1.UserExistsRequest) (*v1.UserExistsResponse, error) {
	return s.uc.UserExists(ctx, req)
}
