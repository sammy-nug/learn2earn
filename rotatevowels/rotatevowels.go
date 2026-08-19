package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Checks if a rune is a vowel (y is NOT a vowel)
func isVowel(c rune) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

// Mirrors vowels in a string
func mirrorVowels(s string) string {
	runes := []rune(s)
	vowels := []rune{}

	// Collect vowels
	for _, r := range runes {
		if isVowel(r) {
			vowels = append(vowels, r)
		}
	}

	// If no vowels, return original string
	if len(vowels) == 0 {
		return s
	}

	// Replace vowels in reverse order
	j := len(vowels) - 1
	for i, r := range runes {
		if isVowel(r) {
			runes[i] = vowels[j]
			j--
		}
	}

	return string(runes)
}

// Prints a string using z01.PrintRune
func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	if len(os.Args) < 2 {
		z01.PrintRune('\n')
		return
	}

	for i, arg := range os.Args[1:] {
		printString(mirrorVowels(arg))
		if i != len(os.Args[1:])-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
