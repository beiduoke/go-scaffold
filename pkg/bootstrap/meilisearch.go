package bootstrap

import (
	"github.com/go-kratos/kratos/v2/log"

	"github.com/beiduoke/go-scaffold/api/common/conf"
	"github.com/meilisearch/meilisearch-go"
)

// NewMeilisearchClient 创建Meilisearch客户端
func NewMeilisearchClient(cfg *conf.Bootstrap, logger *log.Helper) *meilisearch.Client {
	sdb := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:    cfg.Data.Meilisearch.GetHost(),
		APIKey:  cfg.Data.Meilisearch.GetApiKey(),
		Timeout: cfg.Data.Meilisearch.Timeout.AsDuration(),
	})
	_, err := sdb.Health()
	if err != nil {
		logger.Fatalf("failed opening connection to redis %v", err)
	}
	return sdb
}
