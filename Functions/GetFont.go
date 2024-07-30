package reverse

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetFont(lineNum int, font string) string {
	file, err := os.Open(font)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return ""
	}
	defer file.Close()
	content := bufio.NewReader(file)
	x := ""
	for i := 0; i <= lineNum; i++ {
		x, err = content.ReadString('\n')
		if err != nil {
			fmt.Println("ERROR: ", err)
			break
		}
	}
	return strings.TrimRight(x, "\n")
}
