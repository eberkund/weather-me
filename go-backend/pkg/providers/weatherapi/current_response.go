package weatherapi

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
		WindDegree    float64 `json:"wind_degree"`
		WindDirection string  `json:"wind_dir"`
		PressureMb    float64 `json:"pressure_mb"`
		PressureIn    float64 `json:"pressure_in"`
		PrecipMm      float64 `json:"precip_mm"`
		PrecipIn      float64 `json:"precip_in"`
		Humidity      float64 `json:"humidity"`
		Cloud         float64 `json:"cloud"`
		FeelslikeC    float64 `json:"feelslike_c"`
		FeelslikeF    float64 `json:"feelslike_f"`
		WindchillC    float64 `json:"windchill_c"`
		WindchillF    float64 `json:"windchill_f"`
		HeatIndexC    float64 `json:"heatindex_c"`
		HeatIndexF    float64 `json:"heatindex_f"`
		DewPointC     float64 `json:"dewpoint_c"`
		DewPointF     float64 `json:"dewpoint_f"`
		VisibleKms    float64 `json:"vis_km"`
		VisibleMiles  float64 `json:"vis_miles"`
		UV            float64 `json:"uv"`
		GustMph       float64 `json:"gust_mph"`
		GustKph       float64 `json:"gust_kph"`
	} `json:"current"`
}
