package data

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"gorm.io/gorm"
)

const (
	cacheHashKeyMenu = "sys_menu"
)

var _ Cache[*SysMenu] = (*MenuRepo)(nil)

// SetCache 设置菜单缓存
func (r *MenuRepo) SetCache(ctx context.Context, g *SysMenu) error {
	dataStr, err := json.Marshal(g)
	if err != nil {
		r.log.Errorf("菜单缓存失败 %v", err)
		return err
	}
	return r.data.rdb.HSet(ctx, cacheHashKeyMenu, convert.UnitToString(g.ID), dataStr).Err()
}

// GetCache 获取菜单缓存
func (r *MenuRepo) GetCache(ctx context.Context, key string) (sysMenu *SysMenu) {
	dataStr, err := r.data.rdb.HGet(ctx, cacheHashKeyMenu, key).Result()
	if err != nil {
		return nil
	}
	if err := json.Unmarshal([]byte(dataStr), &sysMenu); err != nil {
		r.log.Errorf("缓存反序列化失败 %v", err)
	}
	return sysMenu
}

// DeleteCache 获取菜单缓存
func (r *MenuRepo) DeleteCache(ctx context.Context, key string) error {
	return r.data.rdb.HDel(ctx, cacheHashKeyMenu, key).Err()
}

// ListAllCache 获取全部缓存数据
func (r *MenuRepo) ListAllCache(ctx context.Context) (menus []*SysMenu) {
	if l, _ := r.data.rdb.HLen(ctx, cacheHashKeyMenu).Result(); l > 0 {
		menuMap, _ := r.data.rdb.HGetAll(ctx, cacheHashKeyMenu).Result()
		for _, v := range menuMap {
			sysMenu := SysMenu{}
			err := json.Unmarshal([]byte(v), &sysMenu)
			if err != nil {
				r.log.Errorf("菜单缓存反序列失败 %v", err)
				continue
			}
			menus = append(menus, &sysMenu)
		}
		return menus
	}

	result := r.data.DB(ctx).Find(&menus)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		r.log.Errorf("菜单查询失败 %v", result.Error)
		return nil
	}
	menuMap := make(map[string]string)
	for _, v := range menus {
		menuStr, err := json.Marshal(v)
		if err != nil {
			r.log.Errorf("菜单缓存序列化失败 %v", err)
			continue
		}
		menuMap[convert.UnitToString(v.ID)] = string(menuStr)
	}
	if err := r.data.rdb.HSet(ctx, cacheHashKeyMenu, menuMap).Err(); err != nil {
		r.log.Errorf("菜单缓存失败 %v", err)
	}
	return menus
}
