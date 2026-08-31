package piscine

func AlphaCount(s string) int {
	count := 0
	for _, x := range s {
		if (x >= 'a' && x <= 'z') || (x >= 'A' && x <= 'Z') {
			count++
		}
	}
	return count
}
