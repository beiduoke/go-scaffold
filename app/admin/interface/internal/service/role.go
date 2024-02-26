package service

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/admin/interface/v1"
	"github.com/beiduoke/go-scaffold/api/common/pagination"
	coreV1 "github.com/beiduoke/go-scaffold/api/core/service/v1"
	"github.com/beiduoke/go-scaffold/pkg/middleware/authn"
	"github.com/go-kratos/kratos/v2/log"
)

type RoleService struct {
	v1.RoleServiceHTTPServer

	uc  coreV1.RoleServiceClient
	log *log.Helper
}

func NewRoleService(logger log.Logger, uc coreV1.RoleServiceClient) *RoleService {
	l := log.NewHelper(log.With(logger, "module", "role/service/admin-service"))
	return &RoleService{
		log: l,
		uc:  uc,
	}
}

func (s *RoleService) ListRole(ctx context.Context, req *pagination.PagingRequest) (*coreV1.ListRoleResponse, error) {
	return s.uc.ListRole(ctx, req)
}

func (s *RoleService) GetRole(ctx context.Context, req *coreV1.GetRoleRequest) (*coreV1.Role, error) {
	return s.uc.GetRole(ctx, req)
}

func (s *RoleService) CreateRole(ctx context.Context, req *coreV1.CreateRoleRequest) (*coreV1.CreateRoleResponse, error) {
	// return s.uc.CreateRole(ctx, req)
	authInfo, err := authn.AuthFromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, v1.ErrorAccessForbidden("用户认证失败")
	}

	if req.Role == nil {
		return nil, v1.ErrorBadRequest("错误的参数")
	}

	req.OperatorId = authInfo.UserId
	return s.uc.CreateRole(ctx, req)
}

func (s *RoleService) UpdateRole(ctx context.Context, req *coreV1.UpdateRoleRequest) (*coreV1.UpdateRoleResponse, error) {
	authInfo, err := authn.AuthFromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, v1.ErrorAccessForbidden("用户认证失败")
	}

	if req.Role == nil {
		return nil, v1.ErrorBadRequest("错误的参数")
	}

	req.OperatorId = authInfo.UserId

	return s.uc.UpdateRole(ctx, req)
}

func (s *RoleService) DeleteRole(ctx context.Context, req *coreV1.DeleteRoleRequest) (*coreV1.DeleteRoleResponse, error) {
	authInfo, err := authn.AuthFromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, v1.ErrorAccessForbidden("用户认证失败")
	}

	req.OperatorId = authInfo.UserId

	return s.uc.DeleteRole(ctx, req)
}
