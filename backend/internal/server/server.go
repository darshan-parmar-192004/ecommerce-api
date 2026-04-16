package server

import (
	"backend/internal/database"
	"backend/internal/views"

	"github.com/gofiber/fiber/v3"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "backend",
			AppName:      "backend",
		}),

		db: database.New(),
	}

	views.RegisterRoutes(server.App)

	return server
}
