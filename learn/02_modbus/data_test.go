package main

import (
	"fmt"
	"math"
	"testing"
)

func TestAsd(t *testing.T) {
	a := 0x429b8a3d
	fmt.Println(a)

	f := math.Float32frombits(uint32(a))
	fmt.Println(f)
}
