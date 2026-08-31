package piscine

func IsLower(s string) bool {
	for _, x := range s {
		if x < 'a' || x > 'z' {
			return false
		}
	}
	return true
}
