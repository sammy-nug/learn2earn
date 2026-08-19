package main

import (
	"os"

	"github.com/01-edu/z01"
)

// isVowel checks if a rune is a vowel (y is not a vowel)
func isVowel(c rune) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

// printString prints a string rune by rune
func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		z01.PrintRune('\n')
		return
	}

	// Collect all vowels across all arguments
	var vowels []rune
	for _, arg := range args {
		for _, r := range arg {
			if isVowel(r) {
				vowels = append(vowels, r)
			}
		}
	}

	// Reverse index for vowel replacement
	j := len(vowels) - 1

	for i, arg := range args {
		for _, r := range arg {
			if isVowel(r) && len(vowels) > 0 {
				z01.PrintRune(vowels[j])
				j--
			} else {
				z01.PrintRune(r)
			}
		}
		if i != len(args)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
