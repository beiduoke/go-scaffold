package data

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"gorm.io/gorm"
)

const (
	cacheHashKeyUser     string = "sys_user"
	cacheStringLoginID   string = "login_uuid:%s"
	cacheHashLoginToken  string = "login_token:%d"
	cacheStringLoginUser string = "login_user:%d"
)

var _ Cache[*biz.User] = (*UserRepo)(nil)

type UserLoginInfo struct {
	UUID       string
	Token      string
	Expiration time.Duration
	AuthUser   biz.User
	Info       map[string]interface{}
}

// SetCache 设置用户缓存
func (r *UserRepo) SetCache(ctx context.Context, g *biz.User) error {
	dataStr, err := json.Marshal(g)
	if err != nil {
		r.log.Errorf("用户缓存失败 %v", err)
		return err
	}
	return r.data.rdb.HSet(ctx, cacheHashKeyUser, convert.UnitToString(g.ID), dataStr).Err()
}

// GetCache 获取用户缓存
func (r *UserRepo) GetCache(ctx context.Context, key string) (bizUser *biz.User) {
	dataStr, err := r.data.rdb.HGet(ctx, cacheHashKeyUser, key).Result()
	if err != nil {
		return nil
	}
	if err := json.Unmarshal([]byte(dataStr), &bizUser); err != nil {
		r.log.Errorf("缓存反序列化失败 %v", err)
	}
	return bizUser
}

// DeleteCache 获取用户缓存
func (r *UserRepo) DeleteCache(ctx context.Context, key string) error {
	return r.data.rdb.HDel(ctx, cacheHashKeyUser, key).Err()
}

// ListAllCache 获取全部缓存数据
func (r *UserRepo) ListAllCache(ctx context.Context) (bizUsers []*biz.User) {
	if l, _ := r.data.rdb.HLen(ctx, cacheHashKeyUser).Result(); l > 0 {
		domainMap, _ := r.data.rdb.HGetAll(ctx, cacheHashKeyUser).Result()
		for _, v := range domainMap {
			bizUser := biz.User{}
			err := json.Unmarshal([]byte(v), &bizUser)
			if err != nil {
				r.log.Errorf("用户缓存反序列失败 %v", err)
				continue
			}
			bizUsers = append(bizUsers, &bizUser)
		}
		return bizUsers
	}

	sysUsers := []SysUser{}
	result := r.data.DB(ctx).Find(&sysUsers)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		r.log.Errorf("用户查询失败 %v", result.Error)
		return nil
	}
	userMap := make(map[string]string)
	for _, v := range sysUsers {
		bizUser := r.toBiz(&v)
		marshalUserJson, err := json.Marshal(bizUser)
		if err != nil {
			r.log.Errorf("用户缓存序列化失败 %v", err)
			continue
		}
		userMap[bizUser.GetID()] = string(marshalUserJson)
	}
	if err := r.data.rdb.HSet(ctx, cacheHashKeyUser, userMap).Err(); err != nil {
		r.log.Errorf("用户缓存失败 %v", err)
	}
	return bizUsers
}

// SetLoginCache 设置登录信息
func (r *UserRepo) SetLoginCache(ctx context.Context, info UserLoginInfo) error {
	// 设置登录用户信息
	dataStr, err := json.Marshal(info.AuthUser)
	if err != nil {
		r.log.Errorf("用户缓存序列化失败 %v", err)
		return err
	}
	err = r.data.rdb.Set(ctx, fmt.Sprintf(cacheStringLoginID, info.UUID), dataStr, info.Expiration).Err()
	if err != nil {
		r.log.Errorf("登录用户信息缓存失败 %v", err)
		return err
	}

	return r.data.rdb.Set(ctx, fmt.Sprintf(cacheStringLoginUser, info.AuthUser.ID), info.UUID, info.Expiration).Err()
}

// GetLoginCache 获取登录信息
func (r *UserRepo) GetLoginCache(ctx context.Context, uid uint) (*biz.User, error) {
	result := r.data.rdb.Get(ctx, fmt.Sprintf(cacheStringLoginUser, uid))
	if err := result.Err(); err != nil {
		return nil, err
	}

	result = r.data.rdb.Get(ctx, fmt.Sprintf(cacheStringLoginID, result.Val()))
	authUser := biz.User{}
	if err := json.Unmarshal([]byte(result.Val()), &authUser); err != nil {
		r.log.Errorf("unmarshal login auth user", err)
		return nil, err
	}

	return &authUser, result.Err()
}

// ExistLoginCache 登录信息是否存在
func (r *UserRepo) ExistLoginCache(ctx context.Context, uid uint) bool {
	result := r.data.rdb.Exists(ctx, fmt.Sprintf(cacheStringLoginUser, uid))
	if err := result.Err(); err != nil {
		return false
	}
	return result.Val() == 1
}

// DeleteLoginCache 删除登录信息
func (r *UserRepo) DeleteLoginCache(ctx context.Context, uid uint) error {
	result := r.data.rdb.Get(ctx, fmt.Sprintf(cacheStringLoginUser, uid))
	err := result.Err()
	if err != nil {
		r.log.Errorf("登录用户ID缓存查询失败 %v", err)
		return err
	}
	// 删除缓存信息
	{
		if err = r.data.rdb.Del(ctx, fmt.Sprintf(cacheStringLoginID, result.Val())).Err(); err != nil {
			r.log.Errorf("登录用户ID信息缓存删除失败 %v", err)
		}

		if err = r.data.rdb.Del(ctx, fmt.Sprintf(cacheStringLoginUser, uid)).Err(); err != nil {
			r.log.Errorf("登录用户信息缓存删除失败 %v", err)
		}
	}
	return err
}

// GetLoginIDCache 获取登录信息
func (r *UserRepo) GetLoginIDCache(ctx context.Context, uuid string) (*biz.User, error) {
	result := r.data.rdb.Get(ctx, fmt.Sprintf(cacheStringLoginID, uuid))
	bizUser := biz.User{}
	if err := json.Unmarshal([]byte(result.Val()), &bizUser); err != nil {
		r.log.Errorf("unmarshal login auth user", err)
		return nil, err
	}

	return &bizUser, result.Err()
}
