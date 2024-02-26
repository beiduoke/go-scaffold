package service

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/admin/interface/v1"
	"github.com/beiduoke/go-scaffold/api/common/pagination"
	coreV1 "github.com/beiduoke/go-scaffold/api/core/service/v1"
	"github.com/beiduoke/go-scaffold/pkg/middleware/authn"
	"github.com/go-kratos/kratos/v2/log"
)

type DeptService struct {
	v1.DeptServiceHTTPServer

	uc  coreV1.DeptServiceClient
	log *log.Helper
}

func NewDeptService(logger log.Logger, uc coreV1.DeptServiceClient) *DeptService {
	l := log.NewHelper(log.With(logger, "module", "dept/service/admin-service"))
	return &DeptService{
		log: l,
		uc:  uc,
	}
}

func (s *DeptService) ListDept(ctx context.Context, req *pagination.PagingRequest) (*coreV1.ListDeptResponse, error) {
	return s.uc.ListDept(ctx, req)
}

func (s *DeptService) GetDept(ctx context.Context, req *coreV1.GetDeptRequest) (*coreV1.Dept, error) {
	return s.uc.GetDept(ctx, req)
}

func (s *DeptService) CreateDept(ctx context.Context, req *coreV1.CreateDeptRequest) (*coreV1.CreateDeptResponse, error) {
	// return s.uc.CreateDept(ctx, req)
	authInfo, err := authn.AuthFromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, v1.ErrorAccessForbidden("用户认证失败")
	}

	if req.Dept == nil {
		return nil, v1.ErrorBadRequest("错误的参数")
	}

	req.OperatorId = authInfo.UserId
	return s.uc.CreateDept(ctx, req)
}

func (s *DeptService) UpdateDept(ctx context.Context, req *coreV1.UpdateDeptRequest) (*coreV1.UpdateDeptResponse, error) {
	authInfo, err := authn.AuthFromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, v1.ErrorAccessForbidden("用户认证失败")
	}

	if req.Dept == nil {
		return nil, v1.ErrorBadRequest("错误的参数")
	}

	req.OperatorId = authInfo.UserId

	return s.uc.UpdateDept(ctx, req)
}

func (s *DeptService) DeleteDept(ctx context.Context, req *coreV1.DeleteDeptRequest) (*coreV1.DeleteDeptResponse, error) {
	authInfo, err := authn.AuthFromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, v1.ErrorAccessForbidden("用户认证失败")
	}

	req.OperatorId = authInfo.UserId

	return s.uc.DeleteDept(ctx, req)
}
