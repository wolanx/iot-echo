package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
)

type AuthInfo struct {
	password, username, mqttClientId string
}

func calculateSign(clientId, productKey, deviceName, deviceSecret, timeStamp string) AuthInfo {
	var rawPasswd bytes.Buffer
	rawPasswd.WriteString("clientId" + clientId)
	rawPasswd.WriteString("deviceName")
	rawPasswd.WriteString(deviceName)
	rawPasswd.WriteString("productKey")
	rawPasswd.WriteString(productKey)
	rawPasswd.WriteString("timestamp")
	rawPasswd.WriteString(timeStamp)
	fmt.Println(rawPasswd.String())

	// hmac, use sha1
	mac := hmac.New(sha1.New, []byte(deviceSecret))
	mac.Write([]byte(rawPasswd.String()))
	password := fmt.Sprintf("%02x", mac.Sum(nil))
	fmt.Println(password)
	username := deviceName + "&" + productKey

	var MQTTClientId bytes.Buffer
	MQTTClientId.WriteString(clientId)
	// hmac, use sha1; securemode=2 means TLS connection
	MQTTClientId.WriteString("|securemode=2,_v=paho-go-1.0.0,signmethod=hmacsha1,timestamp=")
	MQTTClientId.WriteString(timeStamp)
	MQTTClientId.WriteString("|")

	auth := AuthInfo{password: password, username: username, mqttClientId: MQTTClientId.String()}
	return auth
}
