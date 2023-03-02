package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/zzsds/go-tools/pkg/password"

	"github.com/beiduoke/go-scaffold/pkg/auth"
)

// AuthUsecase is a User usecase.
type AuthUsecase struct {
	biz           *Biz
	log           *log.Helper
	authenticator auth.Authenticator
}

// NewAuthUsecase new a User usecase.
func NewAuthUsecase(logger log.Logger, biz *Biz, authenticator auth.Authenticator) *AuthUsecase {
	return &AuthUsecase{log: log.NewHelper(logger), authenticator: authenticator, biz: biz}
}

func (ac *AuthUsecase) Token(ctx context.Context, claims auth.AuthClaims, user *User) (string, error) {
	token, err := ac.authenticator.CreateIdentity(ctx, claims)
	if err != nil {
		ac.log.Errorf("token 生成失败 %v", err)
	}

	if err := ac.biz.userRepo.SetLoginCache(ctx, &claims, user); err != nil {
		ac.log.Errorf("token 缓存失败 %v", err)
	}

	if err := ac.biz.userRepo.SetTokenCache(ctx, claims.Subject, token, time.Hour*24); err != nil {
		ac.log.Errorf("token 缓存失败 %v", err)
	}
	return token, nil
}

// Login 登录-密码登录
func (ac *AuthUsecase) Login(ctx context.Context, g *User) (string, error) {
	u, err := ac.biz.userRepo.FindByName(ctx, g.Name)
	if err != nil {
		return "", err
	}
	err = password.Verify(u.Password, g.Password)
	if err != nil {
		return "", err
	}
	return ac.Token(ctx, auth.AuthClaims{Subject: uuid.NewString()}, u)
}

// Login 登录-密码登录
func (ac *AuthUsecase) Register(ctx context.Context, g *User) error {

	return nil
}

// Login 登录-密码登录
func (ac *AuthUsecase) Logout(ctx context.Context) error {

	return nil
}
