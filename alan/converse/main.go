package converse

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var sentence Sentence
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

func determineSubject(foundSubject, subject string) {
	if len(foundSubject) > 0 {
		sentence.Subject = subject
		log.Printf("subject: %#v\n", subject)
	}
}

// to determine meaning
func sentenceStructure(response string) {

	search, _ := regexp.Compile(`[mM]y`)
	determineSubject(search.FindString(response), "actor")

	search, _ = regexp.Compile(`[iI](\'?m)?\s`)
	determineSubject(search.FindString(response), "actor")

	search, _ = regexp.Compile(`[yY]ou\'?r?e?\s`)
	determineSubject(search.FindString(response), "self")

	search, _ = regexp.Compile(`[wW]e\'?r?e?\s`)
	determineSubject(search.FindString(response), "us") // actor & self

	search, _ = regexp.Compile(`[tT]hey\'?r?e?\s`)
	determineSubject(search.FindString(response), "them") // NOT actor || self

	words := strings.Fields(response)
	if len(sentence.Subject) > 0 {
		if len(words) > 1 {
			user.status = words[1]
		}
	} else {
		user.status = words[0]
	}

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
		sentenceStructure(s)
		handleResponse()

	}
}

// public
func Greeting(name string) {
	fmt.Printf("Hello my name is %s.\n\n", name)
}

func AskQ(question string) {
	request(question)
}
