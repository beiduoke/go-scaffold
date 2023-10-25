package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewBiz,
	NewDomainUsecase,
	NewRoleUsecase,
	NewMenuUsecase,
	NewUserUsecase,
	NewDeptUsecase,
	NewAuthUsecase,
	NewPostUsecase,
	NewDictUsecase,
)

type BizType interface {
	GetID() string
}

type Transaction interface {
	InTx(context.Context, func(ctx context.Context) error) error
}

// Biz 公共业务逻辑
type Biz struct {
	log *log.Helper
	// 逻辑事务操作
	tm Transaction
	// 公共
	// 租户数据
	domainRepo DomainRepo
	// 权限数据
	roleRepo RoleRepo
	// 用户数据
	userRepo UserRepo
}

func NewBiz(logger log.Logger, tm Transaction, domainRepo DomainRepo, roleRepo RoleRepo, userRepo UserRepo) *Biz {
	return &Biz{
		log:        log.NewHelper(log.With(logger, "module", "biz/initialize")),
		tm:         tm,
		domainRepo: domainRepo,
		roleRepo:   roleRepo,
		userRepo:   userRepo,
	}
}
