package middleware

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
)

type SecurityLogger struct {
	logger *log.Logger
}

func NewSecurityLogger() *SecurityLogger {
	return &SecurityLogger{
		logger: log.Default(),
	}
}

func (sl *SecurityLogger) LogAuthAttempt(email string, success bool, ip string) {
	sl.logger.Printf("[SECURITY] Auth attempt: email=%s success=%t ip=%s timestamp=%s",
		email, success, ip, time.Now().UTC().Format(time.RFC3339))
}

func (sl *SecurityLogger) LogAuthFailure(email string, reason string, ip string) {
	sl.logger.Printf("[SECURITY] Auth failure: email=%s reason=%s ip=%s timestamp=%s",
		email, reason, ip, time.Now().UTC().Format(time.RFC3339))
}

func (sl *SecurityLogger) LogAuthorizationFailure(customerID string, resource string, ip string) {
	sl.logger.Printf("[SECURITY] Authorization failure: customer_id=%s resource=%s ip=%s timestamp=%s",
		customerID, resource, ip, time.Now().UTC().Format(time.RFC3339))
}

func (sl *SecurityLogger) LogRateLimitExceeded(ip string, endpoint string) {
	sl.logger.Printf("[SECURITY] Rate limit exceeded: ip=%s endpoint=%s timestamp=%s",
		ip, endpoint, time.Now().UTC().Format(time.RFC3339))
}

func (sl *SecurityLogger) LogInvalidToken(customerID string, reason string, ip string) {
	sl.logger.Printf("[SECURITY] Invalid token: customer_id=%s reason=%s ip=%s timestamp=%s",
		customerID, reason, ip, time.Now().UTC().Format(time.RFC3339))
}

var securityLogger = NewSecurityLogger()

func Logging() fiber.Handler {
	return func(c fiber.Ctx) error {

		start := time.Now()

		err := c.Next()

		duration := time.Since(start)

		logData := map[string]interface{}{
			"timestamp":   time.Now().UTC(),
			"method":      c.Method(),
			"path":        c.Path(),
			"status":      c.Response().StatusCode(),
			"duration_ms": duration.Milliseconds(),
			"request_id":  c.Locals("request_id"),
			"ip":          c.IP(),
		}

		if c.Response().StatusCode() == fiber.StatusUnauthorized || c.Response().StatusCode() == fiber.StatusForbidden {
			logData["security_event"] = true
			logData["event_type"] = "access_denied"
		}

		jsonLog, _ := json.Marshal(logData)
		log.Println(string(jsonLog))

		return err
	}
}

func GetSecurityLogger() *SecurityLogger {
	return securityLogger
}
