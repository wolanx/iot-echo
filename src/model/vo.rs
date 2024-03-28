use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize)]
pub struct ConfigYml {
    pub provider: String,
    pub server: ConfigServer,
    pub device: ConfigDevice,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct ConfigServer {
    pub host: String,
    pub tls: String,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct ConfigDevice {
    pub productKey: String,
    pub deviceName: String,
    pub deviceSecret: String,
}
