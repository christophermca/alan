package converse

type User struct {
	firstName string
	lastName  string
	status    string
	age       int
}

type Sentence struct {
	Subject string `json:"subject"`
	Verb    string `json:"verb"`
	Object  string `json:"object"`
}
