package piscine

func IsUpper(s string) bool {
	for _, x := range s {
		if x < 'A' || x > 'Z' {
			return false
		}
	}
	return true
}
