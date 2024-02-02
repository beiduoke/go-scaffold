package service

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/admin/interface/v1"
	coreV1 "github.com/beiduoke/go-scaffold/api/core/service/v1"
	"github.com/beiduoke/go-scaffold/app/admin/interface/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

type AuthService struct {
	v1.AuthServiceHTTPServer

	uc  coreV1.UserServiceClient
	atr *data.AuthTokenRepo
	log *log.Helper
}

func NewAuthService(logger log.Logger, uc coreV1.UserServiceClient, atr *data.AuthTokenRepo) *AuthService {
	l := log.NewHelper(log.With(logger, "module", "user/service/admin-service"))
	return &AuthService{
		log: l,
		uc:  uc,
		atr: atr,
	}
}

// Login 登陆
func (s *AuthService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	if req.GrandType == nil {
		req.GrandType = (*int32)(v1.GrandType_GRAND_TYPE_PASSWORD.Enum())
	}

	var verifyErr error
	switch (v1.GrandType)(req.GetGrandType()) {
	case v1.GrandType_GRAND_TYPE_PASSWORD:
		_, verifyErr = s.uc.VerifyPassword(ctx, &coreV1.VerifyPasswordRequest{
			Name:     req.GetName(),
			Password: req.GetPassword(),
		})
	case v1.GrandType_GRAND_TYPE_CODE:
		break
	default:
	}
	if verifyErr != nil {
		return &v1.LoginResponse{}, verifyErr
	}

	user, err := s.uc.GetUserByName(ctx, &coreV1.GetUserByNameRequest{Name: req.GetName()})
	if err != nil {
		return &v1.LoginResponse{}, err
	}

	token, err := s.atr.GenerateToken(ctx, user)
	if err != nil {
		return &v1.LoginResponse{}, err
	}

	return &v1.LoginResponse{
		TokenType:    "bearer",
		AccessToken:  token,
		RefreshToken: "",
		Id:           user.GetId(),
		Name:         user.GetName(),
	}, nil
}
