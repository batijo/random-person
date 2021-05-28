package handlers

import "strconv"

func getGender(query string) int {
	if query == "" {
		return -1
	}
	gender, err := strconv.Atoi(query)
	if err != nil {
		switch query {
		case "male":
			return 0
		case "female":
			return 1
		default:
			return -1
		}
	}
	return gender
}
