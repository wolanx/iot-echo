package main

import (
	"fmt"
	"os"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"github.com/zx5435/iot-echo/util"
)

func init() {
	util.LoadIni()
}

func main() {
	var (
		productKey   = util.PConfig.AliYun.ProductKey
		deviceName   = util.PConfig.AliYun.DeviceName
		deviceSecret = util.PConfig.AliYun.DeviceSecret
	)
	log.SetFormatter(&log.TextFormatter{})
	var (
		timeStamp          = "1528018257135"
		clientId           = "go_device_id_0001"
		subTopicUserGet    = "/" + productKey + "/" + deviceName + "/user/get"
		pubTopicUserUpdate = "/" + productKey + "/" + deviceName + "/user/update"
	)
	fmt.Println(pubTopicUserUpdate)

	opts := MQTT.NewClientOptions().AddBroker("tls://" + productKey + ".iot-as-mqtt.cn-shanghai.aliyuncs.com:1883")

	auth := util.CalculateSign(clientId, productKey, deviceName, deviceSecret, timeStamp)
	opts.SetClientID(auth.MqttClientId)
	opts.SetUsername(auth.Username)
	opts.SetPassword(auth.Password)
	opts.SetKeepAlive(60 * 2 * time.Second)
	opts.SetDefaultPublishHandler(util.DefaultPublishHandler)

	//tlsconfig := util.NewTLSConfig()
	//opts.SetTLSConfig(tlsconfig)

	c := MQTT.NewClient(opts)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		log.Debug("Connect IoT Cloud Success")
	}
	defer c.Disconnect(250)

	if token := c.Subscribe(subTopicUserGet, 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	} else {
		log.Info("Subscribe topic " + subTopicUserGet + " success")
	}
	defer func() {
		if token := c.Unsubscribe(subTopicUserGet); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		} else {
			println("Unsubscribed.")
		}
	}()

	for i := 1; ; i++ {
		text := fmt.Sprintf("ABC #%d", i)
		//token := c.Publish(pubTopicUserUpdate, 0, false, text)
		//token.Wait()
		fmt.Println("publish msg:", i, text)
		time.Sleep(2 * time.Second)
	}
}
