package alan

import (
	"fmt"
	converse "github.com/christophermca/alan-go/alan/converse"
	userService "github.com/christophermca/alan-go/alan/users"
	"time"
)

func Run() {
	var keys []string
	//To request name
	keys = []string{"firstName", "lastName"}
	user := converse.AskQuestion("what is your name?", keys)
	go userService.GetUser(user["firstName"], user["lastName"])

	// To requesting status
	keys = []string{"status"}
	converse.AskQuestion("How are you?", keys)

	//keep alive hack
	time.Sleep(time.Second)
	fmt.Println("done")
}
