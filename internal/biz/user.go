package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/zzsds/go-tools/pkg/password"
)

// User is a User model.
type User struct {
	ID       uint
	Name     string
	NickName string
	RealName string
	Password string
	Birthday *time.Time
	Gender   int32
	Mobile   string
	Email    string
	State    int32
}

// UserRepo is a Greater repo.
type UserRepo interface {
	// 基准操作
	Save(context.Context, *User) (*User, error)
	Update(context.Context, *User) (*User, error)
	FindByID(context.Context, int64) (*User, error)
	ListAll(context.Context) ([]*User, error)
	// 自定义操作
	FindByName(context.Context, string) (*User, error)
	FindByMobile(context.Context, string) (*User, error)
	FindByEmail(context.Context, string) (*User, error)
	ListByName(context.Context, string) ([]*User, error)
}

// UserUsecase is a User usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
	tm   Transaction
}

// NewUserUsecase new a User usecase.
func NewUserUsecase(repo UserRepo, tm Transaction, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, tm: tm, log: log.NewHelper(logger)}
}

// NamePasswordLogin 用户密码登录
func (uc *UserUsecase) NamePasswordLogin(ctx context.Context, g *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("PasswordLogin: %v", g)
	u, err := uc.repo.FindByName(ctx, g.Name)
	if err != nil {
		return nil, err
	}
	err = password.Verify(g.Password, g.Password)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// MobileSmsLogin 手机验证码登录
func (uc *UserUsecase) MobileSmsLogin(ctx context.Context, g *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("MobileSmsLogin: %v", g)
	return uc.repo.FindByMobile(ctx, g.Mobile)
}

// NamePasswordRegister 用户密码注册
func (uc *UserUsecase) NamePasswordRegister(ctx context.Context, g *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("NamePasswordRegister: %v", g.Name)
	return uc.repo.Save(ctx, g)
}
