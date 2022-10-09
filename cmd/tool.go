package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wolanx/iot-echo/pkg/config"
	"github.com/wolanx/iot-echo/pkg/core"
)

func init() {
	rootCmd.AddCommand(toolCmd)

	toolCmd.AddCommand(dataCmd)
	toolCmd.AddCommand(mqttCmd)
}

var toolCmd = &cobra.Command{
	Use:   "tool",
	Short: "Small tool",
	Long:  "Long description.",
}

var dataCmd = &cobra.Command{
	Use:   "data",
	Short: "Fetch data",
	Long:  "Long description.",
	Run: func(*cobra.Command, []string) {
		config.GetParams().LoadData("sn001")
	},
}

var mqttCmd = &cobra.Command{
	Use:   "mqtt",
	Short: "Short command",
	Long:  "Long description.",
	Run: func(*cobra.Command, []string) {
		core.Run()
	},
}
