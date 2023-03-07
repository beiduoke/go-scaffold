package noop

import (
	"context"

	"github.com/beiduoke/go-scaffold/pkg/auth/authz"
)

var _ authz.Authorized = (*State)(nil)

type State struct{}

func (s State) ProjectsAuthorized(_ context.Context, _ authz.Subjects, _ authz.Action, _ authz.Resource, _ authz.Projects) (authz.Projects, error) {
	return authz.Projects{}, nil
}

func (s State) FilterAuthorizedPairs(_ context.Context, _ authz.Subjects, _ authz.Pairs) (authz.Pairs, error) {
	return authz.Pairs{}, nil
}

func (s State) FilterAuthorizedProjects(_ context.Context, _ authz.Subjects) (authz.Projects, error) {
	return authz.Projects{}, nil
}

func (s State) IsAuthorized(_ context.Context, _ authz.Subject, _ authz.Action, _ authz.Resource, _ authz.Project) (bool, error) {
	return true, nil
}

func (s State) SetPolicies(_ context.Context, _ authz.PolicyMap, _ authz.RoleMap) error {
	return nil
}
