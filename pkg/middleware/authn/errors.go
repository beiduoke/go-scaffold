package authn

import "github.com/go-kratos/kratos/v2/errors"

var (
	ErrWrongContext         = errors.Unauthorized(reason, "wrong context for middleware")
	ErrMissingJwtToken      = errors.Unauthorized(reason, "no jwt token in context")
	ErrExtractSubjectFailed = errors.Unauthorized(reason, "extract subject failed")
)
