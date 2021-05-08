package util

import (
	_ "encoding/json" // test
	"fmt"
	"io/ioutil"
	"log"
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

func createFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(fmt.Sprintf("created file %s", filename))
	file.Close()
}

// GetFile reads userdata from file, if file missing create it
func GetFile(filename string) []byte {
	filepath := fmt.Sprintf("%s/%s/%s.json", usrHome, "alan-data", filename)
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		createFile(filepath)
	}
	return data

}
