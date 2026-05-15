package server

import (
	"backend/internal/cache"
	"backend/internal/constants"
	"backend/internal/database"

	"github.com/gofiber/fiber/v3"
)

type FiberServer struct {
	*fiber.App

	db    database.Service
	cache cache.Service
	stats *cache.CacheStats
}

func New() *FiberServer {
	cacheSvc := cache.New()
	stats := cache.NewStats()

	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: constants.ServerHeader,
			AppName:      constants.AppName,
		}),

		db:    database.New(),
		cache: cacheSvc,
		stats: stats,
	}

	RegisterRoutes(server.App, cacheSvc, stats)

	return server
}

func (s *FiberServer) CloseCache() error {
	return s.cache.Close()
}
