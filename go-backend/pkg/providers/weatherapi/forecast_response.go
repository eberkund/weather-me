package weatherapi

type ForecastResponse struct {
	Location Location `json:"location,omitempty"`
	Current  Current  `json:"current,omitempty"`
	Forecast Forecast `json:"forecast,omitempty"`
}

type Location struct {
	Name           string  `json:"name,omitempty"`
	Region         string  `json:"region,omitempty"`
	Country        string  `json:"country,omitempty"`
	Lat            float64 `json:"lat,omitempty"`
	Lon            float64 `json:"lon,omitempty"`
	TzID           string  `json:"tz_id,omitempty"`
	LocaltimeEpoch float64 `json:"localtime_epoch,omitempty"`
	Localtime      string  `json:"localtime,omitempty"`
}

type Condition struct {
	Text string  `json:"text,omitempty"`
	Icon string  `json:"icon,omitempty"`
	Code float64 `json:"code,omitempty"`
}

type Current struct {
	LastUpdatedEpoch float64   `json:"last_updated_epoch,omitempty"`
	LastUpdated      string    `json:"last_updated,omitempty"`
	TempC            float64   `json:"temp_c,omitempty"`
	TempF            float64   `json:"temp_f,omitempty"`
	IsDay            float64   `json:"is_day,omitempty"`
	Condition        Condition `json:"condition,omitempty"`
	WindMph          float64   `json:"wind_mph,omitempty"`
	WindKph          float64   `json:"wind_kph,omitempty"`
	WindDegree       float64   `json:"wind_degree,omitempty"`
	WindDir          string    `json:"wind_dir,omitempty"`
	PressureMb       float64   `json:"pressure_mb,omitempty"`
	PressureIn       float64   `json:"pressure_in,omitempty"`
	PrecipMm         float64   `json:"precip_mm,omitempty"`
	PrecipIn         float64   `json:"precip_in,omitempty"`
	Humidity         float64   `json:"humidity,omitempty"`
	Cloud            float64   `json:"cloud,omitempty"`
	FeelslikeC       float64   `json:"feelslike_c,omitempty"`
	FeelslikeF       float64   `json:"feelslike_f,omitempty"`
	WindchillC       float64   `json:"windchill_c,omitempty"`
	WindchillF       float64   `json:"windchill_f,omitempty"`
	HeatindexC       float64   `json:"heatindex_c,omitempty"`
	HeatindexF       float64   `json:"heatindex_f,omitempty"`
	DewpointC        float64   `json:"dewpoint_c,omitempty"`
	DewpointF        float64   `json:"dewpoint_f,omitempty"`
	VisKm            float64   `json:"vis_km,omitempty"`
	VisMiles         float64   `json:"vis_miles,omitempty"`
	UV               float64   `json:"uv,omitempty"`
	GustMph          float64   `json:"gust_mph,omitempty"`
	GustKph          float64   `json:"gust_kph,omitempty"`
}

type Day struct {
	MaxTempC          float64   `json:"maxtemp_c,omitempty"`
	MaxTempF          float64   `json:"maxtemp_f,omitempty"`
	MinTempC          float64   `json:"mintemp_c,omitempty"`
	MinTempF          float64   `json:"mintemp_f,omitempty"`
	AvgTempC          float64   `json:"avgtemp_c,omitempty"`
	AvgTempF          float64   `json:"avgtemp_f,omitempty"`
	MaxWindMph        float64   `json:"maxwind_mph,omitempty"`
	MaxWindKph        float64   `json:"maxwind_kph,omitempty"`
	TotalPrecipMm     float64   `json:"totalprecip_mm,omitempty"`
	TotalPrecipIn     float64   `json:"totalprecip_in,omitempty"`
	TotalsnowCm       float64   `json:"totalsnow_cm,omitempty"`
	AvgvisKm          float64   `json:"avgvis_km,omitempty"`
	AvgvisMiles       float64   `json:"avgvis_miles,omitempty"`
	Avghumidity       float64   `json:"avghumidity,omitempty"`
	DailyWillItRain   float64   `json:"daily_will_it_rain,omitempty"`
	DailyChanceOfRain float64   `json:"daily_chance_of_rain,omitempty"`
	DailyWillItSnow   float64   `json:"daily_will_it_snow,omitempty"`
	DailyChanceOfSnow float64   `json:"daily_chance_of_snow,omitempty"`
	Condition         Condition `json:"condition,omitempty"`
	UV                float64   `json:"uv,omitempty"`
}

