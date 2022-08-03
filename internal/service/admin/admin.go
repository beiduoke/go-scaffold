package admin

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/admin/v1"
	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/internal/conf"
	myAuthz "github.com/beiduoke/go-scaffold/internal/pkg/authz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/transport/websocket"
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

// NamePasswordLogin 用户密码登录
func (s *AdminService) NamePasswordLogin(ctx context.Context, in *v1.NamePasswordLoginReq) (*v1.LoginReply, error) {
	res, err := s.uc.NamePasswordLogin(ctx, &biz.User{Name: in.GetName(), Password: in.GetPassword()})
	if err != nil {
		return nil, err
	}
	securityUser := myAuthz.NewSecurityUserData(myAuthz.WithID(string(rune(res.ID))))
	securityUser.CreateAccessJwtToken([]byte(s.ac.ApiKey))
	return &v1.LoginReply{}, nil
}
