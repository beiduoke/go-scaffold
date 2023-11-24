package biz

import (
	"context"
	"time"

	v1 "github.com/beiduoke/go-scaffold/api/saasdesk/service/v1"
	"github.com/go-kratos/kratos/v2/log"
)

type LoginResult struct {
	Token        string
	RefreshToken string
	ExpiresAt    *time.Time
}

// AuthRepo is a Greater repo.
type AuthRepo interface {
	// 用户认证
	Register(context.Context, *v1.RegisterRequest) error
	Logout(context.Context) error
	LoginByPassword(context.Context, *v1.LoginByPasswordRequest) (*v1.LoginResponse, error)
	// 访问用户相关
	GetAuthInfo(context.Context) (*v1.GetAuthInfoResponse, error)
}

// AuthUsecase is a User usecase.
type AuthUsecase struct {
	repo AuthRepo
	log  *log.Helper
}

// NewAuthUsecase new a User usecase.
func NewAuthUsecase(repo AuthRepo, logger log.Logger) *AuthUsecase {
	return &AuthUsecase{repo: repo, log: log.NewHelper(logger)}
}

// Login 登录-密码登录
func (ac *AuthUsecase) LoginByPassword(ctx context.Context, pb *v1.LoginByPasswordRequest) (*v1.LoginResponse, error) {
	return ac.repo.LoginByPassword(ctx, pb)
}

// Login 登录-密码登录
func (ac *AuthUsecase) Register(ctx context.Context, pb *v1.RegisterRequest) error {
	return ac.repo.Register(ctx, pb)
}

// Login 登录-密码登录
func (ac *AuthUsecase) Logout(ctx context.Context) error {
	return ac.repo.Logout(ctx)
}

// GetInfo 用户信息
func (ac *AuthUsecase) GetAuthInfo(ctx context.Context) (*v1.GetAuthInfoResponse, error) {
	return ac.repo.GetAuthInfo(ctx)
}
