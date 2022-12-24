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
	NewResourceUsecase,
	NewMenuUsecase,
	NewAuthUsecase,
	NewUserUsecase,
	NewDepartmentUsecase,
)

type Transaction interface {
	InTx(context.Context, func(ctx context.Context) error) error
}

// Biz 公共业务逻辑
type Biz struct {
	log *log.Helper
	// 逻辑事务操作
	tm Transaction
	// 公共
	// 权限认证
	enforcer casbin.IEnforcer
	// 领域数据
	domainRepo DomainRepo
	// 权限数据
	authorityRepo AuthorityRepo
	// 用户数据
	userRepo UserRepo
}

func NewBiz(logger log.Logger, tm Transaction, enforcer casbin.IEnforcer, domainRepo DomainRepo, authorityRepo AuthorityRepo, userRepo UserRepo) *Biz {
	return &Biz{
		log:           log.NewHelper(log.With(logger, "module", "biz/initialize")),
		tm:            tm,
		enforcer:      enforcer,
		domainRepo:    domainRepo,
		authorityRepo: authorityRepo,
		userRepo:      userRepo,
	}
}
