package config

import (
	"os"
	"time"
)

type Config struct {
	AuthServiceAddr string
	RefreshToketTtl time.Duration
	Domain          string
}

func Load() (*Config, error) {

	refreshToketTtl := os.Getenv("REFRESH_TOKEN_TTL")

	refreshDuration, err := time.ParseDuration(refreshToketTtl)
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		AuthServiceAddr: os.Getenv("AUTH_SERVICE_ADDR"),
		RefreshToketTtl: refreshDuration,
		Domain:          os.Getenv("DOMAIN"),
	}

	return cfg, nil
}
