package token

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type _Token struct {
	config *Config
}

func New(config *Config) API {
	return &_Token{config: config}
}

func (curr *_Token) Create(id string) (string, error) {
	claims := jwt.RegisteredClaims{
		Issuer:    curr.config.Issuer,
		ExpiresAt: jwt.NewNumericDate(time.Unix(time.Now().Unix()+curr.config.ExpiresTime, 0)),
		IssuedAt:  jwt.NewNumericDate(time.Unix(time.Now().Unix(), 0)),
		ID:        id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if tokenStr, err := token.SignedString([]byte(curr.config.SigningKey)); err == nil {
		return tokenStr, nil
	} else {
		return "", err
	}
}

func (curr *_Token) Parse(tokenStr string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(curr.config.SigningKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims.(*jwt.RegisteredClaims), nil
}

func (curr *_Token) FlushTokenWithClaims(claims *jwt.RegisteredClaims) (string, error) {
	claims.ExpiresAt = jwt.NewNumericDate(time.Unix(time.Now().Unix()+curr.config.ExpiresTime, 0))
	claims.IssuedAt = jwt.NewNumericDate(time.Unix(time.Now().Unix(), 0))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if tokenStr, err := token.SignedString([]byte(curr.config.SigningKey)); err == nil {
		return tokenStr, nil
	} else {
		return "", err
	}
}

func (curr *_Token) Issuer() string {
	return curr.config.Issuer
}

func (curr *_Token) SigningKey() string {
	return curr.config.SigningKey
}

func (curr *_Token) ExpiresTime() int64 {
	return curr.config.ExpiresTime
}

func (curr *_Token) BufferTime() int64 {
	return curr.config.BufferTime
}
