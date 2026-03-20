package cache

import (
	"os"
	"sync/atomic"
	"testing"
)

func TestRedisService_New(t *testing.T) {
	t.Run("CreatesNewRedisService", func(t *testing.T) {
		redis := NewRedis()
		if redis == nil {
			t.Error("expected redis service to not be nil")
		}
	})

	t.Run("UsesEnvironmentVariables", func(t *testing.T) {
		oldHost := os.Getenv("REDIS_HOST")
		oldPort := os.Getenv("REDIS_PORT")
		oldEnv := os.Getenv("APP_ENV")

		os.Setenv("REDIS_HOST", "localhost")
		os.Setenv("REDIS_PORT", "6379")
		os.Setenv("APP_ENV", "production")

		redis := NewRedis()
		if redis == nil {
			t.Error("expected redis service to not be nil")
		}

		os.Setenv("REDIS_HOST", oldHost)
		os.Setenv("REDIS_PORT", oldPort)
		os.Setenv("APP_ENV", oldEnv)
	})

	t.Run("ReturnsEmptyServiceInTestEnv", func(t *testing.T) {
		oldEnv := os.Getenv("APP_ENV")
		os.Setenv("APP_ENV", "test")

		redis := NewRedis()
		if redis == nil {
			t.Error("expected redis service to not be nil")
			return
		}
		if redis.Client != nil {
			t.Error("expected nil client in test env")
		}

		os.Setenv("APP_ENV", oldEnv)
	})

	t.Run("UsesDefaultHostWhenNotSet", func(t *testing.T) {
		oldHost := os.Getenv("REDIS_HOST")
		oldEnv := os.Getenv("APP_ENV")

		os.Unsetenv("REDIS_HOST")
		os.Setenv("APP_ENV", "production")

		redis := NewRedis()
		if redis == nil {
			t.Error("expected redis service to not be nil")
		}

		os.Setenv("REDIS_HOST", oldHost)
		os.Setenv("APP_ENV", oldEnv)
	})

	t.Run("UsesDefaultPortWhenNotSet", func(t *testing.T) {
		oldPort := os.Getenv("REDIS_PORT")
		oldEnv := os.Getenv("APP_ENV")

		os.Unsetenv("REDIS_PORT")
		os.Setenv("APP_ENV", "production")

		redis := NewRedis()
		if redis == nil {
			t.Error("expected redis service to not be nil")
		}

		os.Setenv("REDIS_PORT", oldPort)
		os.Setenv("APP_ENV", oldEnv)
	})
}

func TestRedisStats(t *testing.T) {
	t.Run("InitialStats", func(t *testing.T) {
		hits, misses, total := GetStats()

		if total < 0 {
			t.Error("total should be non-negative")
		}
		if hits < 0 {
			t.Error("hits should be non-negative")
		}
		if misses < 0 {
			t.Error("misses should be non-negative")
		}
	})

	t.Run("HitRateCalculation", func(t *testing.T) {
		atomic.StoreInt64(&CacheHits, 2)
		atomic.StoreInt64(&CacheMisses, 1)

		hits, misses, total := GetStats()

		if total != 3 {
			t.Errorf("expected total 3, got %d", total)
		}
		if hits != 2 {
			t.Errorf("expected hits 2, got %d", hits)
		}
		if misses != 1 {
			t.Errorf("expected misses 1, got %d", misses)
		}

		hitRate := float64(hits) / float64(total)
		if hitRate != 2.0/3.0 {
			t.Errorf("expected hit rate %f, got %f", 2.0/3.0, hitRate)
		}

		atomic.StoreInt64(&CacheHits, 0)
		atomic.StoreInt64(&CacheMisses, 0)
	})

	t.Run("ZeroTotal", func(t *testing.T) {
		atomic.StoreInt64(&CacheHits, 0)
		atomic.StoreInt64(&CacheMisses, 0)

		hits, misses, total := GetStats()

		if total != 0 {
			t.Errorf("expected total 0, got %d", total)
		}
		if hits != 0 {
			t.Errorf("expected hits 0, got %d", hits)
		}
		if misses != 0 {
			t.Errorf("expected misses 0, got %d", misses)
		}
	})
}

func TestRecordFunctions(t *testing.T) {
	t.Run("RecordHitIncrements", func(t *testing.T) {
		atomic.StoreInt64(&CacheHits, 0)

		RecordHit()

		finalHits := atomic.LoadInt64(&CacheHits)

		if finalHits != 1 {
			t.Errorf("expected hits to be 1, got %d", finalHits)
		}

		atomic.StoreInt64(&CacheHits, 0)
	})

	t.Run("RecordMissIncrements", func(t *testing.T) {
		atomic.StoreInt64(&CacheMisses, 0)

		RecordMiss()

		finalMisses := atomic.LoadInt64(&CacheMisses)

		if finalMisses != 1 {
			t.Errorf("expected misses to be 1, got %d", finalMisses)
		}

		atomic.StoreInt64(&CacheMisses, 0)
	})

	t.Run("MultipleHitsAndMisses", func(t *testing.T) {
		atomic.StoreInt64(&CacheHits, 0)
		atomic.StoreInt64(&CacheMisses, 0)

		for i := 0; i < 5; i++ {
			RecordHit()
		}
		for i := 0; i < 3; i++ {
			RecordMiss()
		}

		hits, misses, total := GetStats()
		if hits != 5 {
			t.Errorf("expected 5 hits, got %d", hits)
		}
		if misses != 3 {
			t.Errorf("expected 3 misses, got %d", misses)
		}
		if total != 8 {
			t.Errorf("expected 8 total, got %d", total)
		}

		atomic.StoreInt64(&CacheHits, 0)
		atomic.StoreInt64(&CacheMisses, 0)
	})
}
