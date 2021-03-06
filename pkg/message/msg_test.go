package message

import (
	"fmt"
	"testing"

	"github.com/wolanx/iot-echo/pkg/config"
)

func TestAsd(t *testing.T) {
	a, b := GetCpuMem()
	fmt.Println("cpu:", a)
	fmt.Println("mem:", b)
}

func TestGetMetric(t *testing.T) {
	fmt.Println(config.GetMetric())
}

func TestGetDiskPercent(t *testing.T) {
	fmt.Println("disk:", GetDiskPercent())
}
