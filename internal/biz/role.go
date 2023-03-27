package biz

import (
	"context"
	"time"

	pb "github.com/beiduoke/go-scaffold/api/protobuf"
	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
)

// Role is a Role model.
type Role struct {
	ID                uint
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Name              string
	ParentID          uint
	DefaultRouter     string
	Sort              int32
	DataScope         int32
	MenuCheckStrictly int32
	DeptCheckStrictly int32
	State             int32
	Remarks           string
	Users             []*User
	Domains           []*Domain
	Menus             []*Menu
	Resources         []*Resource
	Depts             []*Dept
}

func (g Role) GetID() string {
	return convert.UnitToString(g.ID)
}

// RoleRepo is a Greater repo.
type RoleRepo interface {
	Save(context.Context, *Role) (*Role, error)
	Update(context.Context, *Role) (*Role, error)
	FindByID(context.Context, uint) (*Role, error)
	FindByName(context.Context, string) (*Role, error)
	ListByName(context.Context, string) ([]*Role, error)
	ListByIDs(context.Context, ...uint) ([]*Role, error)
	ListAll(context.Context) ([]*Role, error)
	Delete(context.Context, *Role) error
	ListPage(context.Context, *pagination.Pagination) ([]*Role, int64)
	HandleMenu(context.Context, *Role) error
	ListMenuByIDs(context.Context, ...uint) ([]*Menu, error)
	ListMenuAndParentByIDs(context.Context, ...uint) ([]*Menu, error)
	ListResourceByIDs(context.Context, ...uint) ([]*Resource, error)
	HandleResource(context.Context, *Role) error
	ListDeptByIDs(context.Context, ...uint) ([]*Dept, error)
	HandleDept(context.Context, *Role) error
}

// RoleUsecase is a Role usecase.
type RoleUsecase struct {
	biz *Biz
	log *log.Helper
}

// NewRoleUsecase new a Role usecase.
func NewRoleUsecase(logger log.Logger, biz *Biz) *RoleUsecase {
	return &RoleUsecase{log: log.NewHelper(logger), biz: biz}
}

// Create creates a Role, and returns the new Role.
func (uc *RoleUsecase) Create(ctx context.Context, g *Role) (*Role, error) {
	uc.log.WithContext(ctx).Debugf("Create: %v", g.Name)
	return uc.biz.roleRepo.Save(ctx, g)
}

// ListByIDs 获取指定角色ID集合
func (uc *RoleUsecase) ListByIDs(ctx context.Context, id ...uint) (roles []*Role, err error) {
	roles, _ = uc.biz.roleRepo.ListPage(ctx, &pagination.Pagination{Nopaging: true, OrderBy: map[string]bool{"sort": true}})
	return
}

// Update 修改角色
func (uc *RoleUsecase) Update(ctx context.Context, g *Role) error {
	uc.log.WithContext(ctx).Debugf("UpdateRole: %v", g)

	role, _ := uc.biz.roleRepo.FindByID(ctx, g.ID)
	if role == nil {
		return errors.New("角色未注册")
	}

	if role.Name != g.Name && g.Name != "" {
		name, _ := uc.biz.roleRepo.FindByName(ctx, g.Name)
		if name != nil {
			return errors.New("角色名已存在")
		}
	}

	if g.State <= 0 {
		g.State = int32(pb.RoleState_ROLE_STATE_ACTIVE)
	}
	// 新数据合并到源数据
	// if err := mergo.Merge(role, *g, mergo.WithOverwriteWithEmptyValue); err != nil {
	// 	return errors.Errorf("数据合并失败：%v", err)
	// }
	_, err := uc.biz.roleRepo.Update(ctx, g)
	return err
}

// UpdateState 修改角色状态
func (uc *RoleUsecase) UpdateState(ctx context.Context, g *Role) error {
	uc.log.WithContext(ctx).Debugf("UpdateRoleState: %v", g)

	role, _ := uc.biz.roleRepo.FindByID(ctx, g.ID)
	if role == nil {
		return errors.New("角色未不存在")
	}

	if g.State <= 0 {
		g.State = int32(pb.RoleState_ROLE_STATE_ACTIVE)
	}

	role.State = g.State
	_, err := uc.biz.roleRepo.Update(ctx, role)
	return err
}

