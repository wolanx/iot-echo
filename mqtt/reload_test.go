package mqtt

import (
	"fmt"
	"testing"

	"github.com/zx5435/iot-echo/config"
)

func TestAsda(t *testing.T) {
	fmt.Println(123)
	fmt.Println(config.GetParams())
	config.SaveParamsYaml([]byte("a: 1\nb: 2"))
	fmt.Println(config.GetParams())
}
