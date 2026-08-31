package piscine

func IsAlpha(s string) bool {
	for _, x := range s {
		if (x < 'A' || x > 'Z') && (x < 'a' || x > 'z') && (x < '0' || x > '9') {
			return false
		}
	}
	return true
}
