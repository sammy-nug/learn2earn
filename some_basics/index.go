package piscine

func Index(s string, toFind string) int {
	x := len(s)
	y := len(toFind)
	if y == 0 {
		return 0
	}
	for i := 0; i <= x-y; i++ {
		if s[i:i+y] == toFind {
			return i
		}
	}
	return -1
}
