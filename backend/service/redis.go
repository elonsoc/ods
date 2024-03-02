package service

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisIFace interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
}

type RedisService struct {
	client *redis.Client
}

func initRedis(redisURL string, log LoggerIFace) RedisIFace {
    opts, err := redis.ParseURL(redisURL)
    if err != nil {
        log.Fatal(err)
        return nil
    }

    client := redis.NewClient(opts)

    _, err = client.Ping(context.Background()).Result()
    if err != nil {
        log.Fatal(err)
        return nil
    }

    return &RedisService{client: client}
}

func (r *RedisService) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

func (r *RedisService) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}
