package converse

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// Create user file firstname.lastname.N.json
func CreateUser(name string) {}

func GetUser(firstName string, lastName string) {
	filename, _ := filepath.Abs("alan/converse/users")
	filename += "/" + firstName

	if len(lastName) > 0 {
		filename += filename + "." + lastName
	}

	filename += ".json"

	//TODO open channel for quick file access
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err.Error())
	}

	json.Unmarshal(content, &user)

}
