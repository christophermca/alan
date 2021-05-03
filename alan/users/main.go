package users

import (
	"encoding/json"
	"fmt"
	"github.com/christophermca/alan-go/alan/util"
)

/** TODO CreateUser
*  lookup from file {firstname}.{lastname}.{N}.json
 */

// user Struct
type user struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Status    string `json:"status"`
}

// GetUser return user data from file
func GetUser(firstName string, lastName string) {
	var data user
	rawInput := util.GetFile(firstName)
	err := json.Unmarshal(rawInput, &data)
	if err != nil {
		fmt.Println(err)
		return
	}
}
