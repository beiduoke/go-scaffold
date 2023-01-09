package biz

import (
	"context"
	"time"

	pb "github.com/beiduoke/go-scaffold/api/protobuf"
	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/imdario/mergo"
	"github.com/pkg/errors"
)

// Role is a Role model.
type Role struct {
	ID            uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Name          string
	ParentID      uint
	DefaultRouter string
	Sort          int32
	State         int32
	Remarks       string
	Users         []*User
	Domains       []*Domain
	Menus         []*Menu
	Resources     []*Resource
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
	ListPage(context.Context, pagination.PaginationHandler) ([]*Role, int64)
	HandleMenu(context.Context, *Role) error
	HandleResource(context.Context, *Role) error
	ListMenuByIDs(context.Context, ...uint) ([]*Menu, error)
	ListMenuAndParentByIDs(context.Context, ...uint) ([]*Menu, error)
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
	uc.log.WithContext(ctx).Infof("Create: %v", g.Name)
	return uc.biz.roleRepo.Save(ctx, g)
}

// ListByIDs 获取指定权限角色ID集合
func (uc *RoleUsecase) ListByIDs(ctx context.Context, id ...uint) (roles []*Role, err error) {
	roles, _ = uc.biz.roleRepo.ListPage(ctx, pagination.NewPagination(pagination.WithNopaging(), pagination.WithCondition("id in ?", id)))
	return
}

// Update 修改权限角色
func (uc *RoleUsecase) Update(ctx context.Context, g *Role) error {
	uc.log.WithContext(ctx).Infof("UpdateRole: %v", g)

	role, _ := uc.biz.roleRepo.FindByID(ctx, g.ID)
	if role == nil {
		return errors.New("权限角色未注册")
	}

	if role.Name != g.Name && g.Name != "" {
		name, _ := uc.biz.roleRepo.FindByName(ctx, g.Name)
		if name != nil {
			return errors.New("权限角色名已存在")
		}
	}

	if g.State <= 0 {
		g.State = int32(pb.RoleState_ROLE_STATE_ACTIVE)
	}
	// 新数据合并到源数据
	if err := mergo.Merge(role, *g, mergo.WithOverwriteWithEmptyValue); err != nil {
		return errors.Errorf("数据合并失败：%v", err)
	}
	_, err := uc.biz.roleRepo.Update(ctx, role)
	return err
}

// UpdateState 修改权限角色状态
func (uc *RoleUsecase) UpdateState(ctx context.Context, g *Role) error {
	uc.log.WithContext(ctx).Infof("UpdateRoleState: %v", g)

	role, _ := uc.biz.roleRepo.FindByID(ctx, g.ID)
	if role == nil {
		return errors.New("权限角色未不存在")
	}

	if g.State <= 0 {
		g.State = int32(pb.RoleState_ROLE_STATE_ACTIVE)
	}

	role.State = g.State
	_, err := uc.biz.roleRepo.Update(ctx, role)
	return err
}

// List 权限角色列表全部
func (uc *RoleUsecase) ListAll(ctx context.Context) ([]*Role, int64) {
	uc.log.WithContext(ctx).Infof("RoleList")
	return uc.biz.roleRepo.ListPage(ctx, pagination.NewPagination())
}

// List 权限角色列表分页
func (uc *RoleUsecase) ListPage(ctx context.Context, pageNum, pageSize int32, query map[string]string, order map[string]bool) ([]*Role, int64) {
	uc.log.WithContext(ctx).Infof("RolePage")
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
	return uc.biz.roleRepo.ListPage(ctx, page)
}

// GetID 根据角色ID权限角色
func (uc *RoleUsecase) GetID(ctx context.Context, g *Role) (*Role, error) {
	uc.log.WithContext(ctx).Infof("GetRoleID: %v", g)
	return uc.biz.roleRepo.FindByID(ctx, g.ID)
}

// Delete 根据角色ID删除权限角色
func (uc *RoleUsecase) Delete(ctx context.Context, g *Role) error {
	uc.log.WithContext(ctx).Infof("DeleteRole: %v", g)
	return uc.biz.tm.InTx(ctx, func(ctx context.Context) error {
		if err := uc.biz.roleRepo.Delete(ctx, g); err != nil {
			return err
		}
		_, err := uc.biz.enforcer.DeleteRole(convert.UnitToString(g.ID))
		return err
	})
}

// HandleMenu 获取权限角色菜单
func (uc *RoleUsecase) ListMenuByID(ctx context.Context, g *Role) ([]*Menu, error) {
	uc.log.WithContext(ctx).Infof("ListMenuByIDs: %v", g)
	return uc.biz.roleRepo.ListMenuByIDs(ctx, g.ID)
}

// HandleMenu 绑定菜单
func (uc *RoleUsecase) HandleMenu(ctx context.Context, g *Role) error {
	uc.log.WithContext(ctx).Infof("HandleMenu: %v", g)
	return uc.biz.roleRepo.HandleMenu(ctx, g)
}

// HandleResource 绑定资源
func (uc *RoleUsecase) HandleResource(ctx context.Context, g *Role) error {
	uc.log.WithContext(ctx).Infof("HandleResource: %v", g)
	return uc.biz.roleRepo.HandleResource(ctx, g)
}