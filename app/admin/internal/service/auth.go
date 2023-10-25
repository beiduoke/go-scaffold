package service

import (
	"context"
	"time"

	v1 "github.com/beiduoke/go-scaffold/api/admin/v1"
	"github.com/beiduoke/go-scaffold/app/admin/internal/biz"
	"github.com/beiduoke/go-scaffold/app/admin/internal/pkg/constant"
	"github.com/beiduoke/go-scaffold/app/admin/internal/pkg/middleware/localize"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ v1.AdminServiceServer = (*AdminService)(nil)

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
func (s *AdminService) Logout(ctx context.Context, in *v1.LogoutRequest) (*v1.LogoutResponse, error) {
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
func (s *AdminService) Login(ctx context.Context, in *v1.LoginRequest) (*v1.LoginResponse, error) {
	req := in.GetAuth()
	if in.GetDomain() == "" {
		return nil, v1.ErrorUserRegisterFail("租户不能为空")
	}
	result, err := s.authCase.Login(ctx, &biz.User{Name: req.GetAccount(), Phone: req.GetAccount(), Password: req.GetPassword(), Domain: &biz.Domain{Code: in.GetDomain()}})
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
func (s *AdminService) Register(ctx context.Context, in *v1.RegisterRequest) (*v1.RegisterResponse, error) {
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
