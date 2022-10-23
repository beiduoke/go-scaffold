package biz

import (
	"context"
	"time"

	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
)

// Authority is a Authority model.
type Authority struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}

// AuthorityRepo is a Greater repo.
type AuthorityRepo interface {
	Save(context.Context, *Authority) (*Authority, error)
	Update(context.Context, *Authority) (*Authority, error)
	FindByID(context.Context, uint) (*Authority, error)
	ListByName(context.Context, string) ([]*Authority, error)
	ListAll(context.Context) ([]*Authority, error)
	ListPage(context.Context, pagination.PaginationHandler) ([]*Authority, int64)
}

// AuthorityUsecase is a Authority usecase.
type AuthorityUsecase struct {
	*Biz
	repo AuthorityRepo
}

// NewAuthorityUsecase new a Authority usecase.
func NewAuthorityUsecase(biz *Biz, repo AuthorityRepo) *AuthorityUsecase {
	return &AuthorityUsecase{repo: repo, Biz: biz}
}

// Create creates a Authority, and returns the new Authority.
func (uc *AuthorityUsecase) Create(ctx context.Context, g *Authority) (*Authority, error) {
	uc.log.WithContext(ctx).Infof("Create: %v", g.Name)
	return uc.repo.Save(ctx, g)
}
