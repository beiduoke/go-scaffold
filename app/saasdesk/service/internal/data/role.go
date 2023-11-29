package data

import (
	"context"

	"github.com/beiduoke/go-scaffold/app/saasdesk/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type roleRepo struct {
	data *Data
	log  *log.Helper
}

// NewRoleRepo .
func NewRoleRepo(data *Data, logger log.Logger) biz.RoleRepo {
	return &roleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *roleRepo) Save(ctx context.Context, g *biz.Role) (*biz.Role, error) {
	return g, nil
}

func (r *roleRepo) Update(ctx context.Context, g *biz.Role) (*biz.Role, error) {
	return g, nil
}

func (r *roleRepo) FindByID(context.Context, int64) (*biz.Role, error) {
	return nil, nil
}

func (r *roleRepo) ListByName(context.Context, string) ([]*biz.Role, error) {
	return nil, nil
}

func (r *roleRepo) ListAll(context.Context) ([]*biz.Role, error) {
	return nil, nil
}
