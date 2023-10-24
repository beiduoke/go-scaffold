package data

import (
	author "github.com/beiduoke/go-scaffold/pkg/auth/authn"
	"github.com/go-kratos/kratos/v2/log"
)

type AuthRepo struct {
	data          *Data
	log           *log.Helper
	authenticator author.Authenticator
}

// NewAuthRepo .
func NewAuthRepo(logger log.Logger, data *Data, authenticator author.Authenticator) *AuthRepo {
	return &AuthRepo{
		data:          data,
		log:           log.NewHelper(logger),
		authenticator: authenticator,
	}
}
