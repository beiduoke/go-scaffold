package data

import (
	"context"

	pb "github.com/beiduoke/go-scaffold/api/core/service/v1"
	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) *UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (s *UserRepo) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return &pb.CreateUserResponse{}, nil
}
func (s *UserRepo) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return &pb.UpdateUserResponse{}, nil
}
func (s *UserRepo) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return &pb.DeleteUserResponse{}, nil
}
func (s *UserRepo) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return &pb.GetUserResponse{}, nil
}
func (s *UserRepo) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	return &pb.ListUserResponse{}, nil
}
