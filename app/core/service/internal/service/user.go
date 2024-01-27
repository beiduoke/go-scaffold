package service

import (
	"context"

	"github.com/beiduoke/go-scaffold/api/common/pagination"
	pb "github.com/beiduoke/go-scaffold/api/core/service/v1"
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
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

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return &pb.CreateUserResponse{}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return &pb.UpdateUserResponse{}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return &pb.DeleteUserResponse{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	return &pb.User{}, nil
}
func (s *UserService) GetUserByUserName(ctx context.Context, req *pb.GetUserByUserNameRequest) (*pb.User, error) {
	return &pb.User{}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pagination.PagingRequest) (*pb.ListUserResponse, error) {
	return &pb.ListUserResponse{}, nil
}
func (s *UserService) VerifyPassword(ctx context.Context, req *pb.VerifyPasswordRequest) (*pb.VerifyPasswordResponse, error) {
	return &pb.VerifyPasswordResponse{}, nil
}
func (s *UserService) UserExists(ctx context.Context, req *pb.UserExistsRequest) (*pb.UserExistsResponse, error) {
	return &pb.UserExistsResponse{}, nil
}
