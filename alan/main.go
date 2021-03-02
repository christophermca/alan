package alan

import (
	"alp/alan/converse"
)

type Self struct {
	name string
}

func Run() {
	al := Self{name: "Alan"}
	converse.Greeting(al.name)
	// To requesting status
	// converse.AskQuestion("How are you?")

	//To request name
	converse.AskQuestion("what is your name?", "firstName")
}

func Shutdown() {
	// Shutdown contexts, listeners, and such
}
