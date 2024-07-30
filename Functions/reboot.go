package reverse

import (
	"fmt"
	"os"
)

func R01() {
	content, err := os.ReadFile("Fonts/Reboot.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(string(content))
}
