package config

import (
	"backend/internal/constants"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	Server ServerConfig `envconfig:"SERVER"`
	DB     DBConfig     `envconfig:"DB"`
}

type ServerConfig struct {
	Port int `envconfig:"PORT"`
}

type DBConfig struct {
	Database string `envconfig:"DB_DATABASE"`
	Password string `envconfig:"DB_PASSWORD"`
	Username string `envconfig:"DB_USERNAME"`
	Port     string `envconfig:"DB_PORT"`
	Host     string `envconfig:"DB_HOST"`
	Schema   string `envconfig:"DB_SCHEMA"`
}

type TestConfig struct {
	DB DBConfig
}

func Load() (*AppConfig, error) {
	_ = godotenv.Load()

	var cfg AppConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}

	if cfg.Server.Port == 0 {
		cfg.Server.Port = constants.DefaultPort
	}

	return &cfg, nil
}

func LoadTest(dbHost, dbPort, dbName, dbUser, dbPwd string) (*AppConfig, error) {
	return &AppConfig{
		Server: ServerConfig{Port: constants.DefaultPort},
		DB: DBConfig{
			Host:     dbHost,
			Port:     dbPort,
			Database: dbName,
			Username: dbUser,
			Password: dbPwd,
			Schema:   "public",
		},
	}, nil
}

func (c *AppConfig) GetDSN() string {
	return "postgres://" + c.DB.Username + ":" + c.DB.Password + "@" + c.DB.Host + ":" + c.DB.Port + "/" + c.DB.Database + "?sslmode=disable&search_path=" + c.DB.Schema
}

func GetPort() int {
	_ = godotenv.Load()
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	if port == 0 {
		return constants.DefaultPort
	}
	return port
}
