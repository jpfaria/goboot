package strings

// SliceContains receives a slice of string and checks if contains str
func SliceContains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}

	return false
}
