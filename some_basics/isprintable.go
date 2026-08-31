package piscine

func IsPrintable(s string) bool {
	for _, x := range s {
		if x < 32 || x > 126 {
			return false
		}
	}
	return true
}
