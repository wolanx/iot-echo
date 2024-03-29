package core

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"github.com/wolanx/iot-echo/pkg/config"
	"github.com/wolanx/iot-echo/pkg/mqtt"
	"github.com/wolanx/iot-echo/pkg/util"
)

func Run() {
	cfg := config.GetConfig()
	var (
		productKey      = cfg.Device.ProductKey
		deviceName      = cfg.Device.DeviceName
		deviceSecret    = cfg.Device.DeviceSecret
		topicUserGet    = "/" + productKey + "/" + deviceName + "/user/get"
		topicUserUpdate = "/" + productKey + "/" + deviceName + "/user/update"
	)

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	c := newClient(productKey, deviceName, deviceSecret)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	log.Debug("Connect ok")

	mqtt.Subscribe(c, topicUserGet)
	mqtt.Subscribe(c, "/ota/device/upgrade/"+productKey+"/"+deviceName)            // 固件升级信息下行
	mqtt.Subscribe(c, "/sys/"+productKey+"/"+deviceName+"/thing/config/push")      // 云端主动下推配置信息
	mqtt.Subscribe(c, "/sys/"+productKey+"/"+deviceName+"/thing/config/get_reply") // 云端响应配置信息
	// todo restart

	mqtt.Publish(c, "/ota/device/inform/"+productKey+"/"+deviceName, "0.1.0")
	mqtt.Publish(c, "/ota/device/progress/"+productKey+"/"+deviceName, "10%")
	mqtt.Publish(c, "/ota/device/progress/"+productKey+"/"+deviceName, "20%")
	mqtt.Publish(c, "/ota/device/inform/"+productKey+"/"+deviceName, "0.2.0")

	go func() {
		<-sig
		fmt.Println("exiting...")
		if token := c.Unsubscribe(topicUserGet); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
		}
		c.Disconnect(250)
		fmt.Println("exited.")
		os.Exit(0)
	}()

	for i := 1; ; i++ {
		msgArr := config.GetMetric()
		for idx := range msgArr {
			mqtt.Publish(c, topicUserUpdate, msgArr[idx])
		}
		time.Sleep(30 * time.Second)
	}
}

func newClient(productKey string, deviceName string, deviceSecret string) MQTT.Client {
	cfg := config.GetConfig()
	var url string
	if cfg.Server.Tls {
		url = "tls://" + cfg.Server.Host + ":1883"
	} else {
		url = "tcp://" + cfg.Server.Host + ":1883"
	}
	log.Info(url)

	opt := MQTT.NewClientOptions().AddBroker(url)
	auth := util.CalculateSign("go_device_id_0001", productKey, deviceName, deviceSecret, "1528018257135")
	opt.SetClientID(auth.MqttClientId)
	opt.SetUsername(auth.Username)
	opt.SetPassword(auth.Password)
	opt.SetKeepAlive(1 * 60 * time.Second)
	opt.SetDefaultPublishHandler(mqtt.DefaultPublishHandler)

	if cfg.Server.Tls {
		tlsConfig := util.NewTLSConfig()
		opt.SetTLSConfig(tlsConfig)
	}

	c := MQTT.NewClient(opt)
	return c
}
