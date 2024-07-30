package reverse

import (
	"fmt"
)

// The AsciiPrint function takes two parameters: a string s and
// a font file path font. It returns a slice of strings, which
// represents the ASCII art lines.
func AsciiPrintTerm(s string, font string) []string {
	// If the input string is not "Reboot", the function proceeds to retrieve the
	// font data by calling the GetFont function, passing the font file path.
	// If there is an error during font retrieval, the error is printed,
	// and the function returns nil.
	fontArray, err := GetFontHard(font)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return nil
	}
	// The function initializes an array called charArray using the
	// initializeLines function, which determines the number of lines needed
	// based on the validity of the characters in the input string s.
	charArray := initializeLines(s)
	// Loop through each character in the string
	for i := 0; i < len(s); i++ {
		// Get the valid character from the font.
		// checks if the character is a newline character ('\n'). If it is not
		// a newline character and falls within the valid ASCII range (from 32 to 126),
		// it retrieves the corresponding lines for that character from the font data.
		if s[i] != '\n' && s[i] >= 32 && s[i] <= 126 {
			for linePos, line := range GetCharacter(rune(s[i]), fontArray) {
				charArray[linePos+len(charArray)-8] += line
			}
		} else if s[i] == '\n' {
			// If there is no character after the newline, add 1 line,
			// otherwise add 8.
			if i == 0 || i == len(s)-1 || s[i+1] == '\n' {
				charArray = append(charArray, make([]string, 1)...)
			} else {
				charArray = append(charArray, make([]string, 8)...)
			}
		} else {
			// If the character is neither a valid printable ASCII character
			// nor a newline character, the function prints an error message and
			// returns nil.
			fmt.Println("ERROR: Invalid character")
			return nil
		}
	}
	for _, line := range charArray {
		fmt.Println(line)
	}
	return charArray
}
