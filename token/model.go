package token

import (
	"github.com/golang-jwt/jwt/v4"
)

type Config struct {
	Issuer      string `yaml:"issuer"`
	SigningKey  string `yaml:"signingKey"`
	ExpiresTime int64  `yaml:"expiresTime"`
	BufferTime  int64  `yaml:"bufferTime"`
}

type API interface {
	Create(id string) (token string, err error)
	Parse(token string) (*jwt.RegisteredClaims, error)
	FlushTokenWithClaims(claims *jwt.RegisteredClaims) (token string, err error)

	Issuer() string
	SigningKey() string
	ExpiresTime() int64
	BufferTime() int64
}
