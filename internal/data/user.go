package data

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/internal/conf"
	auth "github.com/beiduoke/go-scaffold/pkg/auth/authn"
	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"github.com/beiduoke/go-scaffold/pkg/util/crypto"
	"github.com/beiduoke/go-scaffold/pkg/util/ip"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

var _ biz.UserRepo = (*UserRepo)(nil)

type UserRepo struct {
	ac            *conf.Auth
	data          *Data
	log           *log.Helper
	menu          MenuRepo
	domain        DomainRepo
	role          RoleRepo
	dept          DeptRepo
	authenticator auth.Authenticator
}

// NewUserRepo .
func NewUserRepo(logger log.Logger, data *Data, ac *conf.Auth, authenticator auth.Authenticator, role biz.RoleRepo, domain biz.DomainRepo, menu biz.MenuRepo, dept biz.DeptRepo) biz.UserRepo {
	return &UserRepo{
		ac:            ac,
		data:          data,
		log:           log.NewHelper(logger),
		role:          *(role.(*RoleRepo)),
		domain:        *(domain.(*DomainRepo)),
		menu:          *(menu.(*MenuRepo)),
		dept:          *(dept.(*DeptRepo)),
		authenticator: authenticator,
	}
}

func (r *UserRepo) toModel(d *biz.User) *SysUser {
	if d == nil {
		return nil
	}
	roles := []SysRole{}
	for _, v := range d.Roles {
		roles = append(roles, *r.role.toModel(v))
	}
	return &SysUser{
		DomainModel: DomainModel{
			ID:        d.ID,
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
		},
		Name:     d.Name,
		Avatar:   d.Avatar,
		NickName: d.NickName,
		RealName: d.RealName,
		Password: d.Password,
		Birthday: d.Birthday,
		Gender:   d.Gender,
		Phone:    d.Phone,
		Email:    d.Email,
		State:    d.State,
		Roles:    roles,
	}
}

func (r *UserRepo) toBiz(d *SysUser) *biz.User {
	if d == nil {
		return nil
	}
	roles := []*biz.Role{}
	for _, v := range d.Roles {
		roles = append(roles, r.role.toBiz(&v))
	}
	return &biz.User{
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
		ID:        d.ID,
		Avatar:    d.Avatar,
		Name:      d.Name,
		NickName:  d.NickName,
		RealName:  d.RealName,
		Password:  d.Password,
		Birthday:  d.Birthday,
		Gender:    d.Gender,
		Phone:     d.Phone,
		Email:     d.Email,
		State:     d.State,
		Roles:     roles,
	}
}

func (r *UserRepo) Save(ctx context.Context, g *biz.User) (*biz.User, error) {
	d := r.toModel(g)
	d.DomainID = r.data.CtxDomainID(ctx)
	// d.ID = uint(r.data.sf.Generate())
	result := r.data.DB(ctx).Omit(clause.Associations).Create(d).Error
	return r.toBiz(d), result
}

func (r *UserRepo) Update(ctx context.Context, g *biz.User) (*biz.User, error) {
	d := r.toModel(g)
	result := r.data.DBD(ctx).Model(d).Select("*").Omit("CreatedAt").Updates(d)
	return r.toBiz(d), result.Error
}

func (r *UserRepo) FindByID(ctx context.Context, id uint) (*biz.User, error) {
	user := SysUser{}
	result := r.data.DBD(ctx).Last(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return r.toBiz(&user), nil
}

func (r *UserRepo) ListAll(ctx context.Context) ([]*biz.User, error) {
	return nil, nil
}

func (r *UserRepo) Delete(ctx context.Context, g *biz.User) error {
	return r.data.DBD(ctx).Delete(r.toModel(g)).Error
}

func (r *UserRepo) ListPage(ctx context.Context, paging *pagination.Pagination) (users []*biz.User, total int64) {
	db := r.data.DBD(ctx).Model(&SysUser{}).Debug()
	sysUsers := []*SysUser{}

	// 查询条件
	if paging.Query != nil {
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
	}

	// 排序
	if createdBy, ok := paging.OrderBy["createdAt"]; ok {
		db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: "created_at"}, Desc: createdBy})
	}

	if idBy, ok := paging.OrderBy["id"]; ok {
		db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: "id"}, Desc: idBy})
	}

	if !paging.Nopaging {
		db = db.Count(&total).Offset(pagination.GetPageOffset(paging.Page, paging.PageSize))
	}

	result := db.Limit(int(paging.PageSize)).Find(&sysUsers)
	if result.Error != nil {
		return nil, 0
	}

	for _, v := range sysUsers {
		users = append(users, r.toBiz(v))
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
	return r.toBiz(&user), nil
}

