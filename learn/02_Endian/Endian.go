package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println(IsLittleEndian())
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
