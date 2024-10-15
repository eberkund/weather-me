use log::LevelFilter;
use envconfig::Envconfig;

#[derive(Envconfig)]
pub struct Config {
    #[envconfig(from = "LOG_LEVEL", default = "info")]
    pub log_level: LevelFilter,

    #[envconfig(from = "SERVER_PORT", default = "8080")]
    pub http_port: u16,
}
