package service

import (
	"context"

	"github.com/beiduoke/go-scaffold/api/common/pagination"
	v1 "github.com/beiduoke/go-scaffold/api/core/service/v1"
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

type RoleService struct {
	v1.UnimplementedRoleServiceServer
	log *log.Helper
	ac  *data.RoleRepo
}

func NewRoleService(logger log.Logger, ac *data.RoleRepo) *RoleService {
	l := log.NewHelper(log.With(logger, "module", "role/service"))
	return &RoleService{
		log: l,
		ac:  ac,
	}
}

func (s *RoleService) CreateRole(ctx context.Context, req *v1.CreateRoleRequest) (*v1.CreateRoleResponse, error) {
	return s.ac.CreateRole(ctx, req)
}

func (s *RoleService) UpdateRole(ctx context.Context, req *v1.UpdateRoleRequest) (*v1.UpdateRoleResponse, error) {
	return s.ac.UpdateRole(ctx, req)
}

func (s *RoleService) DeleteRole(ctx context.Context, req *v1.DeleteRoleRequest) (*v1.DeleteRoleResponse, error) {
	return s.ac.DeleteRole(ctx, req)
}

func (s *RoleService) GetRole(ctx context.Context, req *v1.GetRoleRequest) (*v1.Role, error) {
	return s.ac.GetRole(ctx, req)
}

func (s *RoleService) ListRole(ctx context.Context, req *pagination.PagingRequest) (*v1.ListRoleResponse, error) {
	return s.ac.ListRole(ctx, req)
}
