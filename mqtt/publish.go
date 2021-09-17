package mqtt

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
)

func Publish(c mqtt.Client, topic string, msg string) mqtt.Token {
	token := c.Publish(topic, 0, false, msg)
	token.Wait()

	fmt.Println("publish msg:", msg)
	return token
}

func Subscribe(c mqtt.Client, topic string) mqtt.Token {
	token := c.Subscribe(topic, 0, nil)
	token.Wait()

	log.Info("Subscribe " + topic + " success")
	return token
}
