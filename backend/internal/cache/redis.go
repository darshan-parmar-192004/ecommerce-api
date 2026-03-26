package cache

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

type RedisService struct {
	Client *redis.Client
}

func (r *RedisService) Ping() error {
	if r.Client == nil {
		return fmt.Errorf("redis client not initialized")
	}
	return r.Client.Ping(Ctx).Err()
}

func NewRedis() *RedisService {
	host := os.Getenv("REDIS_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("REDIS_PORT")
	if port == "" {
		port = "6379"
	}

	if os.Getenv("APP_ENV") == "unit_test_no_redis" || os.Getenv("APP_ENV") == "test" {
		return &RedisService{}
	}

	addr := host + ":" + port

	rdb := redis.NewClient(&redis.Options{
		Addr:            addr,
		DialTimeout:     5 * time.Second,
		MaxRetries:      5,
		MinRetryBackoff: 1 * time.Second,
	})

	var err error

	for i := 0; i < 5; i++ {
		err = rdb.Ping(Ctx).Err()
		if err == nil {
			fmt.Println("Successfully connected to Redis at:", addr)
			return &RedisService{Client: rdb}
		}

		fmt.Printf("Attempt %d: Redis not ready at %s (error: %v), retrying in 2s...\n", i+1, addr, err)
		time.Sleep(2 * time.Second)
	}

	fmt.Printf("Final Redis connection failure after retries: %v\n", err)
	return &RedisService{Client: rdb}
}