func (r *UserRepo) FindByPhone(ctx context.Context, s string) (*biz.User, error) {
	user := SysUser{}
	result := r.data.DBD(ctx).Last(&user, "phone = ?", s)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&user), nil
}

func (r *UserRepo) FindByEmail(ctx context.Context, s string) (*biz.User, error) {
	user := SysUser{}
	result := r.data.DBD(ctx).Last(&user, "email = ?", s)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&user), nil
}

func (r *UserRepo) ListByName(ctx context.Context, s string) ([]*biz.User, error) {
	sysUsers, bizUsers := []*SysUser{}, []*biz.User{}
	result := r.data.DBD(ctx).Find(&sysUsers, "name LIKE ?", "%"+s)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, v := range sysUsers {
		bizUsers = append(bizUsers, r.toBiz(v))
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
		bizUsers = append(bizUsers, r.toBiz(v))
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
		bizUsers = append(bizUsers, r.toBiz(v))
	}
	return bizUsers, nil
}

// HandleDomainRole 绑定权限
func (r *UserRepo) HandleRole(ctx context.Context, g *biz.User) error {
	ctxDomain := r.data.CtxAuthUser(ctx).GetDomain()
	rules := make([][]string, 0, len(g.Roles))
	for _, v := range g.Roles {
		rules = append(rules, []string{convert.UnitToString(g.ID), convert.UnitToString(v.ID), ctxDomain, "0"})
		// if _, err := r.data.enforcer.AddRoleForUserInDomain(convert.UnitToString(g.ID), convert.UnitToString(v.ID), domainId); err != nil {
		// 	r.log.Errorf("领域权限绑定失败 %v", err)
		// }
	}
	_, err := r.data.enforcer.AddGroupingPolicies(rules)
	// r.log.Debugf("策略添加 %t %v", success, err)
	return err
}

func (r *UserRepo) ListRoles(ctx context.Context, g *biz.User) ([]*biz.Role, error) {
	rolesIdsStr := r.data.enforcer.GetRolesForUserInDomain(convert.UnitToString(g.ID), convert.UnitToString(g.DomainID))
	rolesIds, sysRoles := make([]uint, 0, len(rolesIdsStr)), make([]SysRole, 0, len(rolesIdsStr))
	for _, v := range rolesIdsStr {
		rolesIds = append(rolesIds, convert.StringToUint(v))
	}
	if len(rolesIds) < 1 {
		return nil, errors.New("未指定角色权限")
	}

	err := r.data.DB(ctx).Where("domain_id = ?", g.DomainID).Find(&sysRoles, rolesIds).Error
	if err != nil {
		return nil, errors.New("角色权限查询失败")
	}
	bizRoles := make([]*biz.Role, 0, len(sysRoles))
	for _, v := range sysRoles {
		bizRoles = append(bizRoles, r.role.toBiz(&v))
	}

	return bizRoles, err
}

