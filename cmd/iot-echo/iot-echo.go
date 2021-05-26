package main

import (
	"bytes"
	"fmt"
	"os"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"github.com/zx5435/iot-echo/util"
)

func main() {
	log.Info("qwe")
	var (
		// set the device info, include product key, device name, and device secret
		// set timestamp, clientid, subscribe topic and publish topic
		timeStamp = "1528018257135"
		clientId  = "go_device_id_0001"
		subTopic  = "/" + productKey + "/" + deviceName + "/user/get"
		pubTopic  = "/" + productKey + "/" + deviceName + "/user/update"
	)

	// set the login broker url
	var rawBroker bytes.Buffer
	rawBroker.WriteString("tls://")
	rawBroker.WriteString(productKey)
	rawBroker.WriteString(".iot-as-mqtt.cn-shanghai.aliyuncs.com:1883")
	opts := MQTT.NewClientOptions().AddBroker(rawBroker.String())

	// calculate the login auth info, and set it into the connection options
	auth := util.CalculateSign(clientId, productKey, deviceName, deviceSecret, timeStamp)
	opts.SetClientID(auth.MqttClientId)
	opts.SetUsername(auth.Username)
	opts.SetPassword(auth.Password)
	opts.SetKeepAlive(60 * 2 * time.Second)
	opts.SetDefaultPublishHandler(f)

	// set the tls configuration
	//tlsconfig := util.NewTLSConfig()
	//opts.SetTLSConfig(tlsconfig)

	// create and start a client using the above ClientOptions
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Print("Connect aliyun IoT Cloud Sucess\n")

	// subscribe to subTopic("/a1Zd7n5yTt8/deng/user/get") and request messages to be delivered
	if token := c.Subscribe(subTopic, 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	fmt.Print("Subscribe topic " + subTopic + " success\n")

	// publish 5 messages to pubTopic("/a1Zd7n5yTt8/deng/user/update")
	for i := 0; i < 5; i++ {
		text := fmt.Sprintf("ABC #%d", i)
		token := c.Publish(pubTopic, 0, false, text)
		fmt.Println("publish msg:", i)
		fmt.Println("publish msg: ", text)
		token.Wait()
		time.Sleep(2 * time.Second)
	}

	// unsubscribe from subTopic("/a1Zd7n5yTt8/deng/user/get")
	if token := c.Unsubscribe(subTopic); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	} else {
		println("exit")
	}

	c.Disconnect(250)
}

// define a function for the default message handler
var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}
