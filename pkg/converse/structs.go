package converse

// User Struct
type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Status    string `json:"status"`
}

// Alan has (name)
type Alan struct {
	name string
}

// CompileSubject ...
type CompileSubject struct {
	regexp  string
	subject string
}

// Sentence ...
type Sentence struct {
	Subject string `json:"subject"`
	Verb    string `json:"verb"`
	Object  string `json:"object"`
}
