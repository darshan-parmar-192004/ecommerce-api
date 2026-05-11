package server

import (
	"backend/internal/constants"
	"backend/internal/database"

	"github.com/gofiber/fiber/v3"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: constants.ServerHeader,
			AppName:      constants.AppName,
		}),

		db: database.New(),
	}

	RegisterRoutes(server.App)

	return server
}
