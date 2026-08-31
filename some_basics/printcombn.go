package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	first := true
	current := make([]rune, n)

	var backtrack func(start int, depth int)

	backtrack = func(start int, depth int) {
		// If we've chosen n digits, print the current combination
		if depth == n {
			if !first {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			first = false

			for i := 0; i < n; i++ {
				z01.PrintRune(current[i])
			}
			return
		}

		// Loop through digits from start to 9
		for i := start; i <= 9; i++ {
			current[depth] = rune('0' + i)
			backtrack(i+1, depth+1)
		}
	}

	backtrack(0, 0)
	z01.PrintRune('\n')
}
