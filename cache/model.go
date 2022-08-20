package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type Config struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type SupportAPI interface {
	CacheTimeoutContext() context.Context
	CacheTimeoutContextWithDuration(t time.Duration) context.Context
}

type API interface {
	redis.Cmdable
	SupportAPI
}
