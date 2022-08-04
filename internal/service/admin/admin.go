package admin

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/admin/v1"
	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/internal/conf"
	"github.com/beiduoke/go-scaffold/internal/pkg/middleware/localize"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/tx7do/kratos-transport/transport/websocket"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ v1.AdminServer = (*AdminService)(nil)

// AdminService is a Admin service.
type AdminService struct {
	v1.UnimplementedAdminServer
	log *log.Helper
	ws  *websocket.Server
	ac  *conf.Auth
	uc  *biz.UserUsecase
}

// NewAdminService new a Admin service.
func NewAdminService(logger log.Logger, ac *conf.Auth, uc *biz.UserUsecase) *AdminService {
	l := log.NewHelper(log.With(logger, "module", "service/admin"))
	return &AdminService{log: l, ac: ac, uc: uc}
}

// Login 登录
func (s *AdminService) Login(ctx context.Context, in *v1.LoginReq) (*v1.LoginReply, error) {
	return &v1.LoginReply{}, nil
	// 使用i18n包进行国际化
	localizer := localize.FromContext(ctx)
	helloMsg, err := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			Description: "login",
			ID:          "Login",
			One:         "Hello {{.Name}} {{.Password}}",
			Other:       "Hello {{.Name}} {{.Password}}",
		},
		TemplateData: map[string]interface{}{
			"Name":     in.Name,
			"Password": in.Password,
		},
	})
	println(helloMsg, 11111)
	if err != nil {
		return nil, err
	}
	res, err := s.uc.NamePasswordLogin(ctx, &biz.User{Name: in.GetName(), Password: in.GetPassword()})
	if err != nil {
		return nil, v1.ErrorUserNotFound("用户 %s 不存在或密码错误", in.GetName())
	}

	token, expiresAt := s.uc.GenerateToken(res)
	return &v1.LoginReply{
		Token:      token,
		ExpireTime: timestamppb.New(expiresAt),
	}, nil
}
