package admin

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/admin/v1"
	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/internal/pkg/middleware/localize"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ v1.AdminServer = (*AdminService)(nil)

var (
	domain       = "domain"
	loginMessage = &i18n.Message{
		Description: "login",
		ID:          "Login",
		One:         "Login {{.Name}} {{.Password}}",
		Other:       "Login {{.Name}} {{.Password}}",
	}
	registerMessage = &i18n.Message{
		Description: "register",
		ID:          "Register",
		One:         "Register {{.Name}} {{.Password}}",
		Other:       "Register {{.Name}} {{.Password}}",
	}
)

// Logout 退出登录
func (s *AdminService) Logout(ctx context.Context, in *emptypb.Empty) (*v1.LogoutReply, error) {
	err := s.authCase.Logout(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.LogoutReply{
		Success: true,
		Message: "退出成功",
	}, nil
}

// MiddlePassLogin 中台密码登录
func (s *AdminService) MiddlePassLogin(ctx context.Context, in *v1.PassLoginReq) (*v1.LoginReply, error) {
	auth := in.GetAuth()
	res, err := s.authCase.MiddlePassLogin(ctx, &biz.User{Name: auth.GetAccount(), Password: auth.GetPassword()})
	if err != nil {
		return nil, v1.ErrorUserLoginFail("账号 %s 登录失败：%v", auth.GetAccount(), err)
	}
	var expireTime *timestamppb.Timestamp
	if res.ExpiresAt != nil {
		expireTime = timestamppb.New(*res.ExpiresAt)
	}
	return &v1.LoginReply{
		Token:      res.Token,
		ExpireTime: expireTime,
	}, nil
}

// PassLogin 密码登录
func (s *AdminService) PassLogin(ctx context.Context, in *v1.PassLoginReq) (*v1.LoginReply, error) {
	auth := in.GetAuth()
	res, err := s.authCase.PassLogin(ctx, &biz.User{Name: auth.GetAccount(), Password: auth.GetPassword()})
	if err != nil {
		return nil, v1.ErrorUserLoginFail("账号 %s 登录失败：%v", auth.GetAccount(), err)
	}
	var expireTime *timestamppb.Timestamp
	if res.ExpiresAt != nil {
		expireTime = timestamppb.New(*res.ExpiresAt)
	}
	return &v1.LoginReply{
		Token:      res.Token,
		ExpireTime: expireTime,
	}, nil
}

// Login 领域登录
func (s *AdminService) LoginDomain(ctx context.Context, in *v1.LoginDomainReq) (*v1.LoginReply, error) {
	auth := in.GetAuth()
	if in.GetDomain() == "" {
		return nil, v1.ErrorUserRegisterFail("领域不能为空")
	}
	res, err := s.authCase.LoginNamePassword(ctx, in.GetDomain(), &biz.User{Name: auth.GetName(), Password: auth.GetPassword()})
	if err != nil {
		return nil, v1.ErrorUserLoginFail("用户 %s 登录失败：%v", auth.GetName(), err)
	}

	localizer := localize.FromContext(ctx)
	_, err = localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: loginMessage,
		TemplateData: map[string]interface{}{
			"Name":     auth.GetName(),
			"Password": auth.GetPassword(),
		},
	})
	if err != nil {
		return nil, err
	}

	var expireTime *timestamppb.Timestamp
	if res.ExpiresAt != nil {
		expireTime = timestamppb.New(*res.ExpiresAt)
	}
	return &v1.LoginReply{
		Token:      res.Token,
		ExpireTime: expireTime,
	}, nil
}

// Register 领域注册
func (s *AdminService) RegisterDomain(ctx context.Context, in *v1.RegisterDomainReq) (*v1.RegisterReply, error) {
	auth := in.GetAuth()
	if in.GetDomain() == "" {
		return nil, v1.ErrorUserRegisterFail("领域不能为空")
	}
	_, err := s.authCase.RegisterNamePassword(ctx, in.GetDomain(), &biz.User{Name: auth.GetName(), Password: auth.GetPassword()})
	if err != nil {
		return nil, v1.ErrorUserRegisterFail("用户 %s 注册失败: %v", auth.GetName(), err.Error())
	}

	localizer := localize.FromContext(ctx)
	_, err = localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: registerMessage,
		TemplateData: map[string]interface{}{
			"Name":     auth.GetName(),
			"Password": auth.GetPassword(),
		},
	})
	if err != nil {
		return nil, err
	}

	return &v1.RegisterReply{
		Success: true,
		Message: "注册成功",
	}, nil
}