// Login 登录
func (r *UserRepo) Login(ctx context.Context, g *biz.User) (*biz.LoginResult, error) {
	sysDomain := SysDomain{}
	if err := r.data.DB(ctx).Last(&sysDomain, "code = ?", g.Domain.Code).Error; err != nil {
		return nil, err
	}
	sysUser := SysUser{}
	result := r.data.DB(ctx).Where("domain_id = ?", sysDomain.ID).Last(&sysUser, "name = ?", g.Name)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Println(g.Password, sysUser.Password)
	if !crypto.CheckPasswordHash(g.Password, sysUser.Password) {
		return nil, errors.New("密码校验失败")
	}

	sysUser.Domain = &sysDomain
	authClaims := auth.AuthClaims{
		Subject: uuid.NewString(),
		Scopes: auth.ScopeSet{
			sysDomain.Code: true,
		},
	}

	token, err := r.authenticator.CreateIdentity(ctx, authClaims)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	r.data.DB(ctx).Model(sysUser).Select("LastLoginAt", "LastLoginIP").Updates(SysUser{
		LastLoginAt: &now,
		LastLoginIP: ip.FormContext(ctx),
	})

	// 判断多点登录
	// 如果已有用户登录设备则踢出，反之
	if !r.ac.Jwt.GetMultipoint() && r.ExistLoginCache(ctx, sysUser.ID) {
		if err := r.DeleteLoginCache(ctx, sysUser.ID); err != nil {
			r.log.Errorf("用户登录缓存删除失败 %v", err)
		}
	}

	bizRoles, err := r.ListRoles(ctx, &biz.User{ID: sysUser.ID, DomainID: sysDomain.ID})
	if err != nil {
		return nil, err
	}
	if len(bizRoles) < 1 {
		return nil, errors.New("未指定角色权限")
	}
	authRoles := make([]AuthRole, 0, len(bizRoles))
	for _, v := range bizRoles {
		authRoles = append(authRoles, AuthRole{
			ID:            v.ID,
			Name:          v.Name,
			DefaultRouter: v.DefaultRouter,
			Sort:          v.Sort,
		})
	}

	loginInfo := UserLoginInfo{
		UUID:  authClaims.Subject,
		Token: token,
		AuthUser: AuthUser{ID: sysUser.ID, DomainID: sysUser.DomainID,
			Name: sysUser.Name, NickName: sysUser.NickName, RealName: sysUser.RealName,
			Avatar: sysUser.Avatar, Birthday: sysUser.Birthday, Gender: sysUser.Gender,
			Phone: sysUser.Phone, Email: sysUser.Email, State: sysUser.State, Remarks: sysUser.Remarks, RoleID: sysUser.LastUseRoleID, Roles: authRoles},
		Expiration: r.ac.Jwt.ExpiresTime.AsDuration(),
	}

	if err := r.SetLoginCache(ctx, loginInfo); err != nil {
		return nil, err
	}

	expires := now.Add(loginInfo.Expiration)
	return &biz.LoginResult{
		Token:     token,
		ExpiresAt: &expires,
	}, nil
}

// Register 注册
func (r *UserRepo) Register(ctx context.Context, g *biz.User) error {
	return nil
}

// Register 注册
func (r *UserRepo) Logout(ctx context.Context) error {
	return r.DeleteLoginCache(ctx, r.data.CtxUserID(ctx))
}

func (r *UserRepo) Info(ctx context.Context) (*biz.User, error) {
	authUser, err := r.GetLoginCache(ctx, r.data.CtxUserID(ctx))
	if err != nil {
		return nil, err
	}
	return &biz.User{
		ID: authUser.ID, Name: authUser.Name,
		Avatar: authUser.Avatar, NickName: authUser.NickName,
		RealName: authUser.RealName, Birthday: authUser.Birthday,
		Gender: authUser.Gender, Phone: authUser.Phone,
		Email: authUser.Email, State: authUser.State,
		DomainID: authUser.DomainID,
		Roles: func(authRoles []AuthRole) (bizRoles []*biz.Role) {
			for _, v := range authRoles {
				bizRoles = append(bizRoles, &biz.Role{
					ID:            v.ID,
					Name:          v.Name,
					Sort:          v.Sort,
					DefaultRouter: v.DefaultRouter,
				})
			}
			return bizRoles
		}(authUser.Roles),
	}, nil
}

func (r *UserRepo) Roles(ctx context.Context) ([]*biz.Role, error) {
	return r.ListRoles(ctx, &biz.User{ID: r.data.CtxUserID(ctx), DomainID: r.data.CtxDomainID(ctx)})
}

func (r *UserRepo) RoleMenus(ctx context.Context) (menus []*biz.Menu, err error) {
	defer func() {
		sort.SliceStable(menus, func(i, j int) bool {
			return int32(menus[i].Sort) < int32(menus[j].Sort)
		})
	}()
	if r.data.HasSuperAdmin(ctx) {
		return r.menu.ListAll(ctx)
	} else if r.data.HasDomainSuperUser(ctx) {
		return r.domain.ListMenuByIDs(ctx, r.data.CtxDomainID(ctx))
	}
	ctxAuthUser := r.data.CtxAuthUser(ctx)
	rolesIdsStr := r.data.enforcer.GetRolesForUserInDomain(ctxAuthUser.GetUser(), ctxAuthUser.GetDomain())
	return r.role.ListMenuByIDs(ctx, convert.ArrayStringToUint(rolesIdsStr)...)
}

func (r *UserRepo) RolePermissions(ctx context.Context) ([]string, error) {
	bizMenus, err := r.RoleMenus(ctx)
	if err != nil {
		return nil, err
	}
	permissions := make([]string, 0)
	for _, v := range bizMenus {
		if v.Permission != "" {
			permissions = append(permissions, v.Permission)
		}
	}
	return permissions, nil
}
