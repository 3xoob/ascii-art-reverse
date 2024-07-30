package reverse

func GetCharacter(c rune, fontArray [95][8]string) []string {
	// Characters start from 32 in the ASCII table,
	// so we need to subtract 32 to get the correct index.
	char := int(c) - 32
	var lines []string

	// Loop through each line in the character and add it to the lines array.
	for i := 0; i <= 7; i++ {
		lines = append(lines, fontArray[char][i])
	}
	return lines
}
