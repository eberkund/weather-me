// go:generate go-enum --marshal --lower

package providers

type ProvidesWeather interface {
	GetWeather() (string, error)
}

// ENUM(OpenWeather)
type WeatherProvider string
