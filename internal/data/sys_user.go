package data

import (
	"context"

	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *UserRepo) Save(ctx context.Context, g *biz.User) (*biz.User, error) {
	return g, nil
}

func (r *UserRepo) Update(ctx context.Context, g *biz.User) (*biz.User, error) {
	return g, nil
}

func (r *UserRepo) FindByID(context.Context, int64) (*biz.User, error) {
	return nil, nil
}

func (r *UserRepo) ListAll(context.Context) ([]*biz.User, error) {
	return nil, nil
}

func (r *UserRepo) FindByName(context.Context, string) (*biz.User, error) {
	return nil, nil
}
func (r *UserRepo) FindByMobile(context.Context, string) (*biz.User, error) {
	return nil, nil
}
func (r *UserRepo) FindByEmail(context.Context, string) (*biz.User, error) {
	return nil, nil
}
func (r *UserRepo) ListByName(context.Context, string) ([]*biz.User, error) {
	return nil, nil
}
