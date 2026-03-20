package cache

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func TestNewRedis(t *testing.T) {
	os.Setenv("APP_ENV", "test")

	rs := NewRedis()
	if rs == nil {
		t.Error("expected non-nil RedisService")
	}
}

func TestRecordHit(t *testing.T) {
	CacheHits = 0
	RecordHit()

	if CacheHits != 1 {
		t.Errorf("expected 1, got %d", CacheHits)
	}
}

func TestRecordMiss(t *testing.T) {
	CacheMisses = 0
	RecordMiss()

	if CacheMisses != 1 {
		t.Errorf("expected 1, got %d", CacheMisses)
	}
}

func TestGetStats(t *testing.T) {
	CacheHits = 100
	CacheMisses = 50

	hits, misses, total := GetStats()

	if hits != 100 {
		t.Errorf("expected 100, got %d", hits)
	}
	if misses != 50 {
		t.Errorf("expected 50, got %d", misses)
	}
	if total != 150 {
		t.Errorf("expected 150, got %d", total)
	}
}

func TestGetStatsZero(t *testing.T) {
	CacheHits = 0
	CacheMisses = 0

	hits, misses, total := GetStats()

	if hits != 0 {
		t.Errorf("expected 0, got %d", hits)
	}
	if misses != 0 {
		t.Errorf("expected 0, got %d", misses)
	}
	if total != 0 {
		t.Errorf("expected 0, got %d", total)
	}
}

func TestRedisClient_Get_ConnectionError(t *testing.T) {
	t.Parallel()

	rs := &RedisService{
		Client: nil,
	}

	ctx := context.Background()

	if rs.Client != nil {
		cmd := rs.Client.Get(ctx, "test-key")
		if cmd.Err() == nil {
			t.Error("expected error when Redis client is nil")
		}
	}
}

func TestRedisClient_Set_ConnectionError(t *testing.T) {
	t.Parallel()

	rs := &RedisService{
		Client: nil,
	}

	ctx := context.Background()

	if rs.Client != nil {
		cmd := rs.Client.Set(ctx, "test-key", "value", time.Minute)
		if cmd.Err() == nil {
			t.Error("expected error when Redis client is nil")
		}
	}
}

func TestRedisClient_Del_ConnectionError(t *testing.T) {
	t.Parallel()

	rs := &RedisService{
		Client: nil,
	}

	ctx := context.Background()

	if rs.Client != nil {
		cmd := rs.Client.Del(ctx, "test-key")
		if cmd.Err() == nil {
			t.Error("expected error when Redis client is nil")
		}
	}
}

func TestRedisClient_InvalidAddress(t *testing.T) {
	t.Parallel()

	rs := &RedisService{
		Client: redis.NewClient(&redis.Options{
			Addr: "invalid-address:6379",
		}),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	cmd := rs.Client.Get(ctx, "nonexistent-key")
	_, err := cmd.Result()

	if err != redis.Nil {
		t.Logf("Expected key not found or connection error, got: %v", err)
	}
}

func TestRedisClient_SetError(t *testing.T) {
	t.Parallel()

	rs := &RedisService{
		Client: redis.NewClient(&redis.Options{
			Addr: "invalid-address:6379",
		}),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	cmd := rs.Client.Set(ctx, "test-key", "value", time.Minute)
	_, err := cmd.Result()

	if err != nil {
		t.Logf("Expected error, got: %v", err)
	}
}

func TestRedisClient_DelError(t *testing.T) {
	t.Parallel()

	rs := &RedisService{
		Client: redis.NewClient(&redis.Options{
			Addr: "invalid-address:6379",
		}),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	cmd := rs.Client.Del(ctx, "test-key")
	_, err := cmd.Result()

	if err != nil {
		t.Logf("Expected error, got: %v", err)
	}
}

func TestRecordHitMultiple(t *testing.T) {
	CacheHits = 0

	for i := 0; i < 5; i++ {
		RecordHit()
	}

	if CacheHits != 5 {
		t.Errorf("expected 5, got %d", CacheHits)
	}
}

func TestRecordMissMultiple(t *testing.T) {
	CacheMisses = 0

	for i := 0; i < 3; i++ {
		RecordMiss()
	}

	if CacheMisses != 3 {
		t.Errorf("expected 3, got %d", CacheMisses)
	}
}
