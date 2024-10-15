package main

import (
	"context"
	"github.com/eberkund/weather-me/gen/weather/v1"
	"github.com/eberkund/weather-me/pkg/config"
	"github.com/eberkund/weather-me/pkg/providers"
	"github.com/eberkund/weather-me/pkg/providers/openweather"
	"github.com/eberkund/weather-me/pkg/providers/weatherapi"
	"github.com/eberkund/weather-me/pkg/server"
	"github.com/eberkund/weather-me/pkg/setup"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	ctx := context.Background()
	c, err := config.FromEnv(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}
	logging(c)
	weatherProviders := map[weatherv1.WeatherProvider]providers.ProvidesWeather{
		weatherv1.WeatherProvider_WEATHER_PROVIDER_OPENWEATHER: openweather.NewOpenWeather(c.OpenWeatherKey),
		weatherv1.WeatherProvider_WEATHER_PROVIDER_WEATHERAPI:  weatherapi.NewWeatherAPI(c.WeatherAPIKey),
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

func logging(config *config.Config) {
	log.Logger = log.With().Caller().Logger()
	if config.Log.Verbose {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	if config.Log.Pretty {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	zerolog.FloatingPointPrecision = 1
}
