package main

import (
	"flag"
	log2 "log"
	"os"
	"strconv"
	"time"
	"unsafe"

	"github.com/goburrow/modbus"
	log "github.com/sirupsen/logrus"
)

var tty *int

// Modbus总结
// https://www.cnblogs.com/iluzhiyong/p/4301192.html
func main() {
	log.SetFormatter(&log.TextFormatter{})
	tty = flag.Int("tty", 2, "tty flag")
	flag.Parse()

	log.Info("IsLittleEndian", IsLittleEndian())

	testGeNi32()
}

func IsLittleEndian() bool {
	var value int32 = 1
	p := unsafe.Pointer(&value)

	pb := (*byte)(p)
	if *pb != 1 {
		return false
	}

	return true
}

func testGeNi32() {
	// tty 4 geni
	address := "/dev/ttyO" + strconv.Itoa(*tty)
	log.Info(address)
	handler := modbus.NewRTUClientHandler(address)
	handler.BaudRate = 38400
	handler.DataBits = 8
	handler.Parity = "N"
	handler.StopBits = 1
	handler.Timeout = 10 * time.Second
	handler.SlaveId = 0x24
	handler.Logger = log2.New(os.Stdout, "test: ", log2.Lshortfile)
	err := handler.Connect()
	doLog1(err)
	defer handler.Close()

	client := modbus.NewClient(handler)
	results, err := ReadByRaw(client, []byte{0x27, 0x07, 0x5A, 0x01, 0x0A, 0x03, 0x52, 0x00, 0x0A, 0x19, 0x36})
	doLog2(err, results)
}

func testMbRtuSlave10() {
	// tty 2 modbus
	address := "/dev/ttyO2"
	log.Info(address)
	handler := modbus.NewRTUClientHandler(address)
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "N"
	handler.StopBits = 1
	handler.Timeout = 10 * time.Second
	handler.SlaveId = 0x0A
	//handler.SlaveId = 0x24
	handler.Logger = log2.New(os.Stdout, "test: ", log2.Lshortfile)
	err := handler.Connect()
	doLog1(err)
	defer handler.Close()

	client := modbus.NewClient(handler)
	// ReadCoils 01
	// ReadDiscreteInputs 02
	// ReadHoldingRegisters 03
	// ReadInputRegisters 04
	results, err := client.ReadHoldingRegisters(6, 2)
	doLog2(err, results)
}

func testMbTcpSlave10() {
	//handler := modbus.NewTCPClientHandler("localhost:502")
	handler := modbus.NewTCPClientHandler("192.168.30.66:502")
	handler.Timeout = 10 * time.Second
	handler.SlaveId = 0x0A
	handler.Logger = log2.New(os.Stdout, "test: ", log2.Lshortfile)
	err := handler.Connect()
	doLog1(err)
	defer handler.Close()

	client := modbus.NewClient(handler)
	results, err := client.ReadInputRegisters(0, 100)
	doLog2(err, results)
}
