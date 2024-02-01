package data

import (
	"context"
	"time"

	v1 "github.com/beiduoke/go-scaffold/api/core/service/v1"
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data/ent"
	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"github.com/go-kratos/kratos/v2/log"
)

func (r *RoleRepo) toProto(in *ent.Role) *v1.Role {
	if in == nil {
		return nil
	}
	return &v1.Role{
		Id:        in.ID,
		Name:      in.Name,
		State:     in.State,
		Remarks:   in.Remark,
		Sort:      in.Sort,
		CreatedAt: convert.TimeValueToString(in.CreatedAt, time.DateTime),
		UpdatedAt: convert.TimeValueToString(in.UpdatedAt, time.DateTime),
	}
}

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

func (s *RoleRepo) CreateRole(ctx context.Context, req *v1.CreateRoleRequest) (*v1.CreateRoleResponse, error) {
	return &v1.CreateRoleResponse{}, nil
}
func (s *RoleRepo) UpdateRole(ctx context.Context, req *v1.UpdateRoleRequest) (*v1.UpdateRoleResponse, error) {
	return &v1.UpdateRoleResponse{}, nil
}
func (s *RoleRepo) DeleteRole(ctx context.Context, req *v1.DeleteRoleRequest) (*v1.DeleteRoleResponse, error) {
	return &v1.DeleteRoleResponse{}, nil
}
func (s *RoleRepo) GetRole(ctx context.Context, req *v1.GetRoleRequest) (*v1.Role, error) {
	return &v1.Role{}, nil
}
func (s *RoleRepo) ListRole(ctx context.Context, req *v1.ListRoleRequest) (*v1.ListRoleResponse, error) {
	return &v1.ListRoleResponse{}, nil
}
