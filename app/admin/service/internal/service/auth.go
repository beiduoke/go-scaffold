package service

import (
	"context"
	"strings"
	"time"

	v1 "github.com/beiduoke/go-scaffold/api/admin/service/v1"
	"github.com/beiduoke/go-scaffold/app/admin/service/internal/biz"
	"github.com/beiduoke/go-scaffold/app/admin/service/internal/conf"
	"github.com/beiduoke/go-scaffold/app/admin/service/internal/pkg/constant"
	"github.com/beiduoke/go-scaffold/app/admin/service/internal/pkg/middleware/localize"
	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ v1.AuthServiceServer = (*AuthService)(nil)

// Service is a  service.
type AuthService struct {
	v1.UnimplementedAuthServiceServer
	ac  *conf.Auth
	log *log.Helper
	// dig        *dig.Container
	authCase *biz.AuthUsecase
	userCase *biz.UserUsecase
}

// NewService new a  service.
func NewAuthService(
	logger log.Logger,
	authCase *biz.AuthUsecase,
	userCase *biz.UserUsecase,
) *AuthService {
	l := log.NewHelper(log.With(logger, "module", "service"))
	return &AuthService{
		log:      l,
		authCase: authCase,
		userCase: userCase,
	}
}

var (
	domain       = "domain"
	loginMessage = &i18n.Message{
		Description: "login",
		ID:          "Login",
		One:         "Login {{.Account}} {{.Password}}",
		Other:       "Login {{.Account}} {{.Password}}",
	}
	registerMessage = &i18n.Message{
		Description: "register",
		ID:          "Register",
		One:         "Register {{.Account}} {{.Password}}",
		Other:       "Register {{.Account}} {{.Password}}",
	}
)

// Logout 退出登录
func (s *AuthService) Logout(ctx context.Context, in *v1.LogoutRequest) (*v1.LogoutResponse, error) {
	err := s.authCase.Logout(ctx)

	if err != nil {
		return nil, err
	}
	return &v1.LogoutResponse{
		Type:    constant.HandleType_success.String(),
		Message: "退出成功",
	}, nil
}

// Login 密码登录
func (s *AuthService) LoginByPassword(ctx context.Context, in *v1.LoginByPasswordRequest) (*v1.LoginResponse, error) {
	req := in.GetAuth()
	if in.GetDomain() == "" {
		return nil, v1.ErrorUserRegisterFail("租户不能为空")
	}
	result, err := s.authCase.LoginByPassword(ctx, req.GetAccount(), req.GetPassword())
	if err != nil {
		return nil, v1.ErrorUserLoginFail("账号 %s 登录失败：%v", req.GetAccount(), err)
	}
	var expiresAt time.Time
	if result.ExpiresAt != nil {
		expiresAt = *result.ExpiresAt
	}
	return &v1.LoginResponse{
		Token:      result.Token,
		ExpireTime: timestamppb.New(expiresAt),
	}, err
}

// Register 租户注册
func (s *AuthService) Register(ctx context.Context, in *v1.RegisterRequest) (*v1.RegisterResponse, error) {
	req := in.GetAuth()
	if in.GetDomain() == "" {
		return nil, v1.ErrorUserRegisterFail("租户不能为空")
	}
	err := s.authCase.Register(ctx, &biz.User{Name: req.GetName(), Password: req.GetPassword(), Domain: &biz.Domain{Code: in.GetDomain()}})
	if err != nil {
		return nil, v1.ErrorUserRegisterFail("用户 %s 注册失败: %v", req.GetName(), err.Error())
	}

	localizer := localize.FromContext(ctx)
	_, err = localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: registerMessage,
		TemplateData: map[string]interface{}{
			"Name":     req.GetName(),
			"Password": req.GetPassword(),
		},
	})
	if err != nil {
		return nil, err
	}

	_, err = localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: loginMessage,
		TemplateData: map[string]interface{}{
			"Account":  req.GetName(),
			"Password": req.GetPassword(),
		},
	})
	if err != nil {
		return nil, err
	}

	return &v1.RegisterResponse{
		Type:    constant.HandleType_success.String(),
		Message: "注册成功",
	}, nil
}

