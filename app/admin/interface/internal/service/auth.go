package service

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/admin/interface/v1"
	coreV1 "github.com/beiduoke/go-scaffold/api/core/service/v1"
	"github.com/go-kratos/kratos/v2/log"
)

type AuthService struct {
	v1.AuthServiceHTTPServer

	uc  coreV1.UserServiceClient
	log *log.Helper
}

func NewAuthService(logger log.Logger, uc coreV1.UserServiceClient) *AuthService {
	l := log.NewHelper(log.With(logger, "module", "user/service/admin-service"))
	return &AuthService{
		log: l,
		uc:  uc,
	}
}

// Login 登陆
func (s *AuthService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	if _, err := s.uc.VerifyPassword(ctx, &coreV1.VerifyPasswordRequest{
		UserName: req.GetUsername(),
		Password: req.GetPassword(),
	}); err != nil {
		return &v1.LoginResponse{}, err
	}

	user, err := s.uc.GetUserByUserName(ctx, &coreV1.GetUserByUserNameRequest{UserName: req.GetUsername()})
	if err != nil {
		return &v1.LoginResponse{}, err
	}

	token, err := s.utuc.GenerateToken(ctx, user)
	if err != nil {
		return &v1.LoginResponse{}, err
	}

	return &v1.LoginResponse{
		TokenType:   "bearer",
		AccessToken: token,
		Id:          user.GetId(),
		Username:    user.GetUserName(),
	}, nil
}
