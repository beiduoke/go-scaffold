package data

import (
	"context"
	"fmt"
	"time"

	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/internal/biz/auth"
)

const cacheKeyToken string = "%d-%d"

func (r *UserRepo) SetLoginCache(ctx context.Context, claims auth.AuthClaims, g *biz.User) error {
	key := fmt.Sprintf(cacheKeyToken, claims.Token())
	result := r.data.rdb.Set(ctx, key, claims.Token, time.Until(claims.ExpiresAt()))
	return result.Err()
}

func (r *UserRepo) SetTokenCache(ctx context.Context, claims *biz.AuthClaims) error {
	key := fmt.Sprintf(cacheKeyToken, claims.ID, claims.Domain)
	result := r.data.rdb.Set(ctx, key, claims.Token, time.Until(*claims.ExpiresAt))
	return result.Err()
}

func (r *UserRepo) GetTokenCache(ctx context.Context, claims *biz.AuthClaims) error {
	key := fmt.Sprintf(cacheKeyToken, claims.ID, claims.Domain)
	result := r.data.rdb.Get(ctx, key)
	return result.Err()
}
