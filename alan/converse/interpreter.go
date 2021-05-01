package converse

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

var sentence Sentence

func determineVerb(response string) {
	list := [2]string{`\w+e d`, `\w+ing`} //{past, present}

	var c string
	for i := 0; i < len(list); i++ {
		search := regexp.MustCompile(list[i])
		c = search.FindString(response)

		fmt.Println(search, c)
	}

}

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
		}
	}

	if len(sentence.Subject) == 0 {
		sentence.Subject = "actor"
	}

	log.Printf(":::__ %s __:::", sentence.Subject)
}

func SentenceStructure(response string) []string {
	/*
		* TODO determine addtional data from sentenceStructure
			- verb
			- object [optional]
		*
	*/
	words := strings.Fields(response)
	go determineSubject(response)
	determineVerb(response)

	return words
}
