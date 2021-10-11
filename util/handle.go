package util

import (
	"fmt"
	"io/ioutil"
	"os"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
)

// DefaultPublishHandler define a function for the default message handler
var DefaultPublishHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	log.Info("topic", msg.Topic())
	fmt.Println("payload =", string(msg.Payload()))
}

func FileGetContents(path string) string {
	file, _ := os.Open(path)
	fileText, _ := ioutil.ReadAll(file)

	return string(fileText)
}
