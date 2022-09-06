package biz

import (
	"context"
	"time"

	"github.com/beiduoke/go-scaffold/internal/conf"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
)

// User is a User model.
type User struct {
	CreatedAt            time.Time
	UpdatedAt            time.Time
	ID                   uint
	Name                 string
	NickName             string
	RealName             string
	Password             string
	Birthday             *time.Time
	Gender               int32
	Mobile               string
	Email                string
	State                int32
	Domains              []Domain
	Authorities          []Authority
	DomainAuthorityUsers []DomainAuthorityUser
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
	ListByMobile(context.Context, string) ([]*User, error)
	ListByEmail(context.Context, string) ([]*User, error)
	ListPage(context.Context, pagination.PaginationHandler) ([]*User, int64)
}

// UserUsecase is a User usecase.
type UserUsecase struct {
	ac         *conf.Auth
	repo       UserRepo
	domainRepo DomainRepo
	log        *log.Helper
	tm         Transaction
}

// NewUserUsecase new a User usecase.
func NewUserUsecase(ac *conf.Auth, repo UserRepo, tm Transaction, logger log.Logger, domainRepo DomainRepo) *UserUsecase {
	return &UserUsecase{ac: ac, repo: repo, tm: tm, log: log.NewHelper(logger), domainRepo: domainRepo}
}

// GetUserByMobile 获取用户手机
func (uc *UserUsecase) GetUserByMobile(ctx context.Context, g *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("MobileSmsLogin: %v", g)
	return uc.repo.FindByMobile(ctx, g.Mobile)
}
