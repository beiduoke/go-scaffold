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
	CreatedAt     time.Time  `json:"createdAt,omitempty" form:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt,omitempty" form:"updatedAt"`
	ID            uint       `json:"id,omitempty" form:"id"`
	Name          string     `json:"name,omitempty" form:"name"`
	Avatar        string     `json:"avatar,omitempty" form:"avatar"`
	NickName      string     `json:"nickName,omitempty" form:"nickName"`
	RealName      string     `json:"realName,omitempty" form:"realName"`
	Password      string     `json:"-,omitempty" form:"-,omitempty"`
	Birthday      *time.Time `json:"birthday,omitempty" form:"birthday"`
	Gender        int32      `json:"gender,omitempty" form:"gender"`
	Phone         string     `json:"phone,omitempty" form:"phone"`
	Email         string     `json:"email,omitempty" form:"email"`
	State         int32      `json:"state,omitempty" form:"state"`
	Remarks       string     `json:"remarks,omitempty" form:"remarks"`
	DeptID        uint       `json:"deptId,omitempty" form:"deptId"`
	Dept          *Dept      `json:"dept,omitempty" form:"dept"`
	DomainID      uint       `json:"domainId,omitempty" form:"domainId"`
	Domain        *Domain    `json:"domain,omitempty" form:"domain"`
	Roles         []*Role    `json:"roles,omitempty" form:"roles"`
	Posts         []*Post    `json:"posts,omitempty" form:"posts"`
	LastUseRoleID uint       `json:"lastUseRoleId,omitempty" form:"lastUseRoleId"`
	LastLoginAt   *time.Time `json:"lastLoginAt,omitempty" form:"lastLoginAt"`
	LastUseRole   *Role      `json:"lastUseRole,omitempty" form:"lastUseRole"`
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
	// 访问用户相关
	AccessInfo(context.Context) (*User, error)
	AccessRoles(context.Context) ([]*Role, error)
	AccessRoleMenus(context.Context) ([]*Menu, error)
	AccessRolePermissions(context.Context) ([]string, error)
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
	return uc.biz.userRepo.Save(ctx, g)
}

// Update 修改用户
func (uc *UserUsecase) Update(ctx context.Context, g *User) error {
	uc.log.WithContext(ctx).Debugf("UpdateUser: %v", g)

	user, _ := uc.biz.userRepo.FindByID(ctx, g.ID)
	if user == nil {
		return errors.New("用户未注册")
	}

	if g.Name != "" && user.Name != g.Name {
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

	if g.Email != "" && user.Email != g.Email {
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
	_, err := uc.biz.userRepo.Update(ctx, g)
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
	return uc.biz.userRepo.Delete(ctx, g)
}

// GetInfo 用户信息
func (ac *UserUsecase) AccessInfo(ctx context.Context) (*User, error) {
	return ac.biz.userRepo.AccessInfo(ctx)
}

// GetRoles 用户角色
func (ac *UserUsecase) AccessRoles(ctx context.Context) ([]*Role, error) {
	return ac.biz.userRepo.AccessRoles(ctx)
}

// GetRoles 用户角色菜单
func (ac *UserUsecase) AccessRoleMenus(ctx context.Context) ([]*Menu, error) {
	return ac.biz.userRepo.AccessRoleMenus(ctx)
}

// GetRoles 用户角色权限
func (ac *UserUsecase) AccessRolePermissions(ctx context.Context) ([]string, error) {
	return ac.biz.userRepo.AccessRolePermissions(ctx)
}

// ListRoles 获取角色列表
func (ac *UserUsecase) ListRoles(ctx context.Context, g *User) (roles []*Role, err error) {
	return ac.biz.userRepo.ListRoles(ctx, g)
}
