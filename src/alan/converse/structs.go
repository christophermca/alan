package converse

type Alan struct {
	name string
}

type CompileSubject struct {
	regexp  string
	subject string
}

type User struct {
	firstname, lastname, status string
}

type Sentence struct {
	Subject string `json:"subject"`
	Verb    string `json:"verb"`
	Object  string `json:"object"`
}
