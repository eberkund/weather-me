package server

import (
	"connectrpc.com/connect"
	"context"
	weatherv1 "github.com/eberkund/weather-me/gen/weather/v1"
	"github.com/eberkund/weather-me/gen/weather/v1/weatherv1connect"
	"github.com/eberkund/weather-me/pkg/providers"
	"github.com/go-faster/errors"
	"github.com/samber/lo"
	"google.golang.org/genproto/googleapis/type/date"
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

func (h *WeatherHandler) GetForecast(ctx context.Context, req *connect.Request[weatherv1.GetForecastRequest]) (*connect.Response[weatherv1.GetForecastResponse], error) {
	provider, err := h.provider(req.Msg.Provider)
	if err != nil {
		return nil, err
	}
	forecast, err := provider.Forecast(ctx, float64(req.Msg.Latitude), float64(req.Msg.Longitude))
	if err != nil {
		return nil, err
	}
	days := lo.Map(forecast.Days, func(item providers.DailyForecast, index int) *weatherv1.DailyForecast {
		return &weatherv1.DailyForecast{
			Temperature: float32(item.Temperature),
			Date: &date.Date{
				Year:  int32(item.Date.Year()),
				Month: int32(item.Date.Month()),
				Day:   int32(item.Date.Day()),
			},
			Condition: item.Condition,
		}
	})
	return &connect.Response[weatherv1.GetForecastResponse]{
		Msg: &weatherv1.GetForecastResponse{
			Days: days,
		},
	}, nil
}

func (h *WeatherHandler) GetCurrent(ctx context.Context, req *connect.Request[weatherv1.GetCurrentRequest]) (*connect.Response[weatherv1.GetCurrentResponse], error) {
	provider, err := h.provider(req.Msg.Provider)
	if err != nil {
		return nil, err
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

func (h *WeatherHandler) provider(provider weatherv1.WeatherProvider) (providers.ProvidesWeather, error) {
	impl, exists := h.providers[provider]
	if !exists {
		return nil, connect.NewError(
			connect.CodeUnimplemented,
			errors.Errorf("provider %s is unimplemented or invalid", provider),
		)
	}
	return impl, nil
}
