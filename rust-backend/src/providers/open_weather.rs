use std::future::Future;

use reqwest;

pub struct OpenWeather {
    client: reqwest::Client,
    key: String,
}

impl OpenWeather {
    pub fn new() -> OpenWeather {
        OpenWeather {
            client: reqwest::Client::new(),
            key: "".to_string(),
        }
    }

    pub async fn get_current(&self, city: &str) -> String {
        let response = self.client
            .get("https://api.weatherapi.com/v1/current.json?key=2f48465f13fb4758ba953406240410&q=London&aqi=no")
            .send()
            .await
            .ok();

        match response {
            Some(value) => {
                let data: String = value.text().await.ok().unwrap();
                return format!("response: {:?}", data).to_string();
            }
            None => (),
        }
        format!("Current weather in {} is sunny", city)
    }

    pub async fn get_4_day(&self, city: &str) -> String {
        format!("Weather in {} will be sunny for the next 4 days", city)
    }
}
