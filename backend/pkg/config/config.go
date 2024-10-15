package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Log            Log
	Port           int      `env:"PORT,default=4000"`
	Origins        []string `env:"ORIGINS,default=*"`
	WeatherAPIKey  string   `env:"WEATHERAPI_KEY,required"`
	OpenWeatherKey string   `env:"OPENWEATHER_KEY,required"`
}

type Log struct {
	Pretty  bool `env:"LOG_PRETTY,default=false"`
	Verbose bool `env:"LOG_VERBOSE,default=false"`
}

func FromEnv(ctx context.Context) (*Config, error) {
	c := new(Config)
	err := envconfig.Process(ctx, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
