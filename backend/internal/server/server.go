package server

import (
	"backend/internal/config"
	"backend/internal/database"

	"github.com/gofiber/fiber/v3"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New(cfg *config.AppConfig) *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "backend",
			AppName:      "backend",
		}),

		db: database.New(cfg),
	}

	RegisterRoutes(server.App)

	return server
}
