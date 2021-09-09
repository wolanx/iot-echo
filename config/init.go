package config

import (
	"fmt"
	"os"

	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Dir string

var v *Model

func init() {
	dir, _ := os.UserHomeDir()
	Dir = dir + "/iot-echo"
	if _, err := os.Stat(Dir); err != nil {
		err := os.Mkdir(Dir, os.ModeDir)
		if err != nil {
			log.Error(err.Error())
		}
	}
}

func GetConfig() *Model {
	if v == nil {
		fmt.Println("config init.")
		maps := viper.AllSettings()
		_ = mapstructure.Decode(maps, &v)
	}
	return v
}
