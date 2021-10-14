package util

import (
	"io/ioutil"
	"os"
)

func FileGetContents(path string) string {
	file, _ := os.Open(path)
	fileText, _ := ioutil.ReadAll(file)

	return string(fileText)
}
