package message

import (
	"fmt"
	"testing"
)

func TestAsd(t *testing.T) {
	a, b := getCpuMem()
	fmt.Println("cpu:", a)
	fmt.Println("mem:", b)
}

func TestGetMetric(t *testing.T) {
	fmt.Println(GetMetric())
}

func TestGetDiskPercent(t *testing.T) {
	fmt.Println("disk:", GetDiskPercent())
}
