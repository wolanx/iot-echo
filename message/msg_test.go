package message

import (
	"fmt"
	"testing"
)

func TestAsd(t *testing.T) {
	a, b := getCpuMem()
	fmt.Println(a)
	fmt.Println(b)
}

func TestGetMetric(t *testing.T) {
	fmt.Println(GetMetric())
}
