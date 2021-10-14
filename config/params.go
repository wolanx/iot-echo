package config

import (
	"encoding/json"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/zx5435/iot-echo/protocol/modbus"
	"github.com/zx5435/iot-echo/util"
	"gopkg.in/yaml.v2"
)

var ParamsIns *Params

func init() {
	NewParams()
}

func NewParams() *Params {
	yml := util.FileGetContents(Dir + "/params.yaml")
	params := &Params{}
	params.Init(yml)
	ParamsIns = params
	return params
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
	Points []Point
}

type Point struct {
	Name     string
	SlaveId  byte
	Address  uint16
	DataType string
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
	for _, attribute := range p.Attributes {
		group, ok := p.DataGroups[attribute.ChannelRefName]
		if !ok {
			log.Error("ChannelRefName not found")
			break
		}
		group.Points = append(group.Points, Point{
			Name:     attribute.Name,
			SlaveId:  byte(attribute.SlaveId),
			Address:  uint16(attribute.Address),
			DataType: attribute.DataType,
		})
	}
}

func (p Params) Print() {
	for _, channel := range p.Channels {
		fmt.Printf("%+v\n", channel)
	}
	for _, attribute := range p.Attributes {
		fmt.Printf("%+v\n", attribute)
	}
	for k, group := range p.DataGroups {
		fmt.Printf("%s: %+v\n", k, group.Points)
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
	for _, group := range p.DataGroups {
		for _, point := range group.Points {
			log.Debugf("%s * %+v", group.Name, point)

			var (
				xVal []byte
				val  interface{}
				err  error
			)

			switch point.DataType {
			case "float":
				xVal, err = group.Client.ReadHoldingRegisters(point.SlaveId, point.Address, 2)
				if err == nil {
					val = util.Byte4ToFloat32(xVal)
				}
			case "int":
				fallthrough
			default:
				xVal, err = group.Client.ReadHoldingRegisters(point.SlaveId, point.Address, 1)
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
