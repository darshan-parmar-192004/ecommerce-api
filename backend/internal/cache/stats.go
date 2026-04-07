package cache

import "sync/atomic"

var CacheHits int64
var CacheMisses int64

func RecordHit() {
	atomic.AddInt64(&CacheHits, 1)
}

func RecordMiss() {
	atomic.AddInt64(&CacheMisses, 1)
}

func GetStats() (hits int64, misses int64, total int64) {
	hits = atomic.LoadInt64(&CacheHits)
	misses = atomic.LoadInt64(&CacheMisses)
	total = hits + misses
	return
}
