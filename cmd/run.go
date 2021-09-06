package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zx5435/iot-echo/web"
)

var testCmd = &cobra.Command{
	Use:   "run",
	Short: "Daemon server",
	Long:  "Long description.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Daemon called")

		web.DefaultWeb()
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
