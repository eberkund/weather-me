
pub struct AccuWeather {}

impl WeatherProvider for AccuWeather {
    fn get_current(&self, city: &str) -> String {
        format!("Current weather in {} is rainy", city)
    }

    fn get_4_day(&self, city: &str) -> String {
        format!("Weather in {} will be rainy for the next 4 days", city)
    }
}