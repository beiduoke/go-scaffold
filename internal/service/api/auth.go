package api

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/server/v1"
	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/internal/pkg/middleware/localize"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ v1.ApiServer = (*ApiService)(nil)

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
func (s *ApiService) Logout(ctx context.Context, in *emptypb.Empty) (*v1.LogoutReply, error) {
	err := s.authCase.Logout(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.LogoutReply{
		Success: true,
		Message: "退出成功",
	}, nil
}

// Login 密码登录
func (s *ApiService) Login(ctx context.Context, in *v1.LoginReq) (*v1.LoginReply, error) {
	req := in.GetAuth()
	if in.GetDomain() == "" {
		return nil, v1.ErrorUserRegisterFail("租户不能为空")
	}
	// server := password.NewServer(auth.Claims(auth.NewAuthClaims(auth.WidthAuthSecurityKey(s.ac.GetApiKey()))))
	// claims, err := server.Login(&password.Data{
	// 	Account:  req.GetAccount(),
	// 	Password: req.GetPassword(),
	// })
	claims, err := s.authCase.Login(ctx, &biz.User{Name: req.GetAccount(), Phone: req.GetAccount(), Password: req.GetAccount(), Domain: &biz.Domain{Code: in.GetDomain()}})
	if err != nil {
		return nil, v1.ErrorUserLoginFail("账号 %s 登录失败：%v", req.GetAccount(), err)
	}

	return &v1.LoginReply{
		Token:      claims.Token(),
		ExpireTime: timestamppb.New(claims.ExpiresAt()),
	}, nil

	// res, err := s.authCase.Login(ctx, &biz.User{Name: req.GetAccount(), Password: req.GetPassword()})
	// if err != nil {
	// 	return nil, v1.ErrorUserLoginFail("账号 %s 登录失败：%v", req.GetAccount(), err)
	// }

	// localizer := localize.FromContext(ctx)
	// _, err = localizer.Localize(&i18n.LocalizeConfig{
	// 	DefaultMessage: loginMessage,
	// 	TemplateData: map[string]interface{}{
	// 		"Account":  req.GetAccount(),
	// 		"Password": req.GetPassword(),
	// 	},
	// })
	// if err != nil {
	// 	return nil, err
	// }
	// var expireTime *timestamppb.Timestamp
	// if res.ExpiresAt != nil {
	// 	expireTime = timestamppb.New(*res.ExpiresAt)
	// }
	// return &v1.LoginReply{
	// 	Token:      res.Token,
	// 	ExpireTime: expireTime,
	// }, nil
}

// Register 租户注册
func (s *ApiService) Register(ctx context.Context, in *v1.RegisterReq) (*v1.RegisterReply, error) {
	auth := in.GetAuth()
	if in.GetDomain() == "" {
		return nil, v1.ErrorUserRegisterFail("租户不能为空")
	}
	err := s.authCase.Register(ctx, &biz.User{Name: auth.GetName(), Password: auth.GetPassword(), Domain: &biz.Domain{Code: in.GetDomain()}})
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
