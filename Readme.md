# Ascii-Art

## Overview

`ascii-art` is a Go program that converts input strings into ASCII art representations using predefined templates. It supports various optional functionalities for file handling, color manipulation, and reversal of the ASCII art process.

## Features

- **Text-to-ASCII Conversion:** Converts input strings into ASCII art using predefined templates.
- **Supports Various Characters:** Handles letters, numbers, spaces, special characters, and newlines (`\n`).

## Optional Features

### Banner Option: `ascii-art-fs`

**Objective:** Allows specifying a banner template through the command line.

### Color Option: `ascii-art-color`

**Objective:** Allows color manipulation of the ASCII art output.
- Colors available:
  - "black"
  - "red"
  - "green"
  - "yellow"
  - "blue"
  - "magenta"
  - "cyan"
  - "white"
  - "pink"
  - "orange"
  - "grey"
  - "purple"
  - "lightblue"
  - "brown"
  - "teal"
  - "lavender"
  - "olive"

### Output Option: `ascii-art-output`

**Objective:** Writes the ASCII art result to a file.

### Reverse Option: `ascii-reverse`

**Objective:** Reverses the ASCII art process by converting a graphic representation back to text.

## Instructions

Run the program with the desired options to generate ASCII art, manipulate colors, handle file outputs, or reverse the ASCII art process.

**Usage:**
- `go run . [STRING]`
- `go run . [STRING] [BANNER]`
- `go run . [OPTION] [STRING] [BANNER]`
- Example:
  - `go run . <text>`
  - `go run . <text> <banner>`
  - `go run . --reverse=<fileName>`
  - `go run . --output=<fileName.txt> <text>`
  - `go run . --output=<fileName.txt> <text> <banner>`
  - `go run . --color=<color> <text>`
  - `go run . --color=<color> <substring to be colored> <text>`
  - `go run . --color=<color> <substring to be colored> <text> <banner>`

**Note: You can run the bash file to test all the examples**
- `bash test.sh`

## Team

You can find our team members by commanding "Ascii Art Team".
  ```sh
  go run . "Ascii Art Team"
  ```
