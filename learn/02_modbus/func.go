package main

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"unsafe"

	log "github.com/sirupsen/logrus"
	"github.com/wolanx/iot-echo/pkg/protocol/crc16"
	"github.com/wolanx/iot-echo/pkg/protocol/modbus"
)

func IsLittleEndian() bool {
	var value int32 = 1
	p := unsafe.Pointer(&value)

	pb := (*byte)(p)
	if *pb != 1 {
		return false
	}

	return true
}

func HandlerInfo(h *modbus.RTUClientHandler) {
	j, _ := json.MarshalIndent(h, "", "  ")
	fmt.Println(string(j))
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

func ReadByRaw(mb modbus.Client, r []byte) (results []byte, err error) {
	request := modbus.ProtocolDataUnit{
		FunctionCode: 7,
		Data:         r,
	}
	response, err := mb.SendRaw(&request)
	if err != nil {
		return
	}
	//count := int(response.Data[0])
	//length := len(response.Data) - 1
	////if count != length {
	////	err = fmt.Errorf("modbus: response data size '%v' does not match count '%v'", length, count)
	////	return
	////}
	results = response.Data
	return
}

// BuildGeniBus crc 从1开始到data
func BuildGeniBus(addr byte, from byte, to byte, data []byte) []byte {
	table := crc16.MakeTable()
	l := len(data) + 2
	checks := []byte{byte(l), from, to}
	checks = append(checks, data...)
	checksum := crc16.Checksum(checks, table)
	crc := []byte{0x00, 0x00}
	binary.BigEndian.PutUint16(crc, checksum)

	ret := []byte{addr, byte(l), from, to}
	ret = append(ret, data...)
	ret = append(ret, crc...)
	return ret
}
