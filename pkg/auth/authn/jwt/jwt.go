package jwt

import (
	"context"
	"log"
	"time"

	"github.com/beiduoke/go-scaffold/pkg/auth/authn"
	"github.com/golang-jwt/jwt/v4"
)

var (
	secretKey = "jwtKey001"
)

type ParseContextToken func(context.Context) (string, error)

type Authenticator struct {
	signingMethod         jwt.SigningMethod
	keyFunc               jwt.Keyfunc
	parseContextTokenFunc ParseContextToken
	expiresAt             time.Duration
}

var _ authn.Authenticator = (*Authenticator)(nil)

type Option func(*Authenticator)

func WithKeyFunc(k jwt.Keyfunc) Option {
	return func(a *Authenticator) {
		a.keyFunc = k
	}
}

func WithSecretKey(s string) Option {
	return func(a *Authenticator) {
		a.keyFunc = func(t *jwt.Token) (interface{}, error) {
			return []byte(s), nil
		}
	}
}

func WithSigMethod(s string) Option {
	return func(a *Authenticator) {
		a.signingMethod = jwt.GetSigningMethod(s)
	}
}

func WithParseContext(p ParseContextToken) Option {
	return func(a *Authenticator) {
		a.parseContextTokenFunc = p
	}
}

func WithExpiresAt(d time.Duration) Option {
	return func(a *Authenticator) {
		a.expiresAt = d
	}
}

func NewAuthenticator(opts ...Option) (authn.Authenticator, error) {
	a := &Authenticator{
		signingMethod: jwt.SigningMethodHS256,
		keyFunc: func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		},
	}

	for _, o := range opts {
		o(a)
	}

	var err error
	if a.parseContextTokenFunc == nil {
		err = authn.ErrInvalidInitJwt
		log.Println(err.Error())
	}

	return a, err
}

func (a *Authenticator) Authenticate(ctx context.Context) (*authn.AuthClaims, error) {
	if a.parseContextTokenFunc == nil {
		return nil, authn.ErrInvalidParseContextFunc
	}

	tokenString, err := a.parseContextTokenFunc(ctx)
	if err != nil {
		return nil, authn.ErrMissingBearerToken
	}

	token, err := a.parseToken(tokenString)
	if err != nil {
		ve, ok := err.(*jwt.ValidationError)
		if !ok {
			return nil, authn.ErrUnauthenticated
		}
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, authn.ErrInvalidToken
		}
		if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return nil, authn.ErrTokenExpired
		}
		return nil, authn.ErrInvalidToken
	}

	if !token.Valid {
		return nil, authn.ErrInvalidToken
	}
	if token.Method != a.signingMethod {
		return nil, authn.ErrUnsupportedSigningMethod
	}
	if token.Claims == nil {
		return nil, authn.ErrInvalidClaims
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, authn.ErrInvalidClaims
	}

	return authn.MapClaimsToAuthClaims(claims)
}

func (a *Authenticator) CreateIdentity(ctx context.Context, claims authn.AuthClaims) (string, error) {
	token := jwt.NewWithClaims(a.signingMethod, authn.AuthClaimsToJwtClaims(claims))
	return a.generateToken(token)
}

func (a *Authenticator) parseToken(token string) (*jwt.Token, error) {
	if a.keyFunc == nil {
		return nil, authn.ErrMissingKeyFunc
	}

	return jwt.Parse(token, a.keyFunc)
}

func (a *Authenticator) generateToken(token *jwt.Token) (string, error) {
	if a.keyFunc == nil {
		return "", authn.ErrMissingKeyFunc
	}

	key, err := a.keyFunc(token)
	if err != nil {
		return "", authn.ErrGetKeyFailed
	}
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return "", authn.ErrSignTokenFailed
	}

	return tokenStr, nil
}
