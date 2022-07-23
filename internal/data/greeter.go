package data

import (
	"context"

	"github.com/bedoke/go-scaffold/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type AdminRepo struct {
	data *Data
	log  *log.Helper
}

// NewAdminRepo .
func NewAdminRepo(data *Data, logger log.Logger) biz.AdminRepo {
	return &AdminRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *AdminRepo) Save(ctx context.Context, g *biz.Admin) (*biz.Admin, error) {
	return g, nil
}

func (r *AdminRepo) Update(ctx context.Context, g *biz.Admin) (*biz.Admin, error) {
	return g, nil
}

func (r *AdminRepo) FindByID(context.Context, int64) (*biz.Admin, error) {
	return nil, nil
}

func (r *AdminRepo) ListByHello(context.Context, string) ([]*biz.Admin, error) {
	return nil, nil
}

func (r *AdminRepo) ListAll(context.Context) ([]*biz.Admin, error) {
	return nil, nil
}
