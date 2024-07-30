package reverse

import (
	"fmt"
	"os"
	"strings"
)

func Printer(str string, color bool, colorNeeded string, font string, ToBeColored string) {
	if str == "\\n" {
		fmt.Println()
		return
	}
	str = strings.ReplaceAll(strings.ReplaceAll(str, "\\t", "    "), `\n`, "\n")
	StrArray := strings.Split(str, "\\n")
	Char := make([]string, 8)
	if color {
		if ColorCodeGetter(colorNeeded) == "" {
			fmt.Println("ERROR: The selected color '", colorNeeded, "' is not valid. Please choose a valid color.")
			os.Exit(0)
		}
	}
	for _, word := range StrArray {
		// if word == "" {
		// 	fmt.Println()
		// 	continue
		// }
		// Handle newlines in the word
		lines := strings.Split(word, "\n")
		for _, line := range lines {
			if line == "" {
				fmt.Println()
				continue
			}
			start := strings.Index(line, ToBeColored)
			index := start
			arr := FindSubstringIndexes(line, ToBeColored)
			for j, char := range line {
				if (char < 32 || char > 126) && char != 10 {
					fmt.Println("ERROR: The input contains invalid characters.")
					os.Exit(0)
				} else {
					charLines := GetChar(string(char), font)
					if color {
						if strings.Contains(line, ToBeColored) && len(ToBeColored) != 1 && j >= index && index < start+len(ToBeColored) {
							for i, line := range charLines {
								line = ColorLine(line, colorNeeded)
								Char[i] += line
							}
							index++
						} else if INTinArray(arr, j) {
							for i, line := range charLines {
								line = ColorLine(line, colorNeeded)
								Char[i] += line
							}
						} else {
							for i, line := range charLines {
								Char[i] += line
							}
						}
					} else {
						for i, line := range charLines {
							Char[i] += line
						}
					}
				}
			}
			for _, line := range Char {
				if color && ToBeColored == "" {
					fmt.Println(line)
				} else {
					fmt.Println(line)
				}
			}
			Char = make([]string, 8)
		}
	}
}
