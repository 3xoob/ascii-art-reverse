package reverse

import (
	"fmt"
	"os"
)

func AsciiPrintFile(s string, font string, fileOutput ...string) []string {

	fontArray, err := GetFontHard(font)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	charArray := initializeLines(s)
	// Loop through each character in the string
	for i := 0; i < len(s); i++ {

		if s[i] != '\n' && s[i] >= 32 && s[i] <= 126 {
			for linePos, line := range GetCharacter(rune(s[i]), fontArray) {
				charArray[linePos+len(charArray)-8] += line
			}
		} else if s[i] == '\n' {

			if i == 0 || i == len(s)-1 || s[i+1] == '\n' {
				charArray = append(charArray, make([]string, 1)...)
			} else {
				charArray = append(charArray, make([]string, 8)...)
			}
		} else {
			fmt.Println("Error: Invalid character")
			return nil
		}
	}
	if len(fileOutput) == 1 {
		file, err := os.Create(fileOutput[0][9:])
		if err != nil {
			fmt.Println(err)
			return nil
		}
		for _, line := range charArray {
			_, err := file.WriteString(line + "\n")
			if err != nil {
				fmt.Println(err)
				file.Close()
				return nil
			}
		}
		file.WriteString("\n")

		err = file.Close()
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}
	return charArray
}
