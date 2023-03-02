package data

import (
	"context"
	"time"

	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/pkg/auth"
)

const cacheKeyToken string = "%d-%d"

func (r *UserRepo) SetLoginCache(ctx context.Context, claims *auth.AuthClaims, g *biz.User) error {
	// result := r.data.rdb.Set(ctx, claims.Token(), claims.Token, time.Until(claims.ExpiresAt()))
	// return result.Err()
	return nil
}

func (r *UserRepo) SetTokenCache(ctx context.Context, uuid string, token string, exp time.Duration) error {
	return r.data.rdb.Set(ctx, uuid, token, exp).Err()
}

func (r *UserRepo) GetTokenCache(ctx context.Context, uuid string) error {
	result := r.data.rdb.Get(ctx, uuid)
	return result.Err()
}
