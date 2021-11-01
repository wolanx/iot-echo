package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zx5435/iot-echo/debug"
)

var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "Short command",
	Long:  "Long description.",
	Run:   debug.Run,
}

func init() {
	rootCmd.AddCommand(debugCmd)
}
