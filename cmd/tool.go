package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zx5435/iot-echo/config"
	"github.com/zx5435/iot-echo/core"
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
		config.GetParams().LoadData()
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
