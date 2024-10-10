package server

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	weatherv1 "github.com/eberkund/weather-me/gen/weather/v1"
	"github.com/eberkund/weather-me/gen/weather/v1/weatherv1connect"
)

var _ weatherv1connect.WeatherServiceHandler = (*WeatherHandler)(nil)

func NewWeatherHandler() *WeatherHandler {
	return &WeatherHandler{}
}

type WeatherHandler struct {
}

func (h *WeatherHandler) GetWeather(ctx context.Context, req *connect.Request[weatherv1.GetWeatherRequest]) (*connect.Response[weatherv1.GetWeatherResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("weather.v1.WeatherService.GetWeather is not implemented"))
}
