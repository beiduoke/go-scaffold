package biz

import (
	"context"
	"time"

	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
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
	biz  *Biz
	log  *log.Helper
	repo AuthorityRepo
}

// NewAuthorityUsecase new a Authority usecase.
func NewAuthorityUsecase(logger log.Logger, biz *Biz, repo AuthorityRepo) *AuthorityUsecase {
	return &AuthorityUsecase{log: log.NewHelper(logger), repo: repo, biz: biz}
}

// Create creates a Authority, and returns the new Authority.
func (uc *AuthorityUsecase) Create(ctx context.Context, g *Authority) (*Authority, error) {
	uc.log.WithContext(ctx).Infof("Create: %v", g.Name)
	return uc.repo.Save(ctx, g)
}

// GetDomainInID 获取指定领域ID集合
func (uc *AuthorityUsecase) ListByIDs(ctx context.Context, id ...uint) (authorities []*Authority, err error) {
	authorities, _ = uc.biz.authorityRepo.ListPage(ctx, pagination.NewPagination(pagination.WithNopaging(), pagination.WithCondition("id in ?", id)))
	return
}
