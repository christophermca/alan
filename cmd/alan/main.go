package alan

import (
	"github.com/christophermca/alan-go/pkg/converse"
	userService "github.com/christophermca/alan-go/pkg/users"
	"log"
)

func Run() {
	var keys []string
	//To request name
	keys = []string{"firstName", "lastName"}
	user := converse.AskQuestion("what is your name?", keys)
	userDTO := make(chan []byte)
	go userService.GetUser(userDTO, user["firstName"], user["lastName"])
	v := <-userDTO
	log.Print(string(v))

	// To requesting status
	keys = []string{"status"}
	converse.AskQuestion("How are you?", keys)

	log.Println("done")
}
