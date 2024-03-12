package service

import (
	"context"
	"time"

	"github.com/elonsoc/ods/src/common"
	"github.com/go-redis/redis/v8"
)

type InMemoryDbIFace interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) error
}

type InMemoryDbService struct {
	client *redis.Client
}

func initInMemoryDb(redisURL string, log common.LoggerIFace) InMemoryDbIFace {
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

	return &InMemoryDbService{client: client}
}

func (r *InMemoryDbService) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

func (r *InMemoryDbService) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *InMemoryDbService) Del(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}
