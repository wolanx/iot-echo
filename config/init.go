package config

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

var v *Model

func GetConfig() *Model {
	fmt.Println("config init.")
	maps := viper.AllSettings()
	_ = mapstructure.Decode(maps, &v)
	return v
}
