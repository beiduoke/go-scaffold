package data

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"gorm.io/gorm"
)

const (
	cacheHashKeyUser          string = "sys_user"
	cacheStringLoginUser      string = "login_user:%s"
	cacheStringLoginUserToken string = "login_user_token:%d"
)

var _ Cache[*SysUser] = (*UserRepo)(nil)

// SetCache 设置用户缓存
func (r *UserRepo) SetCache(ctx context.Context, g *SysUser) error {
	dataStr, err := json.Marshal(g)
	if err != nil {
		r.log.Errorf("用户缓存失败 %v", err)
		return err
	}
	return r.data.rdb.HSet(ctx, cacheHashKeyUser, convert.UnitToString(g.ID), dataStr).Err()
}

// GetCache 获取用户缓存
func (r *UserRepo) GetCache(ctx context.Context, key string) (sysUser *SysUser) {
	dataStr, err := r.data.rdb.HGet(ctx, cacheHashKeyUser, key).Result()
	if err != nil {
		return nil
	}
	if err := json.Unmarshal([]byte(dataStr), &sysUser); err != nil {
		r.log.Errorf("缓存反序列化失败 %v", err)
	}
	return sysUser
}

// DeleteCache 获取用户缓存
func (r *UserRepo) DeleteCache(ctx context.Context, key string) error {
	return r.data.rdb.HDel(ctx, cacheHashKeyUser, key).Err()
}

// ListAllCache 获取全部缓存数据
func (r *UserRepo) ListAllCache(ctx context.Context) (sysUsers []*SysUser) {
	if l, _ := r.data.rdb.HLen(ctx, cacheHashKeyUser).Result(); l > 0 {
		domainMap, _ := r.data.rdb.HGetAll(ctx, cacheHashKeyUser).Result()
		for _, v := range domainMap {
			sysUser := SysUser{}
			err := json.Unmarshal([]byte(v), &sysUser)
			if err != nil {
				r.log.Errorf("用户缓存反序列失败 %v", err)
				continue
			}
			sysUsers = append(sysUsers, &sysUser)
		}
		return sysUsers
	}

	result := r.data.DB(ctx).Find(&sysUsers)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		r.log.Errorf("用户查询失败 %v", result.Error)
		return nil
	}
	domainMap := make(map[string]string)
	for _, v := range sysUsers {
		menuStr, err := json.Marshal(v)
		if err != nil {
			r.log.Errorf("用户缓存序列化失败 %v", err)
			continue
		}
		domainMap[convert.UnitToString(v.ID)] = string(menuStr)
	}
	if err := r.data.rdb.HSet(ctx, cacheHashKeyUser, domainMap).Err(); err != nil {
		r.log.Errorf("用户缓存失败 %v", err)
	}
	return sysUsers
}

func (r *UserRepo) SetLoginCache(ctx context.Context, uuid string, g SysUser) error {
	dataStr, err := json.Marshal(g)
	if err != nil {
		r.log.Errorf("用户缓存失败 %v", err)
		return err
	}
	return r.data.rdb.Set(ctx, fmt.Sprintf(cacheStringLoginUser, uuid), dataStr, 0).Err()
}

func (r *UserRepo) GetLoginCache(ctx context.Context, uuid string) (*SysUser, error) {
	result := r.data.rdb.Get(ctx, fmt.Sprintf(cacheStringLoginUser, uuid))
	if err := result.Err(); err != nil {
		return nil, err
	}
	sysUser := SysUser{}
	if err := result.Scan(&sysUser); err != nil {
		return nil, err
	}

	return &sysUser, result.Err()
}

func (r *UserRepo) SetLoginTokenCache(ctx context.Context, uid uint, token string, exp time.Duration) error {
	return r.data.rdb.Set(ctx, fmt.Sprintf(cacheStringLoginUserToken, uid), token, exp).Err()
}

func (r *UserRepo) ExistLoginTokenCache(ctx context.Context, uid uint) bool {
	result := r.data.rdb.Exists(ctx, fmt.Sprintf(cacheStringLoginUserToken, uid))
	if result.Err() != nil {
		return false
	}

	return result.Val() != 0
}

func (r *UserRepo) ExistLoginCache(ctx context.Context, uid uint) bool {
	result := r.data.rdb.Exists(ctx, convert.UnitToString(uid))
	if err := result.Err(); err != nil {
		return false
	}

	fmt.Println(result.Val())

	return false
}

func (r *UserRepo) DeleteLoginCache(ctx context.Context, uid uint) error {
	key := fmt.Sprintf("login_user_token:%d", uid)
	return r.data.rdb.Del(ctx, key).Err()
}
