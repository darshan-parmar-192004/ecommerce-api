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
	port := os.Getenv("REDIS_PORT")

	addr := fmt.Sprintf("%s:%s", host, port)

	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	return &RedisService{
		Client: rdb,
	}
}