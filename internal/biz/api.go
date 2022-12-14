package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// Api is a Api model.
type Api struct {
	Name string
}

// ApiRepo is a Greater repo.
type ApiRepo interface {
	Save(context.Context, *Api) (*Api, error)
	Update(context.Context, *Api) (*Api, error)
	FindByID(context.Context, int64) (*Api, error)
	ListByName(context.Context, string) ([]*Api, error)
	ListAll(context.Context) ([]*Api, error)
}

// ApiUsecase is a Api usecase.
type ApiUsecase struct {
	repo ApiRepo
	log  *log.Helper
	tm   Transaction
}

// NewApiUsecase new a Api usecase.
func NewApiUsecase(repo ApiRepo, tm Transaction, logger log.Logger) *ApiUsecase {
	return &ApiUsecase{repo: repo, tm: tm, log: log.NewHelper(logger)}
}

// CreateApi creates a Api, and returns the new Api.
func (uc *ApiUsecase) CreateApi(ctx context.Context, g *Api) (*Api, error) {
	uc.log.WithContext(ctx).Infof("CreateApi: %v", g.Name)
	return uc.repo.Save(ctx, g)
}
