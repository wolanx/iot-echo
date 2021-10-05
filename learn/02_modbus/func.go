package main

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"unsafe"

	log "github.com/sirupsen/logrus"
	"github.com/zx5435/iot-echo/protocol/crc16"
	"github.com/zx5435/iot-echo/protocol/modbus"
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
	count := int(response.Data[0])
	length := len(response.Data) - 1
	if count != length {
		err = fmt.Errorf("modbus: response data size '%v' does not match count '%v'", length, count)
		return
	}
	results = response.Data[1:]
	return
}

func BuildGeNiBus(addr byte, from byte, to byte, data []byte) []byte {

	table := crc16.MakeTable(crc16.CRC16_GENIBUS)
	crc := crc16.Checksum([]byte{
		//0x02, 0x07, 0x20, 0x22, 0x25, 0x27, 0x51, 0x98, 0x99,
		0x02, 0x01, 0x3E,
	}, table)
	ret := []byte{addr}
	ret = append(ret, byte(len(data)+2))
	ret = append(ret, from)
	ret = append(ret, to)
	ret = append(ret, data...)
	buf := []byte{0x00, 0x00}
	binary.BigEndian.PutUint16(buf, crc)
	ret = append(ret, buf...)
	return ret
}
