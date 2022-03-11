package com.wolanx.echo.iot.bean;

import org.eclipse.paho.client.mqttv3.MqttClient;
import org.eclipse.paho.client.mqttv3.MqttConnectOptions;
import org.eclipse.paho.client.mqttv3.MqttException;
import org.eclipse.paho.client.mqttv3.persist.MemoryPersistence;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import javax.annotation.Resource;

@Configuration
public class MqttClientConfig {

    @Resource
    MqttConfig config;

    @Bean
    public MqttClient client() throws MqttException {
        String broker = "ssl://" + config.getProductKey() + ".iot-as-mqtt.cn-shanghai.aliyuncs.com" + ":443";
        MemoryPersistence persistence = new MemoryPersistence();
        MqttClient client = new MqttClient(broker, config.getClientId(), persistence);

        MqttConnectOptions options = new MqttConnectOptions();
        options.setCleanSession(true);
        options.setKeepAliveInterval(180);
        options.setUserName(config.getUsername());
        options.setPassword(config.getPassword().toCharArray());
        client.connect(options);
        return client;
    }

}
