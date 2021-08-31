package debug

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zx5435/iot-echo/config"
	"github.com/zx5435/iot-echo/util"
	"os"
	"time"
)

func Run(cmd *cobra.Command, args []string) {
	maps := viper.AllSettings()
	var conf config.Model
	_ = mapstructure.Decode(maps, &conf)
	var (
		productKey   = conf.Device.ProductKey
		deviceName   = conf.Device.DeviceName
		deviceSecret = conf.Device.DeviceSecret
	)
	var (
		timeStamp          = "1528018257135"
		clientId           = "go_device_id_0001"
		subTopicUserGet    = "/" + productKey + "/" + deviceName + "/user/get"
		pubTopicUserUpdate = "/" + productKey + "/" + deviceName + "/user/update"
	)

	// tcp://localhost:1883
	//url := "tls://" + productKey + ".iot-as-mqtt.cn-shanghai.aliyuncs.com:1883"
	url := conf.Server.Host + ":1883"
	if conf.Server.Tls {
		url = "tls://" + url
	} else {
		url = "tcp://" + url
	}
	fmt.Println(url)
	opts := MQTT.NewClientOptions().AddBroker(url)

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
		if i%10 == 0 {
			token := c.Publish(pubTopicUserUpdate, 0, false, text)
			token.Wait()
		}
		fmt.Println("publish msg:", i, text)
		time.Sleep(2 * time.Second)
	}
}
