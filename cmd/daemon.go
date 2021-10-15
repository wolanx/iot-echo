package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/zx5435/iot-echo/config"
)

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
	Run:   DoRestart(),
}

func init() {
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(stopCmd)
	rootCmd.AddCommand(restartCmd)
}

func DoStart(cmd *cobra.Command, args []string) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	appName := strings.TrimLeft(os.Args[0], "./")
	bin := fmt.Sprintf("%s/%s", dir, appName)
	bin = "iot-echo"

	if _, err := os.Stat(config.Dir + "/.lock"); err == nil || os.IsExist(err) {
		log.Warn(".lock is exist")
		return
	}

	command := exec.Command(bin, "run")
	err = command.Start()
	if err != nil {
		log.Error(err.Error())
		return
	}

	pid := command.Process.Pid
	log.Infof("[PID] %d running...\n", pid)
	err = ioutil.WriteFile(config.Dir+"/.lock", []byte(fmt.Sprintf("%d", pid)), 0666)
	if err != nil {
		panic(err)
	}

	fmt.Println("Start is successful")
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

func DoRestart() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		DoStop(cmd, args)
		DoStart(cmd, args)
		fmt.Println("Restart is successful")
	}
}
