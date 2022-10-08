package mqtt

import (
	"fmt"
	"github.com/wolanx/iot-echo/pkg/util"
	"os/exec"
	"strings"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"github.com/wolanx/iot-echo/pkg/config"
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

var DefaultPublishHandler MQTT.MessageHandler = func(c MQTT.Client, msg MQTT.Message) {
	topic := msg.Topic()
	device := config.GetConfig().Device
	topic = strings.Replace(topic, device.ProductKey+"/"+device.DeviceName, "{pk}/{dn}", 1)
	log.Debug("topic = ", topic)
	log.Debug("payload = ", string(msg.Payload()))

	if strings.Contains(topic, "/sys/{pk}/{dn}/rrpc/request/") {
		RRpcHandle(c, topic, string(msg.Payload()))
		return
	}

	switch topic {
	case "/sys/{pk}/{dn}/thing/config/push":
		config.SaveParamsYaml(msg.Payload())
	default:
		log.Warn("miss topic = ", topic)
	}
}

func RRpcHandle(c MQTT.Client, topic string, payload string) {
	device := config.GetConfig().Device
	uuid := topic[strings.LastIndex(topic, "/")+1:]
	topicRet := "/sys/" + device.ProductKey + "/" + device.DeviceName + "/rrpc/response/" + uuid

	switch payload {
	case "LoadConfigInputs":
		yml := util.FileGetContents(config.Dir + "/params.yaml")
		Publish(c, topicRet, yml)
	case "ip addr":
		command := exec.Command("/bin/sh", "-c", "ip addr")
		output, err := command.CombinedOutput()
		if err != nil {
			log.Error(err)
		}
		Publish(c, topicRet, string(output))
	default:
		log.Warn("Unknown cmd " + payload)
		Publish(c, topicRet, "Unknown cmd "+payload)
	}
}
