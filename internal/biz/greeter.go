package biz

import (
	"context"

	v1 "github.com/bedoke/go-scaffold/api/admin/v1"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Admin is a Admin model.
type Admin struct {
	Hello string
}

// AdminRepo is a Greater repo.
type AdminRepo interface {
	Save(context.Context, *Admin) (*Admin, error)
	Update(context.Context, *Admin) (*Admin, error)
	FindByID(context.Context, int64) (*Admin, error)
	ListByHello(context.Context, string) ([]*Admin, error)
	ListAll(context.Context) ([]*Admin, error)
}

// AdminUsecase is a Admin usecase.
type AdminUsecase struct {
	repo AdminRepo
	log  *log.Helper
}

// NewAdminUsecase new a Admin usecase.
func NewAdminUsecase(repo AdminRepo, logger log.Logger) *AdminUsecase {
	return &AdminUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateAdmin creates a Admin, and returns the new Admin.
func (uc *AdminUsecase) CreateAdmin(ctx context.Context, g *Admin) (*Admin, error) {
	uc.log.WithContext(ctx).Infof("CreateAdmin: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
