package reverse

import (
	"strings"
)

func FontChecker(Font string) string {
	Font = strings.ToLower(Font)
	AsciiFonts := []string{"standard", "shadow", "thinkertoy"}
	for _, font := range AsciiFonts {
		if Font == font {
			Font = "Fonts/" + font + ".txt"
			break
		}
	}
	return Font
}
