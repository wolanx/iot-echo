package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Empty",
	Long:  "Long description.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test called")
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
