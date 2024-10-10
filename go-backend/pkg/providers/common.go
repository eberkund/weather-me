//go:generate go-enum --marshal --lower

package providers

import "context"

type ProvidesWeather interface {
	Current(ctx context.Context, city string) (string, error)
	Forecast(ctx context.Context, city string) (string, error)
}

// WeatherProvider is the available weather providers.
// ENUM(
// WeatherAPI
// )
type WeatherProvider string
