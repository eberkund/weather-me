use tonic::{Response,Status};

use crate::gen::weather::v1::{GetWeatherRequest, GetWeatherResponse, WeatherProvider};
use crate::gen::weather::v1::weather_service_server::WeatherService;
use crate::providers::open_weather::OpenWeather;

pub struct WeatherHandler {
    open_weather: OpenWeather,
}

impl Default for WeatherHandler {
    fn default() -> Self {
        WeatherHandler {
            open_weather: OpenWeather::new()
        }
    }
}

// unsafe impl Send for WeatherHandler {}
// unsafe impl Sync for WeatherHandler {}

#[tonic::async_trait]
impl WeatherService for WeatherHandler {
    async fn get_weather(
        &self,
        request: tonic::Request<GetWeatherRequest>,
    ) -> Result<tonic::Response<GetWeatherResponse>, tonic::Status> {

        let weather = self.open_weather.get_current("Toronto").await;
        
        match request.into_inner().provider() {
            WeatherProvider::OpenWeather => Ok(Response::new(GetWeatherResponse {
                weather: weather,
            })),
            _ => Err(Status::invalid_argument("invalid provider specified")),
        }
    }
}
