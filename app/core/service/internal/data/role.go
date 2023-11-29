package data

import (
	"context"

	pb "github.com/beiduoke/go-scaffold/api/core/service/v1"
	"github.com/go-kratos/kratos/v2/log"
)

type RoleRepo struct {
	data *Data
	log  *log.Helper
}

// NewRoleRepo .
func NewRoleRepo(data *Data, logger log.Logger) *RoleRepo {
	return &RoleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (s *RoleRepo) CreateRole(ctx context.Context, req *pb.CreateRoleRequest) (*pb.CreateRoleResponse, error) {
	return &pb.CreateRoleResponse{}, nil
}
func (s *RoleRepo) UpdateRole(ctx context.Context, req *pb.UpdateRoleRequest) (*pb.UpdateRoleResponse, error) {
	return &pb.UpdateRoleResponse{}, nil
}
func (s *RoleRepo) DeleteRole(ctx context.Context, req *pb.DeleteRoleRequest) (*pb.DeleteRoleResponse, error) {
	return &pb.DeleteRoleResponse{}, nil
}
func (s *RoleRepo) GetRole(ctx context.Context, req *pb.GetRoleRequest) (*pb.GetRoleResponse, error) {
	return &pb.GetRoleResponse{}, nil
}
func (s *RoleRepo) ListRole(ctx context.Context, req *pb.ListRoleRequest) (*pb.ListRoleResponse, error) {
	return &pb.ListRoleResponse{}, nil
}
