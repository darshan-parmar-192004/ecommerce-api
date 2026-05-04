package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"backend/internal/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Main entry point for migration CLI
// Usage: go run ./cmd/migrate <command> [args]
// Commands: up, down, redo, status, force <version>, drop
func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		printUsage()
		os.Exit(1)
	}

	dsn := buildDSN()

	// Ensure migrations directory exists
	if _, err := os.Stat("internal/migrations"); os.IsNotExist(err) {
		log.Fatalf("migrations directory not found: %v", err)
	}

	m, err := migrate.New("file://internal/migrations", dsn)
	if err != nil {
		log.Fatalf("failed to create migrate instance: %v", err)
	}
	defer func() {
		if srcErr, dbErr := m.Close(); srcErr != nil || dbErr != nil {
			log.Printf("Warning: failed to close migrate: src=%v, db=%v", srcErr, dbErr)
		}
	}()

	command := args[0]

	switch command {
	case "up":
		migrateUp(m)
	case "down":
		migrateDown(m)
	case "redo":
		rerun(m)
	case "status":
		showStatus(m)
	case "drop":
		dropAll(m)
	case "force":
		forceVersion(m, args)
	default:
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: go run ./cmd/migrate <command> [args]")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("  up            Run all pending migrations")
	fmt.Println("  down          Rollback the last migration")
	fmt.Println("  redo          Rollback and re-run the last migration")
	fmt.Println("  status        Show migration status")
	fmt.Println("  force <ver>   Force database to specific version (for dirty recovery)")
	fmt.Println("  drop          Drop all migrations (destructive!)")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("  go run ./cmd/migrate up             # Run all pending migrations")
	fmt.Println("  go run ./cmd/migrate down           # Rollback 1 migration")
	fmt.Println("  go run ./cmd/migrate redo           # Re-run last migration")
	fmt.Println("  go run ./cmd/migrate status         # Show migration status")
	fmt.Println("  go run ./cmd/migrate force 2        # Force version 2 (dirty recovery)")
	fmt.Println("  go run ./cmd/migrate drop           # Drop all migrations")
}

// migrateUp applies all pending migrations
// Handles dirty database state by forcing last version and retrying
func migrateUp(m *migrate.Migrate) {
	fmt.Println("Running all pending migrations...")
	err := m.Up()
	if err != nil {
		if err.Error() == "no change" {
			fmt.Println("Database is already up to date")
			return
		}
		if strings.Contains(err.Error(), "Dirty database") {
			fmt.Println("Dirty database detected, recovering...")
			// Get current version to force
			version, _, verErr := m.Version()
			if verErr != nil {
				log.Fatalf("failed to get current version: %v", verErr)
			}
			fmt.Printf("Forcing version %d to clear dirty state...\n", version)
			if forceErr := m.Force(int(version)); forceErr != nil {
				log.Fatalf("failed to force version %d: %v", version, forceErr)
			}
			fmt.Println("Retrying migrations...")
			err = m.Up()
			if err != nil && err.Error() != "no change" {
				log.Fatalf("migration up failed: %v", err)
			}
		} else {
			log.Fatalf("migration up failed: %v", err)
		}
	}
	fmt.Println("Migrations applied successfully")
}

// migrateDown rolls back the last migration
func migrateDown(m *migrate.Migrate) {
	fmt.Println("Rolling back last migration...")
	if err := m.Steps(-1); err != nil {
		if err.Error() == "no change" {
			fmt.Println("No migrations to rollback")
			return
		}
		log.Fatalf("migration down failed: %v", err)
	}
	fmt.Println("Last migration rolled back successfully")
}

// rerun rolls back and re-applies the last migration
func rerun(m *migrate.Migrate) {
	fmt.Println("Re-running last migration...")
	// Rollback last migration
	if err := m.Steps(-1); err != nil {
		if err.Error() == "no change" {
			fmt.Println("No migrations to redo")
			return
		}
		log.Fatalf("rollback failed during redo: %v", err)
	}
	// Re-apply last migration
	if err := m.Steps(1); err != nil {
		if err.Error() == "no change" {
			fmt.Println("No migrations to redo")
			return
		}
		log.Fatalf("re-apply failed during redo: %v", err)
	}
	fmt.Println("Last migration re-run successfully")
}

// showStatus shows the current migration status
func showStatus(m *migrate.Migrate) {
	fmt.Println("Migration status:")

	// Get current version
	version, dirty, err := m.Version()
	if err != nil {
		if err.Error() == "no migration" {
			fmt.Println("  No migrations applied yet")
			fmt.Println("  Run 'go run ./cmd/migrate up' to apply all migrations")
			return
		}
		log.Fatalf("failed to get version: %v", err)
	}

	if dirty {
		fmt.Printf("  Version: %d (DIRTY - recovery needed)\n", version)
		fmt.Println("  Run 'go run ./cmd/migrate force <version>' to recover")
	} else {
		fmt.Printf("  Current version: %d\n", version)
	}

	// Check if there are pending migrations
	err = m.Up()
	if err != nil {
		if err.Error() == "no change" {
			fmt.Println("  Database is up to date - no pending migrations")
		} else if strings.Contains(err.Error(), "Dirty database") {
			fmt.Println("  Status: Dirty database (cannot determine pending migrations)")
		} else {
			fmt.Println("  Pending migrations: Unable to determine")
		}
	} else {
		// Rollback the Up() we just ran to check status
		if rollbackErr := m.Steps(-1); rollbackErr != nil {
			fmt.Println("  Warning: Unable to rollback status check")
		}
		fmt.Println("  Pending migrations: Yes - run 'go run ./cmd/migrate up' to apply")
	}
}

// dropAll drops all migrations (destructive - use with caution)
func dropAll(m *migrate.Migrate) {
	fmt.Println("WARNING: Dropping all migrations...")
	if err := m.Drop(); err != nil {
		log.Fatalf("drop failed: %v", err)
	}
	fmt.Println("All migrations dropped successfully")
}

// forceVersion forces database to specific version (for recovery from dirty state)
// Usage: go run ./cmd/migrate force <version>
func forceVersion(m *migrate.Migrate, args []string) {
	if len(args) < 2 {
		log.Fatal("usage: migrate force <version>")
	}
	var version int
	if _, err := fmt.Sscanf(args[1], "%d", &version); err != nil {
		log.Fatalf("invalid version number: %v", err)
	}
	fmt.Printf("Forcing database to version %d...\n", version)
	if err := m.Force(version); err != nil {
		log.Fatalf("force failed: %v", err)
	}
	fmt.Printf("Database forced to version %d successfully\n", version)
}

// buildDSN constructs the database connection string from config
func buildDSN() string {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	return cfg.GetDSN()
}

// checkDBConnection verifies the database connection works
func checkDBConnection(dsn string) error {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	defer func() { _ = db.Close() }()

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}
	return nil
}
