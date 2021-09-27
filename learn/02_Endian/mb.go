package main

import (
	"encoding/hex"

	log "github.com/sirupsen/logrus"
)

func doLog1(err error) {
	if err != nil {
		log.Error(err)
	} else {
		log.Info("connect1 ok")
	}
}

func doLog2(err error, res []byte) {
	if err != nil {
		log.Error(err)
	} else {
		log.Info("connect2 ok")
	}
	log.Info(res)
	log.Info(string(res))
	log.Info(hex.EncodeToString(res))
}
