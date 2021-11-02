package cmd

import (
	"fmt"
	"os/exec"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Service management",
	Long:  "Long description.",
	Run: func(cmd *cobra.Command, args []string) {
		err := sign.Run()
		if err != nil {
			log.Fatal(err)
		}
	},
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Service management",
	Long:  "Long description.",
	Run: func(cmd *cobra.Command, args []string) {
		msg, err := sign.Status()
		err_degrade(err, "status")
		fmt.Println(msg)
	},
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Service management",
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
	Short: "Service management",
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
	Short: "Service management",
	Long:  "Long description.",
	Run:   DoStart,
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Service management",
	Long:  "Long description.",
	Run:   DoStop,
}

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Service management",
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
	err_degrade(err, "start")
}

func DoStop(*cobra.Command, []string) {
	err := sign.Stop()
	err_degrade(err, "stop")
}

func err_degrade(err error, cmd string) {
	if err == nil {
		return
	}
	if err.Error() == `"service" failed: exec: "service": executable file not found in $PATH` {
		command := exec.Command("/etc/init.d/iot-echo", cmd)
		output, err := command.CombinedOutput()
		if err != nil {
			log.Error(err)
		}
		log.Info(string(output))
	} else {
		log.Fatal(err)
	}
}
