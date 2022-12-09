package biz

import (
	"context"
	"fmt"
	"time"

	stdcasbin "github.com/casbin/casbin/v2"
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
	// 用户领域权限操作
	HandleDomain(context.Context, *User) error
	HandleDomainAuthority(context.Context, *User) error
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
	return uc.biz.userRepo.Save(ctx, g)
}

// HandleDomain 绑定领域
func (uc *UserUsecase) HandleDomain(ctx context.Context, g *User) error {
	uc.log.WithContext(ctx).Infof("HandleDomain: %v", g)
	domains := g.Domains
	if len(domains) <= 0 {
		return errors.New("领域未指定")
	}
	return uc.biz.userRepo.HandleDomain(ctx, g)
}

// HandleDomainAuthority 绑定领域权限
func (uc *UserUsecase) HandleDomainAuthority(ctx context.Context, g *User, domainId uint) error {
	uc.log.WithContext(ctx).Infof("HandleDomainAuthority: %v", g)
	authorities := g.Authorities
	if len(authorities) <= 0 {
		return errors.New("权限未指定")
	}

	return uc.biz.userRepo.HandleDomainAuthority(ctx, g)
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
	user, err := uc.biz.userRepo.FindByID(ctx, g.ID)
	if err != nil {
		return nil, err
	}
	return user, err
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
		if err := uc.biz.userRepo.Delete(ctx, g); err != nil {
			return err
		}
		_, err := uc.biz.enforcer.DeleteUser(convert.UnitToString(g.ID))
		return err
	})
}

// GetLastDomain 获取最后切换使用领域
func (ac *UserUsecase) GetLastUseDomain(ctx context.Context, g *User) (*Domain, error) {
	domainPolices := ac.biz.enforcer.GetFilteredGroupingPolicy(0, convert.UnitToString(g.ID))
	if len(domainPolices) < 1 && len(domainPolices[0]) >= 2 {
		return nil, errors.New("领域查询失败")
	}
	lastUseAuthority, lastUseDomain := domainPolices[0][1], domainPolices[0][2]
	fmt.Println(domainPolices)
	for _, policy := range domainPolices {
		if p := policy[len(policy)-1]; p == "1" {
			lastUseAuthority, lastUseDomain = policy[1], policy[2]
			break
		}
	}
	return &Domain{
		ID:                 convert.StringToUint(lastUseDomain),
		DefaultAuthorityID: convert.StringToUint(lastUseAuthority),
	}, nil
	// 暂无用
	// 获取最近一次登录的领域下所有角色
	roles, err := ac.biz.enforcer.(*stdcasbin.SyncedEnforcer).GetNamedRoleManager("g").GetRoles(convert.UnitToString(g.ID), lastUseDomain)
	authorities := make([]uint, 0, len(roles))
	for _, v := range roles {
		authorities = append(authorities, convert.StringToUint(v))
	}
	ac.log.Infof("打印角色列表 %v", authorities)
	return nil, err
}

// GetLastDomain 获取最后切换使用领域
func (ac *UserUsecase) ListDomainAll(ctx context.Context, g *User) ([]*Domain, error) {
	domainPolices := ac.biz.enforcer.GetFilteredGroupingPolicy(0, convert.UnitToString(g.ID))
	if len(domainPolices) < 1 && len(domainPolices[0]) >= 2 {
		return nil, errors.New("领域查询失败")
	}
	userDomainIds, err := ac.biz.enforcer.(*stdcasbin.SyncedEnforcer).GetDomainsForUser(convert.UnitToString(g.ID))
	if err != nil {
		return nil, err
	}

	domainIds := make([]uint, 0, len(userDomainIds))
	for _, v := range userDomainIds {
		domainIds = append(domainIds, convert.StringToUint(v))
	}

	return ac.biz.domainRepo.ListByIDs(ctx, domainIds...)
}

// ListAuthorityID 获取权限角色ID列表
func (ac *UserUsecase) ListAuthorityID(ctx context.Context, g *User) (authorityIds []uint, err error) {
	uidStr := convert.UnitToString(g.ID)
	var rolesIdsStr []string
	if len(g.Domains) < 1 {
		rolesIdsStr, err = ac.biz.enforcer.GetRolesForUser(uidStr, "0")
	} else {
		domainIdStr := convert.UnitToString(g.Domains[0].ID)
		rolesIdsStr = ac.biz.enforcer.GetRolesForUserInDomain(uidStr, domainIdStr)
	}

	if err != nil {
		return nil, err
	}

	rolesIds := make([]uint, 0, len(rolesIdsStr))
	for _, v := range rolesIdsStr {
		rolesIds = append(rolesIds, convert.StringToUint(v))
	}

	return rolesIds, nil
}

// ListAuthorityAll 获取权限角色列表
func (ac *UserUsecase) ListAuthorityAll(ctx context.Context, g *User) (authorities []*Authority, err error) {
	authorityIds, err := ac.ListAuthorityID(ctx, g)
	if err != nil || len(authorityIds) < 1 {
		return authorities, err
	}
	return ac.biz.authorityRepo.ListByIDs(ctx, authorityIds...)
}

// ListUserAuthorityMenuAll 用户权限角色菜单列表(包含权限标识)
func (ac *UserUsecase) ListAuthorityMenuAll(ctx context.Context, g *User) ([]*Menu, error) {
	authorityIds := make([]uint, 0)
	for _, v := range g.Authorities {
		authorityIds = append(authorityIds, v.ID)
	}
	if len(authorityIds) < 1 {
		return nil, errors.Errorf("查询权限不能为空")
	}

	return ac.biz.authorityRepo.ListMenuByIDs(ctx, authorityIds...)
}