// GetUserInfo 用户详情
func (s *AuthService) GetUserInfo(ctx context.Context, in *v1.GetUserInfoRequest) (*v1.GetUserInfoResponse, error) {
	user, err := s.authCase.AccessInfo(ctx)
	if err != nil {
		return nil, v1.ErrorUserNotFound("用户查询失败 %v", err)
	}
	return &v1.GetUserInfoResponse{
		Name:     user.Name,
		NickName: user.NickName,
		RealName: user.RealName,
		Birthday: func() string {
			if user.Birthday != nil {
				return user.Birthday.Format(time.DateOnly)
			}
			return ""
		}(),
		Gender: user.Gender,
		Phone:  user.Phone,
		Email:  user.Email,
		Avatar: user.Avatar,
		State:  user.State,
		Roles: func(roles []*biz.Role) (userRoles []*v1.Role) {
			for _, v := range roles {
				userRoles = append(userRoles, &v1.Role{
					Id:            uint64(v.ID),
					Name:          v.Name,
					DefaultRouter: &v.DefaultRouter,
					Sort:          &v.Sort,
				})
			}
			return userRoles
		}(user.Roles),
	}, nil
}

// ListUserRole 用户角色
func (s *AuthService) ListUserRole(ctx context.Context, in *v1.ListUserRoleRequest) (*v1.ListUserRoleResponse, error) {
	roleModels, err := s.authCase.AccessRoles(ctx)
	if err != nil {
		return nil, v1.ErrorUserRoleFindFail("用户角色失败 %v", err)
	}
	roles := make([]*v1.Role, 0, len(roleModels))
	for _, v := range roleModels {
		roles = append(roles, TransformRole(v))
	}
	return &v1.ListUserRoleResponse{
		Items: roles,
	}, nil
}

// 获取角色菜单路由树形列表
func (s *AuthService) ListUserRoleMenuRouterTree(ctx context.Context, in *v1.ListUserRoleMenuRouterTreeRequest) (*v1.ListUserRoleMenuRouterTreeResponse, error) {
	roleMenus, err := s.authCase.AccessRoleMenus(ctx)
	if err != nil {
		s.log.Debugf("用户菜单查询失败 %v", err)
	}
	treeData := make([]*v1.MenuRouter, 0)
	for _, v := range roleMenus {
		if v.Type == int32(v1.MenuType_MENU_TYPE_ABILITY) {
			continue
		}
		treeData = append(treeData, TransformMenuRouter(v))
	}
	items := convert.ToTree(treeData, in.GetMenuParentId(), func(t *v1.MenuRouter, ts ...*v1.MenuRouter) error {
		redirect := t.Path
		for _, v := range ts {
			if !v.GetMeta().GetHideMenu() {
				if !strings.HasPrefix(redirect, "/") {
					redirect = "/" + redirect
				}
				redirect += "/" + v.Path
				break
			}
		}
		if t.Redirect == nil && redirect != t.Path {
			t.Redirect = &redirect
		}
		t.Children = append(t.Children, ts...)
		return nil
	})
	return &v1.ListUserRoleMenuRouterTreeResponse{
		Items: items,
	}, nil
}

// 获取角色权限列表
func (s *AuthService) ListUserRolePermission(ctx context.Context, in *v1.ListUserRolePermissionRequest) (*v1.ListUserRolePermissionResponse, error) {
	menuModels, _ := s.authCase.AccessRolePermissions(ctx)
	return &v1.ListUserRolePermissionResponse{
		Items: convert.ArrayUnique(menuModels),
	}, nil
}

// 获取角色菜单树形列表
func (s *AuthService) ListUserRoleMenuTree(ctx context.Context, in *v1.ListUserRoleMenuTreeRequest) (*v1.ListUserRoleMenuTreeResponse, error) {
	results, err := s.authCase.AccessRoleMenus(ctx)
	if err != nil {
		s.log.Debugf("用户菜单查询失败 %v", err)
	}
	treeData := make([]*v1.Menu, 0)
	for _, v := range results {
		treeData = append(treeData, TransformMenu(v))
	}
	return &v1.ListUserRoleMenuTreeResponse{
		Items: convert.ToTree(treeData, in.GetMenuParentId(), func(t *v1.Menu, ts ...*v1.Menu) error {
			t.Children = append(t.Children, ts...)
			return nil
		}),
	}, nil
}
