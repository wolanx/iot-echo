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
	Short: "Stop daemon process",
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
		log.Warning("maybe need manual kill it: pkill iot-echo")
		return
	}

	pid := string(pidByte)
	log.Info("[PID] " + pid + " killing...")

	command := exec.Command("kill", pid)
	output, err := command.CombinedOutput()
	if err != nil {
		log.Error(err.Error())
		log.Warning(string(output))
		log.Warning("maybe need manual kill it: pkill iot-echo")
	}
	if err := os.Remove(config.Dir + "/.lock"); err != nil {
		log.Error(err.Error())
		return
	}

	fmt.Println("Stop is successful")
}
