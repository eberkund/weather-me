package providers

import (
	"context"
	"github.com/eberkund/weather-me/swagger"
	"github.com/go-faster/errors"
	"github.com/rs/zerolog/log"
	"strconv"
)

func NewOpenWeather(secret string) *WeatherAPI {
	config := swagger.NewConfiguration()
	client := swagger.NewAPIClient(config)
	return &WeatherAPI{
		client: client,
		secret: secret,
	}
}

type WeatherAPI struct {
	client *swagger.APIClient
	secret string
}

var _ ProvidesWeather = (*WeatherAPI)(nil)

func (o *WeatherAPI) authContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, swagger.ContextAPIKey, swagger.APIKey{
		Key: o.secret,
	})
}

func (o *WeatherAPI) Current(ctx context.Context, city string) (string, error) {
	response, _, err := o.client.APIsApi.RealtimeWeather(o.authContext(ctx), city, nil)
	if err != nil {
		return "", err
	}
	log.Info().
		Str("city", city).
		Any("weather", response).
		Msg("weather requested")
	responseMap, ok := response.(map[string]any)
	if !ok {
		return "", errors.New("could not parse response")
	}
	current, ok := responseMap["current"].(map[string]any)
	if !ok {
		return "", errors.New("could not parse response")
	}
	condition, ok := current["condition"].(map[string]any)
	if !ok {
		return "", errors.New("could not parse response")
	}
	text, ok := condition["text"].(string)
	if !ok {
		return "", errors.New("could not parse response")
	}
	return text, nil
}

func (o *WeatherAPI) Forecast(ctx context.Context, city string) (string, error) {
	response, _, err := o.client.APIsApi.ForecastWeather(o.authContext(ctx), city, 4, nil)
	if err != nil {
		return "", err
	}
	responseMap, ok := response.(map[string]any)
	if !ok {
		return "", errors.New("could not parse response")
	}
	return strconv.Itoa(len(responseMap)), nil
}
