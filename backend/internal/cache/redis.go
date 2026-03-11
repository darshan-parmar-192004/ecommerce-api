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

	addr := os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")

	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	if err := rdb.Ping(Ctx).Err(); err != nil {
		panic(fmt.Sprintf("Redis connection failed: %v", err))
	}

	fmt.Println("Connected to Redis:", addr)

	return &RedisService{
		Client: rdb,
	}
}