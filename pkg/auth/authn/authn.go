package authn

import (
	"context"
)

type Authenticator interface {
	// Authenticate returns a nil error and the AuthClaims info (if available).
	// if the subject is authenticated or a non-nil error with an appropriate error cause otherwise.
	Authenticate(requestContext context.Context) (*AuthClaims, error)

	// CreateIdentity inject user claims into context.
	CreateIdentity(requestContext context.Context, claims AuthClaims) (string, error)
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
	// GetUser returns the user of the token.
	GetUser() string
}

type SecurityUserCreator func(*AuthClaims) SecurityUser
