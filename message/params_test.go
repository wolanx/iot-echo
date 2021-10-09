package message

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/zx5435/iot-echo/config"
)

func TestP(t *testing.T) {
	s := FileGetContents(config.Dir + "/params.yaml")
	//fmt.Println(s)

	params := &config.Params{}
	params.Init(s)
	params.LoadGroup()
	//params.Print()
	params.LoadData()
}

func FileGetContents(path string) string {
	file, _ := os.Open(path)
	fileText, _ := ioutil.ReadAll(file)

	return string(fileText)
}
