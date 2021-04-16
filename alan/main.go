package alan

import (
	"alp/alan/converse"
	"fmt"
	"time"
)

func Run() {
	var keys []string
	//To request name
	keys = []string{"firstName", "lastName"}
	converse.AskQuestion("what is your name?", keys)

	// To requesting status
	keys = []string{"status"}
	converse.AskQuestion("How are you?", keys)

	//keep alive hack
	time.Sleep(time.Second)
	fmt.Println("done")
}
