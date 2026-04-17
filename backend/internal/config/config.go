package config

import (
	"backend/internal/constants"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port     int    `envconfig:"APP_PORT"`
	DBHost   string `envconfig:"APP_DB_HOST"`
	DBPort   string `envconfig:"APP_DB_PORT"`
	DBName   string `envconfig:"APP_DB_NAME"`
	DBUser   string `envconfig:"APP_DB_USER"`
	DBPass   string `envconfig:"APP_DB_PASS"`
	DBSchema string `envconfig:"APP_DB_SCHEMA"`
}

func Load() (*AppConfig, error) {
	_ = godotenv.Load()

	cfg := &AppConfig{
		Port:     constants.DefaultPort,
		DBSchema: "public",
	}

	if v := os.Getenv("APP_PORT"); v != "" {
		if p, err := strconv.Atoi(v); err == nil {
			cfg.Port = p
		}
	}
	if v := os.Getenv("APP_DB_HOST"); v != "" {
		cfg.DBHost = v
	}
	if v := os.Getenv("APP_DB_PORT"); v != "" {
		cfg.DBPort = v
	}
	if v := os.Getenv("APP_DB_NAME"); v != "" {
		cfg.DBName = v
	}
	if v := os.Getenv("APP_DB_USER"); v != "" {
		cfg.DBUser = v
	}
	if v := os.Getenv("APP_DB_PASS"); v != "" {
		cfg.DBPass = v
	}
	if v := os.Getenv("APP_DB_SCHEMA"); v != "" {
		cfg.DBSchema = v
	}

	return cfg, nil
}

func LoadTest(dbHost, dbPort, dbName, dbUser, dbPwd string) (*AppConfig, error) {
	return &AppConfig{
		Port:     constants.DefaultPort,
		DBHost:   dbHost,
		DBPort:   dbPort,
		DBName:   dbName,
		DBUser:   dbUser,
		DBPass:   dbPwd,
		DBSchema: "public",
	}, nil
}

func (c *AppConfig) GetDSN() string {
	return "postgres://" + c.DBUser + ":" + c.DBPass + "@" + c.DBHost + ":" + c.DBPort + "/" + c.DBName + "?sslmode=disable&search_path=" + c.DBSchema
}

func GetPort() int {
	if v := os.Getenv("APP_PORT"); v != "" {
		if p, err := strconv.Atoi(v); err == nil && p > 0 {
			return p
		}
	}
	if v := os.Getenv("PORT"); v != "" {
		if p, err := strconv.Atoi(v); err == nil && p > 0 {
			return p
		}
	}
	return constants.DefaultPort
}
