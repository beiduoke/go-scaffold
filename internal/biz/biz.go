package biz

import (
	"context"

	"github.com/casbin/casbin/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewBiz,
	NewDomainUsecase,
	NewAuthorityUsecase,
	NewApiUsecase,
	NewMenuUsecase,
	NewAuthUsecase,
	NewUserUsecase,
)

type Transaction interface {
	InTx(context.Context, func(ctx context.Context) error) error
}

// Biz 公共业务逻辑
type Biz struct {
	log *log.Helper
	// 逻辑事务操作
	tm Transaction
	// 公共接口
	// 权限认证接口
	enforcer casbin.IEnforcer
	// 领域数据接口
	domainRepo DomainRepo
	// 权限数据接口
	authorityRepo AuthorityRepo
	// 领域数据接口
	userRepo UserRepo
}

func NewBiz(logger log.Logger, tm Transaction, enforcer casbin.IEnforcer, domainRepo DomainRepo, userRepo UserRepo) *Biz {
	return &Biz{
		log:        log.NewHelper(log.With(logger, "module", "biz/initialize")),
		tm:         tm,
		enforcer:   enforcer,
		domainRepo: domainRepo,
		userRepo:   userRepo,
	}
}
