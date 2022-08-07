package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"

	"github.com/saitofun/qkit/base/types"
)

type Jwt struct {
	Issuer  string         `env:""`
	ExpIn   types.Duration `env:""`
	SignKey string         `env:""`
	// Method  SigningMethod  `env:""`
}

func (c *Jwt) SetDefault() {}

func (c *Jwt) Init() {}

func (c *Jwt) GenerateTokenByPayload(payload interface{}) (string, error) {
	claim := &Claims{
		Payload: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(c.ExpIn.Duration())},
			Issuer:    c.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(c.SignKey))
}

func (c *Jwt) ParseToken(v string) (*Claims, error) {
	t, err := jwt.ParseWithClaims(
		v,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(c.SignKey), nil
		},
	)
	if err != nil {
		return nil, err
	}
	if t == nil {
		return nil, ErrNilToken
	}
	claim, ok := t.Claims.(*Claims)
	if !ok || !t.Valid {
		return nil, ErrInvalidClaim
	}
	return claim, nil
}

type Claims struct {
	Payload interface{}
	jwt.RegisteredClaims
}

var (
	ErrNilToken     = errors.New("nil token")
	ErrInvalidClaim = errors.New("invalid claim")
)
