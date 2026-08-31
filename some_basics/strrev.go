package piscine

func StrRev(s string) string {
	y := []rune(s)
	x := len(y)
	for i := 0; i < x/2; i++ {
		y[i], y[x-1-i] = y[x-1-i], y[i]
	}
	return string(y)
}
