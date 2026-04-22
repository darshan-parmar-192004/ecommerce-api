package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"backend/internal/config"
	"backend/internal/constants"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

// Service represents a service that interacts with a database.
type Service interface {
	Health() map[string]string
	Close() error
	DB() *sql.DB
}

type service struct {
	db *sql.DB
}

var dbInstance *service

func (s *service) DB() *sql.DB {
	return s.db
}

func New() Service {
	if dbInstance != nil {
		return dbInstance
	}
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBSchema)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	dbInstance = &service{
		db: db,
	}
	db.SetMaxOpenConns(constants.DBMaxOpenConns)
	db.SetMaxIdleConns(constants.DBMaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(constants.DBConnMaxLifetime) * time.Minute)
	return dbInstance
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
	cfg, _ := config.Load()
	log.Printf("Disconnected from database: %s", cfg.DBName)
	return s.db.Close()
}
