package alan

import (
	"alan/converse"
	"alan/users"
	"fmt"
	"time"
)

func main() {
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
