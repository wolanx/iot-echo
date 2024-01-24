#![allow(non_snake_case)]

extern crate paho_mqtt as mqtt;

use std::{fs, process, thread, time::Duration};

use hex::ToHex;
use hmac_sha256::HMAC;
use mqtt::{ConnectOptions, CreateOptions};
use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize)]
struct ConfigYml {
    provider: String,
    server: ConfigServer,
    device: ConfigDevice,
}

#[derive(Debug, Serialize, Deserialize)]
struct ConfigServer {
    host: String,
    tls: String,
}

#[derive(Debug, Serialize, Deserialize)]
struct ConfigDevice {
    productKey: String,
    deviceName: String,
    deviceSecret: String,
}

// Subscribe to a single topic.
fn subscribe_topic(cli: &mqtt::Client, topic: &str) {
    if let Err(e) = cli.subscribe(topic, 0) {
        println!("Failed to subscribes topic: {:?}", e);
        process::exit(1);
    }
}

// Reconnect
fn try_to_reconnect(cli: &mqtt::Client) -> bool {
    println!("Disconnected. Waiting to retry connection");
    let cnt = 1;
    while cnt != 4 {
        thread::sleep(Duration::from_millis(cnt * 5000));
        if cli.reconnect().is_ok() {
            println!("Reconnected");
            return true;
        }
    }
    println!("Failed to reconnected.");
    false
}

fn getConnOpt(config: &ConfigYml) -> (CreateOptions, ConnectOptions) {
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
    //println!("password ={:02x?}", password);
    let passwd_str = password.encode_hex::<String>();

    // Define options for the create.
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

fn main() {
    let f = fs::read_to_string(r"C:\Users\106006\.iot-echo\config.yaml").unwrap();
    let config: ConfigYml = serde_yaml::from_str(&f).expect("error yaml");
    println!("config = {:#?}", config);
    let (connOpt1, connOpt2) = getConnOpt(&config);

    // Create a mqtt client.
    let insMqtt = mqtt::Client::new(connOpt1).unwrap_or_else(|err| {
        println!("Failed to create mqtt client: {:?}", err);
        process::exit(1);
    });

    // Define consumer
    let rx = insMqtt.start_consuming();

    // Connect and wait for results.
    if let Err(e) = insMqtt.connect(connOpt2) {
        println!("Failed to connect:\n\t{:?}", e);
        process::exit(1);
    }

    // Subscribe to topic "/${productKey}/${deviceName}/user/get"
    let sub_topic = "/".to_string()
        + &config.device.productKey.to_string()
        + &"/".to_string()
        + &config.device.deviceName.to_string()
        + &"/user/get".to_string();
    subscribe_topic(&insMqtt, &sub_topic);
    println!("subscribed to topic:={}", sub_topic);

    // Publish to topic "/${productKey}/${deviceName}/user/get"
    let pub_topic = "/".to_string()
        + &config.device.productKey.to_string()
        + &"/".to_string()
        + &config.device.deviceName.to_string()
        + &"/user/update".to_string();
    let payload = "{\"LightSwitch\":1}".to_string();
    let msg = mqtt::Message::new(pub_topic.clone(), payload.clone(), 0);
    if let Err(e) = insMqtt.publish(msg) {
        println!("Failed to subscribes topic: {:?}", e);
        process::exit(1);
    }
    println!("published to topic:={}", pub_topic.clone());

    println!("start receiving...");

    for message in rx.iter() {
        if let Some(message) = message {
            println!("{}", message);
        }

        if !insMqtt.is_connected() {
            if try_to_reconnect(&insMqtt) {
                //
            } else {
                println!("failed to reconnect...");
                break;
            }
        }
    }

    // Disconnect and exit now.
    if insMqtt.is_connected() {
        println!("Disconnecting");
        insMqtt.disconnect(None).unwrap();
    }
    println!("exit");
}