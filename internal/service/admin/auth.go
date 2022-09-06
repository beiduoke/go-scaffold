package admin

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/admin/v1"
	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/nicksnyder/go-i18n/v2/i18n"
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

// Login 登录
func (s *AdminService) Login(ctx context.Context, in *v1.LoginReq) (*v1.LoginReply, error) {
	auth := in.GetAuth()
	if in.GetDomain() == "" {
		return nil, v1.ErrorUserRegisterFail("Domain不能为空")
	}
	res, err := s.ac.LoginNamePassword(ctx, in.GetDomain(), &biz.User{Name: auth.GetName(), Password: auth.GetPassword()})
	if err != nil {
		return nil, v1.ErrorUserLoginFail("用户 %s 登录失败：%v", auth.GetName(), err)
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

// Register 注册
func (s *AdminService) Register(ctx context.Context, in *v1.RegisterReq) (*v1.RegisterReply, error) {
	auth := in.GetAuth()
	if in.GetDomain() == "" {
		return nil, v1.ErrorUserRegisterFail("Domain不能为空")
	}
	_, err := s.ac.RegisterNamePassword(ctx, in.GetDomain(), &biz.User{Name: auth.GetName(), Password: auth.GetPassword()})
	if err != nil {
		return nil, v1.ErrorUserRegisterFail("用户 %s 注册失败: %v", auth.GetName(), err.Error())
	}

	return &v1.RegisterReply{
		Success: true,
	}, nil
}
