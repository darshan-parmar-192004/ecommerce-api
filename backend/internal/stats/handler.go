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

	hitRate := float64(0)
	missRate := float64(0)

	if total > 0 {
		hitRate = float64(hits) / float64(total)
		missRate = float64(misses) / float64(total)
	}

	return c.JSON(fiber.Map{
		"hits":           hits,
		"misses":         misses,
		"total_requests": total,
		"hit_rate":       hitRate,
		"miss_rate":      missRate,
	})
}
