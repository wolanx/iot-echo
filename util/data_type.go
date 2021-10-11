package util

import "math"

func bytesToUint32(b []byte) uint32 {
	_ = b[3]
	return uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
}

func Byte4ToFloat32(b []byte) float32 {
	return math.Float32frombits(bytesToUint32(b))
}

func Byte2ToInt(b []byte) int {
	_ = b[1]
	u := uint16(b[0])<<8 | uint16(b[1])
	return int(u)
}
