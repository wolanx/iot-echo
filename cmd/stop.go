package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Short command",
	Long:  "Long description.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stop called")
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
