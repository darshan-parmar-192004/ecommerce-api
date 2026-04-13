package unit

import (
	"os"
	"testing"

	"backend/internal/cache"

	"github.com/stretchr/testify/assert"
)

func TestCacheStats(t *testing.T) {
	// Reset stats before each test
	cache.CacheHits = 0
	cache.CacheMisses = 0

	t.Run("initial stats are zero", func(t *testing.T) {
		cache.CacheHits = 0
		cache.CacheMisses = 0

		hits, misses, total := cache.GetStats()
		assert.Equal(t, int64(0), hits)
		assert.Equal(t, int64(0), misses)
		assert.Equal(t, int64(0), total)
	})

	t.Run("record hit increments hits", func(t *testing.T) {
		cache.CacheHits = 0
		cache.CacheMisses = 0

		cache.RecordHit()

		hits, misses, total := cache.GetStats()
		assert.Equal(t, int64(1), hits)
		assert.Equal(t, int64(0), misses)
		assert.Equal(t, int64(1), total)
	})

	t.Run("record miss increments misses", func(t *testing.T) {
		cache.CacheHits = 0
		cache.CacheMisses = 0

		cache.RecordMiss()

		hits, misses, total := cache.GetStats()
		assert.Equal(t, int64(0), hits)
		assert.Equal(t, int64(1), misses)
		assert.Equal(t, int64(1), total)
	})

	t.Run("multiple hits and misses", func(t *testing.T) {
		cache.CacheHits = 0
		cache.CacheMisses = 0

		cache.RecordHit()
		cache.RecordHit()
		cache.RecordMiss()
		cache.RecordHit()
		cache.RecordMiss()

		hits, misses, total := cache.GetStats()
		assert.Equal(t, int64(3), hits)
		assert.Equal(t, int64(2), misses)
		assert.Equal(t, int64(5), total)
	})

	t.Run("stats are cumulative", func(t *testing.T) {
		cache.CacheHits = 0
		cache.CacheMisses = 0

		cache.RecordHit()
		hits1, _, _ := cache.GetStats()
		assert.Equal(t, int64(1), hits1)

		cache.RecordHit()
		hits2, _, _ := cache.GetStats()
		assert.Equal(t, int64(2), hits2)
	})
}

func TestRedisService_NewRedis(t *testing.T) {
	t.Run("returns empty service in test env", func(t *testing.T) {
		origEnv := os.Getenv("APP_ENV")
		_ = os.Setenv("APP_ENV", "test")
		defer func() { _ = os.Setenv("APP_ENV", origEnv) }()

		svc := cache.NewRedis()
		assert.NotNil(t, svc)
		assert.Nil(t, svc.Client)
	})

	t.Run("uses default host and port when env vars not set", func(t *testing.T) {
		origEnv := os.Getenv("APP_ENV")
		origHost := os.Getenv("REDIS_HOST")
		origPort := os.Getenv("REDIS_PORT")

		_ = os.Setenv("APP_ENV", "production")
		_ = os.Setenv("REDIS_HOST", "")
		_ = os.Setenv("REDIS_PORT", "")
		defer func() {
			_ = os.Setenv("APP_ENV", origEnv)
			_ = os.Setenv("REDIS_HOST", origHost)
			_ = os.Setenv("REDIS_PORT", origPort)
		}()

		// This will try to connect to localhost:6379 and fail, but should not panic
		svc := cache.NewRedis()
		assert.NotNil(t, svc)
	})

	t.Run("uses custom host and port from env vars", func(t *testing.T) {
		origEnv := os.Getenv("APP_ENV")
		origHost := os.Getenv("REDIS_HOST")
		origPort := os.Getenv("REDIS_PORT")

		_ = os.Setenv("APP_ENV", "production")
		_ = os.Setenv("REDIS_HOST", "my-redis-host")
		_ = os.Setenv("REDIS_PORT", "6380")
		defer func() {
			_ = os.Setenv("APP_ENV", origEnv)
			_ = os.Setenv("REDIS_HOST", origHost)
			_ = os.Setenv("REDIS_PORT", origPort)
		}()

		// This will try to connect to my-redis-host:6380 and fail, but should not panic
		svc := cache.NewRedis()
		assert.NotNil(t, svc)
	})

	t.Run("empty RedisService struct", func(t *testing.T) {
		svc := cache.RedisService{}
		assert.Nil(t, svc.Client)
	})
}

func TestCacheCtx(t *testing.T) {
	t.Run("Ctx is not nil", func(t *testing.T) {
		assert.NotNil(t, cache.Ctx)
	})
}
