package mqtt

import (
	"fmt"
	"strings"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"github.com/zx5435/iot-echo/config"
)

func Publish(c MQTT.Client, topic string, msg string) MQTT.Token {
	token := c.Publish(topic, 0, false, msg)
	token.Wait()

	fmt.Println("publish msg:", msg)
	return token
}

func Subscribe(c MQTT.Client, topic string) MQTT.Token {
	token := c.Subscribe(topic, 0, nil)
	token.Wait()

	log.Info("Subscribe " + topic + " success")
	return token
}

var DefaultPublishHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	topic := msg.Topic()
	log.Debug("topic = ", topic)
	log.Debug("payload = ", string(msg.Payload()))
	device := config.GetConfig().Device
	topic = strings.Replace(topic, device.ProductKey+"/"+device.DeviceName, "aaa/bbb", 1)
	log.Debug("topic = ", topic)
	// /sys/a1p9xMXq5Nd/iot-echo-903-913332/thing/config/push
	switch topic {
	case "/sys/aaa/bbb/thing/config/push":
	}
}
