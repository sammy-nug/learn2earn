package piscine

func Atoi(s string) int {
	sign := 1
	result := 0
	if len(s) == 0 {
		return 0
	}
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}
	if len(s) == 0 {
		return 0
	}
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return 0
		}
		result = result*10 + int(ch-'0')
	}
	return result * sign
}
