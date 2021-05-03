package util

import (
	"fmt"
	"io/ioutil"
	"os"
)

var usrHome, _ = os.UserHomeDir()

// Exists reports whether the named file or directory exists.
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// GetFile reads file from disk returns []byte
func GetFile(filename string) []byte {
	filepath := fmt.Sprintf("%s/%s/%s.json", usrHome, "alan-data", filename)
	data, _ := ioutil.ReadFile(filepath)
	return data

}
