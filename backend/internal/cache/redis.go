package cache

import (
	"context"

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
		logger.Log.Warnw("Failed to load config for Redis", "error", err)
		return &RedisService{}
	}

	if cfg.AppEnv == "test" {
		return &RedisService{}
	}

	addr := cfg.RedisHost + ":" + cfg.RedisPort

	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	if err := rdb.Ping(Ctx).Err(); err != nil {
		logger.Log.Warnw("Redis connection failed", "error", err)
	}

	logger.Log.Info("Connected to Redis", "address", addr)

	return &RedisService{
		Client: rdb,
	}
}
