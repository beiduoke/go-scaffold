package biz

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"

	pb "github.com/beiduoke/go-scaffold/api/protobuf"
	"github.com/beiduoke/go-scaffold/internal/conf"
	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/imdario/mergo"
	"github.com/zzsds/go-tools/pkg/password"
)

// User is a User model.
type User struct {
	CreatedAt            time.Time
	UpdatedAt            time.Time
	ID                   uint
	Name                 string
	Avatar               string
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
	Delete(context.Context, *User) error
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
	biz *Biz
	log *log.Helper
	ac  *conf.Auth
}

// NewUserUsecase new a User usecase.
func NewUserUsecase(logger log.Logger, biz *Biz, ac *conf.Auth) *UserUsecase {
	return &UserUsecase{log: log.NewHelper(logger), ac: ac, biz: biz}
}

// Create 创建用户
func (uc *UserUsecase) Create(ctx context.Context, g *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("CreateUser: %v", g)
	// 创建用户只能关联一次领域
	domains := g.Domains
	if len(domains) <= 0 {
		return nil, errors.New("领域未指定")
	}
	user, _ := uc.biz.userRepo.FindByName(ctx, g.Name)
	if user != nil && user.Name != "" {
		return nil, errors.New("用户名已存在")
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

	err := uc.biz.tm.InTx(ctx, func(ctx context.Context) error {
		g, err := uc.biz.userRepo.Save(ctx, g)
		if err != nil {
			return err
		}
		for _, domain := range domains {
			if _, err := uc.biz.enforcer.AddRoleForUserInDomain(convert.UnitToString(g.ID), convert.UnitToString(domain.DefaultAuthorityID), convert.UnitToString(domain.ID)); err != nil {
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

	user, _ := uc.biz.userRepo.FindByID(ctx, g.ID)
	if user == nil {
		return errors.New("用户未注册")
	}

	if user.Name != g.Name && g.Name != "" {
		name, _ := uc.biz.userRepo.FindByName(ctx, g.Name)
		if name != nil {
			return errors.New("用户名已存在")
		}
	}

	if user.Mobile != g.Mobile {
		mobile, _ := uc.biz.userRepo.FindByMobile(ctx, g.Mobile)
		if mobile != nil {
			return errors.New("手机号已存在")
		}
	}

	if user.Email != g.Email {
		mobile, _ := uc.biz.userRepo.FindByEmail(ctx, g.Email)
		if mobile != nil {
			return errors.New("邮箱已存在")
		}
	}

	if g.Password != "" {
		password, err := password.Encryption(g.Password)
		if err != nil {
			return errors.Errorf("密码加密失败：%v", err)
		}
		g.Password = password
	}

	if g.State <= 0 {
		g.State = int32(pb.UserState_USER_STATE_ACTIVE)
	}
	// 新数据合并到源数据
	if err := mergo.Merge(user, *g, mergo.WithOverride); err != nil {
		return errors.Errorf("数据合并失败：%v", err)
	}
	_, err := uc.biz.userRepo.Update(ctx, user)
	return err
}

// List 用户列表全部
func (uc *UserUsecase) ListAll(ctx context.Context) ([]*User, int64) {
	uc.log.WithContext(ctx).Infof("UserList")
	return uc.biz.userRepo.ListPage(ctx, pagination.NewPagination())
}

// List 用户列表分页
func (uc *UserUsecase) ListPage(ctx context.Context, pageNum, pageSize int32, query map[string]string, order map[string]bool) ([]*User, int64) {
	uc.log.WithContext(ctx).Infof("UserPage")
	conditions := []pagination.Condition{}
	for k, v := range query {
		conditions = append(conditions, pagination.Condition{Query: k, Args: []interface{}{v}})
	}
	orders := []pagination.Order{}
	for k, v := range order {
		orders = append(orders, pagination.Order{Column: k, Desc: v})
	}

	page := pagination.NewPagination(
		pagination.WithPageNum(pageNum),
		pagination.WithPageSize(pageSize),
		pagination.WithConditions(conditions...),
		pagination.WithOrders(orders...),
	)
	return uc.biz.userRepo.ListPage(ctx, page)
}

// GetID 获取用户ID
func (uc *UserUsecase) GetID(ctx context.Context, g *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("GetUserID: %v", g)
	return uc.biz.userRepo.FindByID(ctx, g.ID)
}

// GetMobile 获取用户手机
func (uc *UserUsecase) GetMobile(ctx context.Context, g *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("GetUserMobile: %v", g)
	return uc.biz.userRepo.FindByMobile(ctx, g.Mobile)
}

// Delete 删除用户
func (uc *UserUsecase) Delete(ctx context.Context, g *User) error {
	uc.log.WithContext(ctx).Infof("DeleteUser: %v", g)
	return uc.biz.tm.InTx(ctx, func(ctx context.Context) error {
		var wg sync.WaitGroup
		if err := uc.biz.userRepo.Delete(ctx, g); err != nil {
			return err
		}
		wg.Add(len(g.Domains))
		for _, v := range g.Domains {
			fmt.Println(v)
			// roles := uc.enforcer.GetRolesForUserInDomain(convert.UnitToString(g.ID), convert.UnitToString(v.DomainID))

			// if _, err := uc.biz.enforcer.DeleteRolesForUserInDomain(convert.UnitToString(g.ID), convert.UnitToString(v.ID)); err != nil {
			// 	uc.log.Errorf("删除删除用户的角色失败：%v", err)
			// }
			wg.Done()
		}
		wg.Wait()
		return nil
	})
}
