package middleware

import (
	"sync"
	"time"

	"github.com/gofiber/fiber/v3"
)

type RateLimiter struct {
	requests map[string][]time.Time
	mu       sync.RWMutex
	limit    int
	window   time.Duration
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	rl := &RateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		window:   window,
	}

	go rl.cleanup()

	return rl
}

func (rl *RateLimiter) cleanup() {
	ticker := time.NewTicker(rl.window)
	for range ticker.C {
		rl.mu.Lock()
		now := time.Now()
		for key, times := range rl.requests {
			var valid []time.Time
			for _, t := range times {
				if now.Sub(t) < rl.window {
					valid = append(valid, t)
				}
			}
			if len(valid) == 0 {
				delete(rl.requests, key)
			} else {
				rl.requests[key] = valid
			}
		}
		rl.mu.Unlock()
	}
}

func (rl *RateLimiter) Allow(key string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	times := rl.requests[key]

	var valid []time.Time
	for _, t := range times {
		if now.Sub(t) < rl.window {
			valid = append(valid, t)
		}
	}

	if len(valid) >= rl.limit {
		rl.requests[key] = valid
		return false
	}

	rl.requests[key] = append(valid, now)
	return true
}

func RateLimiterMiddleware(rl *RateLimiter) fiber.Handler {
	return func(c fiber.Ctx) error {
		key := c.IP()

		if !rl.Allow(key) {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": fiber.Map{
					"code":    "RATE_LIMIT_EXCEEDED",
					"message": "Too many requests. Please try again later.",
				},
			})
		}

		return c.Next()
	}
}

var (
	GeneralRateLimiter = NewRateLimiter(100, time.Minute)
	LoginRateLimiter   = NewRateLimiter(50, time.Minute)
)

func GeneralRateLimit() fiber.Handler {
	return RateLimiterMiddleware(GeneralRateLimiter)
}

func LoginRateLimit() fiber.Handler {
	return RateLimiterMiddleware(LoginRateLimiter)
}

func AuthenticatedRateLimit() fiber.Handler {
	return func(c fiber.Ctx) error {
		customerID := GetCustomerID(c)
		key := customerID
		if key == "" {
			key = c.IP()
		}

		rl := NewRateLimiter(50, time.Minute)

		if !rl.Allow(key) {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": fiber.Map{
					"code":    "RATE_LIMIT_EXCEEDED",
					"message": "Too many requests. Please try again later.",
				},
			})
		}

		return c.Next()
	}
}
