package database

import (
	"backend/internal/config"
	"backend/internal/constants"
	"backend/internal/logger"
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

type Service interface {
	Health() map[string]string
	Close() error
}

type service struct {
	db *sql.DB
}

var dbInstance *service

func New() Service {
	if dbInstance != nil {
		return dbInstance
	}

	cfg, err := config.Load()
	if err != nil {
		logger.Log.Fatalf("failed to load config: %v", err)
	}

	return newWithConfig(cfg)
}

func newWithConfig(cfg *config.AppConfig) *service {
	db, err := sql.Open("pgx", cfg.GetDSN())
	if err != nil {
		logger.Log.Fatalf("failed to open database: %v", err)
	}

	svc := &service{
		db: db,
	}
	dbInstance = svc
	return svc
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	err := s.db.PingContext(ctx)
	if err != nil {
		stats[constants.JSONFieldStatus] = constants.ResponseStatusDown
		stats[constants.JSONFieldError] = fmt.Sprintf("db down: %v", err)
		logger.Log.Errorf("db down: %v", err)
		return stats
	}

	stats[constants.JSONFieldStatus] = constants.ResponseStatusUp
	stats[constants.JSONFieldMessage] = constants.HealthDBUp

	dbStats := s.db.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	if dbStats.OpenConnections > 40 {
		stats[constants.JSONFieldMessage] = constants.HealthDBHeavyLoad
	}

	if dbStats.WaitCount > 1000 {
		stats[constants.JSONFieldMessage] = constants.HealthDBHighWaits
	}

	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats[constants.JSONFieldMessage] = constants.HealthDBIdleClosing
	}

	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats[constants.JSONFieldMessage] = constants.HealthDBLifetimeClose
	}

	return stats
}

func (s *service) Close() error {
	logger.Log.Info("Disconnected from database")
	return s.db.Close()
}
