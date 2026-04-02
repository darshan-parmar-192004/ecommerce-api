package cache

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

type RedisService struct {
	Client *redis.Client
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

	if os.Getenv("APP_ENV") == "test" {
		return &RedisService{}
	}

	addr := host + ":" + port

	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	if err := rdb.Ping(Ctx).Err(); err != nil {
		fmt.Printf("Redis connection failed: %v", err)
	}

	fmt.Println("Connected to Redis:", addr)

	return &RedisService{
		Client: rdb,
	}
}
