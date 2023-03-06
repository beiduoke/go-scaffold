package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/beiduoke/go-scaffold/pkg/auth"
)

type LoginResult struct {
	Token     string
	ExpiresAt *time.Time
}

// AuthUsecase is a User usecase.
type AuthUsecase struct {
	biz *Biz
	log *log.Helper
}

// NewAuthUsecase new a User usecase.
func NewAuthUsecase(logger log.Logger, biz *Biz, authenticator auth.Authenticator) *AuthUsecase {
	return &AuthUsecase{log: log.NewHelper(logger), biz: biz}
}

// Login 登录-密码登录
func (ac *AuthUsecase) Login(ctx context.Context, g *User) (*LoginResult, error) {
	return ac.biz.userRepo.Login(ctx, g)
}

// Login 登录-密码登录
func (ac *AuthUsecase) Register(ctx context.Context, g *User) error {
	return ac.biz.userRepo.Register(ctx, g)
}

// Login 登录-密码登录
func (ac *AuthUsecase) Logout(ctx context.Context) error {

	return nil
}
