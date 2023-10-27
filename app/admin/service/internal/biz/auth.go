package biz

import (
	"context"
	"time"

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
	Register(context.Context, *User) error
	Logout(context.Context) error
	LoginByPassword(ctx context.Context, access, password string) (*LoginResult, error)
	// 访问用户相关
	AccessInfo(context.Context) (*User, error)
	AccessRoles(context.Context) ([]*Role, error)
	AccessRoleMenus(context.Context) ([]*Menu, error)
	AccessRolePermissions(context.Context) ([]string, error)
}

// AuthUsecase is a User usecase.
type AuthUsecase struct {
	biz *Biz
	log *log.Helper
}

// NewAuthUsecase new a User usecase.
func NewAuthUsecase(logger log.Logger, biz *Biz) *AuthUsecase {
	return &AuthUsecase{log: log.NewHelper(logger), biz: biz}
}

// Login 登录-密码登录
func (ac *AuthUsecase) LoginByPassword(ctx context.Context, access, password string) (*LoginResult, error) {
	return ac.biz.authRepo.LoginByPassword(ctx, access, password)
}

// Login 登录-密码登录
func (ac *AuthUsecase) Register(ctx context.Context, g *User) error {
	return ac.biz.authRepo.Register(ctx, g)
}

// Login 登录-密码登录
func (ac *AuthUsecase) Logout(ctx context.Context) error {
	return ac.biz.authRepo.Logout(ctx)
}

// GetInfo 用户信息
func (ac *AuthUsecase) AccessInfo(ctx context.Context) (*User, error) {
	return ac.biz.authRepo.AccessInfo(ctx)
}

// GetRoles 用户角色
func (ac *AuthUsecase) AccessRoles(ctx context.Context) ([]*Role, error) {
	return ac.biz.authRepo.AccessRoles(ctx)
}

// GetRoles 用户角色菜单
func (ac *AuthUsecase) AccessRoleMenus(ctx context.Context) ([]*Menu, error) {
	return ac.biz.authRepo.AccessRoleMenus(ctx)
}

// GetRoles 用户角色权限
func (ac *AuthUsecase) AccessRolePermissions(ctx context.Context) ([]string, error) {
	return ac.biz.authRepo.AccessRolePermissions(ctx)
}
