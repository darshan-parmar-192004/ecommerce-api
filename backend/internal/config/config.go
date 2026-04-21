package config

import (
	"backend/internal/constants"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	Port      int    `envconfig:"BLUEPRINT_PORT" default:"8080"`
	DBHost    string `envconfig:"BLUEPRINT_DB_HOST"`
	DBPort    string `envconfig:"BLUEPRINT_DB_PORT" default:"5432"`
	DBName    string `envconfig:"BLUEPRINT_DB_DATABASE"`
	DBUser    string `envconfig:"BLUEPRINT_DB_USERNAME"`
	DBPass    string `envconfig:"BLUEPRINT_DB_PASSWORD"`
	DBSchema  string `envconfig:"BLUEPRINT_DB_SCHEMA" default:"public"`
	DBDialect string `envconfig:"BLUEPRINT_DB_DIALECT" default:"pgx"`
	RedisHost string `envconfig:"REDIS_HOST" default:"localhost"`
	RedisPort string `envconfig:"REDIS_PORT" default:"6379"`
	AppEnv    string `envconfig:"APP_ENV" default:"development"`
}

var cfg *AppConfig

func Load() (*AppConfig, error) {
	if cfg != nil {
		return cfg, nil
	}

	_ = godotenv.Load()

	cfg = &AppConfig{}

	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, err
	}

	if cfg.Port == 0 {
		cfg.Port = constants.DefaultPort
	}
	if cfg.DBSchema == "" {
		cfg.DBSchema = "public"
	}
	if cfg.DBDialect == "" {
		cfg.DBDialect = "pgx"
	}

	return cfg, nil
}

func LoadTest(dbHost, dbPort, dbName, dbUser, dbPwd string) (*AppConfig, error) {
	return &AppConfig{
		Port:      constants.DefaultPort,
		DBHost:    dbHost,
		DBPort:    dbPort,
		DBName:    dbName,
		DBUser:    dbUser,
		DBPass:    dbPwd,
		DBSchema:  "public",
		DBDialect: "postgres",
	}, nil
}

func (c *AppConfig) GetDSN() string {
	return "postgres://" + c.DBUser + ":" + c.DBPass + "@" + c.DBHost + ":" + c.DBPort + "/" + c.DBName + "?sslmode=disable&search_path=" + c.DBSchema
}

func GetPort() int {
	if cfg != nil && cfg.Port > 0 {
		return cfg.Port
	}
	return constants.DefaultPort
}

func GetDBDialect() string {
	if cfg != nil && cfg.DBDialect != "" {
		return cfg.DBDialect
	}
	return "pgx"
}
