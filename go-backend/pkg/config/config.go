package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Port          int      `env:"PORT,default=4000"`
	Origins       []string `env:"ORIGINS,default=*"`
	WeatherAPIKey string   `env:"WEATHER_API_KEY"`
}

func FromEnv(ctx context.Context) (*Config, error) {
	c := new(Config)
	err := envconfig.Process(ctx, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
