package authz

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthErrorCode int32

const (
	AuthErrorCodeMissingAuthClaims  AuthErrorCode = 2001
	AuthErrorCodeInvalidClaims      AuthErrorCode = 2002
	AuthErrorCodeUnauthorizedAccess AuthErrorCode = 2003
)

var (
	ErrMissingAuthClaims  = status.Error(codes.Code(AuthErrorCodeMissingAuthClaims), "context missing authz claims")
	ErrInvalidClaims      = status.Error(codes.Code(AuthErrorCodeInvalidClaims), "invalid claims")
	ErrUnauthorizedAccess = status.Error(codes.Code(AuthErrorCodeUnauthorizedAccess), "unauthorized access")
)
