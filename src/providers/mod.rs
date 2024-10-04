pub trait WeatherProvider {
    fn get_current(&self, city: &str) -> String;
    fn get_4_day(&self, city: &str) -> String;
}