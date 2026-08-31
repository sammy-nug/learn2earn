package piscine

func ToUpper(s string) string {
	y := []rune(s)
	for i, x := range y {
		if x >= 'a' && x <= 'z' {
			y[i] = x - 32
		}
	}
	return string(y)
}
