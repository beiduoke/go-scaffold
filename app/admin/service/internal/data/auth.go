package data

import (
	"context"
	"errors"
	"sort"
	"time"

	"github.com/beiduoke/go-scaffold/app/admin/service/internal/biz"
	auth "github.com/beiduoke/go-scaffold/pkg/auth/authn"
	"github.com/beiduoke/go-scaffold/pkg/util/crypto"
	"github.com/beiduoke/go-scaffold/pkg/util/ip"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

type AuthRepo struct {
	data          *Data
	log           *log.Helper
	authenticator auth.Authenticator
	menu          MenuRepo
	domain        DomainRepo
	role          RoleRepo
}

// NewAuthRepo .
func NewAuthRepo(logger log.Logger, data *Data, authenticator auth.Authenticator, domainRepo biz.DomainRepo, roleRepo biz.RoleRepo, menuRepo biz.MenuRepo) biz.AuthRepo {
	return &AuthRepo{
		data:          data,
		log:           log.NewHelper(logger),
		authenticator: authenticator,
		menu:          *(menuRepo.(*MenuRepo)),
		domain:        *(domainRepo.(*DomainRepo)),
		role:          *(roleRepo.(*RoleRepo)),
	}
}

// Login 密码登录
func (r *AuthRepo) LoginByPassword(ctx context.Context, access, password string) (*biz.LoginResult, error) {
	var (
		now         = time.Now()
		sysUser     = SysUser{}
		sysRoles    = []SysRole{}
		numSysRoles = 0
	)
	domainId := r.data.CtxDomainID(ctx)
	result := r.data.DB(ctx).Where("domain_id", domainId).Preload("Dept").Preload("Roles").Last(&sysUser, "name = ?", access)
	if result.Error != nil {
		return nil, result.Error
	}
	sysRoles, numSysRoles = sysUser.Roles, len(sysUser.Roles)
	if numSysRoles == 0 {
		return nil, errors.New("未指定角色权限")
	}
	if !crypto.CheckPasswordHash(password, sysUser.Password) {
		return nil, errors.New("密码校验失败")
	}

	authClaims := auth.AuthClaims{
		Subject: uuid.NewString(),
		Scopes:  auth.ScopeSet{
			// domainId: true,
		},
	}
	token, err := r.authenticator.CreateIdentity(ctx, authClaims)
	if err != nil {
		return nil, err
	}
	// 判断多点登录
	// 如果已有用户登录设备则踢出，反之
	if !r.data.cfg.Server.Http.Middleware.Auth.GetMultipoint() && r.ExistLoginCache(ctx, sysUser.ID) {
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
		Expiration: r.data.cfg.Server.Http.Middleware.Auth.ExpiresTime.AsDuration(),
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
func (r *AuthRepo) Register(ctx context.Context, g *biz.User) error {
	return nil
}

// Register 注册
func (r *AuthRepo) Logout(ctx context.Context) error {
	return r.DeleteLoginUUIDCache(ctx, r.data.CtxAuthClaimSubject(ctx))
}

// AccessInfo 访问用户登录缓存信息
func (r *AuthRepo) AccessInfo(ctx context.Context) (*biz.User, error) {
	return r.GetLoginUUIDCache(ctx, r.data.CtxAuthClaimSubject(ctx))
}

// AccessRoles 访问用户角色列表
func (r *AuthRepo) AccessRoles(ctx context.Context) ([]*biz.Role, error) {
	bizUser, err := r.GetLoginUUIDCache(ctx, r.data.CtxAuthClaimSubject(ctx))
	if err != nil {
		return nil, err
	}
	return bizUser.Roles, nil
}

// AccessRoleMenus 访问用户角色菜单列表
func (r *AuthRepo) AccessRoleMenus(ctx context.Context) (bizMenus []*biz.Menu, err error) {
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
func (r *AuthRepo) AccessRolePermissions(ctx context.Context) ([]string, error) {
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
