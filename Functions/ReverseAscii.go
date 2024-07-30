package reverse

import "fmt"

func ReverseFileAscii(file ...string) {
	AsciiContent := ReadLineByLine("Fonts/standard.txt")
	SegmentedAscii := SegmentAscii(AsciiContent)
	Target := ReadLineByLine(file[0][10:])

	content := ""
	for len(Target[0]) > 0 {
		for i, v := range SegmentedAscii {
			if HasPrefix(v, Target) {
				content = content + string(rune(i+32))
				Target = TrimPrefix(len(v[0]), Target)
			}
		}
	}
	fmt.Println(content)
	return
}
