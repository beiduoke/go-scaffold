package authz

import (
	"context"
)

type Authorized interface {
	Authorizer
	Writer
}

type Authorizer interface {
	ProjectsAuthorized(context.Context, Subjects, Action, Resource, Projects) (Projects, error)

	FilterAuthorizedPairs(context.Context, Subjects, Pairs) (Pairs, error)

	FilterAuthorizedProjects(context.Context, Subjects) (Projects, error)

	IsAuthorized(context.Context, Subject, Action, Resource, Project) (bool, error)
}

type Writer interface {
	SetPolicies(context.Context, PolicyMap, RoleMap) error
}

type SecurityUser interface {
	// ParseFromContext parses the user from the context.
	ParseFromContext(ctx context.Context) error
	// GetSubject returns the subject of the token.
	GetSubject() string
	// GetObject returns the object of the token.
	GetObject() string
	// GetAction returns the action of the token.
	GetAction() string
	// GetDomain returns the domain of the token.
	GetDomain() string
	// GetID returns the user of the token.
	GetUser() string
}

type SecurityUserCreator func() SecurityUser
