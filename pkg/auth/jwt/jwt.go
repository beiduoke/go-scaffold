package jwt

import (
	"context"
	"log"
	"time"

	"github.com/beiduoke/go-scaffold/pkg/auth"
	"github.com/beiduoke/go-scaffold/pkg/authz"
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

var _ auth.Authenticator = (*Authenticator)(nil)

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

func NewAuthenticator(opts ...Option) (auth.Authenticator, error) {
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
		err = auth.ErrInvalidInitJwt
		log.Println(err.Error())
	}

	return a, err
}

func (a *Authenticator) Authenticate(ctx context.Context) (*auth.AuthClaims, error) {
	if a.parseContextTokenFunc == nil {
		return nil, auth.ErrInvalidParseContextFunc
	}

	tokenString, err := a.parseContextTokenFunc(ctx)
	if err != nil {
		return nil, auth.ErrMissingBearerToken
	}

	token, err := a.parseToken(tokenString)
	if err != nil {
		ve, ok := err.(*jwt.ValidationError)
		if !ok {
			return nil, auth.ErrUnauthenticated
		}
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, auth.ErrInvalidToken
		}
		if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return nil, auth.ErrTokenExpired
		}
		return nil, auth.ErrInvalidToken
	}

	if !token.Valid {
		return nil, auth.ErrInvalidToken
	}
	if token.Method != a.signingMethod {
		return nil, auth.ErrUnsupportedSigningMethod
	}
	if token.Claims == nil {
		return nil, auth.ErrInvalidClaims
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, auth.ErrInvalidClaims
	}

	authClaims, err := auth.MapClaimsToAuthClaims(claims)
	if err != nil {
		return nil, err
	}

	return authClaims, nil
}

func (a *Authenticator) CreateIdentity(ctx context.Context, claims auth.AuthClaims) (string, error) {
	token := jwt.NewWithClaims(a.signingMethod, auth.AuthClaimsToJwtClaims(claims))

	tokenStr, err := a.generateToken(token)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func (a *Authenticator) parseToken(token string) (*jwt.Token, error) {
	if a.keyFunc == nil {
		return nil, auth.ErrMissingKeyFunc
	}

	return jwt.Parse(token, a.keyFunc)
}

func (a *Authenticator) generateToken(token *jwt.Token) (string, error) {
	if a.keyFunc == nil {
		return "", auth.ErrMissingKeyFunc
	}

	key, err := a.keyFunc(token)
	if err != nil {
		return "", auth.ErrGetKeyFailed
	}
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return "", auth.ErrSignTokenFailed
	}

	return tokenStr, nil
}

func (a Authenticator) Security() authz.SecurityUser {
	return nil
}
