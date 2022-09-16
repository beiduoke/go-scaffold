package sso

type Option func(*options)

type options struct {
	prefix string
}

func WithPrefix(prefix string) Option {
	return func(o *options) {
		o.prefix = prefix
	}
}

type Server interface {
	Valid() error
}
