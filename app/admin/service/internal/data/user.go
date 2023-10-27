package data

import (
	"context"
	"errors"

	"github.com/beiduoke/go-scaffold/app/admin/service/internal/biz"
	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ biz.UserRepo = (*UserRepo)(nil)

type UserRepo struct {
	data   *Data
	log    *log.Helper
	menu   MenuRepo
	domain DomainRepo
	role   RoleRepo
	post   PostRepo
	dept   DeptRepo
}

// NewUserRepo .
func NewUserRepo(logger log.Logger, data *Data, domainRepo biz.DomainRepo, roleRepo biz.RoleRepo, postRepo biz.PostRepo, menuRepo biz.MenuRepo, deptRepo biz.DeptRepo) biz.UserRepo {
	return &UserRepo{
		data:   data,
		log:    log.NewHelper(logger),
		domain: *(domainRepo.(*DomainRepo)),
		role:   *(roleRepo.(*RoleRepo)),
		post:   *(postRepo.(*PostRepo)),
		menu:   *(menuRepo.(*MenuRepo)),
		dept:   *(deptRepo.(*DeptRepo)),
	}
}

func toUserModel(d *biz.User) *SysUser {
	if d == nil {
		return nil
	}
	u := &SysUser{
		DomainModel: DomainModel{ID: d.ID, CreatedAt: d.CreatedAt, UpdatedAt: d.UpdatedAt},
		Name:        d.Name,
		NickName:    d.NickName,
		RealName:    d.RealName,
		Avatar:      d.Avatar,
		Password:    d.Password,
		Birthday:    d.Birthday,
		Gender:      d.Gender,
		Phone:       d.Phone,
		Email:       d.Email,
		DeptID:      d.DeptID,
		State:       d.State,
		Remarks:     d.Remarks,
	}
	for _, v := range d.Roles {
		u.Roles = append(u.Roles, *toRoleModel(v))
	}
	for _, v := range d.Posts {
		u.Posts = append(u.Posts, *toPostModel(v))
	}
	return u
}

func toUserBiz(d *SysUser) *biz.User {
	if d == nil {
		return nil
	}
	u := &biz.User{
		CreatedAt:     d.CreatedAt,
		UpdatedAt:     d.UpdatedAt,
		ID:            d.ID,
		Name:          d.Name,
		Avatar:        d.Avatar,
		NickName:      d.NickName,
		RealName:      d.RealName,
		Password:      d.Password,
		Birthday:      d.Birthday,
		Gender:        d.Gender,
		Phone:         d.Phone,
		Email:         d.Email,
		State:         d.State,
		Remarks:       d.Remarks,
		DeptID:        d.DeptID,
		DomainID:      d.DomainID,
		LastUseRoleID: d.LastUseRoleID,
		LastLoginAt:   d.LastLoginAt,
	}
	for _, v := range d.Roles {
		u.Roles = append(u.Roles, toRoleBiz(&v))
	}
	for _, v := range d.Posts {
		u.Posts = append(u.Posts, toPostBiz(&v))
	}
	if d.Dept != nil {
		u.Dept = toDeptBiz(d.Dept)
	}
	if d.LastUseRole != nil {
		u.LastUseRole = toRoleBiz(d.LastUseRole)
	}
	if d.Domain != nil {
		u.Domain = toDomainBiz(d.Domain)
	}
	return u
}

func (r *UserRepo) Save(ctx context.Context, g *biz.User) (bizUser *biz.User, err error) {
	d := toUserModel(g)
	d.DomainID = r.data.CtxDomainID(ctx)
	// d.ID = uint(r.data.sf.Generate())
	err = r.data.InTx(ctx, func(ctx context.Context) error {
		result := r.data.DB(ctx).Create(d)
		if err = result.Error; err != nil {
			return err
		}
		roleIds := make([]string, 0, len(g.Roles))
		for _, v := range d.Roles {
			roleIds = append(roleIds, convert.UnitToString(v.ID))
		}
		return r.data.CasbinUserSetRole(ctx, r.data.CtxAuthUser(ctx).GetDomain(), g.GetID(), roleIds...)
	})
	return bizUser, err
}

func (r *UserRepo) Update(ctx context.Context, g *biz.User) (*biz.User, error) {
	d := toUserModel(g)
	err := r.data.InTx(ctx, func(ctx context.Context) error {
		userModel := *d
		err := r.data.DB(ctx).Model(&userModel).Association("Roles").Clear()
		if err != nil {
			return err
		}
		err = r.data.DB(ctx).Model(&userModel).Association("Posts").Clear()
		if err != nil {
			return err
		}
		result := r.data.DB(ctx).Model(d).Select("*").Scopes(DBScopesOmitUpdate("LastUseRoleID", "LastLoginAt", "LastLoginIP", "Password", "PasswordSalt")).Updates(d)
		if err := result.Error; err != nil {
			return err
		}
		roleIds := make([]string, 0, len(g.Roles))
		for _, v := range d.Roles {
			roleIds = append(roleIds, convert.UnitToString(v.ID))
		}
		return r.data.CasbinUserSetRole(ctx, r.data.CtxAuthUser(ctx).GetDomain(), g.GetID(), roleIds...)
	})
	return toUserBiz(d), err
}

func (r *UserRepo) FindByID(ctx context.Context, id uint) (*biz.User, error) {
	user := SysUser{}
	result := r.data.DBD(ctx).Last(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return toUserBiz(&user), nil
}

func (r *UserRepo) ListAll(ctx context.Context) ([]*biz.User, error) {
	return nil, nil
}

func (r *UserRepo) Delete(ctx context.Context, g *biz.User) error {
	return r.data.InTx(ctx, func(ctx context.Context) error {
		result := r.data.DB(ctx).Delete(toUserModel(g))
		if err := result.Error; err != nil {
			return err
		}
		_, err := r.data.enforcer.DeleteUser(g.GetID())
		return err
	})
}

func (r *UserRepo) ListPage(ctx context.Context, paging *pagination.Pagination) (users []*biz.User, total int64) {
	db := r.data.DBD(ctx).Model(&SysUser{})
	sysUsers := []*SysUser{}

	// 查询条件
	if name, ok := paging.Query["name"].(string); ok && name != "" {
		if name != "" {
			db = db.Where("name LIKE ?", name+"%")
		}
	}

	if deptId, ok := paging.Query["deptId"].(int32); ok && deptId > 0 {
		deptIds := []uint{uint(deptId)}
		depts, _ := r.dept.ListAll(ctx)
		for _, v := range r.dept.LinkedChildren(depts, uint(deptId)) {
			deptIds = append(deptIds, v.ID)
		}
		db = db.Where("dept_id", deptIds)
	}

	// 排序
	if createdBy, ok := paging.OrderBy["createdAt"]; ok {
		db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: "created_at"}, Desc: createdBy})
	}

	if idBy, ok := paging.OrderBy["id"]; ok {
		db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: "id"}, Desc: idBy})
	}

	// 预加载识别
	if _, ok := paging.Query["preloadPosts"]; ok {
		db = db.Preload("Posts")
	}
	if _, ok := paging.Query["preloadRoles"]; ok {
		db = db.Preload("Roles")
	}
	if _, ok := paging.Query["preloadDept"]; ok {
		db = db.Preload("Dept")
	}

	if !paging.Nopaging {
		db = db.Count(&total).Offset(pagination.GetPageOffset(paging.Page, paging.PageSize))
	}

	result := db.Limit(int(paging.PageSize)).Find(&sysUsers)
	if result.Error != nil {
		return nil, 0
	}

	for _, v := range sysUsers {
		users = append(users, toUserBiz(v))
	}

	if paging.Nopaging {
		total = int64(len(users))
	}

	return users, total
}

