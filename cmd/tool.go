package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zx5435/iot-echo/config"
)

var toolCmd = &cobra.Command{
	Use:   "tool",
	Short: "Small tool",
	Long:  "Long description.",
}

var dataCmd = &cobra.Command{
	Use:   "data",
	Short: "Fetch data",
	Long:  "Long description.",
	Run:   DoToolData,
}

func init() {
	rootCmd.AddCommand(toolCmd)

	toolCmd.AddCommand(dataCmd)
}

func DoToolData(cmd *cobra.Command, args []string) {
	config.ParamsIns.LoadData()
}
