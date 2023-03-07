package casbin

import (
	"context"

	stdcasbin "github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"

	"github.com/beiduoke/go-scaffold/pkg/auth/authz"
)

var _ authz.Authorized = (*State)(nil)

type State struct {
	model    model.Model
	policy   persist.Adapter
	enforcer stdcasbin.IEnforcer
	watcher  persist.Watcher
	projects authz.Projects

	wildcardItem              string
	authorizedProjectsMatcher string
}

func NewAuthorized(_ context.Context, opts ...Option) (*State, error) {
	s := State{
		policy:                    newAdapter(),
		projects:                  authz.Projects{},
		wildcardItem:              "*",
		authorizedProjectsMatcher: "g(r.sub, p.sub, p.dom) && (keyMatch(r.dom, p.dom) || p.dom == '*')",
		watcher:                   nil,
	}

	for _, opt := range opts {
		opt(&s)
	}

	var err error

	if s.model == nil {
		s.model, err = model.NewModelFromString(DefaultRestfullWithRoleModel)
		if err != nil {
			return nil, err
		}
	}

	if s.enforcer == nil {
		s.enforcer, err = stdcasbin.NewSyncedEnforcer(s.model, s.policy)
	}
	if err != nil {
		return nil, err
	}

	if s.watcher != nil {
		err = s.enforcer.SetWatcher(s.watcher)
		s.watcher.SetUpdateCallback(func(v string) {
			s.enforcer.LoadPolicy()
		})
	}

	return &s, err
}

func (s *State) ProjectsAuthorized(_ context.Context, subjects authz.Subjects, action authz.Action, resource authz.Resource, projects authz.Projects) (authz.Projects, error) {
	result := make(authz.Projects, 0, len(projects))

	var err error
	var allowed bool
	for _, project := range projects {
		for _, subject := range subjects {
			if allowed, err = s.enforcer.Enforce(string(subject), string(resource), string(action), string(project)); err != nil {
				//fmt.Println(allowed, err)
				return nil, err
			} else if allowed {
				result = append(result, project)
			}
		}
	}

	return result, nil
}

func (s *State) FilterAuthorizedPairs(_ context.Context, subjects authz.Subjects, pairs authz.Pairs) (authz.Pairs, error) {
	result := make(authz.Pairs, 0, len(pairs))

	project := authz.Project(s.wildcardItem)

	var err error
	var allowed bool
	for _, p := range pairs {
		for _, subject := range subjects {
			if allowed, err = s.enforcer.Enforce(string(subject), string(p.Resource), string(p.Action), string(project)); err != nil {
				//fmt.Println(allowed, err)
				return nil, err
			} else if allowed {
				result = append(result, p)
			}
		}
	}
	return result, nil
}

func (s *State) FilterAuthorizedProjects(_ context.Context, subjects authz.Subjects) (authz.Projects, error) {
	result := make(authz.Projects, 0, len(s.projects))

	resource := authz.Resource(s.wildcardItem)
	action := authz.Action(s.wildcardItem)

	var err error
	var allowed bool
	for _, project := range s.projects {
		for _, subject := range subjects {
			if allowed, err = s.enforcer.EnforceWithMatcher(s.authorizedProjectsMatcher, string(subject), string(resource), string(action), string(project)); err != nil {
				//fmt.Println(allowed, err)
				return nil, err
			} else if allowed {
				result = append(result, project)
			}
		}
	}

	return result, nil
}

func (s *State) IsAuthorized(_ context.Context, subject authz.Subject, action authz.Action, resource authz.Resource, project authz.Project) (bool, error) {
	if len(project) == 0 {
		project = authz.Project(s.wildcardItem)
	}

	var err error
	var allowed bool
	if allowed, err = s.enforcer.Enforce(string(subject), string(resource), string(action), string(project)); err != nil {
		//fmt.Println(allowed, err)
		return false, err
	} else if allowed {
		return true, nil
	}
	return false, nil
}

func (s *State) SetPolicies(_ context.Context, policyMap authz.PolicyMap, _ authz.RoleMap) error {
	// s.policy.SetPolicies(policyMap)
	// s.policy.SavePolicy()
	err := s.enforcer.LoadPolicy()
	//fmt.Println(err, s.enforcer.GetAllSubjects(), s.enforcer.GetAllRoles())

	projects, ok := policyMap["projects"]
	if ok {
		switch t := projects.(type) {
		case authz.Projects:
			s.projects = t
		}
	}

	return err
}
