package data

import (
	"context"

	"github.com/beiduoke/go-scaffold-single/internal/biz"
	"github.com/beiduoke/go-scaffold-single/pkg/util/convert"
	"github.com/beiduoke/go-scaffold-single/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm/clause"
)

type RoleRepo struct {
	data *Data
	log  *log.Helper
	menu MenuRepo
	dept DeptRepo
}

// NewRoleRepo .
func NewRoleRepo(logger log.Logger, data *Data, menuRepo biz.MenuRepo, deptRepo biz.DeptRepo) biz.RoleRepo {
	return &RoleRepo{
		data: data,
		log:  log.NewHelper(logger),
		menu: *(menuRepo.(*MenuRepo)),
		dept: *(deptRepo.(*DeptRepo)),
	}
}

func toRoleModel(d *biz.Role) *SysRole {
	if d == nil {
		return nil
	}
	sysData := &SysRole{
		DomainModel:       DomainModel{},
		Name:              d.Name,
		ParentID:          d.ParentID,
		DefaultRouter:     d.DefaultRouter,
		Sort:              d.Sort,
		DataScope:         d.DataScope,
		MenuCheckStrictly: d.MenuCheckStrictly,
		DeptCheckStrictly: d.DeptCheckStrictly,
		State:             d.State,
		Remarks:           d.Remarks,
	}

	sysData.ID = d.ID
	sysData.CreatedAt = d.CreatedAt
	sysData.CreatedAt = d.UpdatedAt
	return sysData
}

func toRoleBiz(d *SysRole) *biz.Role {
	if d == nil {
		return nil
	}
	return &biz.Role{
		ID:                d.ID,
		CreatedAt:         d.CreatedAt,
		UpdatedAt:         d.UpdatedAt,
		Name:              d.Name,
		ParentID:          d.ParentID,
		DefaultRouter:     d.DefaultRouter,
		Sort:              d.Sort,
		DataScope:         d.DataScope,
		MenuCheckStrictly: d.MenuCheckStrictly,
		DeptCheckStrictly: d.DeptCheckStrictly,
		State:             d.State,
		Remarks:           d.Remarks,
	}
}

func (r *RoleRepo) Save(ctx context.Context, g *biz.Role) (*biz.Role, error) {
	d := toRoleModel(g)
	d.DomainID = r.data.CtxDomainID(ctx)
	result := r.data.DB(ctx).Omit(clause.Associations).Create(d).Error
	return toRoleBiz(d), result
}

func (r *RoleRepo) Update(ctx context.Context, g *biz.Role) (*biz.Role, error) {
	d := toRoleModel(g)
	result := r.data.DBD(ctx).Model(d).Omit("UpdatedAt", "DomainID").Updates(d)
	return toRoleBiz(d), result.Error
}

func (r *RoleRepo) FindByName(ctx context.Context, s string) (*biz.Role, error) {
	role := SysRole{}
	result := r.data.DBD(ctx).Last(&role, "name = ?", s)
	if result.Error != nil {
		return nil, result.Error
	}
	return toRoleBiz(&role), nil
}

func (r *RoleRepo) FindByID(ctx context.Context, id uint) (*biz.Role, error) {
	role := SysRole{}
	result := r.data.DBD(ctx).Last(&role, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return toRoleBiz(&role), nil
}

func (r *RoleRepo) ListByIDs(ctx context.Context, id ...uint) (roles []*biz.Role, err error) {
	db := r.data.DBD(ctx).Model(&SysRole{})
	sysRoles := []*SysRole{}

	err = db.Find(&sysRoles, id).Error
	if err != nil {
		return roles, err
	}
	for _, v := range sysRoles {
		roles = append(roles, toRoleBiz(v))
	}
	return
}

func (r *RoleRepo) ListByName(ctx context.Context, name string) ([]*biz.Role, error) {
	sysRoles, bizRoles := []*SysRole{}, []*biz.Role{}
	result := r.data.DBD(ctx).Find(&sysRoles, "name LIKE ?", "%"+name)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, v := range sysRoles {
		bizRoles = append(bizRoles, toRoleBiz(v))
	}
	return bizRoles, nil
}

func (r *RoleRepo) Delete(ctx context.Context, g *biz.Role) error {
	return r.data.InTx(ctx, func(ctx context.Context) error {
		result := r.data.DB(ctx).Delete(toRoleModel(g))
		if err := result.Error; err != nil {
			return err
		}
		success, err := r.data.enforcer.DeleteRole(g.GetID())
		if success {
			_, err = r.data.enforcer.DeletePermission(g.GetID())
		}
		return err
	})
}

func (r *RoleRepo) ListAll(ctx context.Context) ([]*biz.Role, error) {
	return nil, nil
}

func (r *RoleRepo) ListPage(ctx context.Context, paging *pagination.Pagination) (roles []*biz.Role, total int64) {
	db := r.data.DBD(ctx).Model(&SysRole{})
	sysRoles := []*SysRole{}
	// 查询条件
	if name, ok := paging.Query["name"].(string); ok && name != "" {
		if name != "" {
			db = db.Where("name LIKE ?", name+"%")
		}
	}
	// 排序
	if sortBy, ok := paging.OrderBy["sort"]; ok {
		db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: "sort"}, Desc: sortBy})
	}

	if createdBy, ok := paging.OrderBy["createdAt"]; ok {
		db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: "created_at"}, Desc: createdBy})
	}

	if idBy, ok := paging.OrderBy["id"]; ok {
		db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: "id"}, Desc: idBy})
	}

	if !paging.Nopaging {
		db = db.Count(&total).Offset(pagination.GetPageOffset(paging.Page, paging.PageSize))
	}

	result := db.Limit(int(paging.PageSize)).Find(&sysRoles)
	if result.Error != nil {
		return nil, 0
	}

	for _, v := range sysRoles {
		roles = append(roles, toRoleBiz(v))
	}

	if paging.Nopaging {
		total = int64(len(roles))
	}

	return roles, total
}

