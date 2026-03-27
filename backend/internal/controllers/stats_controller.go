package controllers

import (
	"backend/internal/services"

	"github.com/gofiber/fiber/v3"
)

type StatsController struct{}

func NewStatsController() *StatsController {
	return &StatsController{}
}

func (ctrl *StatsController) CacheStats(c fiber.Ctx) error {
	hits, misses, total := services.GetStats()

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
