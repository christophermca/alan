package converse

func initialize_userMap() map[string]string {
	actor := make(map[string]string)
	actor["firstName"] = ""
	actor["lastName"] = ""
	actor["status"] = ""

	return actor
}
