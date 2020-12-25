package main

import (
	"alan/converse"
)

type Self struct {
	name string
}

func main() {
	al := Self{name: "Alan"}
	converse.Greeting(al.name)
}
