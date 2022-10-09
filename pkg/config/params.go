package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/wolanx/iot-echo/pkg/core/calc"
	"github.com/wolanx/iot-echo/pkg/protocol/modbus"
	"github.com/wolanx/iot-echo/pkg/util"
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
	Namespace string `yaml:"namespace"` // 名字空间，sn_xx，后缀
	Name      string
	Value     string // 计算值

	// 脚本获取
	Script string

	// 协议相关
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

func (p *Params) LoadData(sn string) map[string]map[string]interface{} {
	map2 := make(map[string]map[string]interface{})

	ts := int32(time.Now().Unix())
	cpuPct, memPct := util.GetCpuMem()

	for _, point := range p.Attributes {
		snAll := getSnAll(sn, point.Namespace)
		map1 := make(map[string]interface{})
		map1["ts"] = ts
		map1["sn"] = snAll
		map1["cpu"] = util.Less6(cpuPct)
		map1["mem"] = util.Less6(memPct)
		map2[snAll] = map1
	}

	for _, point := range p.Attributes {
		snAll := getSnAll(sn, point.Namespace)
		var val interface{}
		map1 := map2[snAll]

		if point.Value != "" {
			// 计算值
			val = calc.Calc(point.Value, map1)
		} else if point.Script != "" {
			// 脚本获取
			valStr := util.RunShell(point.Script)
			var err interface{}
			val, err = strconv.ParseFloat(strings.TrimSpace(valStr), 64)
			if err != nil {
				fmt.Println("script error in: ", valStr)
				continue
			}
		} else {
			// 协议值
			gName := point.ChannelRefName
			client := p.DataGroups[gName].Client
			log.Debugf("%s * %+v", gName, point)

			var (
				xVal []byte
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
		}
		map1[point.Name] = val
	}

	j, _ := json.MarshalIndent(map2, "", "  ")
	fmt.Println(string(j))
	return map2
}

func getSnAll(sn string, suffix string) string {
	if suffix == "" {
		return sn
	} else {
		return sn + "_" + suffix
	}
}

func GetMetric() []string {
	map2 := GetParams().LoadData(GetConfig().Device.DeviceName)

	ret := []string{}
	for _, map1 := range map2 {
		aaa, _ := json.Marshal(map1)
		ret = append(ret, string(aaa))
	}

	return ret
}
