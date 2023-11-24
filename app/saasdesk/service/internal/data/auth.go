package data

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/saasdesk/service/v1"
	"github.com/beiduoke/go-scaffold/app/saasdesk/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type authRepo struct {
	data *Data
	log  *log.Helper
}

// NewAuthRepo .
func NewAuthRepo(data *Data, logger log.Logger) biz.AuthRepo {
	return &authRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *authRepo) Register(ctx context.Context, pb *v1.RegisterRequest) error {
	return nil
}

func (r *authRepo) Logout(ctx context.Context) error {
	return nil
}

func (r *authRepo) LoginByPassword(ctx context.Context, pb *v1.LoginByPasswordRequest) (*v1.LoginResponse, error) {
	return nil, nil
}

func (r *authRepo) GetAuthInfo(context.Context) (*v1.GetAuthInfoResponse, error) {
	return nil, nil
}
