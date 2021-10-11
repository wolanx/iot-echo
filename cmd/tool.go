package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zx5435/iot-echo/config"
	"github.com/zx5435/iot-echo/util"
)

var toolCmd = &cobra.Command{
	Use:   "tool",
	Short: "Start daemon process",
	Long:  "Long description.",
	Run:   DoTool,
}

func init() {
	rootCmd.AddCommand(toolCmd)
}

func DoTool(cmd *cobra.Command, args []string) {
	s := util.FileGetContents(config.Dir + "/params.yaml")
	//fmt.Println(s)

	params := &config.Params{}
	params.Init(s)
	params.LoadGroup()
	//params.Print()
	params.LoadData()
}
