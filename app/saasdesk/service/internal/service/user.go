package service

import (
	"context"

	pb "github.com/beiduoke/go-scaffold/api/saasdesk/service/v1"
	"github.com/beiduoke/go-scaffold/app/saasdesk/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	uc  biz.UserRepo
	log *log.Helper
}

func NewUserService(logger log.Logger, uc biz.UserRepo) *UserService {
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
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return &pb.GetUserResponse{}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	return &pb.ListUserResponse{}, nil
}
