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

func init() {
	rootCmd.AddCommand(startCmd)
}

func DoStart(cmd *cobra.Command, args []string) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	appName := strings.TrimLeft(os.Args[0], "./")
	bin := fmt.Sprintf("%s/%s", dir, appName)

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
