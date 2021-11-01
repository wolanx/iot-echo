package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start daemon process",
	Long:  "Long description.",
	Run: func(cmd *cobra.Command, args []string) {
		err := sign.Run()
		if err != nil {
			return
		}
	},
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Start daemon process",
	Long:  "Long description.",
	Run: func(cmd *cobra.Command, args []string) {
		msg, err := sign.Status()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(msg)
	},
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Start daemon process",
	Long:  "Long description.",
	Run: func(cmd *cobra.Command, args []string) {
		err := sign.Install()
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Start daemon process",
	Long:  "Long description.",
	Run: func(cmd *cobra.Command, args []string) {
		err := sign.Uninstall()
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start daemon process",
	Long:  "Long description.",
	Run:   DoStart,
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop daemon process",
	Long:  "Long description.",
	Run:   DoStop,
}

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
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(removeCmd)
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(stopCmd)
	rootCmd.AddCommand(restartCmd)
}

func DoStart(*cobra.Command, []string) {
	err := sign.Start()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func DoStop(*cobra.Command, []string) {
	err := sign.Stop()
	if err != nil {
		fmt.Println(err)
		return
	}
}
