package main

import (
	"fmt"
	"os"
	reverse "reverse/Functions"
	"strings"
)

func main() {
	colorFlag := false
	filename := "standard"
	colorNeeded := ""
	ToBeColored := ""
	if len(os.Args) == 1 {
		fmt.Println("ERROR:")
		fmt.Println("Usage: go run . [STRING]")
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("EX: go run . <text>")
		fmt.Println("EX: go run . <text> <banner>")
		fmt.Println("EX: go run . --reverse=<fileName>")
		fmt.Println("EX: go run . --output=<fileName.txt> <text>")
		fmt.Println("EX: go run . --output=<fileName.txt> <text> <banner>")
		fmt.Println("EX: go run . --color=<color> <text>")
		fmt.Println("EX: go run . --color=<color> <substring to be colored> <text>")
		fmt.Println("EX: go run . --color=<color> <substring to be colored> <text> <banner>")
		os.Exit(0)
	}
	if string(os.Args[1]) == "Ascii Art Team" {
		reverse.R01()
		os.Exit(0)
	}
	if len(os.Args[1:]) == 4 && strings.HasPrefix(os.Args[1], "--color=") {
		colorNeeded = strings.TrimPrefix(os.Args[1], "--color=")
		colorFlag = true
		//  0			1					2					3				4
		//"go run . --color=<color> <letters to be colored> \"something\"    font type"
		filename = reverse.FontChecker(os.Args[4])
		str := os.Args[3]
		ToBeColored = os.Args[2]
		if !strings.Contains(str, ToBeColored) {
			fmt.Println("ERROR: The provided string does not contain the specified letters.")
			os.Exit(0)
		}
		if ToBeColored == "" {
			ToBeColored = os.Args[3]
		}
		// printing and coloring
		reverse.Printer(str, colorFlag, colorNeeded, filename, ToBeColored)
		os.Exit(0)
	} else if len(os.Args[1:]) == 3 {
		if strings.HasPrefix(os.Args[1], "--output=") {
			reverse.AsciiPrintFile(strings.ReplaceAll(strings.ReplaceAll(os.Args[2], "\\t", "    "), `\n`, "\n"), os.Args[3], os.Args[1])
			os.Exit(0)
		} else if strings.HasPrefix(os.Args[1], "--color=") && (os.Args[3] != "standard" && os.Args[3] != "shadow" && os.Args[3] != "thinkertoy") {
			//  0			1				2						3
			//"go run . --color=<color> <letters to be colored> \"something\""
			colorNeeded = strings.TrimPrefix(os.Args[1], "--color=")
			colorFlag = true
			str := os.Args[3]
			ToBeColored = os.Args[2]
			if !strings.Contains(str, ToBeColored) {
				fmt.Println("ERROR: The provided string does not contain the specified letters.")
				os.Exit(0)
			}
			if ToBeColored == "" {
				ToBeColored = os.Args[3]
			}
			// printing and coloring with standard font
			filename = reverse.FontChecker(filename)
			reverse.Printer(str, colorFlag, colorNeeded, filename, ToBeColored)
			os.Exit(0)
		} else if strings.HasPrefix(os.Args[1], "--color=") && (os.Args[3] == "standard" || os.Args[3] == "shadow" || os.Args[3] == "thinkertoy") {
			//  0			1				2						3
			//"go run . --color=<color>  \"something\" 			font type"
			colorNeeded = strings.TrimPrefix(os.Args[1], "--color=")
			colorFlag = true
			filename = reverse.FontChecker(os.Args[3])
			str := os.Args[2]
			ToBeColored = os.Args[2]
			// printing and full coloring with selected font
			reverse.Printer(str, colorFlag, colorNeeded, filename, ToBeColored)
			os.Exit(0)
		} else {
			fmt.Println("ERROR:")
			fmt.Println("Usage: go run . [STRING]")
			fmt.Println("Usage: go run . [STRING] [BANNER]")
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
			fmt.Println("EX: go run . <text>")
			fmt.Println("EX: go run . <text> <banner>")
			fmt.Println("EX: go run . --reverse=<fileName>")
			fmt.Println("EX: go run . --output=<fileName.txt> <text>")
			fmt.Println("EX: go run . --output=<fileName.txt> <text> <banner>")
			fmt.Println("EX: go run . --color=<color> <text>")
			fmt.Println("EX: go run . --color=<color> <substring to be colored> <text>")
			fmt.Println("EX: go run . --color=<color> <substring to be colored> <text> <banner>")
			os.Exit(0)
		}
	} else if len(os.Args[1:]) == 2 {
		FONT := strings.ToLower(os.Args[2])
		if strings.HasPrefix(os.Args[1], "--output=") {
			reverse.AsciiPrintFile(strings.ReplaceAll(strings.ReplaceAll(os.Args[2], "\\t", "    "), `\n`, "\n"), "standard", os.Args[1])
			os.Exit(0)
		} else if strings.HasPrefix(os.Args[1], "--color=") {
			//  0			1					2
			//"go run .  --color=<color>  \"something\"
			colorNeeded = strings.TrimPrefix(os.Args[1], "--color=")
			colorFlag = true
			str := os.Args[2]
			ToBeColored = os.Args[2]
			// printing and coloring with standard font
			filename = reverse.FontChecker(filename)
			reverse.Printer(str, colorFlag, colorNeeded, filename, ToBeColored)
			os.Exit(0)

		} else if FONT == "standard" || FONT == "shadow" || FONT == "thinkertoy" {
			reverse.AsciiPrintTerm(strings.ReplaceAll(strings.ReplaceAll(os.Args[1], "\\t", "    "), `\n`, "\n"), os.Args[2])
			os.Exit(0)
		} else {
			fmt.Println("ERROR:")
			fmt.Println("Usage: go run . [STRING]")
			fmt.Println("Usage: go run . [STRING] [BANNER]")
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
			fmt.Println("EX: go run . <text>")
			fmt.Println("EX: go run . <text> <banner>")
			fmt.Println("EX: go run . --reverse=<fileName>")
			fmt.Println("EX: go run . --output=<fileName.txt> <text>")
			fmt.Println("EX: go run . --output=<fileName.txt> <text> <banner>")
			fmt.Println("EX: go run . --color=<color> <text>")
			fmt.Println("EX: go run . --color=<color> <substring to be colored> <text>")
			fmt.Println("EX: go run . --color=<color> <substring to be colored> <text> <banner>")
			os.Exit(0)
		}
	} else if len(os.Args[1:]) == 1 {
		if strings.HasPrefix(os.Args[1], "--reverse=") {
			reverse.ReverseFileAscii(os.Args[1])
			os.Exit(0)
		} else {
			reverse.AsciiPrintTerm(strings.ReplaceAll(strings.ReplaceAll(os.Args[1], "\\t", "    "), `\n`, "\n"), "standard")
			os.Exit(0)
		}
	} else {
		fmt.Println("ERROR:")
		fmt.Println("Usage: go run . [STRING]")
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("EX: go run . <text>")
		fmt.Println("EX: go run . <text> <banner>")
		fmt.Println("EX: go run . --reverse=<fileName>")
		fmt.Println("EX: go run . --output=<fileName.txt> <text>")
		fmt.Println("EX: go run . --output=<fileName.txt> <text> <banner>")
		fmt.Println("EX: go run . --color=<color> <text>")
		fmt.Println("EX: go run . --color=<color> <substring to be colored> <text>")
		fmt.Println("EX: go run . --color=<color> <substring to be colored> <text> <banner>")
		os.Exit(0)
	}

}
