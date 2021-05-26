package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {
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
	auth := calculateSign(clientId, productKey, deviceName, deviceSecret, timeStamp)
	opts.SetClientID(auth.mqttClientId)
	opts.SetUsername(auth.username)
	opts.SetPassword(auth.password)
	opts.SetKeepAlive(60 * 2 * time.Second)
	opts.SetDefaultPublishHandler(f)

	// set the tls configuration
	//tlsconfig := NewTLSConfig()
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
		fmt.Println("publish msg:", i)
		text := fmt.Sprintf("ABC #%d", i)
		token := c.Publish(pubTopic, 0, false, text)
		fmt.Println("publish msg: ", text)
		token.Wait()
		time.Sleep(2 * time.Second)
	}

	// unsubscribe from subTopic("/a1Zd7n5yTt8/deng/user/get")
	if token := c.Unsubscribe(subTopic); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	c.Disconnect(250)
}

func NewTLSConfig() *tls.Config {
	// Import trusted certificates from CAfile.pem.
	// Alternatively, manually add CA certificates to default openssl CA bundle.
	certpool := x509.NewCertPool()
	pemCerts, err := ioutil.ReadFile("./x509/root.pem")
	if err != nil {
		fmt.Println("0. read file error, game over!!")

	}

	certpool.AppendCertsFromPEM(pemCerts)

	// Create tls.Config with desired tls properties
	return &tls.Config{
		// RootCAs = certs used to verify server cert.
		RootCAs: certpool,
		// ClientAuth = whether to request cert from server.
		// Since the server is set up for SSL, this happens
		// anyways.
		ClientAuth: tls.NoClientCert,
		// ClientCAs = certs used to validate client cert.
		ClientCAs: nil,
		// InsecureSkipVerify = verify that cert contents
		// match server. IP matches what is in cert etc.
		InsecureSkipVerify: false,
		// Certificates = list of certs client sends to server.
		// Certificates: []tls.Certificate{cert},
	}
}

// define a function for the default message handler
var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}
