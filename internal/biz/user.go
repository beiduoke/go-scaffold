package biz

import (
	"context"
	"errors"
	"time"

	pb "github.com/beiduoke/go-scaffold/api/protobuf"
	"github.com/beiduoke/go-scaffold/internal/conf"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/zzsds/go-tools/pkg/password"
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
	Domains              []*Domain
	Authorities          []*Authority
	DomainAuthorityUsers []*DomainAuthorityUser
}

// UserRepo is a Greater repo.
type UserRepo interface {
	// 基准操作
	Save(context.Context, *User) (*User, error)
	Update(context.Context, *User) (*User, error)
	FindByID(context.Context, uint) (*User, error)
	ListAll(context.Context) ([]*User, error)
	// 自定义操作
	FindByName(context.Context, string) (*User, error)
	FindByMobile(context.Context, string) (*User, error)
	FindByEmail(context.Context, string) (*User, error)
	ListByName(context.Context, string) ([]*User, error)
	ListByMobile(context.Context, string) ([]*User, error)
	ListByEmail(context.Context, string) ([]*User, error)
	ListPage(context.Context, pagination.PaginationHandler) ([]*User, int64)

	// 缓存操作
	SetTokenCache(context.Context, AuthClaims) error
	GetTokenCache(context.Context, AuthClaims) error
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

// CreateUser 创建用户
func (uc *UserUsecase) Create(ctx context.Context, g *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("CreateUser: %v", g)
	// 创建用户只能关联一次领域
	domains := g.Domains
	if len(domains) <= 0 {
		return nil, errors.New("未绑定领域")
	}
	user, _ := uc.repo.FindByName(ctx, g.Name)
	if user != nil && user.Name != "" {
		return nil, errors.New("已注册")
	}

	if g.Password != "" {
		password, err := password.Encryption(g.Password)
		if err != nil {
			return nil, errors.New("密码加密失败")
		}
		g.Password = password
	}

	if g.State <= 0 {
		g.State = int32(pb.UserState_USER_STATE_ACTIVE)
	}

	err := uc.tm.InTx(ctx, func(ctx context.Context) error {
		g, err := uc.repo.Save(ctx, g)
		if err != nil {
			return err
		}
		for _, domain := range domains {
			if err := uc.domainRepo.SaveAuthorityForUserInDomain(ctx, g.ID, domain.DefaultAuthorityID, domain.ID); err != nil {
				uc.log.Errorf("领域权限绑定失败 %v", err)
			}
		}
		return nil
	})
	return g, err
}

// Update 修改用户
func (uc *UserUsecase) Update(ctx context.Context, g *User) error {
	uc.log.WithContext(ctx).Infof("UpdateUser: %v", g)

	user, _ := uc.repo.FindByID(ctx, g.ID)
	if user == nil {
		return errors.New("用户未注册")
	}

	if user.Name != g.Name {
		name, _ := uc.repo.FindByName(ctx, g.Name)
		if name != nil {
			return errors.New("用户名已存在")
		}
	}

	if user.Mobile != g.Mobile {
		mobile, _ := uc.repo.FindByMobile(ctx, g.Mobile)
		if mobile != nil {
			return errors.New("手机号已存在")
		}
	}

	if user.Email != g.Email {
		mobile, _ := uc.repo.FindByEmail(ctx, g.Email)
		if mobile != nil {
			return errors.New("邮箱已存在")
		}
	}

	if g.Password != "" {
		password, err := password.Encryption(g.Password)
		if err != nil {
			return errors.New("密码加密失败")
		}
		g.Password = password
	}

	if g.State <= 0 {
		g.State = int32(pb.UserState_USER_STATE_ACTIVE)
	}

	err := uc.tm.InTx(ctx, func(ctx context.Context) error {
		_, err := uc.repo.Update(ctx, g)
		return err
	})
	return err
}

// List 用户列表
func (uc *UserUsecase) List(ctx context.Context) ([]*User, int64) {
	uc.log.WithContext(ctx).Infof("List")
	return uc.repo.ListPage(ctx, pagination.NewPagination())
}

// GetID 获取用户ID
func (uc *UserUsecase) GetID(ctx context.Context, g *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("GetID: %v", g)
	return uc.repo.FindByID(ctx, g.ID)
}

// GetMobile 获取用户手机
func (uc *UserUsecase) GetMobile(ctx context.Context, g *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("GetMobile: %v", g)
	return uc.repo.FindByMobile(ctx, g.Mobile)
}
