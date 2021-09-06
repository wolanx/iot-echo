package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/zx5435/iot-echo/config"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Short command",
	Long:  "Long description.",
	Run:   DoStop,
}

func init() {
	rootCmd.AddCommand(stopCmd)
}

func DoStop(cmd *cobra.Command, args []string) {
	pidByte, err := ioutil.ReadFile(fmt.Sprintf("%s/.lock", config.Dir))
	if err != nil {
		log.Error(err.Error())
	} else {
		pid := string(pidByte)
		fmt.Println("pid", pid)

		command := exec.Command("kill", pid)
		err = command.Start()
		if err != nil {
			fmt.Print(err.Error())
		}
		fmt.Println(`maybe need manual kill it.
	ps -ef | grep iot-echo
	kill $pid`)
		err := os.Remove(config.Dir + "/.lock")
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	fmt.Println("Stop is successful")
}
