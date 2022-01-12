package cmd

import (
	"github.com/kardianos/service"
	log "github.com/sirupsen/logrus"
	"github.com/wolanx/iot-echo/pkg/core"
)

var sign service.Service

func init() {
	svcConfig := &service.Config{
		Name:        "iot-echo",
		DisplayName: "iot-echo",
		Description: "My Echo Service",
		Arguments:   []string{"run"},
	}

	pg := &daemon{}
	var err error
	sign, err = service.New(pg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
}

type daemon struct {
}

func (p *daemon) Start(s service.Service) error {
	log.Info("start")
	// 真正开始干事
	go core.DefaultWeb()
	go core.Run()
	return nil
}

func (p *daemon) Stop(s service.Service) error {
	log.Info("stop")
	return nil
}
