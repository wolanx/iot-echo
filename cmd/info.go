package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zx5435/iot-echo/config"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show description",
	Long:  "Long description.",
	Run: func(cmd *cobra.Command, args []string) {
		cfb, _ := json.MarshalIndent(config.GetConfig(), "", "\t")
		fmt.Println(string(cfb))
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
