package service

import (
	"context"

	pb "github.com/beiduoke/go-scaffold/api/saasdesk/service/v1"
	"github.com/beiduoke/go-scaffold/app/saasdesk/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type RoleService struct {
	pb.UnimplementedRoleServiceServer
	log *log.Helper
	ac  biz.RoleRepo
}

func NewRoleService(logger log.Logger, ac biz.RoleRepo) *RoleService {
	l := log.NewHelper(log.With(logger, "module", "role/service"))
	return &RoleService{
		log: l,
		ac:  ac,
	}
}

func (s *RoleService) CreateRole(ctx context.Context, req *pb.CreateRoleRequest) (*pb.CreateRoleResponse, error) {
	return &pb.CreateRoleResponse{}, nil
}
func (s *RoleService) UpdateRole(ctx context.Context, req *pb.UpdateRoleRequest) (*pb.UpdateRoleResponse, error) {
	return &pb.UpdateRoleResponse{}, nil
}
func (s *RoleService) DeleteRole(ctx context.Context, req *pb.DeleteRoleRequest) (*pb.DeleteRoleResponse, error) {
	return &pb.DeleteRoleResponse{}, nil
}
func (s *RoleService) GetRole(ctx context.Context, req *pb.GetRoleRequest) (*pb.GetRoleResponse, error) {
	return &pb.GetRoleResponse{}, nil
}
func (s *RoleService) ListRole(ctx context.Context, req *pb.ListRoleRequest) (*pb.ListRoleResponse, error) {
	return &pb.ListRoleResponse{}, nil
}
