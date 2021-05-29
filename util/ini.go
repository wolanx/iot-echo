package util

import (
	"io/ioutil"
	"os/user"

	log "github.com/sirupsen/logrus"
	"gopkg.in/gcfg.v1"
)

var PConfig PConfigModel

type PConfigModel struct {
	AliYun struct {
		ProductKey   string
		DeviceName   string
		DeviceSecret string
	}
}

func LoadIni() {
	current, _ := user.Current()
	filename := current.HomeDir + "/.iot-echo/iot-echo.ini"
	_, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Error(err)
	}
	err2 := gcfg.ReadFileInto(&PConfig, filename)
	if err2 != nil {
		println(err2)
	}
}
