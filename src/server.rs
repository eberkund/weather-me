use crate::gen::weather::v1::{GetWeatherRequest, GetWeatherResponse};
use crate::gen::weather::v1::weather_service_server::WeatherService;
use std::result::Result;

#[derive(Debug, Default)]
pub struct WeatherHandler {
    
}

#[tonic::async_trait]
impl WeatherService for WeatherHandler {
    async fn get_weather(
        &self,
        request: tonic::Request<GetWeatherRequest>,
    ) -> Result<tonic::Response<GetWeatherResponse>, tonic::Status> {
        println!("Got a request: {:?}", request);
        Ok(tonic::Response::new(GetWeatherResponse {
            weather: "sunny".into(),
        }))
    }
}
