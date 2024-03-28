use crate::model::vo::ConfigYml;
use hmac_sha256::HMAC;
use paho_mqtt::connect_options::ConnectOptions;
use paho_mqtt::create_options::CreateOptions;
use std::time::Duration;
use hex::ToHex;

pub fn getConnOpt(config: &ConfigYml) -> (CreateOptions, ConnectOptions) {
    let host = format!("tcp://{}:1883", config.server.host);
    let product_key = config.device.productKey.clone();
    let device_name = config.device.deviceName.clone();
    let device_secret = config.device.deviceSecret.as_bytes();

    let keep_alive_s = 60;
    let client_id = product_key.to_string() + &".".to_string() + &device_name.to_string();
    let timestamp = "2524608000000";

    // 1.Calculate user name
    let user_name = device_name.to_string() + &"&".to_string() + &product_key.to_string();
    // 2.Calculate the extended clientId
    let extended_client_id = product_key.to_string()
        + &".".to_string()
        + &device_name.to_string()
        + &"|timestamp=".to_string()
        + &timestamp.to_string()
        + &",lan=RUST,_v=1.0.0,securemode=2,signmethod=hmacsha256,ext=3|".to_string();
    // 3.Calculate the password from product key, device name, device secret
    let sign_src = "clientId".to_string()
        + &client_id.to_string()
        + &"deviceName".to_string()
        + &device_name.to_string()
        + &"productKey".to_string()
        + &product_key.to_string()
        + &"timestamp".to_string()
        + &timestamp.to_string();
    let password = HMAC::mac(&sign_src.into_bytes(), device_secret);
    //log::info!("password ={:02x?}", password);
    let passwd_str = password.encode_hex::<String>();

    // Define options
    let connOpt1 = mqtt::CreateOptionsBuilder::new()
        .server_uri(host)
        .client_id(extended_client_id.to_string())
        .finalize();

    // Define connection options.
    let connOpt2 = mqtt::ConnectOptionsBuilder::new()
        .keep_alive_interval(Duration::from_secs(keep_alive_s))
        .clean_session(false)
        .user_name(user_name)
        .password(passwd_str)
        .finalize();

    return (connOpt1, connOpt2);
}
