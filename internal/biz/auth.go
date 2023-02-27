package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/zzsds/go-tools/pkg/password"

	"github.com/beiduoke/go-scaffold/internal/conf"
	"github.com/beiduoke/go-scaffold/internal/pkg/auth"
)

// AuthUsecase is a User usecase.
type AuthUsecase struct {
	biz *Biz
	ac  *conf.Auth
	log *log.Helper
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
