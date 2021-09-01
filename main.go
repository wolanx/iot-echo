package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/zx5435/iot-echo/cmd"
)

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{})
}

func main() {
	cmd.Execute()
}
