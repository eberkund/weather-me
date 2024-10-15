mod config;
mod gen;
mod server;
mod providers;

extern crate pretty_logger;

use anyhow;
use config::Config;
use envconfig::Envconfig;
use gen::weather::v1::weather_service_server::WeatherServiceServer;
use log::info;
use tonic::transport::Server;
use providers::open_weather::OpenWeather;

#[tokio::main]
async fn main() -> anyhow::Result<()> {
    let config = Config::init_from_env()?;
    femme::with_level(config.log_level);
    
    let addr: std::net::SocketAddr = "0.0.0.0:8080".parse()?;
    let weather = server::WeatherHandler::default();
    let service = WeatherServiceServer::new(weather);
    let web = tonic_web::enable(service);

    // let descriptors = tonic::include_file_descriptor_set!("descriptors.binpb");
    let reflection = tonic_reflection::server::Builder::configure()
        .register_encoded_file_descriptor_set(include_bytes!("../descriptors.binpb"))
        .build_v1()?;


    info!("starting server on {}", addr);
    Server::builder()
        .accept_http1(true)
        .add_service(web)
        .add_service(reflection)
        .serve(addr)
        .await?;
    Ok(())
}
