package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/zx5435/iot-echo/core/calc"
	"github.com/zx5435/iot-echo/message"
	"github.com/zx5435/iot-echo/protocol/modbus"
	"github.com/zx5435/iot-echo/util"
	"gopkg.in/yaml.v2"
)

var ins *Params

func init() {
	LoadParams()
}

func GetParams() *Params {
	return ins
}

func LoadParams() {
	yml := util.FileGetContents(Dir + "/params.yaml")
	params := &Params{}
	params.Init(yml)
	ins = params
}

func SaveParamsYaml(s []byte) {
	err := ioutil.WriteFile(Dir+"/params.yaml", s, 0666)
	if err != nil {
		log.Error(err)
	}
	LoadParams()
}

type Params struct {
	Channels   []Channel
	Attributes []Attribute
	DataGroups map[string]*DataGroup
}

type Channel struct {
	Name     string
	Network  string // rtu tcp
	Endpoint string
	Protocol string // modbus geNi
}

type Attribute struct {
	Name           string
	Value          string
	ChannelRefName string `yaml:"channelRefName"`
	SlaveId        int    `yaml:"slaveId"`
	Address        int
	DataType       string `yaml:"dataType"`
}

type Client struct {
	packager    modbus.Packager
	transporter modbus.Transporter
}

type DataGroup struct {
	Name   string
	Client modbus.Client
}

func (p *Params) Init(s string) {
	err := yaml.Unmarshal([]byte(s), p)
	if err != nil {
		log.Error(err)
		return
	}

	p.DataGroups = make(map[string]*DataGroup)
	for _, channel := range p.Channels {
		p.DataGroups[channel.Name] = &DataGroup{
			Name:   channel.Name,
			Client: createClientByChannel(channel),
		}
	}
}

func createClientByChannel(c Channel) modbus.Client {
	switch c.Network {
	case "rtu":
		handler := modbus.NewRTUClientHandler(c.Endpoint)
		handler.BaudRate = 9600
		handler.DataBits = 8
		handler.Parity = "N"
		handler.StopBits = 1
		handler.Timeout = 10 * time.Second
		handler.SlaveId = 0x0A
		client := modbus.NewClient(handler)
		return client
	case "tcp":
		handler := modbus.NewTCPClientHandler(c.Endpoint)
		client := modbus.NewClient(handler)
		return client
	}
	return nil
}

func (p *Params) LoadData() map[string]interface{} {
	ret := make(map[string]interface{})

	cpuPct, memPct := message.GetCpuMem()
	ret["cpu"] = message.F2(cpuPct)
	ret["mem"] = message.F2(memPct)

	for _, point := range p.Attributes {
		if point.Value != "" {
			ret[point.Name] = calc.Calc(point.Value, ret)
		} else {
			gName := point.ChannelRefName
			client := p.DataGroups[gName].Client
			log.Debugf("%s * %+v", gName, point)

			var (
				xVal []byte
				val  interface{}
				err  error
			)

			switch point.DataType {
			case "float":
				xVal, err = client.ReadHoldingRegisters(byte(point.SlaveId), uint16(point.Address), 2)
				if err == nil {
					val = util.Byte4ToFloat32(xVal)
				}
			case "int":
				fallthrough
			default:
				xVal, err = client.ReadHoldingRegisters(byte(point.SlaveId), uint16(point.Address), 1)
				if err == nil {
					val = util.Byte2ToInt(xVal)
				}
			}
			if err != nil {
				log.Warn(err)
			}
			log.Infof("%s [% x] = %v", point.Name, xVal, val)
			ret[point.Name] = val
		}
	}

	j, _ := json.MarshalIndent(ret, "", "  ")
	fmt.Println(string(j))
	return ret
}

func GetMetric() string {
	//arr := make(map[string]interface{})
	arr := GetParams().LoadData()

	arr["ts"] = int32(time.Now().Unix())
	arr["sn"] = GetConfig().Device.DeviceName
	ret, _ := json.Marshal(arr)
	return string(ret)
}
