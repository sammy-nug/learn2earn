package main

import (
	"fmt"
	"os"
	"strings"
)

var colors = map[string]string{
	"black":   "\033[30m",
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"white":   "\033[37m",
	"reset":   "\033[0m",
}

// ✅ Proper function (OUTSIDE main)
func GenerateASCII(input string, fontPath string, color string) (string, error) {

	if input == "" {
		return "", nil
	}

	var builder strings.Builder

	data, err := os.ReadFile(fontPath)
	if err != nil {
		return "", err
	}

	// Normalize line endings
	normalized := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(normalized, "\n")

	input = strings.ReplaceAll(input, "\\n", "\n")

	colorCode, hasColor := colors[color]

	for _, segment := range strings.Split(input, "\n") {

		if segment == "" {
			builder.WriteString("\n")
			continue
		}

		for row := 0; row < 8; row++ {

			for _, char := range segment {

				index := int(char) - 32
				if index < 0 || index > 94 {
					continue
				}

				lineIndex := index*9 + 1 + row
				if lineIndex >= len(lines) {
					return "", fmt.Errorf("font file malformed")
				}

				line := lines[lineIndex]

				if hasColor {
					builder.WriteString(colorCode + line + colors["reset"])
				} else {
					builder.WriteString(line)
				}
			}

			builder.WriteString("\n")
		}
	}

	return builder.String(), nil
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . [--color=<color>] \"text\"")
		return
	}

	color := ""
	input := ""

	// Parse arguments
	if strings.HasPrefix(os.Args[1], "--color=") {
		color = strings.TrimPrefix(os.Args[1], "--color=")

		if len(os.Args) < 3 {
			fmt.Println("Error: missing text input")
			return
		}

		input = os.Args[2]
	} else {
		input = os.Args[1]
	}

	// ✅ Call the function (not define it!)
	result, err := GenerateASCII(input, "thinkertoy.txt", color)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Print(result)
}