package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	FilePath string `envconfig:"FILE_PATH"`
}

func Load() (*AppConfig, error) {
	godotenv.Load()

	var cfg AppConfig
	err := envconfig.Process("app", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
