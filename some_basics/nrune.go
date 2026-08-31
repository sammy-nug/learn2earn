package piscine

func NRune(s string, n int) rune {
	if n <= 0 {
		return 0
	}
	x := 0
	for _, r := range s {
		x++
		if x == n {
			return r
		}
	}
	return 0
}
