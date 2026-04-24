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

	"github.com/doug-martin/goqu/v9"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

// Service represents a service that interacts with a database.
type Service interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() map[string]string

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error
}

type service struct {
	db     *sql.DB
	gdb    *goqu.Database
	dbName string
}

var (
	dbInstance *service
)

func New(cfg *config.AppConfig) Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}

	connStr := cfg.GetDSN()
	sqlDB, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	// Initialize goqu ORM with dialect from config (default "postgres")
	gdb := goqu.New(cfg.DBDialect, sqlDB)

	dbInstance = &service{
		db:     sqlDB,
		gdb:    gdb,
		dbName: cfg.DBName,
	}
	return dbInstance
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// Ping the database
	err := s.db.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Printf("db down: %v", err) // Log the error and terminate the program
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = constants.HealthDBUp

	// Get database stats (like open connections, in use, idle, etc.)
	dbStats := s.db.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	// Evaluate stats to provide a health message
	if dbStats.OpenConnections > constants.DBPoolMaxConnectionsThreshold {
		stats["message"] = constants.HealthDBHeavyLoad
	}

	if dbStats.WaitCount > constants.DBPoolWaitCountThreshold {
		stats["message"] = constants.HealthDBHighWaits
	}

	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = constants.HealthDBIdleClosing
	}

	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = constants.HealthDBLifetimeClose
	}

	return stats
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", s.dbName)
	return s.db.Close()
}
