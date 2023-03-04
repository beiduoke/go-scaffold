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
	cacheHashKeyUser     string = "sys_user"
	cacheStringLoginID   string = "login_uuid:%s"
	cacheHashLoginToken  string = "login_token:%d"
	cacheStringLoginUser string = "login_user:%d"
)

var _ Cache[*SysUser] = (*UserRepo)(nil)

type UserLoginInfo struct {
	UUID       string
	Token      string
	Expiration time.Duration
	User       SysUser
	Info       map[string]interface{}
}

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

// // GetLoginTokenCache 获取登录Token
// func (r *UserRepo) GetLoginTokenCache(ctx context.Context, token string, field string) (*UserLoginInfo, error) {
// 	result := r.data.rdb.HGet(ctx, fmt.Sprintf(cacheHashLoginToken, token), field)
// 	if err := result.Err(); err != nil {
// 		return nil, err
// 	}

// 	return nil, result.Err()
// }

// // SetLoginTokenCache 设置登录Token
// func (r *UserRepo) SetLoginTokenCache(ctx context.Context, token string, value UserLoginInfo) error {
// 	return r.data.rdb.HSet(ctx, fmt.Sprintf(cacheHashLoginToken, token), token).Err()
// }

// // ExistLoginTokenCache 登录Token是否存在
// func (r *UserRepo) ExistLoginTokenCache(ctx context.Context, uid uint) bool {
// 	result := r.data.rdb.Exists(ctx, fmt.Sprintf(cacheHashLoginToken, uid))
// 	if result.Err() != nil {
// 		return false
// 	}

// 	return result.Val() != 0
// }

// // DeleteLoginTokenCache 删除登录Token
// func (r *UserRepo) DeleteLoginTokenCache(ctx context.Context, uid uint) error {
// 	return r.data.rdb.Del(ctx, fmt.Sprintf(cacheHashLoginToken, uid)).Err()
// }

// SetLoginCache 设置登录信息
func (r *UserRepo) SetLoginCache(ctx context.Context, info UserLoginInfo) error {

	// 设置登录用户UUID 用于判断用户是否登录
	err := r.data.rdb.Set(ctx, fmt.Sprintf(cacheStringLoginID, info.UUID), info.Token, info.Expiration).Err()
	if err != nil {
		r.log.Errorf("用户登录UUID缓存失败 %v", err)
		return err
	}

	// 设置登录用户Token 用于记录用户登录记录
	err = r.data.rdb.HSet(ctx, fmt.Sprintf(cacheHashLoginToken, info.User.ID), map[string]interface{}{
		"UUID":     info.UUID,
		"UID":      info.User.ID,
		"DomainID": info.User.Domain.ID,
	}).Err()
	if err != nil {
		r.log.Errorf("用户token缓存失败 %v", err)
		return err
	}
	// 设置登录用户信息
	dataStr, err := json.Marshal(info.User)
	if err != nil {
		r.log.Errorf("用户缓存失败 %v", err)
		return err
	}
	return r.data.rdb.Set(ctx, fmt.Sprintf(cacheStringLoginUser, info.User.ID), dataStr, info.Expiration).Err()
}

// GetLoginCache 获取登录信息
func (r *UserRepo) GetLoginCache(ctx context.Context, uid uint) (*SysUser, error) {
	result := r.data.rdb.Get(ctx, fmt.Sprintf(cacheStringLoginUser, uid))
	if err := result.Err(); err != nil {
		return nil, err
	}
	sysUser := SysUser{}
	if err := result.Scan(&sysUser); err != nil {
		return nil, err
	}

	return &sysUser, result.Err()
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
	return r.data.rdb.Del(ctx, fmt.Sprintf(cacheStringLoginUser, uid)).Err()
}
