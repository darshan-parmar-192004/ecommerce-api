package cache

import (
	"context"
	"time"

	"backend/internal/config"
	"backend/internal/logger"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

type RedisService struct {
	Client *redis.Client
}

func NewRedis() *RedisService {
	cfg, err := config.Load()
	if err != nil {
		cfg = &config.AppConfig{}
	}

	if cfg.AppEnv == "test" {
		return &RedisService{}
	}

	addr := cfg.RedisHost + ":" + cfg.RedisPort

	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	if err := rdb.Ping(Ctx).Err(); err != nil {
		logger.Log.Errorf("Redis connection failed: %v", err)
		return &RedisService{} // no client
	}

	logger.Log.Infof("Connected to Redis: %s", addr)

	return &RedisService{
		Client: rdb,
	}
}

func (r *RedisService) Get(key string) (string, error) {
	if r.Client == nil {
		return "", redis.Nil
	}
	return r.Client.Get(Ctx, key).Result()
}

func (r *RedisService) GetWithTTL(key string, ttl time.Duration) (string, error) {
	if r.Client == nil {
		return "", redis.Nil
	}
	val, err := r.Client.Get(Ctx, key).Result()
	if err == redis.Nil {
		return "", err
	}
	if err != nil {
		return "", err
	}
	r.Client.Expire(Ctx, key, ttl)
	return val, nil
}

func (r *RedisService) Set(key string, value interface{}, ttl time.Duration) error {
	if r.Client == nil {
		return nil
	}
	return r.Client.Set(Ctx, key, value, ttl).Err()
}

func (r *RedisService) Delete(key string) error {
	if r.Client == nil {
		return nil
	}
	return r.Client.Del(Ctx, key).Err()
}

func (r *RedisService) DeletePattern(pattern string) error {
	if r.Client == nil {
		return nil
	}
	iter := r.Client.Scan(Ctx, 0, pattern, 0).Iterator()
	for iter.Next(Ctx) {
		if err := r.Client.Del(Ctx, iter.Val()).Err(); err != nil {
			return err
		}
	}
	return iter.Err()
}
