package reverse

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ColorLine(line, color string) string {
	colorCode := ColorCodeGetter(color)
	return fmt.Sprintf("%s%s%s", colorCode, line, "\033[0m")
}

func ColorSubstring(line string, substring string, colorCode string) string {
	coloredLine := strings.ReplaceAll(line, substring, fmt.Sprintf("%s%s%s", colorCode, substring, "\033[0m"))
	return coloredLine
}

func HexToRGB(colorCode string) (int, int, int) {
	var r, g, b int
	fmt.Sscanf(colorCode, "#%2x%2x%2x", &r, &g, &b)
	return r, g, b
}

func extractRGBValues(color string) []int {
	var rgb []int
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(color, -1)
	for _, match := range matches {
		value, _ := strconv.Atoi(match)
		rgb = append(rgb, value)
	}
	return rgb
}

func HSLToRGB(color string) []int {
	var h, s, l float64
	fmt.Sscanf(color, "hsl(%f, %f%%, %f%%)", &h, &s, &l)
	h /= 360.0
	s /= 100.0
	l /= 100.0

	if s == 0 {
		value := uint8(l * 255)
		return []int{int(value), int(value), int(value)}
	}

	var r, g, b float64
	var q float64

	if l < 0.5 {
		q = l * (1 + s)
	} else {
		q = l + s - (l * s)
	}

	p := 2*l - q
	r = hueToRGB(p, q, h+1.0/3.0)
	g = hueToRGB(p, q, h)
	b = hueToRGB(p, q, h-1.0/3.0)

	red := uint8(r * 255)
	green := uint8(g * 255)
	blue := uint8(b * 255)

	return []int{int(red), int(green), int(blue)}
}

func hueToRGB(p, q, t float64) float64 {
	if t < 0 {
		t += 1
	} else if t > 1 {
		t -= 1
	}

	if t < 1.0/6.0 {
		return p + (q-p)*6*t
	} else if t < 1.0/2.0 {
		return q
	} else if t < 2.0/3.0 {
		return p + (q-p)*6*(2.0/3.0-t)
	} else {
		return p
	}
}
