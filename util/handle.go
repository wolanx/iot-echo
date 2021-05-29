package util

import (
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
)

// DefaultPublishHandler define a function for the default message handler
var DefaultPublishHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	log.Info("topic", msg.Topic())
	fmt.Println("payload =", string(msg.Payload()))
}