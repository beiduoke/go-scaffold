package auth

type Options struct {
	auth AuthClaims
}

type Option func(*Options)

func WithAuthClaims(auth AuthClaims) Option {
	return func(o *Options) {
		o.auth = auth
	}
}
