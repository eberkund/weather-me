package openweather

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/eberkund/weather-me/pkg/providers"
	"github.com/go-faster/errors"
	"net/http"
)

var _ providers.ProvidesWeather = (*OpenWeather)(nil)

type OpenWeather struct {
	secret string
	client *http.Client
}

func NewOpenWeather(secret string) *OpenWeather {
	return &OpenWeather{
		secret: secret,
		client: &http.Client{},
	}
}

func (o *OpenWeather) Current(ctx context.Context, lat float64, lng float64) (*providers.CurrentResponse, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s", lat, lng, o.secret)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "invalid request")
	}
	response, err := o.client.Do(request)
	if err != nil {
		return nil, err
	}
	decoded := new(CurrentResponse)
	err = json.NewDecoder(response.Body).Decode(decoded)
	if err != nil {
		return nil, err
	}
	return &providers.CurrentResponse{
		Temperature: decoded.Main.Temp - 273.15,
		Humidity:    int(decoded.Main.Humidity),
		UVIndex:     4,
		Visibility:  int(decoded.Visibility / 1000),
	}, nil
}

func (o *OpenWeather) Forecast(ctx context.Context, lat float64, lng float64) (*providers.ForecastResponse, error) {
	return nil, errors.New("not implemented")
}