type Astro struct {
	Sunrise          string  `json:"sunrise,omitempty"`
	Sunset           string  `json:"sunset,omitempty"`
	Moonrise         string  `json:"moonrise,omitempty"`
	Moonset          string  `json:"moonset,omitempty"`
	MoonPhase        string  `json:"moon_phase,omitempty"`
	MoonIllumination float64 `json:"moon_illumination,omitempty"`
	IsMoonUp         float64 `json:"is_moon_up,omitempty"`
	IsSunUp          float64 `json:"is_sun_up,omitempty"`
}

type Hour struct {
	TimeEpoch    float64   `json:"time_epoch,omitempty"`
	Time         string    `json:"time,omitempty"`
	TempC        float64   `json:"temp_c,omitempty"`
	TempF        float64   `json:"temp_f,omitempty"`
	IsDay        float64   `json:"is_day,omitempty"`
	Condition    Condition `json:"condition,omitempty"`
	WindMph      float64   `json:"wind_mph,omitempty"`
	WindKph      float64   `json:"wind_kph,omitempty"`
	WindDegree   float64   `json:"wind_degree,omitempty"`
	WindDir      string    `json:"wind_dir,omitempty"`
	PressureMb   float64   `json:"pressure_mb,omitempty"`
	PressureIn   float64   `json:"pressure_in,omitempty"`
	PrecipMm     float64   `json:"precip_mm,omitempty"`
	PrecipIn     float64   `json:"precip_in,omitempty"`
	SnowCm       float64   `json:"snow_cm,omitempty"`
	Humidity     float64   `json:"humidity,omitempty"`
	Cloud        float64   `json:"cloud,omitempty"`
	FeelsLikeC   float64   `json:"feelslike_c,omitempty"`
	FeelsLikeF   float64   `json:"feelslike_f,omitempty"`
	WindChillC   float64   `json:"windchill_c,omitempty"`
	WindChillF   float64   `json:"windchill_f,omitempty"`
	HeatIndexC   float64   `json:"heatindex_c,omitempty"`
	HeatIndexF   float64   `json:"heatindex_f,omitempty"`
	DewPointC    float64   `json:"dewpoint_c,omitempty"`
	DewPointF    float64   `json:"dewpoint_f,omitempty"`
	WillItRain   float64   `json:"will_it_rain,omitempty"`
	ChanceOfRain float64   `json:"chance_of_rain,omitempty"`
	WillItSnow   float64   `json:"will_it_snow,omitempty"`
	ChanceOfSnow float64   `json:"chance_of_snow,omitempty"`
	VisibleKm    float64   `json:"vis_km,omitempty"`
	VisibleMiles float64   `json:"vis_miles,omitempty"`
	GustMph      float64   `json:"gust_mph,omitempty"`
	GustKph      float64   `json:"gust_kph,omitempty"`
	UV           float64   `json:"uv,omitempty"`
}

type ForecastDay struct {
	Date      string  `json:"date,omitempty"`
	DateEpoch float64 `json:"date_epoch,omitempty"`
	Day       Day     `json:"day,omitempty"`
	Astro     Astro   `json:"astro,omitempty"`
	Hour      []Hour  `json:"hour,omitempty"`
}

type Forecast struct {
	ForecastDay []ForecastDay `json:"forecastday,omitempty"`
}
