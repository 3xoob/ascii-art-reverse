package reverse

import (
	"bufio"
	"fmt"
	"os"
)

// ReadLineByLine opens a file specified by filePath and reads its content line by line.
func ReadLineByLine(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(0)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(0)
	}

	return lines
}

// TrimPrefix removes the first 'n' characters from each string in the 'words' slice.
func TrimPrefix(n int, words []string) []string {
	for i, word := range words[0 : len(words)-1] {
		words[i] = word[n:]
	}
	return words
}

// HasPrefix checks if the prefix of each string in 'words' matches the corresponding string in 'prefixes'.
func HasPrefix(prefixes, words []string) bool {
	found := true
	if len(prefixes[0]) > len(words[0]) {
		return false
	}
	for i, v := range words[0 : len(words)-1] {
		if prefixes[i] != v[:len(prefixes[i])] {
			found = false
		}
	}
	return found
}

// SegmentTemplate splits the 'template' slice into segments of 8 strings each, starting from the second string.
func SegmentAscii(AsciiLetters []string) [][]string {
	var segments [][]string
	for i := 0; i < len(AsciiLetters)-1; i = i + 9 {
		temp := AsciiLetters[i+1 : i+9]
		segments = append(segments, temp)
	}
	return segments
}
