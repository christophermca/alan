package converse

import (
	"bufio"
	"fmt"
	"os"
)

var user User

func handleResponse() {
	var context string

	switch sentence.Subject {
	case "actor":
		context = "you're"
	case "us":
		context = "we're"
	case "them":
		context = "they're"
	}

	res := fmt.Sprintf("im glad %s %q\n", context, user.status)
	fmt.Printf("<Alan>: %s\n", res)
}

func request(question string) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("<Alan>: " + question)

	if scanner.Scan() {
		s := scanner.Text()
		/*
			* TODO determine addtional data from sentenceStructure
				- verb
				- object [optional]
			*
		*/
		SentenceStructure(s)
		handleResponse()

	}
}

// Public API
func Greeting(name string) {
	fmt.Printf("Hello my name is %s.\n\n", name)
}

func AskQ(question string) {
	request(question)
}

//func Respond(answer string) {}
