package piscine

func Compare(a, b string) int {
	x := []rune(a)
	y := []rune(b)
	for i := 0; i < len(x) && i < len(y); i++ {
		if x[i] < y[i] {
			return -1
		} else if x[i] > y[i] {
			return 1
		}
	}
	if len(x) < len(y) {
		return -1
	} else if len(x) > len(y) {
		return 1
	}
	return 0
}
