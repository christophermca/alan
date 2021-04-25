package converse

import (
	"bufio"
	"fmt"
	"os"
)

var user map[string]string

/*
 * TODO build
 * - lookup datafile if userObject exists
 * - TODO pattern for saved file names
 *  - do this a cleaner way.
 **/

func handleResponse() {
	fmt.Printf("\n<Alan>: %s\n", user)
}

func request(question string) []string {
	scanner := bufio.NewScanner(os.Stdin)
	var words []string
	if scanner.Scan() {
		words = SentenceStructure(scanner.Text())
	}
	return words
}

func (al *Alan) greeting() {
	fmt.Printf("Hello my name is %s.\n\n", al.name)
}

func init() {
	al := &Alan{name: "Alan"}
	al.greeting()
	user = make(map[string]string)
}

// Public API
func AskQuestion(question string, keys []string) {
	fmt.Println("<Alan>: " + question)
	words := request(question)

	if keys == nil {
		return
	}

	switch keys[0] {
	case "status":
		user[keys[0]] = words[0]
	case "firstName":
		user[keys[0]] = words[0]
		if keys[1] == "lastname" {
			user[keys[1]] = words[1]
		}

		// NOTE possibly abstract to a new location.
		go GetUser(user["firstName"], user["lastName"])
	}
}
