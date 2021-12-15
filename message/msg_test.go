package message

import (
	"fmt"
	"testing"

	"github.com/zx5435/iot-echo/config"
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
