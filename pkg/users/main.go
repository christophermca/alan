package users

import (
	"fmt"
	"github.com/christophermca/alan-go/pkg/util"
)

/** TODO CreateUser
*  lookup from file {firstname}.{lastname}.{N}.json
 */

func init() {
	// User Struct
	type User struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Status    string `json:"status"`
	}
}

// GetUser return user data from file
func GetUser(firstName string, lastName string) {
	// get jsonfile or create it
	var name string
	if len(lastName) > 0 {
		name = fmt.Sprintf("%s.%s", firstName, lastName)
	} else {
		name = firstName
	}
	rawInput := util.GetFile(name)
	fmt.Printf("rawinput: %s\n", rawInput)
}
