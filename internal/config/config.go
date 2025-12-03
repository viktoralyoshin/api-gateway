package config

import (
	"os"
)

type Config struct {
	AuthServiceAddr string
}

func Load() (*Config, error) {

	cfg := &Config{
		AuthServiceAddr: os.Getenv("AUTH_SERVICE_ADDR"),
	}

	return cfg, nil
}
