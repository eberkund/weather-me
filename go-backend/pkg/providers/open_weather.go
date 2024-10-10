package providers

var _ ProvidesWeather = (*OpenWeather)(nil)

type OpenWeather struct{}

// GetWeather implements ProvidesWeather.
func (o *OpenWeather) GetWeather() (string, error) {
	panic("unimplemented")
}