// List 角色列表全部
func (uc *RoleUsecase) ListAll(ctx context.Context) ([]*Role, int64) {
	uc.log.WithContext(ctx).Debugf("RoleList")
	return uc.biz.roleRepo.ListPage(ctx, &pagination.Pagination{Nopaging: true, OrderBy: map[string]bool{"sort": true}})
}

// List 角色列表分页
func (uc *RoleUsecase) ListPage(ctx context.Context, paging *pagination.Pagination) ([]*Role, int64) {
	uc.log.WithContext(ctx).Debugf("RolePage")
	return uc.biz.roleRepo.ListPage(ctx, paging)
}

// GetID 根据角色ID角色
func (uc *RoleUsecase) GetID(ctx context.Context, g *Role) (*Role, error) {
	uc.log.WithContext(ctx).Debugf("GetRoleID: %v", g)
	return uc.biz.roleRepo.FindByID(ctx, g.ID)
}

// Delete 根据角色ID删除角色
func (uc *RoleUsecase) Delete(ctx context.Context, g *Role) error {
	uc.log.WithContext(ctx).Debugf("DeleteRole: %v", g)
	return uc.biz.tm.InTx(ctx, func(ctx context.Context) error {
		if err := uc.biz.roleRepo.Delete(ctx, g); err != nil {
			return err
		}
		_, err := uc.biz.enforcer.DeleteRole(g.GetID())
		return err
	})
}

// HandleMenu 获取角色菜单
func (uc *RoleUsecase) ListMenuByID(ctx context.Context, g *Role) ([]*Menu, error) {
	uc.log.WithContext(ctx).Debugf("ListMenuByIDs: %v", g)
	return uc.biz.roleRepo.ListMenuByIDs(ctx, g.ID)
}

// HandleMenu 绑定菜单
func (uc *RoleUsecase) HandleMenu(ctx context.Context, g *Role) error {
	uc.log.WithContext(ctx).Debugf("HandleMenu: %v", g)
	return uc.biz.roleRepo.HandleMenu(ctx, g)
}

// HandleMenu 获取角色菜单
func (uc *RoleUsecase) ListResourceByID(ctx context.Context, g *Role) ([]*Resource, error) {
	uc.log.WithContext(ctx).Debugf("ListMenuByIDs: %v", g)
	return uc.biz.roleRepo.ListResourceByIDs(ctx, g.ID)
}

// HandleResource 绑定资源
func (uc *RoleUsecase) HandleResource(ctx context.Context, g *Role) error {
	uc.log.WithContext(ctx).Debugf("HandleResource: %v", g)
	return uc.biz.roleRepo.HandleResource(ctx, g)
}

// HandleMenu 获取角色菜单
func (uc *RoleUsecase) ListDeptByID(ctx context.Context, g *Role) ([]*Dept, error) {
	uc.log.WithContext(ctx).Debugf("ListMenuByIDs: %v", g)
	return uc.biz.roleRepo.ListDeptByIDs(ctx, g.ID)
}

// HandleMenu 获取角色数据范围
func (uc *RoleUsecase) GetDataScopeByID(ctx context.Context, g *Role) (*Role, error) {
	uc.log.WithContext(ctx).Debugf("GetDataScopeByID: %v", g)
	role, err := uc.biz.roleRepo.FindByID(ctx, g.ID)
	if err != nil {
		return nil, err
	}
	role.Depts, _ = uc.biz.roleRepo.ListDeptByIDs(ctx, role.ID)
	return role, err
}

// HandleDept 绑定数据范围
func (uc *RoleUsecase) HandleDataScope(ctx context.Context, g *Role) error {
	uc.log.WithContext(ctx).Debugf("HandleDataScope: %v", g)
	role, err := uc.biz.roleRepo.FindByID(ctx, g.ID)
	if err != nil {
		return err
	}

	_, err = uc.biz.roleRepo.Update(ctx, role)
	if err != nil {
		return err
	}

	if len(g.Depts) > 0 {
		return uc.biz.roleRepo.HandleDept(ctx, g)
	}
	return err
}
