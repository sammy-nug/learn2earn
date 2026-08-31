package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		if n == -9223372036854775808 {
			for _, r := range "9223372036854775808" {
				z01.PrintRune(r)
			}
			return
		}
		n = -n
	}

	printDigits(n)
}

func printDigits(n int) {
	if n > 0 {
		printDigits(n / 10)
		z01.PrintRune(rune(n%10 + '0'))
	}
}
