package stats

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestCacheStats(t *testing.T) {
	t.Run("ReturnsCacheStatistics", func(t *testing.T) {
		handler := &Handler{}
		app := fiber.New()
		app.Get("/stats/cache", handler.CacheStats)

		req := httptest.NewRequest(http.MethodGet, "/stats/cache", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusOK {
			t.Fatalf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("ReturnsJSONWithHits", func(t *testing.T) {
		handler := &Handler{}
		app := fiber.New()
		app.Get("/stats/cache", handler.CacheStats)

		req := httptest.NewRequest(http.MethodGet, "/stats/cache", nil)
		resp, _ := app.Test(req)

		var result map[string]interface{}
		err := json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if _, ok := result["hits"]; !ok {
			t.Fatalf("expected hits field in response")
		}
	})

	t.Run("ReturnsJSONWithMisses", func(t *testing.T) {
		handler := &Handler{}
		app := fiber.New()
		app.Get("/stats/cache", handler.CacheStats)

		req := httptest.NewRequest(http.MethodGet, "/stats/cache", nil)
		resp, _ := app.Test(req)

		var result map[string]interface{}
		err := json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if _, ok := result["misses"]; !ok {
			t.Fatalf("expected misses field in response")
		}
	})

	t.Run("ReturnsJSONWithTotalRequests", func(t *testing.T) {
		handler := &Handler{}
		app := fiber.New()
		app.Get("/stats/cache", handler.CacheStats)

		req := httptest.NewRequest(http.MethodGet, "/stats/cache", nil)
		resp, _ := app.Test(req)

		var result map[string]interface{}
		err := json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if _, ok := result["total_requests"]; !ok {
			t.Fatalf("expected total_requests field in response")
		}
	})

	t.Run("ReturnsJSONWithHitRate", func(t *testing.T) {
		handler := &Handler{}
		app := fiber.New()
		app.Get("/stats/cache", handler.CacheStats)

		req := httptest.NewRequest(http.MethodGet, "/stats/cache", nil)
		resp, _ := app.Test(req)

		var result map[string]interface{}
		err := json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if _, ok := result["hit_rate"]; !ok {
			t.Fatalf("expected hit_rate field in response")
		}
	})

	t.Run("ReturnsJSONWithMissRate", func(t *testing.T) {
		handler := &Handler{}
		app := fiber.New()
		app.Get("/stats/cache", handler.CacheStats)

		req := httptest.NewRequest(http.MethodGet, "/stats/cache", nil)
		resp, _ := app.Test(req)

		var result map[string]interface{}
		err := json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if _, ok := result["miss_rate"]; !ok {
			t.Fatalf("expected miss_rate field in response")
		}
	})
}

func TestCacheStatsValues(t *testing.T) {
	t.Run("HitRateCalculation", func(t *testing.T) {
		app := fiber.New()
		handler := &Handler{}
		app.Get("/stats/cache", handler.CacheStats)

		req := httptest.NewRequest(http.MethodGet, "/stats/cache", nil)
		resp, _ := app.Test(req)

		var result map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		hitRate, ok := result["hit_rate"].(float64)
		if ok {
			if hitRate < 0.0 || hitRate > 1.0 {
				t.Fatalf("expected hit_rate between 0 and 1, got %f", hitRate)
			}
		}
	})

	t.Run("MissRateCalculation", func(t *testing.T) {
		app := fiber.New()
		handler := &Handler{}
		app.Get("/stats/cache", handler.CacheStats)

		req := httptest.NewRequest(http.MethodGet, "/stats/cache", nil)
		resp, _ := app.Test(req)

		var result map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		missRate, ok := result["miss_rate"].(float64)
		if ok {
			if missRate < 0.0 || missRate > 1.0 {
				t.Fatalf("expected miss_rate between 0 and 1, got %f", missRate)
			}
		}
	})

	t.Run("RatesSumToOne", func(t *testing.T) {
		app := fiber.New()
		handler := &Handler{}
		app.Get("/stats/cache", handler.CacheStats)

		req := httptest.NewRequest(http.MethodGet, "/stats/cache", nil)
		resp, _ := app.Test(req)

		var result map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		hitRate, hitOk := result["hit_rate"].(float64)
		missRate, missOk := result["miss_rate"].(float64)

		if hitOk && missOk && (hitRate > 0 || missRate > 0) {
			sum := hitRate + missRate
			if sum < 0.999 || sum > 1.001 {
				t.Fatalf("expected rates to sum to 1, got %f", sum)
			}
		}
	})
}

func TestStatsRoutes(t *testing.T) {
	t.Run("CacheStatsRoute", func(t *testing.T) {
		app := fiber.New()
		handler := &Handler{}
		app.Get("/stats/cache", handler.CacheStats)

		req := httptest.NewRequest(http.MethodGet, "/stats/cache", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Fatalf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("AlternativeCacheRoute", func(t *testing.T) {
		app := fiber.New()
		handler := &Handler{}
		app.Get("/cache/stats", handler.CacheStats)

		req := httptest.NewRequest(http.MethodGet, "/cache/stats", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Fatalf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestStatsContentType(t *testing.T) {
	t.Run("ReturnsJSONContentType", func(t *testing.T) {
		app := fiber.New()
		handler := &Handler{}
		app.Get("/stats/cache", handler.CacheStats)

		req := httptest.NewRequest(http.MethodGet, "/stats/cache", nil)
		resp, _ := app.Test(req)

		contentType := resp.Header.Get("Content-Type")
		if contentType == "" {
			t.Fatalf("expected content type, got empty")
		}
	})
}

func TestStatsInvalidMethods(t *testing.T) {
	t.Run("GetReturnsOK", func(t *testing.T) {
		app := fiber.New()
		handler := &Handler{}
		app.Get("/stats/cache", handler.CacheStats)

		req := httptest.NewRequest(http.MethodGet, "/stats/cache", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Fatalf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestCacheStatsResponseFormat(t *testing.T) {
	t.Run("AllExpectedFieldsPresent", func(t *testing.T) {
		app := fiber.New()
		handler := &Handler{}
		app.Get("/stats/cache", handler.CacheStats)

		req := httptest.NewRequest(http.MethodGet, "/stats/cache", nil)
		resp, _ := app.Test(req)

		var result map[string]interface{}
		err := json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedFields := []string{"hits", "misses", "total_requests", "hit_rate", "miss_rate"}
		for _, field := range expectedFields {
			if _, ok := result[field]; !ok {
				t.Fatalf("expected field %s in response", field)
			}
		}
	})
}

func TestNewHandler(t *testing.T) {
	t.Run("CreatesHandlerInstance", func(t *testing.T) {
		handler := NewHandler()
		if handler == nil {
			t.Fatalf("expected handler, got nil")
		}
	})
}
