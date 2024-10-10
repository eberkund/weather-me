package main

import (
	"context"

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
	weatherHandler := server.NewWeatherHandler()
	s := server.NewServer(weatherHandler, c.Port, c.Origins)
	app := setup.New(s)
	log.Info().Int("port", c.Port).Msg("starting service")
	cause, err := app.WaitWithErrors()
	log.Error().
		AnErr("cause", cause).
		AnErr("exit", err).
		Msg("exited")
}
