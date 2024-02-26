package service

import (
	"context"

	"github.com/beiduoke/go-scaffold/api/common/pagination"
	v1 "github.com/beiduoke/go-scaffold/api/core/service/v1"
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

type MenuService struct {
	v1.UnimplementedMenuServiceServer
	log *log.Helper
	ac  *data.MenuRepo
}

func NewMenuService(logger log.Logger, ac *data.MenuRepo) *MenuService {
	l := log.NewHelper(log.With(logger, "module", "menu/service"))
	return &MenuService{
		log: l,
		ac:  ac,
	}
}

func (s *MenuService) CreateMenu(ctx context.Context, req *v1.CreateMenuRequest) (*v1.CreateMenuResponse, error) {
	return s.ac.CreateMenu(ctx, req)
}

func (s *MenuService) UpdateMenu(ctx context.Context, req *v1.UpdateMenuRequest) (*v1.UpdateMenuResponse, error) {
	return s.ac.UpdateMenu(ctx, req)
}

func (s *MenuService) DeleteMenu(ctx context.Context, req *v1.DeleteMenuRequest) (*v1.DeleteMenuResponse, error) {
	return s.ac.DeleteMenu(ctx, req)
}

func (s *MenuService) GetMenu(ctx context.Context, req *v1.GetMenuRequest) (*v1.Menu, error) {
	return s.ac.GetMenu(ctx, req)
}

func (s *MenuService) ListMenu(ctx context.Context, req *pagination.PagingRequest) (*v1.ListMenuResponse, error) {
	return s.ac.ListMenu(ctx, req)
}
