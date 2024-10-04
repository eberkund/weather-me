

pub struct OpenWeather {}

impl WeatherProvider for OpenWeather {
    fn get_current(&self, city: &str) -> String {
        format!("Current weather in {} is sunny", city)
    }

    fn get_4_day(&self, city: &str) -> String {
        format!("Weather in {} will be sunny for the next 4 days", city)
    }
}