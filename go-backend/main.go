package main

import (
	"context"
	weatherv1 "github.com/eberkund/weather-me/gen/weather/v1"
	"github.com/eberkund/weather-me/pkg/providers"

	"github.com/eberkund/weather-me/pkg/config"
	"github.com/eberkund/weather-me/pkg/server"
	"github.com/eberkund/weather-me/pkg/setup"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx := context.Background()
	c, err := config.FromEnv(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}
	weatherProviders := map[weatherv1.WeatherProvider]providers.ProvidesWeather{
		weatherv1.WeatherProvider_WEATHER_PROVIDER_OPEN_WEATHER: providers.NewOpenWeather(c.WeatherAPIKey),
	}
	weatherHandler := server.NewWeatherHandler(weatherProviders)
	s := server.NewServer(weatherHandler, c.Port, c.Origins)
	app := setup.New(s)
	log.Info().Int("port", c.Port).Msg("starting service")
	cause, err := app.WaitWithErrors()
	log.Error().
		AnErr("cause", cause).
		AnErr("exit", err).
		Msg("exited")
}
