package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/takama/daemon"
)

const (
	name        = "iot-echo"
	description = "My Echo Service"
)

var dependencies []string

var stdout, stderr *log.Logger

func init() {
	stdout = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	stderr = log.New(os.Stderr, "", log.Ldate|log.Ltime)
}

// cd Z:\www\src\github.com\zx5435\iot-echo\learn
func main() {
	srv, err := daemon.New(name, description, daemon.SystemDaemon, dependencies...)
	if err != nil {
		stderr.Println("Error: ", err)
		os.Exit(1)
	}
	service := &Service{srv}
	status, err := service.Manage()
	if err != nil {
		stderr.Println(status, "\nError: ", err)
		os.Exit(1)
	}
	fmt.Println(status)
}

// Accept a client connection and collect it in a channel
func acceptConnection() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello World!"))
}

type Service struct {
	daemon.Daemon
}

// Manage by daemon commands or run the daemon
func (service *Service) Manage() (string, error) {
	usage := "Usage: install | remove | start | stop | status"

	// if received any kind of command, do it
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "install":
			return service.Install()
		case "remove":
			return service.Remove()
		case "start":
			return service.Start()
		case "stop":
			return service.Stop()
		case "status":
			return service.Status()
		default:
			status, err := service.Status()
			if err != nil {
				return "", err
			}
			stdout.Println(status)

			stdout.Println("	this body")

			sig := make(chan os.Signal, 1)
			done := make(chan bool, 1)
			signal.Notify(sig, os.Interrupt, os.Kill, syscall.SIGTERM)

			go func() {
				go acceptConnection()
				<-sig
				done <- true
			}()

			stdout.Println("awaiting sig")
			<-done
			stdout.Println("exiting")
			return usage, nil
		}
	}

	return "nono", nil
}
