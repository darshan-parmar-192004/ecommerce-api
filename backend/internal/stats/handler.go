package stats

import (
	"backend/internal/cache"

	"github.com/gofiber/fiber/v3"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) CacheStats(c fiber.Ctx) error {
	hits, misses, total := cache.GetStats()
	var hitRate float64
	if total > 0 {
		hitRate = float64(hits) / float64(total) * 100
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"hits":     hits,
		"misses":   misses,
		"total":    total,
		"hit_rate": hitRate,
	})
}
