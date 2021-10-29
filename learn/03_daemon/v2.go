package main

import (
	"log"
	"os"
	"time"

	"github.com/kardianos/service"
)

type program struct {
}

func (p *program) Start(s service.Service) error {

	go p.run()
	return nil
}

func (p *program) run() {

	for {

		time.Sleep(time.Second)
		log.Println("running")
	}
}

func (p *program) Stop(s service.Service) error {

	return nil
}

func init() {

	f, err := os.Create("d:/gowinservice.txt")
	if err != nil {

		log.Fatal(err)
	}
	log.SetOutput(f)
}

func main() {

	svcConfig := &service.Config{

		Name:        "GoService",
		DisplayName: "GoServiceDis",
		Description: "windows service form golang",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {

		log.Fatal(err)
	}

	if len(os.Args) > 1 {

		if os.Args[1] == "install" {

			s.Install()
			log.Println("服务安装成功")
			return
		}

		if os.Args[1] == "remove" {

			s.Uninstall()
			log.Println("服务卸载成功")
			return
		}
	}

	if err = s.Run(); err != nil {

		log.Fatal(err)
	}
}
