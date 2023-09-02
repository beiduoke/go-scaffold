package data

import (
	"context"
	"errors"
	"sort"
	"time"

	"github.com/beiduoke/go-scaffold/api/common/conf"
	"github.com/beiduoke/go-scaffold/internal/biz"
	auth "github.com/beiduoke/go-scaffold/pkg/auth/authn"
	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"github.com/beiduoke/go-scaffold/pkg/util/crypto"
	"github.com/beiduoke/go-scaffold/pkg/util/ip"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ biz.UserRepo = (*UserRepo)(nil)

type UserRepo struct {
	cfg           *conf.Bootstrap
	data          *Data
	log           *log.Helper
	menu          MenuRepo
	domain        DomainRepo
	role          RoleRepo
	post          PostRepo
	dept          DeptRepo
	authenticator auth.Authenticator
}

// NewUserRepo .
func NewUserRepo(logger log.Logger, cfg *conf.Bootstrap, data *Data, authenticator auth.Authenticator, domainRepo biz.DomainRepo, roleRepo biz.RoleRepo, postRepo biz.PostRepo, menuRepo biz.MenuRepo, deptRepo biz.DeptRepo) biz.UserRepo {
	return &UserRepo{
		data:          data,
		log:           log.NewHelper(logger),
		domain:        *(domainRepo.(*DomainRepo)),
		role:          *(roleRepo.(*RoleRepo)),
		post:          *(postRepo.(*PostRepo)),
		menu:          *(menuRepo.(*MenuRepo)),
		dept:          *(deptRepo.(*DeptRepo)),
		authenticator: authenticator,
	}
}

func (r *UserRepo) toModel(d *biz.User) *SysUser {
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
		u.Roles = append(u.Roles, *r.role.toModel(v))
	}
	for _, v := range d.Posts {
		u.Posts = append(u.Posts, *r.post.toModel(v))
	}
	return u
}

func (r *UserRepo) toBiz(d *SysUser) *biz.User {
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
		u.Roles = append(u.Roles, r.role.toBiz(&v))
	}
	for _, v := range d.Posts {
		u.Posts = append(u.Posts, r.post.toBiz(&v))
	}
	if d.Dept != nil {
		u.Dept = r.dept.toBiz(d.Dept)
	}
	if d.LastUseRole != nil {
		u.LastUseRole = r.role.toBiz(d.LastUseRole)
	}
	if d.Domain != nil {
		u.Domain = r.domain.toBiz(d.Domain)
	}
	return u
}

