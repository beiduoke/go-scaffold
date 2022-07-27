package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// Role is a Role model.
type Role struct {
	Name string
}

// RoleRepo is a Greater repo.
type RoleRepo interface {
	Save(context.Context, *Role) (*Role, error)
	Update(context.Context, *Role) (*Role, error)
	FindByID(context.Context, int64) (*Role, error)
	ListByName(context.Context, string) ([]*Role, error)
	ListAll(context.Context) ([]*Role, error)
}

// RoleUsecase is a Role usecase.
type RoleUsecase struct {
	repo RoleRepo
	log  *log.Helper
	tm   Transaction
}

// NewRoleUsecase new a Role usecase.
func NewRoleUsecase(repo RoleRepo, tm Transaction, logger log.Logger) *RoleUsecase {
	return &RoleUsecase{repo: repo, tm: tm, log: log.NewHelper(logger)}
}

// CreateRole creates a Role, and returns the new Role.
func (uc *RoleUsecase) CreateRole(ctx context.Context, g *Role) (*Role, error) {
	uc.log.WithContext(ctx).Infof("CreateRole: %v", g.Name)
	return uc.repo.Save(ctx, g)
}
