package test

import (
	"fmt"
	"github.com/wolanx/iot-echo/pkg/util"
	"testing"

	"github.com/wolanx/iot-echo/pkg/config"
)

func TestAsd(t *testing.T) {
	a, b := util.GetCpuMem()
	fmt.Println("cpu:", a)
	fmt.Println("mem:", b)
}

func TestGetMetric(t *testing.T) {
	fmt.Println(config.GetMetric())
}

func TestGetDiskPercent(t *testing.T) {
	fmt.Println("disk:", util.GetDiskPercent())
}
