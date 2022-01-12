package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/wolanx/iot-echo/pkg/util"
	"golang.org/x/crypto/ssh"
)

func main() {
	var firstName string
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName)
	fmt.Println(firstName)

	//auths := []ssh.AuthMethod{ssh.Password("123456")}
	auths := []ssh.AuthMethod{ssh.PublicKeys(loadByPem())}
	client, err := ssh.Dial("tcp", "47.100.105.217:22", &ssh.ClientConfig{
		User:            "root",
		Auth:            auths,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		log.Fatalf("SSH dial error: %s", err.Error())
	}

	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("new session error: %s", err.Error())
	}
	defer session.Close()

	var out bytes.Buffer
	session.Stdout = &out
	session.Stderr = &out
	if err := session.Run("curl -sfL http://ccm-perf.cd81f591dfeeb4a4d977da58456d29adc.cn-shanghai.alicontainer.com/static/remote.txt | sh -"); err != nil {
		panic("Failed to run: " + err.Error())
	}
	fmt.Println(out.String())

	fmt.Println("Enter any key to exit.")
	fmt.Scanln()
}

func loadByPem() ssh.Signer {
	homeDir, _ := os.UserHomeDir()
	pem := util.FileGetContents(homeDir + "/Desktop/key_hogan.pem")
	signer, _ := ssh.ParsePrivateKey([]byte(pem))
	return signer
}
