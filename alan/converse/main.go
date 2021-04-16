package converse

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

var user map[string]string

func fileExist(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func determineUser(userData map[string]string) {
	/**
	* TODO build
	* - lookup datafile if userObject exists
	* - TODO pattern for saved file names
	*  - do this a cleaner way.
	 */
	filename, _ := filepath.Abs("alan/converse/users")
	filename += "/" + userData["firstName"]

	if len(userData["lastname"]) > 0 {
		filename += filename + "." + userData["lastname"]
	}

	filename += ".json"

	if fileExist(filename) {
		//TODO load data set to channel
	} else {
		fmt.Printf("%s was not found\n", filename)
	}

}

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
	// res := fmt.Sprintf("Hello %s %s", user["firstName"], user["lastName"])

	// res = fmt.Sprintf("im glad %s %s \n", context, user["status"])
	//fmt.Printf("\n<Alan>: %s\n", res)
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
		determineUser(user)
	}
}
