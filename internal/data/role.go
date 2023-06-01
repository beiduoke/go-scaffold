package data

import (
	"context"
	"encoding/json"

	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
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
func NewRoleRepo(logger log.Logger, data *Data, menu biz.MenuRepo, dept biz.DeptRepo) biz.RoleRepo {
	return &RoleRepo{
		data: data,
		log:  log.NewHelper(logger),
		menu: *(menu.(*MenuRepo)),
		dept: *(dept.(*DeptRepo)),
	}
}

func (r *RoleRepo) toModel(d *biz.Role) *SysRole {
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

func (r *RoleRepo) toBiz(d *SysRole) *biz.Role {
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
	d := r.toModel(g)
	d.DomainID = r.data.CtxDomainID(ctx)
	result := r.data.DB(ctx).Omit(clause.Associations).Create(d).Error
	return r.toBiz(d), result
}

func (r *RoleRepo) Update(ctx context.Context, g *biz.Role) (*biz.Role, error) {
	d := r.toModel(g)
	result := r.data.DBD(ctx).Model(d).Omit("UpdatedAt", "DomainID").Updates(d)
	return r.toBiz(d), result.Error
}

func (r *RoleRepo) FindByName(ctx context.Context, s string) (*biz.Role, error) {
	role := SysRole{}
	result := r.data.DBD(ctx).Last(&role, "name = ?", s)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&role), nil
}

func (r *RoleRepo) FindByID(ctx context.Context, id uint) (*biz.Role, error) {
	role := SysRole{}
	result := r.data.DBD(ctx).Last(&role, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&role), nil
}

func (r *RoleRepo) ListByIDs(ctx context.Context, id ...uint) (roles []*biz.Role, err error) {
	db := r.data.DBD(ctx).Model(&SysRole{})
	sysRoles := []*SysRole{}

	err = db.Find(&sysRoles, id).Error
	if err != nil {
		return roles, err
	}
	for _, v := range sysRoles {
		roles = append(roles, r.toBiz(v))
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
		bizRoles = append(bizRoles, r.toBiz(v))
	}
	return bizRoles, nil
}

func (r *RoleRepo) Delete(ctx context.Context, g *biz.Role) error {
	return r.data.DBD(ctx).Delete(r.toModel(g)).Error
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
	if createdBy, ok := paging.OrderBy["createdAt"]; ok {
		db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: "created_at"}, Desc: createdBy})
	}

	if idBy, ok := paging.OrderBy["id"]; ok {
		db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: "id"}, Desc: idBy})
	}

	if sortBy, ok := paging.OrderBy["sort"]; ok {
		db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: "sort"}, Desc: sortBy})
	}

	if !paging.Nopaging {
		db = db.Count(&total).Offset(pagination.GetPageOffset(paging.Page, paging.PageSize))
	}

	result := db.Limit(int(paging.PageSize)).Find(&sysRoles)
	if result.Error != nil {
		return nil, 0
	}

	for _, v := range sysRoles {
		roles = append(roles, r.toBiz(v))
	}

	if paging.Nopaging {
		total = int64(len(roles))
	}

	return roles, total
}

// 处理角色菜单
func (r *RoleRepo) HandleMenu(ctx context.Context, g *biz.Role) error {
	sysRole := r.toModel(g)
	err := r.data.DB(ctx).Model(sysRole).Association("Menus").Clear()
	if err != nil {
		return err
	}
	sysRoleMenus := []SysRoleMenu{}
	for _, v := range g.Menus {
		menuButtons := make([]uint, 0, len(v.Buttons))
		for _, m := range v.Buttons {
			menuButtons = append(menuButtons, m.ID)
		}
		buttons, _ := json.Marshal(menuButtons)
		menuParameters := make([]uint, 0, len(v.Parameters))
		for _, m := range v.Parameters {
			menuParameters = append(menuParameters, m.ID)
		}
		parameters, _ := json.Marshal(menuParameters)
		sysRoleMenus = append(sysRoleMenus, SysRoleMenu{
			RoleID:        g.ID,
			MenuID:        v.ID,
			MenuButton:    string(buttons),
			MenuParameter: string(parameters),
		})
	}
	return r.data.DB(ctx).Model(&SysRoleMenu{}).CreateInBatches(&sysRoleMenus, len(g.Menus)).Error
}

