package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"

	pb "github.com/beiduoke/go-scaffold/api/protobuf"
	"github.com/beiduoke/go-scaffold/internal/conf"
	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/zzsds/go-tools/pkg/password"
)

// User is a User model.
type User struct {
	CreatedAt       time.Time
	UpdatedAt       time.Time
	ID              uint
	Name            string
	Avatar          string
	NickName        string
	RealName        string
	Password        string
	Birthday        *time.Time
	Gender          int32
	Phone           string
	Email           string
	State           int32
	DeptID          uint
	Dept            *Dept
	DomainID        uint
	Domain          *Domain
	Roles           []*Role
	Posts           []*Post
	DomainRoleUsers []*DomainRoleUser
}

func (g User) GetID() string {
	return convert.UnitToString(g.ID)
}

func (g User) GetDomainID() string {
	return convert.UnitToString(g.DomainID)
}

// UserRepo is a Greater repo.
type UserRepo interface {
	// 用户认证
	Login(context.Context, *User) (*LoginResult, error)
	Register(context.Context, *User) error
	Logout(context.Context) error
	// 当前用户相关
	Info(context.Context) (*User, error)
	Roles(context.Context) ([]*Role, error)
	RoleMenus(context.Context) ([]*Menu, error)
	RolePermissions(context.Context) ([]string, error)
	// 基础操作
	Save(context.Context, *User) (*User, error)
	Update(context.Context, *User) (*User, error)
	FindByID(context.Context, uint) (*User, error)
	ListAll(context.Context) ([]*User, error)
	Delete(context.Context, *User) error
	// 自定义操作
	FindByName(context.Context, string) (*User, error)
	FindByPhone(context.Context, string) (*User, error)
	FindByEmail(context.Context, string) (*User, error)
	ListByName(context.Context, string) ([]*User, error)
	ListByPhone(context.Context, string) ([]*User, error)
	ListByEmail(context.Context, string) ([]*User, error)
	ListPage(context.Context, *pagination.Pagination) ([]*User, int64)
	// 用户关联
	ListRoles(context.Context, *User) ([]*Role, error)
	// 用户领域权限操作
	HandleRole(context.Context, *User) error
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
func (uc *UserUsecase) Create(ctx context.Context, g *User) (user *User, err error) {
	uc.log.WithContext(ctx).Debugf("CreateUser: %v", g)
	user, _ = uc.biz.userRepo.FindByName(ctx, g.Name)
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
	err = uc.biz.tm.InTx(ctx, func(ctx context.Context) error {
		if user, err = uc.biz.userRepo.Save(ctx, g); err != nil {
			return err
		}
		return uc.biz.userRepo.HandleRole(ctx, g)
	})
	return user, err
}

// HandleRole 绑定领域权限
func (uc *UserUsecase) HandleRole(ctx context.Context, g *User) error {
	uc.log.WithContext(ctx).Debugf("HandleRole: %v", g)
	roles := g.Roles
	if len(roles) <= 0 {
		return errors.New("权限未指定")
	}

	return uc.biz.userRepo.HandleRole(ctx, g)
}

// Update 修改用户
func (uc *UserUsecase) Update(ctx context.Context, g *User) error {
	uc.log.WithContext(ctx).Debugf("UpdateUser: %v", g)

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

	if user.Phone != g.Phone {
		phone, _ := uc.biz.userRepo.FindByPhone(ctx, g.Phone)
		if phone != nil {
			return errors.New("手机号已存在")
		}
	}

	if user.Email != g.Email {
		phone, _ := uc.biz.userRepo.FindByEmail(ctx, g.Email)
		if phone != nil {
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
	// if err := mergo.Merge(user, *g, mergo.WithOverride); err != nil {
	// 	return errors.Errorf("数据合并失败：%v", err)
	// }
	err := uc.biz.tm.InTx(ctx, func(ctx context.Context) error {
		if _, err := uc.biz.userRepo.Update(ctx, g); err != nil {
			return err
		}
		return uc.biz.userRepo.HandleRole(ctx, g)
	})
	return err
}

// List 用户列表全部
func (uc *UserUsecase) ListAll(ctx context.Context) ([]*User, int64) {
	uc.log.WithContext(ctx).Debugf("UserList")
	return uc.biz.userRepo.ListPage(ctx, &pagination.Pagination{Nopaging: true, OrderBy: map[string]bool{"sort": true}})
}

// List 用户列表分页
func (uc *UserUsecase) ListPage(ctx context.Context, paging *pagination.Pagination) ([]*User, int64) {
	uc.log.WithContext(ctx).Debugf("UserPage")
	return uc.biz.userRepo.ListPage(ctx, paging)
}

// GetID 获取用户ID
func (uc *UserUsecase) GetID(ctx context.Context, g *User) (*User, error) {
	uc.log.WithContext(ctx).Debugf("GetUserID: %v", g)
	user, err := uc.biz.userRepo.FindByID(ctx, g.ID)
	if err != nil {
		return nil, err
	}
	return user, err
}

// GetPhone 获取用户手机
func (uc *UserUsecase) GetPhone(ctx context.Context, g *User) (*User, error) {
	uc.log.WithContext(ctx).Debugf("GetUserPhone: %v", g)
	return uc.biz.userRepo.FindByPhone(ctx, g.Phone)
}

// GetName 获取用户名
func (uc *UserUsecase) GetName(ctx context.Context, g *User) (*User, error) {
	uc.log.WithContext(ctx).Debugf("GetUserName: %v", g)
	return uc.biz.userRepo.FindByName(ctx, g.Name)
}

// Delete 删除用户
func (uc *UserUsecase) Delete(ctx context.Context, g *User) error {
	uc.log.WithContext(ctx).Debugf("DeleteUser: %v", g)
	return uc.biz.tm.InTx(ctx, func(ctx context.Context) error {
		if err := uc.biz.userRepo.Delete(ctx, g); err != nil {
			return err
		}

		_, err := uc.biz.enforcer.DeleteUser(g.GetID())
		return err
	})
}

// GetInfo 用户信息
func (ac *UserUsecase) Info(ctx context.Context) (*User, error) {
	return ac.biz.userRepo.Info(ctx)
}

// GetRoles 用户角色
func (ac *UserUsecase) Roles(ctx context.Context) ([]*Role, error) {
	return ac.biz.userRepo.Roles(ctx)
}

// GetRoles 用户角色菜单
func (ac *UserUsecase) RoleMenus(ctx context.Context) ([]*Menu, error) {
	return ac.biz.userRepo.RoleMenus(ctx)
}

// GetRoles 用户角色权限
func (ac *UserUsecase) RolePermissions(ctx context.Context) ([]string, error) {
	return ac.biz.userRepo.RolePermissions(ctx)
}

// ListPostID 获取岗位ID列表
func (ac *UserUsecase) ListPostID(ctx context.Context, g *User) (postIds []uint, err error) {

	return postIds, nil
}

// ListRoleID 获取角色ID列表
func (ac *UserUsecase) ListRoleID(ctx context.Context, g *User) (roleIds []uint, err error) {
	if g.DomainID <= 0 {
		return nil, errors.New("领域不能为空")
	}
	for _, v := range ac.biz.enforcer.GetRolesForUserInDomain(g.GetID(), g.GetDomainID()) {
		roleIds = append(roleIds, convert.StringToUint(v))
	}
	return roleIds, nil
}

// ListRoleAll 获取角色列表
func (ac *UserUsecase) ListRoleAll(ctx context.Context, g *User) (roles []*Role, err error) {
	roleIds, err := ac.ListRoleID(ctx, g)
	if err != nil || len(roleIds) < 1 {
		return roles, err
	}
	return ac.biz.roleRepo.ListByIDs(ctx, roleIds...)
}

// ListRoleMenu 用户角色菜单列表(包含权限标识)
func (ac *UserUsecase) ListRoleMenu(ctx context.Context, g *User) ([]*Menu, error) {
	roleIds, err := ac.ListRoleID(ctx, g)
	if err != nil {
		return nil, errors.Errorf("用户角色查询失败 %v", err)
	}

	roleIdsRes := make([]uint, len(roleIds))
	if len(g.Roles) > 1 {
		for _, v := range g.Roles {
			for _, a := range roleIds {
				if a == v.ID {
					roleIdsRes = append(roleIdsRes, v.ID)
					break
				}
			}
		}
	} else {
		roleIdsRes = roleIds
	}

	if len(roleIdsRes) < 1 {
		return nil, nil
	}

	return ac.biz.roleRepo.ListMenuAndParentByIDs(ctx, roleIdsRes...)
}
