package data

import (
	"context"

	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type ApiRepo struct {
	data *Data
	log  *log.Helper
}

// NewApiRepo .
func NewApiRepo(data *Data, logger log.Logger) biz.ApiRepo {
	return &ApiRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *ApiRepo) Save(ctx context.Context, g *biz.Api) (*biz.Api, error) {
	return g, nil
}

func (r *ApiRepo) Update(ctx context.Context, g *biz.Api) (*biz.Api, error) {
	return g, nil
}

func (r *ApiRepo) FindByID(context.Context, int64) (*biz.Api, error) {
	return nil, nil
}

func (r *ApiRepo) ListByName(context.Context, string) ([]*biz.Api, error) {
	return nil, nil
}

func (r *ApiRepo) ListAll(context.Context) ([]*biz.Api, error) {
	return nil, nil
}
