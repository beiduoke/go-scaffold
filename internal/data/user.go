package data

import (
	"context"

	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
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

func (r *UserRepo) toModel(g *biz.User) *SysUser {
	if g == nil {
		return nil
	}
	return &SysUser{
		Model: gorm.Model{
			ID:        g.ID,
			CreatedAt: g.CreatedAt,
			UpdatedAt: g.UpdatedAt,
		},
		Name:     g.Name,
		NickName: g.NickName,
		RealName: g.RealName,
		Password: g.Password,
		Birthday: g.Birthday,
		Gender:   g.Gender,
		Mobile:   g.Mobile,
		Email:    g.Email,
		State:    g.State,
	}
}

func (r *UserRepo) toBiz(u *SysUser) *biz.User {
	if u == nil {
		return nil
	}
	return &biz.User{
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		ID:        u.ID,
		Name:      u.Name,
		NickName:  u.NickName,
		RealName:  u.RealName,
		Password:  u.Password,
		Birthday:  u.Birthday,
		Gender:    u.Gender,
		Mobile:    u.Mobile,
		Email:     u.Email,
		State:     u.State,
	}
}

func (r *UserRepo) Save(ctx context.Context, g *biz.User) (*biz.User, error) {
	user := r.toModel(g)
	result := r.data.DB(ctx).Create(user)
	return r.toBiz(user), result.Error
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

func (r *UserRepo) FindByName(ctx context.Context, s string) (*biz.User, error) {
	user := SysUser{}
	result := r.data.DB(ctx).First(&user, "name = ?", s)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&user), nil
}
func (r *UserRepo) FindByMobile(ctx context.Context, s string) (*biz.User, error) {
	return nil, nil
}
func (r *UserRepo) FindByEmail(ctx context.Context, s string) (*biz.User, error) {
	return nil, nil
}
func (r *UserRepo) ListByName(ctx context.Context, s string) ([]*biz.User, error) {
	return nil, nil
}
