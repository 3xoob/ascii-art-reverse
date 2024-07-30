package reverse

func GetChar(word string, Font string) []string {
	lines := make([]string, 8)
	for i := 0; i < 8; i++ {
		for _, ch := range word {
			charLines := GetFont(1+int(ch-32)*9+i, Font)
			lines[i] += charLines
		}
	}
	return lines
}
