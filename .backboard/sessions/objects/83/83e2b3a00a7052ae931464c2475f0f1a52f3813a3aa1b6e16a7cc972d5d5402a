package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	// If no arguments: produce NO output at all
	if len(args) == 0 {
		return
	}

	upper := false

	// Check for --upper flag
	if len(args) > 0 && args[0] == "--upper" {
		upper = true
		args = args[1:]
	}

	for _, arg := range args {
		// Convert string to int manually
		n, ok := atoi(arg)
		if !ok || n < 1 || n > 26 {
			fmt.Print(" ")
			continue
		}

		// Get letter
		var letter rune
		if upper {
			letter = rune('A' + n - 1)
		} else {
			letter = rune('a' + n - 1)
		}

		fmt.Printf("%c", letter)
	}

	fmt.Println()
}

// Manual atoi using only runes
func atoi(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}

	n := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0, false
		}

		n = n*10 + int(r-'0')
	}

	return n, true
}
