package admin

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/admin/v1"
	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/internal/conf"
	"github.com/beiduoke/go-scaffold/internal/pkg/websocket"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ v1.AdminServer = (*AdminService)(nil)

var (
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

// 使用i18n包进行国际化
// localizer := localize.FromContext(ctx)
// fmt.Println(localizer)
// helloMsg, err := localizer.Localize(&i18n.LocalizeConfig{
// 	DefaultMessage: loginMessage,
// 	TemplateData: map[string]interface{}{
// 		"Name":     in.Name,
// 		"Password": in.Password,
// 	},
// })
// println(helloMsg, 11111)
// if err != nil {
// 	return nil, err
// }

// AdminService is a Admin service.
type AdminService struct {
	v1.UnimplementedAdminServer
	ws  *websocket.WebsocketService
	log *log.Helper
	ac  *conf.Auth
	uc  *biz.UserUsecase
}

// NewAdminService new a Admin service.
func NewAdminService(logger log.Logger, ac *conf.Auth, ws *websocket.WebsocketService, uc *biz.UserUsecase) *AdminService {
	l := log.NewHelper(log.With(logger, "module", "service/admin"))
	return &AdminService{log: l, ac: ac, uc: uc, ws: ws}
}

// Login 登录
func (s *AdminService) Login(ctx context.Context, in *v1.LoginReq) (*v1.LoginReply, error) {
	res, err := s.uc.NamePasswordLogin(ctx, &biz.User{Name: in.GetName(), Password: in.GetPassword()})
	if err != nil {
		return nil, v1.ErrorUserLoginFail("用户 %s 不存在或密码错误", in.GetName())
	}

	token, expiresAt := s.uc.GenerateToken(res)
	return &v1.LoginReply{
		Token:      token,
		ExpireTime: timestamppb.New(expiresAt),
	}, nil
}

// Register 注册
func (s *AdminService) Register(ctx context.Context, in *v1.RegisterReq) (*v1.RegisterReply, error) {
	_, err := s.uc.NamePasswordRegister(ctx, &biz.User{Name: in.GetName(), Password: in.GetPassword()})
	if err != nil {
		return nil, v1.ErrorUserRegisterFail("用户 %s 注册失败", in.GetName())
	}

	return &v1.RegisterReply{
		Success: true,
	}, nil
}
