package database

import (
	"backend/internal/config"
	"backend/internal/constants"
	"backend/internal/logger"
	"context"
	"database/sql"
	"strconv"
	"time"

	"backend/internal/config"
	"backend/internal/constants"
	"backend/internal/logger"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

type Service interface {
	Health() map[string]string
	Close() error
	DB() *sql.DB
}

type service struct {
	db *sql.DB
}

var dbInstance Service

func New() Service {
	if dbInstance != nil {
		return dbInstance
	}

	cfg, err := config.Load()
	if err != nil {
		logger.Log.Fatalf("failed to load config: %v", err)
	}

	connStr := cfg.GetDSN()
	db, err := sql.Open(cfg.DBDialect, connStr)
	if err != nil {
		logger.Log.Fatalf("failed to open database: %v", err)
	}

	db.SetMaxOpenConns(constants.DBMaxOpenConns)
	db.SetMaxIdleConns(constants.DBMaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(constants.DBConnMaxLifetime) * time.Minute)

	dbInstance = &service{
		db: db,
	}

	logger.Log.Infof("Database connection established: %s/%s", cfg.DBHost, cfg.DBName)
	return dbInstance
}

func (s *service) DB() *sql.DB {
	return s.db
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	err := s.db.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = err.Error()
		return stats
	}

	stats["status"] = "up"
	stats["message"] = "Database is healthy"

	dbStats := s.db.Stats()

	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()

	return stats
}

func (s *service) Close() error {
	logger.Log.Info("Closing database connection")
	return s.db.Close()
}
