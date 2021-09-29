package main

import (
	"fmt"
	"math"
	"testing"
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
