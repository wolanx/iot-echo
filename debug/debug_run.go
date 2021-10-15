package debug

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/zx5435/iot-echo/config"
	"github.com/zx5435/iot-echo/message"
	"github.com/zx5435/iot-echo/mqtt"
	"github.com/zx5435/iot-echo/util"
)

func Run(cmd *cobra.Command, args []string) {
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

	c := NewClient(productKey, deviceName, deviceSecret)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	log.Debug("Connect ok")

	mqtt.Subscribe(c, topicUserGet)
	mqtt.Subscribe(c, "/ota/device/upgrade/"+productKey+"/"+deviceName)
	mqtt.Subscribe(c, "/sys/"+productKey+"/"+deviceName+"/thing/config/push")
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
		msg1 := message.GetMetric()
		mqtt.Publish(c, topicUserUpdate, msg1)
		time.Sleep(10 * time.Second)
	}
}

func NewClient(productKey string, deviceName string, deviceSecret string) MQTT.Client {
	cfg := config.GetConfig()
	var url string
	if cfg.Provider == "iothub-echo" {
		if cfg.Server.Tls {
			url = "tls://" + cfg.Server.Host + ":1883"
		} else {
			url = "tcp://" + cfg.Server.Host + ":1883"
		}
	} else {
		url = "tls://" + cfg.Device.ProductKey + ".iot-as-mqtt.cn-shanghai.aliyuncs.com:1883"
	}
	fmt.Println(url)

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
