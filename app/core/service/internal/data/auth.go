package data

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/core/service/v1"
	"github.com/beiduoke/go-scaffold/pkg/auth/authz"

	"github.com/go-kratos/kratos/v2/log"
)

type AuthRepo struct {
	data       *Data
	log        *log.Helper
	authorized authz.Authorized
}

// NewAuthRepo .
func NewAuthRepo(data *Data, authorized authz.Authorized, logger log.Logger) *AuthRepo {
	return &AuthRepo{
		data:       data,
		log:        log.NewHelper(logger),
		authorized: authorized,
	}
}

func (r *AuthRepo) Register(ctx context.Context, pb *v1.RegisterRequest) (*v1.RegisterResponse, error) {
	return nil, nil
}

func (r *AuthRepo) IsAuthorized(ctx context.Context, pb *v1.IsAuthorizedRequest) (*v1.IsAuthorizedResponse, error) {
	allowed, err := r.authorized.IsAuthorized(ctx, authz.Subject(pb.Subject), authz.Action(pb.Action), authz.Resource(pb.Resource), authz.Project(pb.Project))
	if err != nil {
		return nil, err
	}
	if !allowed {
		return nil, authz.ErrUnauthorizedAccess
	}
	return nil, err
}
