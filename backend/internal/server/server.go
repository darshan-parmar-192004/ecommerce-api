package server

import (
	"os"

	"backend/internal/database"
	"backend/internal/repositories"
	"backend/internal/services"

	"github.com/gofiber/fiber/v3"
)

type FiberServer struct {
	*fiber.App

	db        repositories.Service
	cache     services.RedisService
	jwtSecret string
}

func New() *FiberServer {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "your-super-secret-jwt-key-change-in-production"
	}

	_ = database.GetDB()

	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "backend",
			AppName:      "backend",
		}),

		db:        repositories.New(),
		cache:     *services.NewRedis(),
		jwtSecret: jwtSecret,
	}

	return server
}
