package piscine

func LastRune(s string) rune {
	var l rune
	for _, r := range s {
		l = r
	}
	return l
}
