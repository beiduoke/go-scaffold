package data

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/core/service/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type AuthRepo struct {
	data *Data
	log  *log.Helper
}

// NewAuthRepo .
func NewAuthRepo(data *Data, logger log.Logger) *AuthRepo {

	return &AuthRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *AuthRepo) Register(ctx context.Context, pb *v1.RegisterRequest) (*v1.RegisterResponse, error) {
	return nil, nil
}
