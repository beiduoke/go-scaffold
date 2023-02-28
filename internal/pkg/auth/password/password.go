package password

import (
	"github.com/beiduoke/go-scaffold/internal/pkg/auth"
)

var _ auth.Auth[*Data] = (*password)(nil)

type password struct {
	options auth.Options
}

type Data struct {
	Account  string
	Password string
}

func NewServer(opts ...auth.Option) auth.Auth[*Data] {
	p := &password{}

	for _, o := range opts {
		o(&p.options)
	}
	return p
}

func (p *password) Options() auth.Options {
	return p.options
}

func (p *password) Init(opts ...auth.Option) error {
	for _, o := range opts {
		o(&p.options)
	}
	return nil
}
func (p *password) String() string {
	return "Password"
}

func (p *password) Login(d *Data) (auth.AuthClaims, error) {
	// o := p.options

	// if err = passwordutil.Verify(d.Password, u.Password); err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

func (p *password) Register(d *Data) error {
	return nil
}

func (p *password) Logout() error {
	return nil
}
