package main

import (
	"encoding/hex"
	"fmt"

	"github.com/goburrow/modbus"
	log "github.com/sirupsen/logrus"
)

func ReadByRaw(mb modbus.Client, r []byte) (results []byte, err error) {
	request := modbus.ProtocolDataUnit{
		FunctionCode: 7,
		Data:         r,
	}
	response, err := mb.Send2(&request)
	if err != nil {
		return
	}
	count := int(response.Data[0])
	length := len(response.Data) - 1
	if count != length {
		err = fmt.Errorf("modbus: response data size '%v' does not match count '%v'", length, count)
		return
	}
	results = response.Data[1:]
	return
}

func doLog1(err error) {
	if err != nil {
		log.Error(err)
	} else {
		log.Info("connect1 ok")
	}
}

func doLog2(err error, res []byte) {
	if err != nil {
		log.Error(err)
	} else {
		log.Info("connect2 ok")
	}
	log.Info(res)
	log.Info(string(res))
	log.Info(hex.EncodeToString(res))
}
