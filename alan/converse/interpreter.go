package converse

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

var sentence Sentence

func determineSubject(response string) {
	var list [5]CompileSubject
	list[0] = CompileSubject{`[mM]y`, "actor"}
	list[1] = CompileSubject{`[iI](\'?m)?\s`, "actor"}
	list[2] = CompileSubject{`[yY]ou\'?r?e?\s`, "self"}
	list[3] = CompileSubject{`[wW]e\'?r?e?\s`, "us"}
	list[4] = CompileSubject{`[tT]hey\'?r?e?\s`, "them"}

	for i := 0; i < len(list); i++ {
		search, _ := regexp.Compile(list[i].regexp)

		if len(search.FindString(response)) > 0 {
			sentence.Subject = list[i].subject
			log.Printf("subject: %#v\n", sentence.Subject)
		}
	}

	if len(sentence.Subject) == 0 {
		sentence.Subject = "actor"

	}
}

func SentenceStructure(response string) {
	var words []string
	determineSubject(response)

	// DEBUGGING below
	words = strings.Fields(response)
	fmt.Print(":::__" + sentence.Subject + "__:::")

	if len(sentence.Subject) > 0 {
		if len(words) > 1 {
			user.status = words[1]
		} else {
			user.status = words[0]
		}
	}
}
