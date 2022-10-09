package util

import (
	"fmt"
	"os/exec"
)

func RunShell(s string) string {
	cmd := exec.Command("sh", "-c", s)
	//cmd := exec.Command("lua", "C:\\Users\\106006\\Desktop\\www\\gimc\\src\\backend\\temp\\mock-sh\\hvac.lua", "a2")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(output)
}
