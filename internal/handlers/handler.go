package handlers

func Handler(list, path string) {

	if list == "or" {
		HandleOR()
	}

	if list == "exceltodb" {
		PostToDB(path)
	}

}
