package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zx5435/iot-echo/config"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Short command",
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
	fmt.Println(bin) // /www/iot-echo/iot-echo

	command := exec.Command(bin, "run")
	err = command.Start()
	if err != nil {
		panic(err)
	}

	pid := command.Process.Pid
	fmt.Printf("%s start, [PID] %d running...\n", bin, pid)
	err = ioutil.WriteFile(config.Dir+"/.lock", []byte(fmt.Sprintf("%d", pid)), 0666)
	if err != nil {
		panic(err)
	}
}
