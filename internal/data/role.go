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
	menu biz.MenuRepo
}

// NewRoleRepo .
func NewRoleRepo(logger log.Logger, data *Data, menu biz.MenuRepo) biz.RoleRepo {
	return &RoleRepo{
		data: data,
		log:  log.NewHelper(logger),
		menu: menu,
	}
}

func (r *RoleRepo) toModel(d *biz.Role) *SysRole {
	if d == nil {
		return nil
	}
	sysData := &SysRole{
		Name:          d.Name,
		DefaultRouter: d.DefaultRouter,
		State:         d.State,
		ParentID:      d.ParentID,
		Remarks:       d.Remarks,
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
		ID:            d.ID,
		CreatedAt:     d.CreatedAt,
		UpdatedAt:     d.UpdatedAt,
		Name:          d.Name,
		ParentID:      d.ParentID,
		DefaultRouter: d.DefaultRouter,
		Sort:          d.Sort,
		State:         d.State,
		Remarks:       d.Remarks,
	}
}

func (r *RoleRepo) Save(ctx context.Context, g *biz.Role) (*biz.Role, error) {
	d := r.toModel(g)
	d.DomainID = r.data.DomainID(ctx)
	result := r.data.DB(ctx).Omit(clause.Associations).Create(d).Error
	return r.toBiz(d), result
}

func (r *RoleRepo) Update(ctx context.Context, g *biz.Role) (*biz.Role, error) {
	d := r.toModel(g)
	result := r.data.DBD(ctx).Model(d).Select("*").Omit("CreatedAt").Updates(d)
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

func (r *RoleRepo) ListPage(ctx context.Context, handler pagination.PaginationHandler) (roles []*biz.Role, total int64) {
	db := r.data.DBD(ctx).Model(&SysRole{})
	sysRoles := []*SysRole{}
	// 查询条件
	for _, v := range handler.GetConditions() {
		db = db.Where(v.Query, v.Args...)
	}
	// 排序
	for _, v := range handler.GetOrders() {
		db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: v.Column}, Desc: v.Desc})
	}

	if !handler.GetNopaging() {
		db = db.Count(&total).Offset(handler.GetPageOffset())
	}

	result := db.Limit(int(handler.GetPageSize())).Find(&sysRoles)
	if result.Error != nil {
		return nil, 0
	}

	for _, v := range sysRoles {
		roles = append(roles, r.toBiz(v))
	}

	if handler.GetNopaging() {
		total = int64(len(roles))
	}

	return roles, total
}

func (r *RoleRepo) HandleMenu(ctx context.Context, g *biz.Role) error {
	sysRole := r.toModel(g)
	err := r.data.DB(ctx).Debug().Model(sysRole).Association("Menus").Clear()
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

// 处理角色绑定
func (r *RoleRepo) HandleResource(ctx context.Context, g *biz.Role) error {
	domain := r.data.Domain(ctx)

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
	// 删除角色域下所有权限
	for _, v := range r.data.enforcer.GetPermissionsForUser(role, domain) {
		if _, err := r.data.enforcer.DeletePermissionForUser(role, v[1:]...); err != nil {
			r.log.Errorf("删除casbin角色领域下权限失败 %v", err)
		}
	}
	// 根据最新资源重新绑定
	rules := make([][]string, 0, len(g.Resources))
	for _, v := range g.Resources {
		rules = append(rules, []string{domain, v.Path, v.Method})
	}
	_, err := r.data.enforcer.AddPermissionsForUser(role, rules...)
	return err
}

// 获取指定权限菜单列表
func (r *RoleRepo) ListMenuByIDs(ctx context.Context, ids ...uint) ([]*biz.Menu, error) {
	var roleMenus []*SysRoleMenu
	db := r.data.DBD(ctx).Model(&SysRoleMenu{})
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

// 获取指定权限菜单列表-返回父级菜单
func (r *RoleRepo) ListMenuAndParentByIDs(ctx context.Context, ids ...uint) ([]*biz.Menu, error) {
	var roleMenus []*SysRoleMenu
	db := r.data.DB(ctx).Model(&SysRoleMenu{}).Debug()
	result := db.Find(&roleMenus, "sys_role_id in ?", ids)
	if err := result.Error; err != nil {
		return nil, err
	}
	bizAllMenus, _ := r.menu.ListAll(ctx)
	menuIds := []uint{}
	for _, v := range roleMenus {
		menuIds = append(menuIds, v.MenuID)
	}
	return menuRecursiveParent(bizAllMenus, menuIds...), nil
}
