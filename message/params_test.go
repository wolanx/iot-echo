package message

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/zx5435/iot-echo/config"
	"github.com/zx5435/iot-echo/util"
)

func TestP(t *testing.T) {
	s := util.FileGetContents(config.Dir + "/params.yaml")
	//fmt.Println(s)

	params := &config.Params{}
	params.Init(s)
	params.LoadGroup()
	//params.Print()
	data := params.LoadData()
	j, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(j))
}
