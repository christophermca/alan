package converse

import (
	"bufio"
	"fmt"
	"os"
)

var user User

func handleResponse() {
	// var context string
	// switch sentence.Subject {
	// case "actor":
	// 	context = "you're"
	// case "us":
	// 	context = "we're"
	// case "them":
	// 	context = "they're"
	// }
	// res := fmt.Sprintf("im glad %s %q\n", context, user.status)

	res := fmt.Sprintf("Hello %s %s", user.firstName, user.lastName)
	fmt.Printf("\n<Alan>: %s\n", res)
}

func request(question string) []string {
	scanner := bufio.NewScanner(os.Stdin)
	var words []string
	if scanner.Scan() {
		words = SentenceStructure(scanner.Text())
	}
	return words
}

// Public API
func Greeting(name string) {
	fmt.Printf("Hello my name is %s.\n\n", name)
}

func AskQuestion(question string, key string) {
	words := request(question)
	if len(words) == 1 {
		user = User{"chris"}
	}
	// else if len(words) > 1 {
	// 	user.lastName = words[1]
	// }

	handleResponse()
}
