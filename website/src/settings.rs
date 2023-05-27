use config::{Config, ConfigError, Environment, File};
use serde_derive::Deserialize;

const CONFIG_FILE_PATH: &str = "./config/config.toml";

#[derive(Debug, Deserialize)]
#[allow(unused)]
pub struct Server {
    pub url: String,
    pub port: u16,
}

#[derive(Debug, Deserialize)]
#[allow(unused)]
pub struct Log {
    pub level: String,
}

#[derive(Debug, Deserialize)]
#[allow(unused)]
pub struct Settings {
    pub server: Server,
    pub log: Log,
}

impl Settings {
    pub fn new() -> Result<Self, ConfigError> {
        let s = Config::builder()
            // Start off by merging in the "default" configuration file
            .add_source(File::with_name(CONFIG_FILE_PATH))
            // Add in settings from the environment (with a prefix of APP)
            // Eg.. `APP_DEBUG=1 ./target/app` would set the `debug` key
            .add_source(Environment::with_prefix("app"))
            .build()?;
        s.try_deserialize()
    }
}
