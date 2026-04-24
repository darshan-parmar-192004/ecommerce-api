package database

import (
	"context"
	"log"
	"testing"
	"time"

	"backend/internal/config"
	"backend/internal/constants"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var testCfg *config.AppConfig

func mustStartPostgresContainer() (func(context.Context, ...testcontainers.TerminateOption) error, error) {
	var (
		dbName = "database"
		dbPwd  = "password"
		dbUser = "user"
	)

	dbContainer, err := postgres.Run(
		context.Background(),
		"postgres:latest",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPwd),
		testcontainers.WithWaitStrategy(
			wait.ForListeningPort("5432/tcp").
				WithStartupTimeout(10*time.Second)),
	)
	if err != nil {
		return nil, err
	}

	dbHost, err := dbContainer.Host(context.Background())
	if err != nil {
		return dbContainer.Terminate, err
	}

	dbPort, err := dbContainer.MappedPort(context.Background(), "5432/tcp")
	if err != nil {
		return dbContainer.Terminate, err
	}

	// Build test config
	testCfg = &config.AppConfig{
		Port:      8080,
		DBHost:    dbHost,
		DBPort:    dbPort.Port(),
		DBName:    dbName,
		DBUser:    dbUser,
		DBPass:    dbPwd,
		DBSchema:  "public",
		DBDialect: "postgres",
	}

	return dbContainer.Terminate, err
}

func TestMain(m *testing.M) {
	teardown, err := mustStartPostgresContainer()
	if err != nil {
		log.Fatalf("could not start postgres container: %v", err)
	}

	m.Run()

	if teardown != nil && teardown(context.Background()) != nil {
		log.Fatalf("could not teardown postgres container: %v", err)
	}
}

func TestNew(t *testing.T) {
	srv := New(testCfg)
	if srv == nil {
		t.Fatal("New() returned nil")
	}
}

func TestHealth(t *testing.T) {
	srv := New(testCfg)

	stats := srv.Health()

	if stats["status"] != "up" {
		t.Fatalf("expected status to be up, got %s", stats["status"])
	}

	if _, ok := stats["error"]; ok {
		t.Fatalf("expected error not to be present")
	}

	if stats["message"] != constants.HealthDBUp {
		t.Fatalf("expected message to be '%s', got %s", constants.HealthDBUp, stats["message"])
	}
}

func TestClose(t *testing.T) {
	srv := New(testCfg)

	if srv.Close() != nil {
		t.Fatalf("expected Close() to return nil")
	}
}
