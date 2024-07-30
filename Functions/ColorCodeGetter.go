package reverse

import (
	"fmt"
	"regexp"
	"strings"
)

func ColorCodeGetter(C string) string {
	if strings.HasPrefix(C, "#") {
		//Hex
		red, green, blue := HexToRGB(C)
		return fmt.Sprintf("\033[38;2;%d;%d;%dm", red, green, blue)

	} else if match, _ := regexp.MatchString(`^rgb\(\d{1,3},\s*\d{1,3},\s*\d{1,3}\)$`, C); match {
		//RGB
		rgb := extractRGBValues(C)
		return fmt.Sprintf("\033[38;2;%d;%d;%dm", rgb[0], rgb[1], rgb[2])

	} else if match, _ := regexp.MatchString(`^hsl\(\d{1,3},\s*\d{1,3}%,\s*\d{1,3}%\)$`, C); match {
		//HSL
		rgb := HSLToRGB(C)
		return fmt.Sprintf("\033[38;2;%d;%d;%dm", rgb[0], rgb[1], rgb[2])
	}
	switch strings.ToLower(C) {
	// ANSI
	case "black":
		return "\033[30m"
	case "red":
		return "\033[31m"
	case "green":
		return "\033[32m"
	case "yellow":
		return "\033[33m"
	case "blue":
		return "\033[34m"
	case "magenta":
		return "\033[35m"
	case "cyan":
		return "\033[36m"
	case "white":
		return "\033[37m"
	case "pink":
		return "\033[95m"
	case "orange":
		return "\033[38;5;208m"
	case "grey":
		return "\033[90m"
	case "purple":
		return "\033[38;5;165m"
	case "lightblue":
		return "\033[94m"
	case "brown":
		return "\033[38;5;130m"
	case "teal":
		return "\033[38;5;37m"
	case "lavender":
		return "\033[38;5;183m"
	case "olive":
		return "\033[38;5;58m"
	default:
		return ""
	}
}
