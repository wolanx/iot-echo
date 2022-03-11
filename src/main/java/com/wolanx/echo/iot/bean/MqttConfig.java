package com.wolanx.echo.iot.bean;

import com.wolanx.echo.iot.util.CryptoUtil;
import lombok.Getter;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Component;
import org.yaml.snakeyaml.Yaml;

import java.io.File;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.util.Map;

@Slf4j
@Component
public class MqttConfig {

    @Getter
    private String productKey = "";

    @Getter
    private String deviceName = "";

    @Getter
    private String deviceSecret = "";

    @Getter
    private String username = "";

    @Getter
    private String password = "";

    @Getter
    private String clientId = "";

    public MqttConfig() throws FileNotFoundException {
        File dumpFile = new File(System.getProperty("user.home") + "/.iot-echo/config.yaml");
        Yaml yaml = new Yaml();
        Map<String, Map<String, String>> res = yaml.load(new FileReader(dumpFile));

        this.productKey = res.get("device").get("productKey");
        this.deviceName = res.get("device").get("deviceName");
        this.deviceSecret = res.get("device").get("deviceSecret");

        this.calculate();
    }

    public void calculate() {
        if (productKey == null || deviceName == null || deviceSecret == null) {
            return;
        }

        try {
            String timestamp = Long.toString(System.currentTimeMillis());
            String pwd = "clientId" + productKey + "." + deviceName + "deviceName" + deviceName + "productKey" + productKey + "timestamp" + timestamp;

            this.username = deviceName + "&" + productKey;
            this.password = CryptoUtil.hmacSha256(pwd, deviceSecret);
            this.clientId = productKey + "." + deviceName + "|" + "timestamp=" + timestamp + ",_v=paho-java-1.0.0,securemode=2,signmethod=hmacsha256|";
        } catch (Exception e) {
            e.printStackTrace();
        }

        log.debug("username: " + getUsername());
        log.debug("password: " + getPassword());
    }

}