// 处理角色资源
func (r *RoleRepo) HandleResource(ctx context.Context, g *biz.Role) error {
	ctxDomain := r.data.CtxAuthUser(ctx).GetDomain()

	var apiRepo = ResourceRepo{}
	var resources []SysResource
	for _, v := range g.Resources {
		resources = append(resources, *apiRepo.toModel(v))
	}

	sysRole := r.toModel(g)
	if err := r.data.DB(ctx).Model(sysRole).Debug().Association("Resources").Replace(&resources); err != nil {
		return err
	}

	role := convert.UnitToString(g.ID)
	// 删除角色域下所有角色
	for _, v := range r.data.enforcer.GetPermissionsForUser(role, ctxDomain) {
		if _, err := r.data.enforcer.DeletePermissionForUser(role, v[1:]...); err != nil {
			r.log.Errorf("删除casbin角色领域下角色失败 %v", err)
		}
	}
	// 根据最新资源重新绑定
	rules := make([][]string, 0, len(g.Resources))
	for _, v := range g.Resources {
		rules = append(rules, []string{ctxDomain, v.Path, v.Method})
	}
	_, err := r.data.enforcer.AddPermissionsForUser(role, rules...)
	return err
}

// 获取指定角色菜单列表
func (r *RoleRepo) ListMenuByIDs(ctx context.Context, ids ...uint) ([]*biz.Menu, error) {
	var roleMenus []*SysRoleMenu
	db := r.data.DB(ctx).Model(&SysRoleMenu{})
	result := db.Find(&roleMenus, "sys_role_id in ?", ids)
	if err := result.Error; err != nil {
		return nil, err
	}
	bizAllMenus, err := r.menu.ListAll(ctx)
	bizMenus := make([]*biz.Menu, 0, len(roleMenus))
	for _, menu := range bizAllMenus {
		for _, authMenu := range roleMenus {
			if authMenu.MenuID == menu.ID {
				bizMenus = append(bizMenus, menu)
				continue
			}
		}
	}
	return bizMenus, err
}

// 获取指定角色部门列表
func (r *RoleRepo) ListDeptByIDs(ctx context.Context, ids ...uint) ([]*biz.Dept, error) {
	sysDepts, bizDepts := []*SysDept{}, []*biz.Dept{}
	result := r.data.DBD(ctx).Joins("right join sys_role_depts on sys_role_depts.sys_dept_id = sys_depts.id").Where("sys_role_depts.sys_role_id", ids).Find(&sysDepts)
	if err := result.Error; err != nil {
		return nil, err
	}
	deptRepo := r.dept
	for _, d := range sysDepts {
		bizDepts = append(bizDepts, deptRepo.toBiz(d))
	}

	return bizDepts, nil
}

// 获取指定角色部门列表
func (r *RoleRepo) ListResourceByIDs(ctx context.Context, ids ...uint) ([]*biz.Resource, error) {
	return nil, nil
}

// 获取指定角色菜单列表-返回父级菜单
func (r *RoleRepo) ListMenuAndParentByIDs(ctx context.Context, ids ...uint) ([]*biz.Menu, error) {
	var sysRoleMenus []*SysRoleMenu
	db := r.data.DB(ctx).Model(&SysRoleMenu{})
	result := db.Find(&sysRoleMenus, "sys_role_id in ?", ids)
	if err := result.Error; err != nil {
		return nil, err
	}
	bizAllMenus, _ := r.menu.ListAll(ctx)
	menuIds := []uint{}
	for _, v := range sysRoleMenus {
		menuIds = append(menuIds, v.MenuID)
	}
	return menuRecursiveParent(bizAllMenus, menuIds...), nil
}

// 绑定角色部门
func (r *RoleRepo) HandleDept(ctx context.Context, g *biz.Role) error {
	sysRole, sysDepts := r.toModel(g), []SysDept{}
	for _, v := range g.Depts {
		sysDepts = append(sysDepts, *r.dept.toModel(v))
	}
	return r.data.DB(ctx).Model(sysRole).Association("Depts").Replace(&sysDepts)
}
