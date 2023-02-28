package jwt

import (
	"context"
	"log"

	"github.com/beiduoke/go-scaffold/pkg/authn"
	"github.com/golang-jwt/jwt/v4"
)

var (
	key = "jwtKey001"
)

type ParseContext func(context.Context) (string, error)

type Authenticator struct {
	signingMethod    jwt.SigningMethod
	keyFunc          jwt.Keyfunc
	parseContextFunc ParseContext
}

var _ authn.Authenticator = (*Authenticator)(nil)

type Option func(*Authenticator)

func WithKeyFunc(k jwt.Keyfunc) Option {
	return func(a *Authenticator) {
		a.keyFunc = k
	}
}

func WithSigMethod(s string) Option {
	return func(a *Authenticator) {
		a.signingMethod = jwt.GetSigningMethod(s)
	}
}

func WithParseContext(p ParseContext) Option {
	return func(a *Authenticator) {
		a.parseContextFunc = p
	}
}

func NewAuthenticator(opts ...Option) (authn.Authenticator, error) {
	auth := &Authenticator{
		signingMethod: jwt.SigningMethodHS256,
		keyFunc: func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		},
	}

	for _, o := range opts {
		o(auth)
	}

	if auth.parseContextFunc == nil {
		log.Fatal(authn.ErrInvalidInitJwt)
	}

	return auth, nil
}

func (a *Authenticator) Authenticate(ctx context.Context) (*authn.AuthClaims, error) {
	if a.parseContextFunc == nil {
		return nil, authn.ErrInvalidParseContextFunc
	}

	tokenString, err := a.parseContextFunc(ctx)
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

	authClaims, err := authn.MapClaimsToAuthClaims(claims)
	if err != nil {
		return nil, err
	}

	return authClaims, nil
}

func (a *Authenticator) CreateIdentity(ctx context.Context, claims authn.AuthClaims) (string, error) {
	token := jwt.NewWithClaims(a.signingMethod, authn.AuthClaimsToJwtClaims(claims))

	tokenStr, err := a.generateToken(token)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
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
