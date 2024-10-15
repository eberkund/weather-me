package weatherapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	weatherv1 "github.com/eberkund/weather-me/gen/weather/v1"
	"github.com/eberkund/weather-me/pkg/providers"
	"github.com/eberkund/weather-me/swagger"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"
	"strings"
	"time"
)

func NewWeatherAPI(secret string) *WeatherAPI {
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

var _ providers.ProvidesWeather = (*WeatherAPI)(nil)

func (o *WeatherAPI) authContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, swagger.ContextAPIKey, swagger.APIKey{
		Key: o.secret,
	})
}

func (o *WeatherAPI) Current(ctx context.Context, lat float64, lng float64) (*providers.CurrentResponse, error) {
	response, _, err := o.client.APIsApi.RealtimeWeather(
		o.authContext(ctx),
		fmt.Sprintf("%f,%f", lat, lng),
		nil,
	)
	if err != nil {
		return nil, err
	}
	log.Info().
		Any("weather", response).
		Msg("weather requested")
	buf := new(bytes.Buffer)
	decoded := new(RealtimeWeatherResponse)
	err = json.NewEncoder(buf).Encode(response)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(buf).Decode(decoded)
	if err != nil {
		return nil, err
	}
	return &providers.CurrentResponse{
		Temperature: decoded.Current.TempC,
		Humidity:    int(decoded.Current.Humidity),
		UVIndex:     int(decoded.Current.UV),
		Visibility:  int(decoded.Current.VisibleKms),
	}, nil
}

func (o *WeatherAPI) Forecast(ctx context.Context, lat float64, lng float64) (*providers.ForecastResponse, error) {
	response, _, err := o.client.APIsApi.ForecastWeather(
		o.authContext(ctx),
		fmt.Sprintf("%f,%f", lat, lng),
		7,
		nil,
	)
	if err != nil {
		return nil, err
	}
	log.Info().Msg("forecast requested")
	buf := new(bytes.Buffer)
	decoded := new(ForecastResponse)
	err = json.NewEncoder(buf).Encode(response)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(buf).Decode(decoded)
	if err != nil {
		return nil, err
	}
	days := lo.Map(decoded.Forecast.ForecastDay, func(day ForecastDay, i int) providers.DailyForecast {
		date, err := time.Parse("2006-01-02", day.Date)
		if err != nil {
			log.Error().Err(err).Msg("failed to parse date")
		}
		return providers.DailyForecast{
			Temperature: int(day.Day.AvgTempC),
			Date:        date,
			Condition:   normalizeCondition(day.Day.Condition.Text),
		}
	})
	return &providers.ForecastResponse{
		Days: days,
	}, nil
}

func normalizeCondition(condition string) weatherv1.Condition {
	switch strings.TrimSpace(condition) {
	case "Overcast":
		return weatherv1.Condition_OVERCAST
	case "Cloudy":
		return weatherv1.Condition_CLOUDY
	case "Partly Cloudy":
		return weatherv1.Condition_PARTLY_CLOUDY
	case "Sunny":
		return weatherv1.Condition_SUNNY
	case "Clear":
		return weatherv1.Condition_CLEAR
	case "Mist":
		return weatherv1.Condition_MIST
	case "Moderate rain":
		fallthrough
	case "Heavy rain":
		fallthrough
	case "Rainy":
		fallthrough
	case "Patchy rain nearby":
		return weatherv1.Condition_RAINY
	default:
		return weatherv1.Condition_CONDITION_UNSPECIFIED
	}
}
