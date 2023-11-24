package biz

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/saasdesk/service/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrRoleNotFound is user not found.
	ErrRoleNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Role is a Role model.
type Role struct {
	Hello string
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
}

// NewRoleUsecase new a Role usecase.
func NewRoleUsecase(repo RoleRepo, logger log.Logger) *RoleUsecase {
	return &RoleUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateRole creates a Role, and returns the new Role.
func (uc *RoleUsecase) CreateRole(ctx context.Context, g *Role) (*Role, error) {
	uc.log.WithContext(ctx).Infof("CreateRole: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
