package auth

import (
	"time"
)

type AuthTypes[T any] map[interface{}]Auth[T]

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
	CreateToken(string) error
	Token() string
	ExpiresAt() time.Time
}

// AuthUser 认证用户
type AuthUser struct {
	ID       uint
	Name     string
	Avatar   string
	NickName string
	RealName string
	Password string
	Birthday *time.Time
	Gender   int32
	Phone    string
	Email    string
	State    int32
}
