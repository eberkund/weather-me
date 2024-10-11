package providers

import "context"

type CurrentResponse struct {
	Temperature float64
	Humidity    int
	UVIndex     int
	Visibility  int
}

type DailyForecast struct {
	Temperature int
}

type ForecastResponse struct {
	Days []DailyForecast
}

type ProvidesWeather interface {
	Current(ctx context.Context, lat float64, lng float64) (*CurrentResponse, error)
	Forecast(ctx context.Context, lat float64, lng float64) (*ForecastResponse, error)
}
