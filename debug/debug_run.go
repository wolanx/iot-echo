package debug

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/zx5435/iot-echo/config"
	"github.com/zx5435/iot-echo/message"
	"github.com/zx5435/iot-echo/util"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	} else {
		log.Debug("Connect IoT Cloud Success")
	}

	if token := c.Subscribe(topicUserGet, 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	} else {
		log.Info("Subscribe topic " + topicUserGet + " success")
	}

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
		msg := message.GetMetric()
		token := c.Publish(topicUserUpdate, 0, false, msg)
		token.Wait()
		fmt.Println("publish msg:", msg)
		time.Sleep(3 * time.Second)
	}
}

func NewClient(productKey string, deviceName string, deviceSecret string) MQTT.Client {
	cfg := config.GetConfig()
	url := cfg.Server.Host + ":1883"
	if cfg.Server.Tls {
		url = "tls://" + url
	} else {
		url = "tcp://" + url
	}
	fmt.Println(url)

	opt := MQTT.NewClientOptions().AddBroker(url)
	auth := util.CalculateSign("go_device_id_0001", productKey, deviceName, deviceSecret, "1528018257135")
	opt.SetClientID(auth.MqttClientId)
	opt.SetUsername(auth.Username)
	opt.SetPassword(auth.Password)
	opt.SetKeepAlive(1 * 60 * time.Second)
	opt.SetDefaultPublishHandler(util.DefaultPublishHandler)

	//tlsconfig := util.NewTLSConfig()
	//opt.SetTLSConfig(tlsconfig)

	c := MQTT.NewClient(opt)
	return c
}
