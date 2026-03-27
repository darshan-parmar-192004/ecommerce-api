package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

type connection struct {
	db *sql.DB
}

var instance *connection

func GetDB() *sql.DB {
	if instance != nil {
		return instance.db
	}

	database := os.Getenv("BLUEPRINT_DB_DATABASE")
	password := os.Getenv("BLUEPRINT_DB_PASSWORD")
	username := os.Getenv("BLUEPRINT_DB_USERNAME")
	port := os.Getenv("BLUEPRINT_DB_PORT")
	host := os.Getenv("BLUEPRINT_DB_HOST")
	schema := os.Getenv("BLUEPRINT_DB_SCHEMA")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s",
		username, password, host, port, database, schema)

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}

	instance = &connection{db: db}
	instance.db.SetMaxOpenConns(20)
	instance.db.SetMaxIdleConns(5)
	instance.db.SetConnMaxLifetime(30 * time.Minute)

	return instance.db
}

func Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	err := instance.db.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = err.Error()
		return stats
	}

	stats["status"] = "up"
	stats["message"] = "Database is healthy"

	dbStats := instance.db.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()

	return stats
}

func Close() error {
	log.Printf("Disconnected from database")
	return instance.db.Close()
}
