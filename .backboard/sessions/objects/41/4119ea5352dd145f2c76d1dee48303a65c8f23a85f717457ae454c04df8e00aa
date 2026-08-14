package main

import (
	"fmt"
	"os"
	"strings"
)
//	Check to make sure we have exactly two arguments
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . \"text\"")
		return
	}

	input := os.Args[1]

	fontData, err := os.ReadFile("thinkertoy.txt")
	if err != nil {
		fmt.Println("Error: could not read font file 'standard.txt'")
		return
	}

	// Normalize line endings to handle Windows-formatted files
	normalized := strings.ReplaceAll(string(fontData), "\r\n", "\n")
	lines := strings.Split(normalized, "\n")

	// Handle escaped newline sequences in input
	input = strings.ReplaceAll(input, "\\n", "\n")

	// Each character in the font file occupies 9 lines:
	// 1 blank separator + 8 rows of art
	// Characters start at ASCII 32 (space), so index = char - 32
	// Line for a given character row = index*9 + 1 + row
	for _, segment := range strings.Split(input, "\n") {
		for row := 0; row < 8; row++ {
			for _, char := range segment {
				index := int(char) - 32
				if index < 0 || index > 94 {
					// Skip non-printable or out-of-range characters
					continue
				}
				lineIndex := index*9 + 1 + row
				if lineIndex >= len(lines) {
					fmt.Println("Error: font file appears to be incomplete or malformed")
					return
				}
				fmt.Print(lines[lineIndex])
			}
			fmt.Println()
		}
	}
}
