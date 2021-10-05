package main

import (
	"fmt"
	"math"
	"testing"

	"github.com/zx5435/iot-echo/protocol/crc16"
)

func TestAsd(t *testing.T) {
	//a := 0x429b8a3d // 77.77
	a := 0xa0207439
	fmt.Println(a)

	f := math.Float32frombits(uint32(a))
	fmt.Println(f)
}

func TestZxczxc(t *testing.T) {
	//var a float32 = 77.77 // 429b8a3d
	var a float32 = 79.92126 // 429fd7af
	fmt.Println(a)

	f := math.Float32bits(a)
	fmt.Printf("%x\n", f)
}

func TestCrc(t *testing.T) {
	// 27 05 20 01     02 01 3E    46 50
	//a := []byte{
	//	0x27, 0x07, 0x20, 0x01,
	//	0x02, 0xc3, 0x02, 0x10, 0x1a,
	//	0x90, 0x1c,
	//}
	a := []byte{
		0x27, 0x0b, 0x21, 0x04,
		0x02, 0x07, 0x20, 0x22, 0x25, 0x27, 0x51, 0x98, 0x99,
		0xcc, 0x5a,
	}
	//a := []byte{
	//	27 05 20 01 02 01 3E 46 22
	//	0x27, 0x05, 0x20, 0x01,
	//	0x02, 0x01, 0x3E,
	//	0x46, 0x22,
	//}
	fmt.Printf("%x\n", a)

	table := crc16.MakeTable(crc16.CRC16_GENIBUS)
	crc := crc16.Checksum([]byte{
		//0x02, 0x07, 0x20, 0x22, 0x25, 0x27, 0x51, 0x98, 0x99,
		0x02, 0x01, 0x3E,
	}, table)
	fmt.Printf("%x\n", crc)
}

func TestBuildGeNiBus(t *testing.T) {
	fmt.Printf("% x\n", BuildGeNiBus(0x27, 0x20, 0x01, []byte{0x02, 0x01, 0x3E}))
}