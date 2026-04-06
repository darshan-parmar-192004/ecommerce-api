package unit

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/internal/cache"
	"backend/internal/stats"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestStatsHandler_NewHandler(t *testing.T) {
	t.Run("creates handler", func(t *testing.T) {
		h := stats.NewHandler()
		assert.NotNil(t, h)
	})
}

func TestStatsHandler_CacheStats(t *testing.T) {
	t.Run("returns cache stats with zero values", func(t *testing.T) {
		// Reset stats
		cache.CacheHits = 0
		cache.CacheMisses = 0

		app := fiber.New()
		h := stats.NewHandler()
		app.Get("/stats", h.CacheStats)

		req := httptest.NewRequest(http.MethodGet, "/stats", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		respStr := string(respBody)
		assert.Contains(t, respStr, `"hits"`)
		assert.Contains(t, respStr, `"misses"`)
		assert.Contains(t, respStr, `"total"`)
		assert.Contains(t, respStr, `"hit_rate"`)
	})

	t.Run("returns cache stats with hits", func(t *testing.T) {
		// Reset and set stats
		cache.CacheHits = 10
		cache.CacheMisses = 5

		app := fiber.New()
		h := stats.NewHandler()
		app.Get("/stats", h.CacheStats)

		req := httptest.NewRequest(http.MethodGet, "/stats", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		respStr := string(respBody)
		assert.Contains(t, respStr, `"hits":10`)
		assert.Contains(t, respStr, `"misses":5`)
		assert.Contains(t, respStr, `"total":15`)
		assert.Contains(t, respStr, `"hit_rate":66.666`)
	})

	t.Run("returns cache stats with only misses", func(t *testing.T) {
		// Reset and set stats
		cache.CacheHits = 0
		cache.CacheMisses = 10

		app := fiber.New()
		h := stats.NewHandler()
		app.Get("/stats", h.CacheStats)

		req := httptest.NewRequest(http.MethodGet, "/stats", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		respStr := string(respBody)
		assert.Contains(t, respStr, `"hits":0`)
		assert.Contains(t, respStr, `"misses":10`)
		assert.Contains(t, respStr, `"total":10`)
		assert.Contains(t, respStr, `"hit_rate":0`)
	})

	t.Run("returns cache stats with 100% hit rate", func(t *testing.T) {
		// Reset and set stats
		cache.CacheHits = 100
		cache.CacheMisses = 0

		app := fiber.New()
		h := stats.NewHandler()
		app.Get("/stats", h.CacheStats)

		req := httptest.NewRequest(http.MethodGet, "/stats", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		respStr := string(respBody)
		assert.Contains(t, respStr, `"hits":100`)
		assert.Contains(t, respStr, `"misses":0`)
		assert.Contains(t, respStr, `"total":100`)
		assert.Contains(t, respStr, `"hit_rate":100`)
	})

	t.Run("hit rate is 0 when total is 0", func(t *testing.T) {
		// Reset stats
		cache.CacheHits = 0
		cache.CacheMisses = 0

		app := fiber.New()
		h := stats.NewHandler()
		app.Get("/stats", h.CacheStats)

		req := httptest.NewRequest(http.MethodGet, "/stats", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		respStr := string(respBody)
		assert.Contains(t, respStr, `"hit_rate":0`)
	})
}

func TestStatsHandler_Struct(t *testing.T) {
	t.Run("Handler struct", func(t *testing.T) {
		h := stats.Handler{}
		assert.NotNil(t, h)
	})
}
