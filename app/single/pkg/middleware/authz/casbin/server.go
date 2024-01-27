package casbin

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/beiduoke/go-scaffold-single/pkg/auth/authn"
	stdcasbin "github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
)

// 请求权限
// authnM.Server(
// authnM.WithDomainSupport(),
// enforcer 指定则无需传 model 以及 policy
// casbinM.WithCasbinModel(m),
// casbinM.WithCasbinPolicy(policy),
// authnM.WithCasbinEnforcer(enforcer),
// authnM.WithSecurityUserCreator(securityUserCreator),
// 	authorizer,
// ),

type contextKey string

const (
	ModelContextKey        contextKey = "CasbinModel"
	PolicyContextKey       contextKey = "CasbinPolicy"
	EnforcerContextKey     contextKey = "CasbinEnforcer"
	SecurityUserContextKey contextKey = "CasbinSecurityUser"

	reason string = "FORBIDDEN"

	defaultRBACModel = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && keyMatch(r.obj, p.obj) && (r.act == p.act || p.act == "*")
`
)

var (
	ErrSecurityUserCreatorMissing = errors.Forbidden(reason, "SecurityUserCreator is required")
	ErrEnforcerMissing            = errors.Forbidden(reason, "Enforcer is missing")
	ErrSecurityParseFailed        = errors.Forbidden(reason, "Security Info fault")
	ErrUnauthorized               = errors.Forbidden(reason, "Unauthorized Access")
)

type Option func(*options)

type options struct {
	enableDomain           bool
	autoLoadPolicy         bool
	autoLoadPolicyInterval time.Duration
	securityUserCreator    authn.SecurityUserCreator
	model                  model.Model
	policy                 persist.Adapter
	watcher                persist.Watcher
	enforcer               stdcasbin.IEnforcer
}

// WithDomainSupport  enable domain support
func WithDomainSupport() Option {
	return func(o *options) {
		o.enableDomain = true
	}
}

// WithWatcher Set Watcher for Casbin
func WithWatcher(watcher persist.Watcher) Option {
	return func(o *options) {
		o.watcher = watcher
	}
}

// WithAutoLoadPolicy enable policy auto load option
func WithAutoLoadPolicy(auto bool, per time.Duration) Option {
	return func(o *options) {
		o.autoLoadPolicy = auto
		o.autoLoadPolicyInterval = per
	}
}

func WithSecurityUserCreator(securityUserCreator authn.SecurityUserCreator) Option {
	return func(o *options) {
		o.securityUserCreator = securityUserCreator
	}
}

func WithCasbinModel(model model.Model) Option {
	return func(o *options) {
		o.model = model
	}
}

func WithCasbinPolicy(policy persist.Adapter) Option {
	return func(o *options) {
		o.policy = policy
	}
}

func WithCasbinEnforcer(enforcer stdcasbin.IEnforcer) Option {
	return func(o *options) {
		o.enforcer = enforcer
	}
}

// loadRbacModel 加载RBAC模型
func loadRbacModel() (model.Model, error) {
	return model.NewModelFromString(defaultRBACModel)
}

func Server(opts ...Option) middleware.Middleware {
	o := &options{
		enableDomain: false,
	}
	for _, opt := range opts {
		opt(o)
	}

	if o.enforcer == nil {
		if o.model == nil {
			o.model, _ = loadRbacModel()
		}
		if o.policy == nil {
			o.policy = fileadapter.NewAdapter("configs/casbin/policy.csv")
		}

		o.enforcer, _ = stdcasbin.NewSyncedEnforcer(o.model, o.policy)
		if o.enforcer != nil && o.watcher != nil {
			o.watcher.SetUpdateCallback(func(s string) {
				o.enforcer.LoadPolicy()
			})
			_ = o.enforcer.SetWatcher(o.watcher)
		}
		// set autoload policy
		// if o.enforcer != nil && o.autoLoadPolicy && o.autoLoadPolicyInterval > time.Duration(0) {
		// 	if !o.enforcer.IsAutoLoadingRunning() {
		// 		o.enforcer.StartAutoLoadPolicy(o.autoLoadPolicyInterval)
		// 	}
		// }
	}

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			var (
				allowed bool
				err     error
			)

			if o.enforcer == nil {
				return nil, ErrEnforcerMissing
			}
			claims, success := authn.AuthClaimsFromContext(ctx)
			if !success {
				return nil, ErrSecurityUserCreatorMissing
			}
			fmt.Println(claims)
			// return handler(ctx, req)

			if o.securityUserCreator == nil {
				return nil, ErrSecurityUserCreatorMissing
			}

			securityUser := o.securityUserCreator(&authn.AuthClaims{})
			if err := securityUser.ParseFromContext(ctx); err != nil {
				return nil, ErrSecurityParseFailed
			}

			ctx = context.WithValue(ctx, SecurityUserContextKey, securityUser)
			if o.enableDomain {
				allowed, err = o.enforcer.Enforce(securityUser.GetSubject(), securityUser.GetDomain(), securityUser.GetObject(), securityUser.GetAction())
			} else {
				allowed, err = o.enforcer.Enforce(securityUser.GetSubject(), securityUser.GetObject(), securityUser.GetAction())
			}
			// fmt.Printf("权限：%s \n", securityUser.GetSubject())
			// fmt.Printf("租户：%s \n", securityUser.GetDomain())
			// fmt.Printf("资源：%s \n", securityUser.GetObject())
			// fmt.Printf("方法：%s \n", securityUser.GetAction())
			// fmt.Printf("允许访问 %t \n", allowed)
			// fmt.Printf("校验状态 %v \n", err)
			if err != nil {
				return nil, err
			}
			if !allowed {
				return nil, ErrUnauthorized
			}
			return handler(ctx, req)
		}
	}
}

func Client(opts ...Option) middleware.Middleware {
	o := &options{
		enableDomain:        false,
		securityUserCreator: nil,
	}
	for _, opt := range opts {
		opt(o)
	}

	if o.enforcer == nil {
		if o.model == nil {
			o.model, _ = loadRbacModel()
		}
		if o.policy == nil {
			o.policy = fileadapter.NewAdapter("configs/casbin/policy.csv")
		}
		enforcer, err := stdcasbin.NewEnforcer(o.model, o.policy)
		if err != nil {
			log.Fatalf("failed casbin enforcer %v", err)
		}
		o.enforcer = enforcer
	}

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			return handler(ctx, req)
		}
	}
}

// SecurityUserFromContext extract SecurityUser from context
func SecurityUserFromContext(ctx context.Context) (authn.SecurityUser, bool) {
	user, ok := ctx.Value(SecurityUserContextKey).(authn.SecurityUser)
	return user, ok
}
