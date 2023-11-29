package service

import (
	"context"

	pb "github.com/beiduoke/go-scaffold/api/saasdesk/service/v1"
	"github.com/beiduoke/go-scaffold/app/saasdesk/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type AuthService struct {
	pb.UnimplementedAuthServiceServer
	log *log.Helper
	ac  biz.AuthRepo
}

func NewAuthService(logger log.Logger, ac biz.AuthRepo) *AuthService {
	l := log.NewHelper(log.With(logger, "module", "auth/service"))
	return &AuthService{
		log: l,
		ac:  ac,
	}
}

func (s *AuthService) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	return &pb.LogoutResponse{}, nil
}
func (s *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return &pb.RegisterResponse{}, nil
}
func (s *AuthService) LoginByPassword(ctx context.Context, req *pb.LoginByPasswordRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{}, nil
}
func (s *AuthService) LoginBySms(ctx context.Context, req *pb.LoginBySmsRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{}, nil
}
func (s *AuthService) LoginByEmail(ctx context.Context, req *pb.LoginByEmailRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{}, nil
}
func (s *AuthService) GetAuthInfo(ctx context.Context, req *pb.GetAuthInfoRequest) (*pb.GetAuthInfoResponse, error) {
	return &pb.GetAuthInfoResponse{}, nil
}
func (s *AuthService) GetAuthProfile(ctx context.Context, req *pb.GetAuthProfileRequest) (*pb.GetAuthProfileResponse, error) {
	return &pb.GetAuthProfileResponse{}, nil
}
