package main

//#cgo CFLAGS: -I./
//#cgo LDFLAGS: -L./ -lddd
//#include "libddd.h"
import "C"

func main() {
	C.nbgo()
}