func (r *UserRepo) FindByName(ctx context.Context, s string) (*biz.User, error) {
	user := SysUser{}
	result := r.data.DBD(ctx).Last(&user, "name = ?", s)
	if result.Error != nil {
		return nil, result.Error
	}
	return toUserBiz(&user), nil
}

func (r *UserRepo) FindByPhone(ctx context.Context, s string) (*biz.User, error) {
	user := SysUser{}
	result := r.data.DBD(ctx).Last(&user, "phone = ?", s)
	if result.Error != nil {
		return nil, result.Error
	}
	return toUserBiz(&user), nil
}

func (r *UserRepo) FindByEmail(ctx context.Context, s string) (*biz.User, error) {
	user := SysUser{}
	result := r.data.DBD(ctx).Last(&user, "email = ?", s)
	if result.Error != nil {
		return nil, result.Error
	}
	return toUserBiz(&user), nil
}

func (r *UserRepo) ListByName(ctx context.Context, s string) ([]*biz.User, error) {
	sysUsers, bizUsers := []*SysUser{}, []*biz.User{}
	result := r.data.DBD(ctx).Find(&sysUsers, "name LIKE ?", "%"+s)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, v := range sysUsers {
		bizUsers = append(bizUsers, toUserBiz(v))
	}
	return bizUsers, nil
}

func (r *UserRepo) ListByPhone(ctx context.Context, s string) ([]*biz.User, error) {
	sysUsers, bizUsers := []*SysUser{}, []*biz.User{}
	result := r.data.DBD(ctx).Find(&sysUsers, "phone LIKE ?", "%"+s)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, v := range sysUsers {
		bizUsers = append(bizUsers, toUserBiz(v))
	}
	return bizUsers, nil
}

func (r *UserRepo) ListByEmail(ctx context.Context, s string) ([]*biz.User, error) {
	sysUsers, bizUsers := []*SysUser{}, []*biz.User{}
	result := r.data.DBD(ctx).Find(&sysUsers, "email LIKE ?", "%"+s)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, v := range sysUsers {
		bizUsers = append(bizUsers, toUserBiz(v))
	}
	return bizUsers, nil
}

// HandleDomainRole 绑定权限
func (r *UserRepo) HandleRole(ctx context.Context, g *biz.User) error {
	sysUser := toUserModel(g)
	err := r.data.DB(ctx).Model(&sysUser).Debug().Association("Roles").Replace(&sysUser.Roles)
	if err != nil {
		return err
	}
	ctxDomain := r.data.CtxAuthUser(ctx).GetDomain()
	if _, err := r.data.enforcer.DeleteRolesForUserInDomain(g.GetID(), ctxDomain); err != nil {
		return err
	}
	rules := make([][]string, 0, len(g.Roles))
	for _, r := range g.Roles {
		rules = append(rules, []string{g.GetID(), r.GetID(), ctxDomain})
	}
	_, err = r.data.enforcer.AddGroupingPolicies(rules)
	return err
}

// ListRoles 指定用户角色列表
func (r *UserRepo) ListRoles(ctx context.Context, g *biz.User) ([]*biz.Role, error) {
	sysUser := SysUser{}
	result := r.data.DB(ctx).Preload("Roles").Last(&sysUser, g.ID)
	if err := result.Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	bizRoles := make([]*biz.Role, 0)
	for _, v := range sysUser.Roles {
		bizRoles = append(bizRoles, toRoleBiz(&v))
	}
	return bizRoles, nil
}
