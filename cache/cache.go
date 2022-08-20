package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type _Cache struct {
	*redis.Client
}

func (curr *_Cache) CacheTimeoutContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	return ctx
}

func (curr *_Cache) CacheTimeoutContextWithDuration(t time.Duration) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), t)
	return ctx
}

func New(config *Config) (API, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}
	return &_Cache{client}, nil
}
