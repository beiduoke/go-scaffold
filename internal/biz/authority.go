package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// Authority is a Authority model.
type Authority struct {
	Name string
}

// AuthorityRepo is a Greater repo.
type AuthorityRepo interface {
	Save(context.Context, *Authority) (*Authority, error)
	Update(context.Context, *Authority) (*Authority, error)
	FindByID(context.Context, int64) (*Authority, error)
	ListByName(context.Context, string) ([]*Authority, error)
	ListAll(context.Context) ([]*Authority, error)
}

// AuthorityUsecase is a Authority usecase.
type AuthorityUsecase struct {
	repo AuthorityRepo
	log  *log.Helper
	tm   Transaction
}

// NewAuthorityUsecase new a Authority usecase.
func NewAuthorityUsecase(repo AuthorityRepo, tm Transaction, logger log.Logger) *AuthorityUsecase {
	return &AuthorityUsecase{repo: repo, tm: tm, log: log.NewHelper(logger)}
}

// CreateAuthority creates a Authority, and returns the new Authority.
func (uc *AuthorityUsecase) CreateAuthority(ctx context.Context, g *Authority) (*Authority, error) {
	uc.log.WithContext(ctx).Infof("CreateAuthority: %v", g.Name)
	return uc.repo.Save(ctx, g)
}
