package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	ctx    context.Context
	client *redis.Client
}

func NewRedisClient(ctx context.Context, opts *redis.Options) Cache {
	return &RedisCache{
		ctx:    ctx,
		client: redis.NewClient(opts),
	}
}

func (c RedisCache) Set(key string, value interface{}, expiration time.Duration) error {
	return c.client.Set(c.ctx, key, value, expiration).Err()
}

func (c RedisCache) Get(key string) (string, error) {
	return c.client.Get(c.ctx, key).Result()
}

func (c RedisCache) Delete(key string) (int64, error) {
	return c.client.Del(c.ctx, key).Result()
}
