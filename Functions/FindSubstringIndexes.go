package reverse

import "strings"

func FindSubstringIndexes(word, substring string) []int {
	Ind := []int{}
	startIndex := 0

	for {
		index := strings.Index(word[startIndex:], substring)
		if index == -1 {
			break
		}
		for i := 0; i < len(substring); i++ {
			Ind = append(Ind, startIndex+index+i)
		}
		startIndex += index + len(substring)
	}

	return Ind
}

