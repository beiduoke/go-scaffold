package service

import (
	"context"

	v1 "github.com/bedoke/go-scaffold/api/admin/v1"
	"github.com/bedoke/go-scaffold/internal/biz"
)

// AdminService is a Admin service.
type AdminService struct {
	v1.UnimplementedAdminServer

	uc *biz.AdminUsecase
}

// NewAdminService new a Admin service.
func NewAdminService(uc *biz.AdminUsecase) *AdminService {
	return &AdminService{uc: uc}
}

// SayHello implements admin.AdminServer.
func (s *AdminService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateAdmin(ctx, &biz.Admin{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
