package users

import (
	_ "encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

/** TODO CreateUser
*  lookup from file {firstname}.{lastname}.{N}.json
 */

func GetUser(firstName string, lastName string) {
	filename, _ := filepath.Abs("alan/users/data")
	filename += "/" + firstName

	if len(lastName) > 0 {
		filename += filename + "." + lastName
	}

	filename += ".json"

	//TODO open channel for quick file access
	_, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err.Error())
	}

	//json.Unmarshal(content, &user)

}
