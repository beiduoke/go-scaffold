package data

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/beiduoke/go-scaffold-single/internal/biz"
	"github.com/beiduoke/go-scaffold-single/pkg/util/convert"
	"gorm.io/gorm"
)

const (
	cacheHashKeyMenu = "sys_menu"
)

var _ Cache[*biz.Menu] = (*MenuRepo)(nil)

// SetCache 设置菜单缓存
func (r *MenuRepo) SetCache(ctx context.Context, g *biz.Menu) error {
	dataStr, err := json.Marshal(g)
	if err != nil {
		r.log.Errorf("菜单缓存失败 %v", err)
		return err
	}
	return r.data.rdb.HSet(ctx, cacheHashKeyMenu, convert.UnitToString(g.ID), dataStr).Err()
}

// GetCache 获取菜单缓存
func (r *MenuRepo) GetCache(ctx context.Context, key string) (bizMenu *biz.Menu) {
	dataStr, err := r.data.rdb.HGet(ctx, cacheHashKeyMenu, key).Result()
	if err != nil {
		return nil
	}
	if err := json.Unmarshal([]byte(dataStr), &bizMenu); err != nil {
		r.log.Errorf("缓存反序列化失败 %v", err)
	}
	return bizMenu
}

// DeleteCache 获取菜单缓存
func (r *MenuRepo) DeleteCache(ctx context.Context, key string) error {
	return r.data.rdb.HDel(ctx, cacheHashKeyMenu, key).Err()
}

// ListAllCache 获取全部缓存数据
func (r *MenuRepo) ListAllCache(ctx context.Context) (bizMenus []*biz.Menu) {
	if l, _ := r.data.rdb.HLen(ctx, cacheHashKeyMenu).Result(); l > 0 {
		menuMap, _ := r.data.rdb.HGetAll(ctx, cacheHashKeyMenu).Result()
		for _, v := range menuMap {
			bizMenu := biz.Menu{}
			err := json.Unmarshal([]byte(v), &bizMenu)
			if err != nil {
				r.log.Errorf("菜单缓存反序列失败 %v", err)
				continue
			}
			bizMenus = append(bizMenus, &bizMenu)
		}
		return bizMenus
	}

	sysMenus := []SysMenu{}
	result := r.data.DB(ctx).Find(&sysMenus)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		r.log.Errorf("菜单查询失败 %v", result.Error)
		return nil
	}
	menuMap := make(map[string]string)
	for _, v := range sysMenus {
		bizMenu := toMenuBiz(&v)
		marshalUserJson, err := json.Marshal(bizMenu)
		if err != nil {
			r.log.Errorf("菜单缓存序列化失败 %v", err)
			continue
		}
		menuMap[bizMenu.GetID()] = string(marshalUserJson)
	}
	if err := r.data.rdb.HSet(ctx, cacheHashKeyMenu, menuMap).Err(); err != nil {
		r.log.Errorf("菜单缓存失败 %v", err)
	}
	return bizMenus
}
