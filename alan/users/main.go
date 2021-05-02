package users

import (
	"fmt"
	"io/ioutil"
)

/** TODO CreateUser
*  lookup from file {firstname}.{lastname}.{N}.json
 */

// GetUser ...
func GetUser(firstName string, lastName string) {
	//TODO get user/data/ location
	filename := "data"

	// get file firstName.lastname
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

}
