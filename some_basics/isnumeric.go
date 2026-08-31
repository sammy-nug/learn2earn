package piscine

func IsNumeric(s string) bool {
	for _, x := range s {
		if x < '0' || x > '9' {
			return false
		}
	}
	return true
}
