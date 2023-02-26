package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/zzsds/go-tools/pkg/password"

	"github.com/beiduoke/go-scaffold/internal/biz/auth"
	"github.com/beiduoke/go-scaffold/internal/conf"
)

var _ auth.AuthUserRepo = (*AuthUsecase)(nil)

// AuthUsecase is a User usecase.
type AuthUsecase struct {
	biz *Biz
	ac  *conf.Auth
	log *log.Helper
}

// NewAuthUsecase new a User usecase.
func NewAuthUsecase(logger log.Logger, biz *Biz, ac *conf.Auth) *AuthUsecase {
	return &AuthUsecase{log: log.NewHelper(logger), ac: ac, biz: biz}
}

func (ac *AuthUsecase) ExistUserByName(ctx context.Context, name string) bool {
	u, err := ac.biz.userRepo.FindByName(ctx, name)
	if err == nil && u != nil {
		return true
	}
	return false
}

func (ac *AuthUsecase) ExistUserByPhone(ctx context.Context, phone string) bool {
	u, err := ac.biz.userRepo.FindByPhone(ctx, phone)
	if err == nil && u != nil {
		return true
	}
	return false
}

func (ac *AuthUsecase) FindUserByName(ctx context.Context, name string) (user auth.AuthUser, err error) {
	u, err := ac.biz.userRepo.FindByName(ctx, name)
	if err == nil {
		return user, err
	}
	return auth.AuthUser{
		ID:       u.ID,
		Name:     u.Name,
		Avatar:   u.Avatar,
		NickName: u.NickName,
		RealName: u.RealName,
		Password: u.Password,
		Birthday: u.Birthday,
		Gender:   u.Gender,
		Phone:    u.Phone,
		Email:    u.Email,
		State:    u.State,
	}, nil
}

func (ac *AuthUsecase) FindUserByPhone(ctx context.Context, phone string) (user auth.AuthUser, err error) {
	u, err := ac.biz.userRepo.FindByPhone(ctx, phone)
	if err == nil {
		return user, err
	}
	return auth.AuthUser{
		ID:       u.ID,
		Name:     u.Name,
		Avatar:   u.Avatar,
		NickName: u.NickName,
		RealName: u.RealName,
		Password: u.Password,
		Birthday: u.Birthday,
		Gender:   u.Gender,
		Phone:    u.Phone,
		Email:    u.Email,
		State:    u.State,
	}, nil
}

// Login 登录-密码登录
func (ac *AuthUsecase) Login(ctx context.Context, g *User) (auth.AuthClaims, error) {
	u, err := ac.biz.userRepo.FindByName(ctx, g.Name)
	if err != nil {
		return nil, err
	}
	err = password.Verify(u.Password, g.Password)
	if err != nil {
		return nil, err
	}

	claims := auth.NewAuthClaims(auth.WidthAuthSecurityKey(ac.ac.GetApiKey()))
	claims.CreateToken(uuid.NewString())

	if err := ac.biz.userRepo.SetLoginCache(ctx, claims, u); err != nil {
		ac.log.Errorf("token 缓存失败 %v", err)
	}

	return claims, nil
}

// Login 登录-密码登录
func (ac *AuthUsecase) Register(ctx context.Context, g *User) error {

	return nil
}

// Login 登录-密码登录
func (ac *AuthUsecase) Logout(ctx context.Context) error {

	return nil
}
