package main

import (
	"fmt"
	"io/ioutil"
	"os/user"

	log "github.com/sirupsen/logrus"
	"gopkg.in/gcfg.v1"
)

func main() {
	current, _ := user.Current()
	filename := current.HomeDir + "/.iot-echo/iot-echo.ini"
	_, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Error(err)
	}

	iniCtx := struct {
		AliYun struct {
			ProductKey string
			DeviceName string
			DeviceSecret string
		}
	}{}
	err2 := gcfg.ReadFileInto(&iniCtx, filename)
	if err2 != nil {
		println(err2)
	}
	fmt.Println(iniCtx.AliYun)
	fmt.Println(iniCtx.AliYun.DeviceName)
	fmt.Println()
}
