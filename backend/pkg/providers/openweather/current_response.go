package openweather

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Weather struct {
	ID          float64 `json:"id"`
	Main        string  `json:"main"`
	Description string  `json:"description"`
	Icon        string  `json:"icon"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  float64 `json:"pressure"`
	Humidity  float64 `json:"humidity"`
	SeaLevel  float64 `json:"sea_level"`
	GrndLevel float64 `json:"grnd_level"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
	Gust  float64 `json:"gust"`
}

type Clouds struct {
	All float64 `json:"all"`
}

type Sys struct {
	Sunrise float64 `json:"sunrise"`
	Sunset  float64 `json:"sunset"`
}

type CurrentResponse struct {
	Coord      Coord     `json:"coord"`
	Weather    []Weather `json:"weather"`
	Base       string    `json:"base"`
	Main       Main      `json:"main"`
	Visibility float64   `json:"visibility"`
	Wind       Wind      `json:"wind"`
	Clouds     Clouds    `json:"clouds"`
	Dt         float64   `json:"dt"`
	Sys        Sys       `json:"sys"`
	Timezone   float64   `json:"timezone"`
	ID         float64   `json:"id"`
	Name       string    `json:"name"`
	Cod        float64   `json:"cod"`
}
