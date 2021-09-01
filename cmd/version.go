package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the iot-echo version information",
	Long:  "Long description.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("iot-echo version 0.0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
