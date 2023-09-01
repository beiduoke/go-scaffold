package bootstrap

import (
	"github.com/go-kratos/kratos/v2/log"

	"github.com/redis/go-redis/v9"

	"github.com/beiduoke/go-scaffold/api/common/conf"
)

// NewRedisClient 创建Redis客户端
func NewRedisClient(cfg *conf.Bootstrap, logger *log.Helper) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         cfg.Data.Redis.Addr,
		Password:     cfg.Data.Redis.Password,
		DB:           int(cfg.Data.Redis.Db),
		DialTimeout:  cfg.Data.Redis.DialTimeout.AsDuration(),
		WriteTimeout: cfg.Data.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  cfg.Data.Redis.ReadTimeout.AsDuration(),
	})
	if rdb == nil {
		logger.Fatalf("failed opening connection to redis")
		return nil
	}
	// rdb.AddHook(redis.ProcessHook())

	return rdb
}
