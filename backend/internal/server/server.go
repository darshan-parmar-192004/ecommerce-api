package server

import (
	"github.com/gofiber/fiber/v3"

	"backend/internal/database"
	"backend/internal/cache"
)

type FiberServer struct {
	*fiber.App

	db database.Service
	cache cache.RedisService
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "backend",
			AppName:      "backend",
		}),

		db: database.New(),
		cache: *cache.NewRedis(),
	}

	return server
}
