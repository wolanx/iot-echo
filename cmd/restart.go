package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Stop and start",
	Long:  "Long description.",
	Run: func(cmd *cobra.Command, args []string) {
		DoStop(cmd, args)
		DoStart(cmd, args)
		fmt.Println("Restart is successful")
	},
}

func init() {
	rootCmd.AddCommand(restartCmd)
}
