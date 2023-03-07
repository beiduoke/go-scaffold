package casbin

import (
	_ "embed"

	stdcasbin "github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
)

//go:embed model/rbac.conf
var DefaultRbacModel string

//go:embed model/rbac_with_domains.conf
var DefaultRbacWithDomainModel string

//go:embed model/abac.conf
var DefaultAbacModel string

//go:embed model/acl.conf
var DefaultAclModel string

//go:embed model/restfull.conf
var DefaultRestfullModel string

//go:embed model/restfull_with_role.conf
var DefaultRestfullWithRoleModel string

type Option func(*State)

func WithModel(model model.Model) Option {
	return func(s *State) {
		s.model = model
	}
}

func WithStringModel(str string) Option {
	return func(s *State) {
		s.model, _ = model.NewModelFromString(str)
	}
}

func WithFileModel(path string) Option {
	return func(s *State) {
		s.model, _ = model.NewModelFromFile(path)
	}
}

func WithPolicyAdapter(policy persist.Adapter) Option {
	return func(s *State) {
		s.policy = policy
	}
}

func WithEnforcer(enforcer stdcasbin.IEnforcer) Option {
	return func(s *State) {
		s.enforcer = enforcer
	}
}
