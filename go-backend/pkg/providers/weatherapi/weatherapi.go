package weatherapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/eberkund/weather-me/pkg/providers"
	"github.com/eberkund/weather-me/swagger"
	"github.com/go-faster/errors"
	"github.com/rs/zerolog/log"
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

type RealtimeWeatherResponse struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Latitude       float64 `json:"lat"`
		Longitude      float64 `json:"lon"`
		TzID           string  `json:"tz_id"`
		LocalTimeEpoch int     `json:"localtime_epoch"`
		LocalTime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph       float64 `json:"wind_mph"`
		WindKph       float64 `json:"wind_kph"`
		WindDegree    int     `json:"wind_degree"`
		WindDirection string  `json:"wind_dir"`
		PressureMb    int     `json:"pressure_mb"`
		PressureIn    float64 `json:"pressure_in"`
		PrecipMm      float64 `json:"precip_mm"`
		PrecipIn      int     `json:"precip_in"`
		Humidity      int     `json:"humidity"`
		Cloud         int     `json:"cloud"`
		FeelslikeC    float64 `json:"feelslike_c"`
		FeelslikeF    float64 `json:"feelslike_f"`
		WindchillC    float64 `json:"windchill_c"`
		WindchillF    float64 `json:"windchill_f"`
		HeatIndexC    float64 `json:"heatindex_c"`
		HeatIndexF    float64 `json:"heatindex_f"`
		DewPointC     float64 `json:"dewpoint_c"`
		DewPointF     float64 `json:"dewpoint_f"`
		VisibleKms    int     `json:"vis_km"`
		VisibleMiles  int     `json:"vis_miles"`
		UV            int     `json:"uv"`
		GustMph       float64 `json:"gust_mph"`
		GustKph       float64 `json:"gust_kph"`
	} `json:"current"`
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
		Humidity:    decoded.Current.Humidity,
		UVIndex:     decoded.Current.UV,
		Visibility:  decoded.Current.VisibleKms,
	}, nil
}

func (o *WeatherAPI) Forecast(ctx context.Context, lat float64, lng float64) (*providers.ForecastResponse, error) {
	response, _, err := o.client.APIsApi.ForecastWeather(
		o.authContext(ctx),
		fmt.Sprintf("%f,%f", lat, lng),
		4,
		nil,
	)
	if err != nil {
		return nil, err
	}
	_, ok := response.(map[string]any)
	if !ok {
		return nil, errors.New("could not parse response")
	}
	return &providers.ForecastResponse{}, nil
}
