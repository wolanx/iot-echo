package config

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/zx5435/iot-echo/protocol/modbus"
	"gopkg.in/yaml.v2"
)

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

type DataGroup struct {
	Name    string
	Handler *modbus.RTUClientHandler
	Points  []string
}

func (p *Params) Init(s string) {
	err := yaml.Unmarshal([]byte(s), p)
	if err != nil {
		log.Error(err)
		return
	}
}

func (p Params) Print() {
	for _, channel := range p.Channels {
		fmt.Printf("%+v\n", channel)
	}
	for _, attribute := range p.Attributes {
		fmt.Printf("%+v\n", attribute)
	}
	for k, v := range p.DataGroups {
		fmt.Println(k, v.Points)
	}
}

func (p *Params) LoadGroup() {
	p.DataGroups = make(map[string]*DataGroup)
	for _, channel := range p.Channels {
		p.DataGroups[channel.Name] = &DataGroup{
			Name:    channel.Name,
			Handler: &modbus.RTUClientHandler{},
		}
	}
	for _, attribute := range p.Attributes {
		group, ok := p.DataGroups[attribute.ChannelRefName]
		if !ok {
			log.Error("ChannelRefName not found")
			break
		}
		group.Points = append(group.Points, attribute.Name)
	}
}
