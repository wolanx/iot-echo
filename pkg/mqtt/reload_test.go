package mqtt

import (
	"fmt"
	"testing"

	"github.com/wolanx/iot-echo/pkg/config"
)

func TestAsda(t *testing.T) {
	fmt.Println(123)
	fmt.Println(config.GetParams())
	config.SaveParamsYaml([]byte("a: 1\nb: 2"))
	fmt.Println(config.GetParams())
}

func BenchmarkPublish(b *testing.B) {
	sum := 0
	for i := 0; i < b.N; i++ {
		sum += i
	}
	fmt.Println(sum)
}
