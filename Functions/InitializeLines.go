package reverse

func initializeLines(s string) []string {
	// Initialize the charArray with the correct amount of lines,
	// 0 if a character is invalid, otherwise 8
	charArray := make([]string, 0)
	for _, c := range s {
		if c >= 32 && c <= 126 {
			charArray = make([]string, 8)
			break
		}
	}
	return charArray
}
