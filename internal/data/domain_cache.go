package data

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"gorm.io/gorm"
)

const (
	cacheHashKeyDomain = "sys_domain"
)

var _ Cache[*SysDomain] = (*DomainRepo)(nil)

// SetCache 设置领域缓存
func (r *DomainRepo) SetCache(ctx context.Context, g *SysDomain) error {
	dataStr, err := json.Marshal(g)
	if err != nil {
		r.log.Errorf("领域缓存失败 %v", err)
		return err
	}
	return r.data.rdb.HSet(ctx, cacheHashKeyDomain, convert.UnitToString(g.ID), dataStr).Err()
}

// GetCache 获取领域缓存
func (r *DomainRepo) GetCache(ctx context.Context, key string) (sysDomain *SysDomain) {
	dataStr, err := r.data.rdb.HGet(ctx, cacheHashKeyDomain, key).Result()
	if err != nil {
		return nil
	}
	if err := json.Unmarshal([]byte(dataStr), &sysDomain); err != nil {
		r.log.Errorf("缓存反序列化失败 %v", err)
	}
	return sysDomain
}

// DeleteCache 获取领域缓存
func (r *DomainRepo) DeleteCache(ctx context.Context, key string) error {
	return r.data.rdb.HDel(ctx, cacheHashKeyDomain, key).Err()
}

// ListAllCache 获取全部缓存数据
func (r *DomainRepo) ListAllCache(ctx context.Context) (sysDomains []*SysDomain) {
	if l, _ := r.data.rdb.HLen(ctx, cacheHashKeyDomain).Result(); l > 0 {
		domainMap, _ := r.data.rdb.HGetAll(ctx, cacheHashKeyDomain).Result()
		for _, v := range domainMap {
			sysDomain := SysDomain{}
			err := json.Unmarshal([]byte(v), &sysDomain)
			if err != nil {
				r.log.Errorf("领域缓存反序列失败 %v", err)
				continue
			}
			sysDomains = append(sysDomains, &sysDomain)
		}
		return sysDomains
	}

	result := r.data.DB(ctx).Find(&sysDomains)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		r.log.Errorf("领域查询失败 %v", result.Error)
		return nil
	}
	domainMap := make(map[string]string)
	for _, v := range sysDomains {
		menuStr, err := json.Marshal(v)
		if err != nil {
			r.log.Errorf("领域缓存序列化失败 %v", err)
			continue
		}
		domainMap[convert.UnitToString(v.ID)] = string(menuStr)
	}
	if err := r.data.rdb.HSet(ctx, cacheHashKeyDomain, domainMap).Err(); err != nil {
		r.log.Errorf("领域缓存失败 %v", err)
	}
	return sysDomains
}
