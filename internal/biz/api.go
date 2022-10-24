package biz

import (
	"context"
	"time"

	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
)

// Api is a Api model.
type Api struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Path        string
	Method      string
	Group       string
	Description string
}

// ApiRepo is a Greater repo.
type ApiRepo interface {
	Save(context.Context, *Api) (*Api, error)
	Update(context.Context, *Api) (*Api, error)
	FindByID(context.Context, uint) (*Api, error)
	ListByName(context.Context, string) ([]*Api, error)
	ListAll(context.Context) ([]*Api, error)
	ListPage(context.Context, pagination.PaginationHandler) ([]*Api, int64)
}

// ApiUsecase is a Api usecase.
type ApiUsecase struct {
	biz  *Biz
	log  *log.Helper
	repo ApiRepo
}

// NewApiUsecase new a Api usecase.
func NewApiUsecase(logger log.Logger, biz *Biz, repo ApiRepo) *ApiUsecase {
	return &ApiUsecase{log: log.NewHelper(logger), repo: repo, biz: biz}
}

// Create creates a Api, and returns the new Api.
func (uc *ApiUsecase) Create(ctx context.Context, g *Api) (*Api, error) {
	uc.log.WithContext(ctx).Infof("Create: %v", g.Name)
	return uc.repo.Save(ctx, g)
}

// GetID 获取用户ID
func (uc *ApiUsecase) GetID(ctx context.Context, g *Api) (*Api, error) {
	uc.log.WithContext(ctx).Infof("GetID: %v", g)
	return uc.repo.FindByID(ctx, g.ID)
}
