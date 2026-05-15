package cache

import (
	"context"
	"fmt"
	"time"

	"backend/internal/config"
	"backend/internal/logger"

	"github.com/redis/go-redis/v9"
)

type Service interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string, ttl time.Duration) error
	Del(ctx context.Context, key string) error
	DelPattern(ctx context.Context, pattern string) error
	Close() error
}

type service struct {
	client *redis.Client
}

var cacheInstance Service

func New() Service {
	if cacheInstance != nil {
		return cacheInstance
	}

	cfg, err := config.Load()
	if err != nil {
		logger.Log.Fatalf("failed to load config: %v", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		logger.Log.Warnf("redis connection failed: %v (cache operations will fail open)", err)
	}

	logger.Log.Infof("redis connected: %s:%s", cfg.RedisHost, cfg.RedisPort)

	cacheInstance = &service{
		client: client,
	}
	return cacheInstance
}

func (s *service) Get(ctx context.Context, key string) (string, error) {
	return s.client.Get(ctx, key).Result()
}

func (s *service) Set(ctx context.Context, key string, value string, ttl time.Duration) error {
	return s.client.Set(ctx, key, value, ttl).Err()
}

func (s *service) Del(ctx context.Context, key string) error {
	return s.client.Del(ctx, key).Err()
}

func (s *service) DelPattern(ctx context.Context, pattern string) error {
	iter := s.client.Scan(ctx, 0, pattern, 0).Iterator()
	var keys []string
	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}
	if err := iter.Err(); err != nil {
		return err
	}
	if len(keys) > 0 {
		return s.client.Del(ctx, keys...).Err()
	}
	return nil
}

func (s *service) Close() error {
	if logger.Log != nil {
		logger.Log.Info("Closing redis connection")
	}
	return s.client.Close()
}

