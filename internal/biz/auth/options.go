package auth

import (
	"github.com/go-kratos/kratos/v2/log"

	"github.com/beiduoke/go-scaffold/internal/conf"
)

type Option func(*Options)

// AuthUsecase is a User usecase.
type Options struct {
	Name   string
	Config *conf.Auth
	Log    *log.Helper
}

func newOptions(opt ...Option) Options {
	opts := Options{}

	for _, o := range opt {
		o(&opts)
	}
	return opts
}

// Server name
func Name(n string) Option {
	return func(o *Options) {
		o.Name = n
	}
}

// Biz 业务聚合
// func Biz(b biz.Biz) Option {
// 	return func(o *Options) {
// 		o.Biz = b
// 	}
// }

// Config 配置
func Config(c *conf.Auth) Option {
	return func(o *Options) {
		o.Config = c
	}
}

// Log 日志
func Log(l *log.Helper) Option {
	return func(o *Options) {
		o.Log = l
	}
}
