package com.wolanx.echo.iot;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.sun.management.OperatingSystemMXBean;
import com.wolanx.echo.iot.bean.MqttConfig;
import lombok.extern.slf4j.Slf4j;
import org.eclipse.paho.client.mqttv3.MqttClient;
import org.eclipse.paho.client.mqttv3.MqttException;
import org.eclipse.paho.client.mqttv3.MqttMessage;
import org.springframework.boot.ApplicationArguments;
import org.springframework.boot.ApplicationRunner;
import org.springframework.stereotype.Component;

import javax.annotation.Resource;
import java.lang.management.ManagementFactory;
import java.util.LinkedHashMap;

@Slf4j
@Component
public class MqttApplicationRunner implements ApplicationRunner {

    @Resource
    MqttConfig config;

    @Resource
    MqttClient client;

    private static final OperatingSystemMXBean mxBean = (OperatingSystemMXBean) ManagementFactory.getOperatingSystemMXBean();

    @Override
    public void run(ApplicationArguments args) {
        try {
            boolean connected = client.isConnected();
            log.info("connected: " + connected);

            //Paho Mqtt 消息订阅
//            String topicReply = "/sys/" + productKey + "/" + deviceName + "/thing/event/property/post_reply";
//            client.subscribe(topicReply, new Mqtt3PostPropertyMessageListener());
//            System.out.println("subscribe: " + topicReply);

            while (true) {
                sendMsg();
                Thread.sleep(15 * 1000);
            }

//            client.disconnect();
//            System.out.println("Disconnected");
//            System.exit(0);
        } catch (MqttException e) {
            e.printStackTrace();
            System.out.println("reason " + e.getReasonCode());
            System.out.println("msg " + e.getMessage());
            System.out.println("loc " + e.getLocalizedMessage());
            System.out.println("cause " + e.getCause());
            System.out.println("excep " + e);
        } catch (InterruptedException | JsonProcessingException e) {
            e.printStackTrace();
        }
    }

    private void sendMsg() throws MqttException, JsonProcessingException {
        String topic = "/" + config.getProductKey() + "/" + config.getDeviceName() + "/user/update";
        LinkedHashMap<String, Object> m = new LinkedHashMap<>();
        m.put("ts", (int) (System.currentTimeMillis() / 1000));
        m.put("sn", config.getDeviceName());
        m.put("cpu", getCpu());
        m.put("mem", getMem());

        String content = new ObjectMapper().writeValueAsString(m);
        MqttMessage message = new MqttMessage(content.getBytes());
        message.setQos(0);

        log.debug("publish: " + content);
        client.publish(topic, message);
    }

    private static int getCpu() {
        double cpuLoad = mxBean.getSystemCpuLoad();
        return (int) (cpuLoad * 100);
    }

    private static int getMem() {
        double mem1 = mxBean.getFreePhysicalMemorySize();
        double mem2 = mxBean.getTotalPhysicalMemorySize();
        double value = mem1 / mem2;

        return (int) ((1 - value) * 100);
    }

}
