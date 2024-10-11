package server

import (
	"connectrpc.com/connect"
	"context"
	weatherv1 "github.com/eberkund/weather-me/gen/weather/v1"
	"github.com/eberkund/weather-me/gen/weather/v1/weatherv1connect"
	"github.com/eberkund/weather-me/pkg/providers"
	"github.com/go-faster/errors"
	"github.com/rs/zerolog/log"
)

var _ weatherv1connect.WeatherServiceHandler = (*WeatherHandler)(nil)

func NewWeatherHandler(providers map[weatherv1.WeatherProvider]providers.ProvidesWeather) *WeatherHandler {
	return &WeatherHandler{
		providers: providers,
	}
}

type WeatherHandler struct {
	providers map[weatherv1.WeatherProvider]providers.ProvidesWeather
}

func (h *WeatherHandler) GetForecast(ctx context.Context, c *connect.Request[weatherv1.GetForecastRequest]) (*connect.Response[weatherv1.GetForecastResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (h *WeatherHandler) GetCurrent(ctx context.Context, req *connect.Request[weatherv1.GetCurrentRequest]) (*connect.Response[weatherv1.GetCurrentResponse], error) {
	log.Info().Msg("get weather request")
	provider, implemented := h.providers[req.Msg.Provider]
	if !implemented {
		log.Info().Msg("not implemented")
		return nil, connect.NewError(
			connect.CodeUnimplemented,
			errors.Errorf("provider %s not implemented", req.Msg.Provider),
		)
	}
	current, err := provider.Current(ctx, float64(req.Msg.Latitude), float64(req.Msg.Longitude))
	if err != nil {
		return nil, err
	}
	return &connect.Response[weatherv1.GetCurrentResponse]{
		Msg: &weatherv1.GetCurrentResponse{
			Temperature: float32(current.Temperature),
			Humidity:    float32(current.Humidity),
			UvIndex:     float32(current.UVIndex),
			Visibility:  int32(current.Visibility),
		},
	}, nil
}
