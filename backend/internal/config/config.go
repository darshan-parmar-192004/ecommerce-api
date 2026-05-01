package config

import (
	"fmt"
	"os"
	"strconv"

	"backend/internal/constants"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	Port int      `envconfig:"APP_PORT"`
	DB   DBConfig `envconfig:"-"`
}

type DBConfig struct {
	Host     string `envconfig:"DB_HOST"`
	Port     string `envconfig:"DB_PORT"`
	Username string `envconfig:"DB_USERNAME"`
	Password string `envconfig:"DB_PASSWORD"`
	DBName   string `envconfig:"DB_NAME"`
	Schema   string `envconfig:"DB_SCHEMA"`
}

// Load loads configuration from .env file or OS environment using envconfig
func Load() (*AppConfig, error) {
	_ = godotenv.Load()

	var cfg AppConfig

	// envconfig automatically maps environment variables to struct fields
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	if cfg.Port == 0 {
		cfg.Port = constants.DefaultPort
	}

	if cfg.DB.Schema == "" {
		cfg.DB.Schema = "public"
	}

	return &cfg, nil
}

func LoadTest(dbHost, dbPort, dbName, dbUser, dbPwd string) (*AppConfig, error) {
	return &AppConfig{
		Port: constants.DefaultPort,
		DB: DBConfig{
			Host:     dbHost,
			Port:     dbPort,
			DBName:   dbName,
			Username: dbUser,
			Password: dbPwd,
			Schema:   "public",
		},
	}, nil
}

func (c *AppConfig) GetDSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s",
		c.DB.Username, c.DB.Password, c.DB.Host, c.DB.Port, c.DB.DBName, c.DB.Schema,
	)
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
