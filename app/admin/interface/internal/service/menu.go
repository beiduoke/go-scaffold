package service

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/admin/interface/v1"
	"github.com/beiduoke/go-scaffold/api/common/pagination"
	coreV1 "github.com/beiduoke/go-scaffold/api/core/service/v1"
	"github.com/beiduoke/go-scaffold/pkg/middleware/authn"
	"github.com/go-kratos/kratos/v2/log"
)

type MenuService struct {
	v1.MenuServiceHTTPServer

	uc  coreV1.MenuServiceClient
	log *log.Helper
}

func NewMenuService(logger log.Logger, uc coreV1.MenuServiceClient) *MenuService {
	l := log.NewHelper(log.With(logger, "module", "menu/service/admin-service"))
	return &MenuService{
		log: l,
		uc:  uc,
	}
}

func (s *MenuService) ListMenu(ctx context.Context, req *pagination.PagingRequest) (*coreV1.ListMenuResponse, error) {
	return s.uc.ListMenu(ctx, req)
}

func (s *MenuService) GetMenu(ctx context.Context, req *coreV1.GetMenuRequest) (*coreV1.Menu, error) {
	return s.uc.GetMenu(ctx, req)
}

func (s *MenuService) CreateMenu(ctx context.Context, req *coreV1.CreateMenuRequest) (*coreV1.CreateMenuResponse, error) {
	// return s.uc.CreateMenu(ctx, req)
	authInfo, err := authn.AuthFromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, v1.ErrorAccessForbidden("用户认证失败")
	}

	if req.Menu == nil {
		return nil, v1.ErrorBadRequest("错误的参数")
	}

	req.OperatorId = authInfo.UserId
	return s.uc.CreateMenu(ctx, req)
}

func (s *MenuService) UpdateMenu(ctx context.Context, req *coreV1.UpdateMenuRequest) (*coreV1.UpdateMenuResponse, error) {
	authInfo, err := authn.AuthFromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, v1.ErrorAccessForbidden("用户认证失败")
	}

	if req.Menu == nil {
		return nil, v1.ErrorBadRequest("错误的参数")
	}

	req.OperatorId = authInfo.UserId

	return s.uc.UpdateMenu(ctx, req)
}

func (s *MenuService) DeleteMenu(ctx context.Context, req *coreV1.DeleteMenuRequest) (*coreV1.DeleteMenuResponse, error) {
	authInfo, err := authn.AuthFromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, v1.ErrorAccessForbidden("用户认证失败")
	}

	req.OperatorId = authInfo.UserId

	return s.uc.DeleteMenu(ctx, req)
}
