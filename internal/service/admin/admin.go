package admin

import (
	"context"

	v1 "github.com/bedoke/go-scaffold/api/admin/v1"
	"github.com/bedoke/go-scaffold/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/transport/websocket"
)

// AdminService is a Admin service.
type AdminService struct {
	v1.UnimplementedAdminServer
	log *log.Helper
	ws  *websocket.Server

	uc *biz.UserUsecase
}

// NewAdminService new a Admin service.
func NewAdminService(logger log.Logger, uc *biz.UserUsecase) *AdminService {
	l := log.NewHelper(log.With(logger, "module", "service/admin"))
	return &AdminService{log: l, uc: uc}
}

// SayHello implements admin.AdminServer.
func (s *AdminService) Login(ctx context.Context, in *v1.LoginReq) (*v1.User, error) {
	_, err := s.uc.CreateUser(ctx, &biz.User{Name: in.GetUserName()})
	if err != nil {
		return nil, err
	}
	return &v1.User{}, nil
}
