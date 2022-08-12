package web

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/web/v1"
	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/transport/websocket"
)

// WebService is a Web service.
type WebService struct {
	v1.UnimplementedWebServer
	log *log.Helper
	ws  *websocket.Server

	uc *biz.UserUsecase
}

// NewWebService new a Web service.
func NewWebService(logger log.Logger, uc *biz.UserUsecase) *WebService {
	l := log.NewHelper(log.With(logger, "module", "service/web"))
	return &WebService{log: l, uc: uc}
}

// SayHello implements web.WebServer.
func (s *WebService) Login(ctx context.Context, in *v1.LoginReq) (*v1.User, error) {
	_, err := s.uc.NamePasswordLogin(ctx, "", &biz.User{Name: in.GetUserName()})
	if err != nil {
		return nil, err
	}
	return &v1.User{}, nil
}
