package main

import (
	"flag"
	"fmt"
	log2 "log"
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	_ "github.com/wolanx/iot-echo/pkg/log"
	"github.com/wolanx/iot-echo/pkg/protocol/modbus"
)

var tty *int

// Modbus总结
// https://www.cnblogs.com/iluzhiyong/p/4301192.html
func main() {
	tty = flag.Int("tty", 2, "tty flag")
	flag.Parse()

	log.Info("IsLittleEndian", IsLittleEndian())

	testGeni32(4)
}

func testModbusRtuSlave10() {
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
	results, err := client.ReadHoldingRegisters(0x0A, 6, 2)
	doLog2(err, results)
}

func testModbusTcpSlave10() {
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

func testGeni32(tty int) {
	// tty 4 Geni
	address := "/dev/ttyO" + strconv.Itoa(tty)
	handler := modbus.NewRTUClientHandler(address)
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "N"
	handler.StopBits = 2 // Geni 是 2
	handler.Timeout = 10 * time.Second
	handler.SlaveId = 0x01
	handler.Logger = log2.New(os.Stdout, "test: ", log2.Lshortfile)
	HandlerInfo(handler)
	err := handler.Connect()
	doLog1(err)
	defer handler.Close()

	client := modbus.NewClient(handler)

	dd := BuildGeniBus(0x27, 0x20, 0x01, []byte{0x02, 0x01, 0x3E})
	fmt.Printf("% x\n", dd)
	results, err := ReadByRaw(client, dd) // SysDischargePressure 79.92

	doLog2(err, results)
}
