package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 { // checking if it is negative
		return
	}
	var x [20]int
	y := 0
	if n == 0 { // checking if it is a number
		x[0] = 0
		y = 1
	} else {
		for n > 0 {
			x[y] = n % 10
			n /= 10
			y++
		}
	} // sort in assending order
	for i := 0; i < y; i++ {
		for j := i + 1; j < y; j++ {
			if x[j] < x[i] {
				x[i], x[j] = x[j], x[i]
			}
		}
	}
	for i := 0; i < y; i++ {
		z01.PrintRune(rune(x[i] + '0'))
	}
}
