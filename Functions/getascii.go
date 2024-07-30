package reverse

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 95 is the amount of characters in the font and 8 is the amount of lines in the font.

func GetFontHard(fontToUse string) ([95][8]string, error) {
	var fontArray = [95][8]string{}
	AsciiFonts := []string{"standard", "shadow", "thinkertoy"}
	fontToUse = strings.ToLower(fontToUse)
	for _, font := range AsciiFonts {
		if fontToUse == font {
			fontToUse = "Fonts/" + font + ".txt"
			break
		}
	}
	file, err := os.Open(fontToUse)
	if err != nil {
		fmt.Print("ERROR: ")
		return [95][8]string{}, err
	}

	scanner := bufio.NewScanner(file)
	currentChar := -1 // Font starts with an empty line, so we need to start at -1
	currentLine := 0

	// Loop through each line in the font file.
	for scanner.Scan() {
		// If the line is empty, we are at the start of a new character.
		if scanner.Text() == "" {
			currentChar++
			currentLine = 0
		} else {
			// Map the line to the current character in the font array.
			fontArray[currentChar][currentLine] = scanner.Text()
			currentLine++
		}
	}
	eror := scanner.Err()
	if eror != nil {
		file.Close()
		return [95][8]string{}, err
	}

	err = file.Close()
	return fontArray, err
}