// 处理角色菜单
func (r *RoleRepo) HandleMenu(ctx context.Context, g *biz.Role) error {
	var (
		menuIds  = make([]uint, 0)
		sysRole  = SysRole{}
		sysMenus = make([]*SysMenu, 0)
		policies = make([]CasbinPolicy, 0)
	)
	if len(g.Menus) > 0 {
		for _, v := range g.Menus {
			sysMenus = append(sysMenus, toMenuModel(v))
			menuIds = append(menuIds, v.ID)
		}
		bizMenus, _ := r.menu.ListByIDs(ctx, menuIds...)
		for _, v := range bizMenus {
			if v.ApiResource != "" {
				policies = append(policies, CasbinPolicy{ID: convert.UnitToString(v.ID), Resource: v.ApiResource})
			}
		}
	}
	return r.data.InTx(ctx, func(ctx context.Context) error {
		sysRole.ID = g.ID
		err := r.data.DB(ctx).Model(&sysRole).Association("Menus").Replace(sysMenus)
		if err != nil {
			return err
		}
		err = r.data.CasbinRoleSetPolicy(ctx, r.data.CtxAuthUser(ctx).GetDomain(), g.GetID(), policies...)
		return err
	})
}

// 获取指定角色部门列表
func (r *RoleRepo) ListDeptByIDs(ctx context.Context, ids ...uint) ([]*biz.Dept, error) {
	sysDepts, bizDepts := []*SysDept{}, []*biz.Dept{}
	result := r.data.DBD(ctx).Joins("right join sys_role_depts on sys_role_depts.sys_dept_id = sys_depts.id").Where("sys_role_depts.sys_role_id", ids).Find(&sysDepts)
	if err := result.Error; err != nil {
		return nil, err
	}
	for _, d := range sysDepts {
		bizDepts = append(bizDepts, toDeptBiz(d))
	}
	return bizDepts, nil
}

func (r *RoleRepo) ListMenuIDByIDs(ctx context.Context, ids ...uint) []uint {
	var roleMenuIds []uint
	result := r.data.DB(ctx).Table("sys_role_menus").Where("sys_role_id", ids).Pluck("sys_menu_id", &roleMenuIds)
	err := result.Error
	if err != nil {
		r.log.Error(err)
		return nil
	}
	return roleMenuIds
}

// 获取指定角色菜单列表
func (r *RoleRepo) ListMenuByIDs(ctx context.Context, ids ...uint) ([]*biz.Menu, error) {
	roleMenuIds := r.ListMenuIDByIDs(ctx, ids...)
	bizAllMenus, err := r.menu.ListAll(ctx)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}
	bizMenus := make([]*biz.Menu, 0, len(roleMenuIds))
	for _, menu := range bizAllMenus {
		for _, menuId := range roleMenuIds {
			if menuId == menu.ID {
				bizMenus = append(bizMenus, menu)
				continue
			}
		}
	}
	return bizMenus, err
}

// ListMenuAndParentByIDs
func (r *RoleRepo) ListMenuAndParentByIDs(ctx context.Context, ids ...uint) ([]*biz.Menu, error) {
	bizAllMenus, err := r.menu.ListAll(ctx)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}
	return menuRecursiveParent(bizAllMenus, r.ListMenuIDByIDs(ctx, ids...)...), nil
}

// 绑定角色部门
func (r *RoleRepo) HandleDept(ctx context.Context, g *biz.Role) error {
	sysRole, sysDepts := toRoleModel(g), []SysDept{}
	for _, v := range g.Depts {
		sysDepts = append(sysDepts, *toDeptModel(v))
	}
	return r.data.DB(ctx).Model(sysRole).Association("Depts").Replace(&sysDepts)
}
