package stats

import (
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"backend/internal/services"

	"github.com/gofiber/fiber/v3"
)

func TestCacheStats_Detailed(t *testing.T) {
	t.Run("ZeroStats", func(t *testing.T) {
		handler := NewHandler()

		app := fiber.New()
		app.Get("/stats", handler.CacheStats)

		req := httptest.NewRequest(http.MethodGet, "/stats", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("NonZeroStats", func(t *testing.T) {
		services.RecordHit()
		services.RecordMiss()
		services.RecordHit()

		handler := NewHandler()

		app := fiber.New()
		app.Get("/stats", handler.CacheStats)

		req := httptest.NewRequest(http.MethodGet, "/stats", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestStatsEndpoints(t *testing.T) {
	t.Run("CacheStatsEndpoint", func(t *testing.T) {
		handler := NewHandler()

		app := fiber.New()
		app.Get("/stats/cache", handler.CacheStats)

		req := httptest.NewRequest(http.MethodGet, "/stats/cache", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestCacheHitMissReset(t *testing.T) {
	atomic.StoreInt64(&services.CacheHits, 0)
	atomic.StoreInt64(&services.CacheMisses, 0)

	_, _, total := services.GetStats()
	if total != 0 {
		t.Errorf("expected total 0, got %d", total)
	}

	services.RecordHit()
	services.RecordMiss()

	hits, misses, total := services.GetStats()
	if hits != 1 || misses != 1 || total != 2 {
		t.Errorf("expected 1 hit, 1 miss, total 2, got %d hits, %d misses, total %d", hits, misses, total)
	}
}
