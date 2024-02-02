package data

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/beiduoke/go-scaffold/pkg/auth/authn"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"

	v1 "github.com/beiduoke/go-scaffold/api/admin/interface/v1"
	coreV1 "github.com/beiduoke/go-scaffold/api/core/service/v1"
)

const userTokenKeyPrefix = "admin_ut_"

type AuthTokenRepo struct {
	data          *Data
	log           *log.Helper
	authenticator authn.Authenticator
}

func NewAuthTokenRepo(data *Data, authenticator authn.Authenticator, logger log.Logger) *AuthTokenRepo {
	l := log.NewHelper(log.With(logger, "module", "user-token/repo/admin-service"))
	return &AuthTokenRepo{
		data:          data,
		log:           l,
		authenticator: authenticator,
	}
}

func (r *AuthTokenRepo) createAccessJwtToken(_ string, userId uint32) string {
	principal := authn.AuthClaims{
		Subject: strconv.FormatUint(uint64(userId), 10),
		Scopes:  make(authn.ScopeSet),
	}

	signedToken, err := r.authenticator.CreateIdentity(context.Background(), principal)
	if err != nil {
		return ""
	}

	return signedToken
}

func (r *AuthTokenRepo) GenerateToken(ctx context.Context, user *coreV1.User) (string, error) {
	token := r.createAccessJwtToken(user.GetName(), user.GetId())
	if token == "" {
		return "", errors.New("create token failed")
	}

	err := r.setToken(ctx, user.GetId(), token)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *AuthTokenRepo) GetToken(ctx context.Context, userId uint32) string {
	return r.getToken(ctx, userId)
}

func (r *AuthTokenRepo) RemoveToken(ctx context.Context, userId uint32) error {
	validToken := r.getToken(ctx, userId)
	if validToken == "" {
		return v1.ErrorAuthTokenNotExist("令牌不存在")
	}

	return r.deleteToken(ctx, userId)
}

func (r *AuthTokenRepo) RemoveUserToken(ctx context.Context, userId uint32) error {
	validToken := r.getToken(ctx, userId)
	if validToken == "" {
		return v1.ErrorAuthTokenNotExist("令牌不存在")
	}

	return r.deleteToken(ctx, userId)
}

func (r *AuthTokenRepo) setToken(ctx context.Context, userId uint32, token string) error {
	key := fmt.Sprintf("%s%d", userTokenKeyPrefix, userId)
	return r.data.rdb.Set(ctx, key, token, 0).Err()
}

func (r *AuthTokenRepo) getToken(ctx context.Context, userId uint32) string {
	key := fmt.Sprintf("%s%d", userTokenKeyPrefix, userId)
	result, err := r.data.rdb.Get(ctx, key).Result()
	if err != nil {
		if err != redis.Nil {
			r.log.Errorf("get redis user token failed: %s", err.Error())
		}
		return ""
	}
	return result
}

func (r *AuthTokenRepo) deleteToken(ctx context.Context, userId uint32) error {
	key := fmt.Sprintf("%s%d", userTokenKeyPrefix, userId)
	return r.data.rdb.Del(ctx, key).Err()
}

func (r *AuthTokenRepo) setAuthUser(ctx context.Context, g *coreV1.User) error {

	return nil
}

func (r *AuthTokenRepo) getAuthUser(ctx context.Context, token string) (*coreV1.User, error) {
	return nil, nil
}

func (r *AuthTokenRepo) deleteAuthUser(ctx context.Context, userId uint32) error {
	key := fmt.Sprintf("%s%d", userTokenKeyPrefix, userId)
	return r.data.rdb.Del(ctx, key).Err()
}
