package converse

import (
	"bufio"
	"fmt"
	"os"
)

var actor map[string]string

/*
 * TODO build
 * - lookup datafile if actorObject exists
 * - TODO pattern for saved file names
 *  - do this a cleaner way.
 **/

func handleResponse() {
	fmt.Printf("\n<Alan>: %s\n", actor)
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
	actor = make(map[string]string)
}

// Public API
func AskQuestion(question string, keys []string) map[string]string {
	fmt.Println("<Alan>: " + question)
	words := request(question)

	if keys == nil {
		return actor
	}

	switch keys[0] {
	case "firstName":
		actor[keys[0]] = words[0]
		if keys[1] == "lastname" {
			actor[keys[1]] = words[1]
		}
	}
	return actor
}
