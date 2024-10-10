package server

import (
	"context"
	"github.com/eberkund/weather-me/pkg/providers"
	"github.com/go-faster/errors"

	"connectrpc.com/connect"
	weatherv1 "github.com/eberkund/weather-me/gen/weather/v1"
	"github.com/eberkund/weather-me/gen/weather/v1/weatherv1connect"
)

var _ weatherv1connect.WeatherServiceHandler = (*WeatherHandler)(nil)

func NewWeatherHandler() *WeatherHandler {
	return &WeatherHandler{}
}

type WeatherHandler struct {
	providers map[weatherv1.WeatherProvider]providers.ProvidesWeather
}

func (h *WeatherHandler) GetWeather(ctx context.Context, req *connect.Request[weatherv1.GetWeatherRequest]) (*connect.Response[weatherv1.GetWeatherResponse], error) {
	provider, implemented := h.providers[req.Msg.Provider]
	if !implemented {
		return nil, connect.NewError(
			connect.CodeUnimplemented,
			errors.Errorf("provider %s not implemented", req.Msg.Provider),
		)
	}
	current, err := provider.Current(ctx, req.Msg.City)
	if err != nil {
		return nil, err
	}
	return &connect.Response[weatherv1.GetWeatherResponse]{
		Msg: &weatherv1.GetWeatherResponse{
			Weather: current,
		},
	}, nil
}
