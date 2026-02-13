package server

import (
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

}
