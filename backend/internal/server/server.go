package server

import (
	"os"

	"github.com/gofiber/fiber/v3"

	"backend/internal/cache"
	"backend/internal/database"
)

type FiberServer struct {
	*fiber.App

	db        database.Service
	cache     cache.RedisService
	jwtSecret string
}

func New() *FiberServer {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "your-super-secret-jwt-key-change-in-production"
	}

	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "backend",
			AppName:      "backend",
		}),

		db:        database.New(),
		cache:     *cache.NewRedis(),
		jwtSecret: jwtSecret,
	}

	return server
}
