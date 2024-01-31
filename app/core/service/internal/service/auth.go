package service

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/core/service/v1"
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

type AuthService struct {
	v1.UnimplementedAuthServiceServer
	log *log.Helper
	ac  *data.AuthRepo
}

func NewAuthService(logger log.Logger, ac *data.AuthRepo) *AuthService {
	l := log.NewHelper(log.With(logger, "module", "auth/service"))
	return &AuthService{
		log: l,
		ac:  ac,
	}
}

func (s *AuthService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterResponse, error) {
	return s.ac.Register(ctx, req)
}

func (s *AuthService) IsAuthorized(ctx context.Context, req *v1.IsAuthorizedRequest) (*v1.IsAuthorizedResponse, error) {
	return s.ac.IsAuthorized(ctx, req)
}
