package providers

import (
	"net/http"
)

func NewOpenWeather() *OpenWeather {
	return &OpenWeather{
		http: http.Client{},
	}
}

type OpenWeather struct {
	http http.Client
}

var _ ProvidesWeather = (*OpenWeather)(nil)

func (o *OpenWeather) Current(city string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (o *OpenWeather) Forecast(city string) (string, error) {
	//TODO implement me
	panic("implement me")
}
