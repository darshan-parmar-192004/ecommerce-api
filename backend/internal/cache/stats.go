package cache

import (
	"sync/atomic"
)

type CacheStats struct {
	hitCount    atomic.Uint64
	missCount   atomic.Uint64
	totalCount  atomic.Uint64
}

func NewStats() *CacheStats {
	return &CacheStats{}
}

func (s *CacheStats) RecordHit() {
	s.hitCount.Add(1)
	s.totalCount.Add(1)
}

func (s *CacheStats) RecordMiss() {
	s.missCount.Add(1)
	s.totalCount.Add(1)
}

func (s *CacheStats) Stats() map[string]interface{} {
	hits := s.hitCount.Load()
	misses := s.missCount.Load()
	total := s.totalCount.Load()

	var hitRate, missRate float64
	if total > 0 {
		hitRate = float64(hits) / float64(total)
		missRate = float64(misses) / float64(total)
	}

	return map[string]interface{}{
		"hit_count":     hits,
		"miss_count":    misses,
		"total_requests": total,
		"hit_rate":      hitRate,
		"miss_rate":     missRate,
	}
}
