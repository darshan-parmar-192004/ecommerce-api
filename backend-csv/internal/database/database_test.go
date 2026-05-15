package database

import (
	"context"
	"database/sql"
	"log"
	"testing"
	"time"

	"backend/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var testDB *sql.DB

func mustStartPostgresContainer() (func(context.Context, ...testcontainers.TerminateOption) error, *config.AppConfig, error) {
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
		return nil, nil, err
	}

	dbHost, err := dbContainer.Host(context.Background())
	if err != nil {
		return dbContainer.Terminate, nil, err
	}

	dbPort, err := dbContainer.MappedPort(context.Background(), "5432/tcp")
	if err != nil {
		return dbContainer.Terminate, nil, err
	}

	cfg, err := config.LoadTest(dbHost, dbPort.Port(), dbName, dbUser, dbPwd)
	if err != nil {
		return dbContainer.Terminate, nil, err
	}

	return dbContainer.Terminate, cfg, nil
}

func TestMain(m *testing.M) {
	teardown, cfg, err := mustStartPostgresContainer()
	if err != nil {
		log.Fatalf("could not start postgres container: %v", err)
	}

	db, err := sql.Open("pgx", cfg.GetDSN())
	if err != nil {
		log.Fatalf("could not open database: %v", err)
	}
	testDB = db

	code := m.Run()

	if testDB != nil {
		if err := testDB.Close(); err != nil {
			log.Fatalf("failed to close test db: %v", err)
		}
	}

	if teardown != nil && teardown(context.Background()) != nil {
		log.Fatalf("could not teardown postgres container: %v", err)
	}

	if code != 0 {
		log.Fatalf("tests failed: %v", code)
	}
}

func TestNew(t *testing.T) {
	srv := New()
	if srv == nil {
		t.Fatal("New() returned nil")
	}
}

func TestHealth(t *testing.T) {
	srv := New()

	stats := srv.Health()

	if stats["status"] != "up" {
		t.Fatalf("expected status to be up, got %s", stats["status"])
	}

	if _, ok := stats["error"]; ok {
		t.Fatalf("expected error not to be present")
	}

	if stats["message"] != "It's healthy" {
		t.Fatalf("expected message to be 'It's healthy', got %s", stats["message"])
	}
}

func TestClose(t *testing.T) {
	t.Skip("Skipping close test - instance is shared")
}
