package cmd

import (
	"github.com/kardianos/service"
	log "github.com/sirupsen/logrus"
	"github.com/zx5435/iot-echo/web"
)

var sign service.Service

type daemon struct {
}

func (p *daemon) Start(s service.Service) error {
	log.Info("start")
	go func() {
		web.DefaultWeb()
	}()
	return nil
}

func (p *daemon) Stop(s service.Service) error {
	log.Info("stop")
	return nil
}

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