func (r *UserRepo) Save(ctx context.Context, g *biz.User) (bizUser *biz.User, err error) {
	d := r.toModel(g)
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
	d := r.toModel(g)
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
	return r.toBiz(d), err
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
	return r.data.InTx(ctx, func(ctx context.Context) error {
		result := r.data.DB(ctx).Delete(r.toModel(g))
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
	sysUser := r.toModel(g)
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
		bizRoles = append(bizRoles, r.role.toBiz(&v))
	}
	return bizRoles, nil
}

// Login 登录
func (r *UserRepo) Login(ctx context.Context, g *biz.User) (*biz.LoginResult, error) {
	var (
		now         = time.Now()
		sysDomain   = SysDomain{}
		sysUser     = SysUser{}
		sysRoles    = []SysRole{}
		numSysRoles = 0
	)
	if err := r.data.DB(ctx).Last(&sysDomain, "code = ?", g.Domain.Code).Error; err != nil {
		return nil, err
	}
	result := r.data.DB(ctx).Where("domain_id", sysDomain.ID).Preload("Dept").Preload("Roles").Last(&sysUser, "name = ?", g.Name)
	if result.Error != nil {
		return nil, result.Error
	}
	sysRoles, numSysRoles = sysUser.Roles, len(sysUser.Roles)
	if numSysRoles == 0 {
		return nil, errors.New("未指定角色权限")
	}
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
	// 判断多点登录
	// 如果已有用户登录设备则踢出，反之
	if !r.cfg.Server.Http.Middleware.Auth.GetMultipoint() && r.ExistLoginCache(ctx, sysUser.ID) {
		if err := r.DeleteLoginCache(ctx, sysUser.ID); err != nil {
			r.log.Errorf("用户登录缓存删除失败 %v", err)
		}
	}
	if sysUser.LastUseRoleID <= 0 {
		sysUser.LastUseRoleID = sysRoles[numSysRoles-1].ID
	}
	authRoles := make([]*biz.Role, 0, numSysRoles)
	for _, v := range sysRoles {
		if v.ID == sysUser.LastUseRoleID {
			sysUser.LastUseRole = &v
		}
		authRoles = append(authRoles, &biz.Role{
			ID:            v.ID,
			Name:          v.Name,
			DefaultRouter: v.DefaultRouter,
			Sort:          v.Sort,
		})
	}
	loginInfo := UserLoginInfo{
		UUID:  authClaims.Subject,
		Token: token,
		AuthUser: biz.User{ID: sysUser.ID, DomainID: sysUser.DomainID,
			Name: sysUser.Name, NickName: sysUser.NickName, RealName: sysUser.RealName,
			Avatar: sysUser.Avatar, Birthday: sysUser.Birthday, Gender: sysUser.Gender,
			Phone: sysUser.Phone, Email: sysUser.Email, State: sysUser.State,
			Remarks: sysUser.Remarks, LastUseRoleID: sysUser.LastUseRoleID,
			Roles: authRoles, DeptID: sysUser.DeptID,
			LastUseRole: func() *biz.Role {
				if role := sysUser.LastUseRole; role != nil {
					return &biz.Role{ID: role.ID, Name: role.Name}
				}
				return nil
			}(),
			Dept: func() *biz.Dept {
				if dept := sysUser.Dept; dept != nil {
					return &biz.Dept{ID: dept.ID, Name: dept.Name}
				}
				return nil
			}(),
			LastLoginAt: &now,
		},
		Expiration: r.cfg.Server.Http.Middleware.Auth.ExpiresTime.AsDuration(),
	}
	err = r.data.DB(ctx).Model(sysUser).Select("LastLoginAt", "LastLoginIP", "LastUseRoleID").Updates(SysUser{
		LastLoginAt:   &now,
		LastLoginIP:   ip.FormContext(ctx),
		LastUseRoleID: sysUser.LastUseRoleID,
	}).Error
	if err != nil {
		return nil, err
	}
	if err = r.SetLoginCache(ctx, loginInfo); err != nil {
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
	return r.DeleteLoginUUIDCache(ctx, r.data.CtxAuthClaimSubject(ctx))
}

// AccessInfo 访问用户登录缓存信息
func (r *UserRepo) AccessInfo(ctx context.Context) (*biz.User, error) {
	return r.GetLoginUUIDCache(ctx, r.data.CtxAuthClaimSubject(ctx))
}

// AccessRoles 访问用户角色列表
func (r *UserRepo) AccessRoles(ctx context.Context) ([]*biz.Role, error) {
	bizUser, err := r.GetLoginUUIDCache(ctx, r.data.CtxAuthClaimSubject(ctx))
	if err != nil {
		return nil, err
	}
	return bizUser.Roles, nil
}

// AccessRoleMenus 访问用户角色菜单列表
func (r *UserRepo) AccessRoleMenus(ctx context.Context) (bizMenus []*biz.Menu, err error) {
	defer func() {
		sort.SliceStable(bizMenus, func(i, j int) bool {
			return int32(bizMenus[i].Sort) < int32(bizMenus[j].Sort)
		})
	}()
	bizAllMenus, err := r.menu.ListAll(ctx)
	var accessMenuIds = []uint{}
	if r.data.HasSystemSuperAdmin(ctx) {
		for _, v := range bizAllMenus {
			accessMenuIds = append(accessMenuIds, v.ID)
		}
	} else if r.data.HasDomainSuperUser(ctx) {
		accessMenuIds = r.domain.ListMenuIDByIDs(ctx, r.data.CtxDomainID(ctx))
	} else {
		accessMenuIds = r.role.ListMenuIDByIDs(ctx, r.data.CtxRoleID(ctx))
	}
	return menuRecursiveParent(bizAllMenus, accessMenuIds...), nil
}

// AccessRolePermissions 访问用户角色的权限列表
func (r *UserRepo) AccessRolePermissions(ctx context.Context) ([]string, error) {
	bizMenus, err := r.AccessRoleMenus(ctx)
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
