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
	converse.AskQ("How are you?")
}

func Shutdown() {
	// Shutdown contexts, listeners, and such
}
