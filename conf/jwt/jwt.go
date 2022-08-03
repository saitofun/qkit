package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/saitofun/qkit/base/types"
)

type Jwt struct {
	Issuer  string         `env:""`
	ExpIn   types.Duration `env:""`
	Method  SigningMethod  `env:""`
	SignKey string         `env:""`
}

func (c *Jwt) SetDefault() {}

func (c *Jwt) Init() {}

func (c *Jwt) GenerateTokenByPayload(payload interface{}) (string, error) {
	claim := &Claims{
		pl: payload,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(c.ExpIn.Duration()).Unix(),
			Issuer:    c.Issuer,
		},
	}
	token := jwt.NewWithClaims(SigningMethods[c.Method], claim)
	return token.SignedString(c.SignKey)
}

func (c *Jwt) ParseToken(v string) (*Claims, error) {
	t, err := jwt.ParseWithClaims(
		v,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return c.SignKey, nil
		},
	)
	if err != nil {
		return nil, err
	}
	if t == nil {
		return nil, ErrNilToken
	}
	claim, ok := t.Claims.(*Claims)
	if !ok {
		return nil, ErrInvalidClaim
	}
	return claim, nil
}

type Claims struct {
	pl interface{}
	jwt.StandardClaims
}

var (
	ErrNilToken     = errors.New("nil token")
	ErrInvalidClaim = errors.New("invalid claim")
)
