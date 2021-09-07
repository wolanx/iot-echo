package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zx5435/iot-echo/web"
)

var testCmd = &cobra.Command{
	Use:   "run",
	Short: "Run daemon process",
	Long:  "Long description.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run called")

		web.DefaultWeb()
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
