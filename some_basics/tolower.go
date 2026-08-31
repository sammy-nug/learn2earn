package piscine

func ToLower(s string) string {
	y := []rune(s)
	for i, x := range y {
		if x >= 'A' && x <= 'Z' {
			y[i] = x + 32
		}
	}
	return string(y)
}
