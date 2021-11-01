package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/kardianos/service"
)

type program struct {
}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}

func (p *program) Stop(s service.Service) error {
	return nil
}

func init() {
	f, err := os.Create("z:/log.txt")
	if err != nil {

		log.Fatal(err)
	}
	log.SetOutput(f)
}
func sayHello(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello World!"))
}

func main() {
	svcConfig := &service.Config{
		Name:        "GoService",
		DisplayName: "GoServiceDis",
		Description: "windows service form golang",
		Arguments:   []string{"run"},
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "install":
			fmt.Println(s.Install())
			log.Println("服务安装成功")
		case "remove":
			fmt.Println(s.Uninstall())
			log.Println("服务卸载成功")
		case "start":
			fmt.Println(s.Start())
		case "stop":
			fmt.Println(s.Stop())
		case "run":
			fmt.Println(s.Run())
		case "status":
			fmt.Println(s.Status())
		}
	}
}
