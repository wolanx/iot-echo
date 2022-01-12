module github.com/wolanx/iot-echo

go 1.14

replace golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190411191339-88737f569e3a

require (
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/antlr/antlr4/runtime/Go/antlr v0.0.0-20211213210530-5d6a78255383
	github.com/eclipse/paho.mqtt.golang v1.3.5
	github.com/goburrow/serial v0.1.0
	github.com/kardianos/service v1.2.0
	github.com/mitchellh/mapstructure v1.4.1
	github.com/shirou/gopsutil v3.21.7+incompatible
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.8.1
	github.com/tklauser/go-sysconf v0.3.9 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	gopkg.in/yaml.v2 v2.4.0
)
