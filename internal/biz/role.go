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
	ID                uint      `json:"id,omitempty" form:"id,omitempty"`
	CreatedAt         time.Time `json:"createdAt,omitempty" form:"createdAt,omitempty"`
	UpdatedAt         time.Time `json:"updatedAt,omitempty" form:"updatedAt,omitempty"`
	Name              string    `json:"name,omitempty" form:"name,omitempty"`
	ParentID          uint      `json:"parentID,omitempty" form:"parentID,omitempty"`
	DefaultRouter     string    `json:"defaultRouter,omitempty" form:"defaultRouter,omitempty"`
	Sort              int32     `json:"sort,omitempty" form:"sort,omitempty"`
	DataScope         int32     `json:"dataScope,omitempty" form:"dataScope,omitempty"`
	MenuCheckStrictly int32     `json:"menuCheckStrictly,omitempty" form:"menuCheckStrictly,omitempty"`
	DeptCheckStrictly int32     `json:"deptCheckStrictly,omitempty" form:"deptCheckStrictly,omitempty"`
	State             int32     `json:"state,omitempty" form:"state,omitempty"`
	Remarks           string    `json:"remarks,omitempty" form:"remarks,omitempty"`
	Users             []*User   `json:"users,omitempty" form:"users,omitempty"`
	Domains           []*Domain `json:"domains,omitempty" form:"domains,omitempty"`
	Menus             []*Menu   `json:"menus,omitempty" form:"menus,omitempty"`
	Depts             []*Dept   `json:"depts,omitempty" form:"depts,omitempty"`
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
	// 指定角色ID菜单列表
	ListMenuByIDs(context.Context, ...uint) ([]*Menu, error)
	// 指定角色ID菜单及父级菜单列表
	ListMenuAndParentByIDs(context.Context, ...uint) ([]*Menu, error)
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
	uc.log.WithContext(ctx).Debugf("RoleCreate: %v", g.Name)
	return uc.biz.roleRepo.Save(ctx, g)
}

// ListByIDs 获取指定角色ID集合
func (uc *RoleUsecase) ListByIDs(ctx context.Context, id ...uint) (roles []*Role, err error) {
	roles, _ = uc.biz.roleRepo.ListPage(ctx, &pagination.Pagination{Nopaging: true, OrderBy: map[string]bool{"sort": true}})
	return
}

// Update 修改角色
func (uc *RoleUsecase) Update(ctx context.Context, g *Role) error {
	uc.log.WithContext(ctx).Debugf("RoleUpdate: %v", g)

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
	uc.log.WithContext(ctx).Debugf("RoleUpdateRoleState: %v", g)

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
	uc.log.WithContext(ctx).Debugf("RoleListAll")
	return uc.biz.roleRepo.ListPage(ctx, pagination.NewPagination(
		pagination.WithNopaging(),
		pagination.WithOrderBy(map[string]bool{"sort": false, "id": true}),
	))
}

// List 角色列表分页
func (uc *RoleUsecase) ListPage(ctx context.Context, paging *pagination.Pagination) ([]*Role, int64) {
	uc.log.WithContext(ctx).Debugf("RoleListPage")
	return uc.biz.roleRepo.ListPage(ctx, paging)
}

// GetID 根据角色ID角色
func (uc *RoleUsecase) GetID(ctx context.Context, g *Role) (*Role, error) {
	uc.log.WithContext(ctx).Debugf("RoleGetID: %v", g)
	return uc.biz.roleRepo.FindByID(ctx, g.ID)
}

// Delete 根据角色ID删除角色
func (uc *RoleUsecase) Delete(ctx context.Context, g *Role) error {
	uc.log.WithContext(ctx).Debugf("RoleDelete: %v", g)
	return uc.biz.roleRepo.Delete(ctx, g)
}

// ListMenuByID 获取角色菜单
func (uc *RoleUsecase) ListMenuByID(ctx context.Context, g *Role) ([]*Menu, error) {
	uc.log.WithContext(ctx).Debugf("RoleListMenuByID: %v", g)
	return uc.biz.roleRepo.ListMenuByIDs(ctx, g.ID)
}

// HandleMenu 绑定菜单
func (uc *RoleUsecase) HandleMenu(ctx context.Context, g *Role) error {
	uc.log.WithContext(ctx).Debugf("RoleHandleMenu: %v", g)
	return uc.biz.roleRepo.HandleMenu(ctx, g)
}

// ListDeptByID 获取角色部门列表
func (uc *RoleUsecase) ListDeptByID(ctx context.Context, g *Role) ([]*Dept, error) {
	uc.log.WithContext(ctx).Debugf("RoleListDeptByID: %v", g)
	return uc.biz.roleRepo.ListDeptByIDs(ctx, g.ID)
}

// GetDataScopeByID 获取角色数据范围
func (uc *RoleUsecase) GetDataScopeByID(ctx context.Context, g *Role) (*Role, error) {
	uc.log.WithContext(ctx).Debugf("RoleGetDataScopeByID: %v", g)
	role, err := uc.biz.roleRepo.FindByID(ctx, g.ID)
	if err != nil {
		return nil, err
	}
	if role.DataScope == int32(pb.RoleDataScope_ROLE_DATA_SCOPE_DEPT_CUSTOM) {
		role.Depts, _ = uc.biz.roleRepo.ListDeptByIDs(ctx, role.ID)
	}
	return role, err
}

// HandleDept 绑定数据范围
func (uc *RoleUsecase) HandleDataScope(ctx context.Context, g *Role) error {
	uc.log.WithContext(ctx).Debugf("RoleHandleDataScope: %v", g)
	role, err := uc.biz.roleRepo.FindByID(ctx, g.ID)
	if err != nil {
		return err
	}
	role.DataScope = g.DataScope
	_, err = uc.biz.roleRepo.Update(ctx, role)
	if err != nil {
		return err
	}
	if g.DataScope == int32(pb.RoleDataScope_ROLE_DATA_SCOPE_DEPT_CUSTOM) && len(g.Depts) > 0 {
		return uc.biz.roleRepo.HandleDept(ctx, g)
	}
	return err
}
