package auth

import (
	"time"
)

// Auth 多登录方式进行资源封装
type Auth[T any] interface {
	Options() Options
	Init(...Option) error
	String() string
	Login(T) (AuthClaims, error)
	Register(T) error
	Logout() error
}

type AuthClaims interface {
	Token() string
	ExpiresAt() time.Time
}
