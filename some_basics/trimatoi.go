package piscine

func TrimAtoi(s string) int {
	sign := 1
	total := 0
	Insign := false
	Isnub := false

	for i := 0; i < len(s); i++ {
		x := s[i]

		if x == '-' && !Isnub && !Insign {
			sign = -1
			Insign = true
			continue

		}
		if x >= '0' && x <= '9' {
			Isnub = true
			total = total*10 + int(x-'0')
		}
	}
	if !Isnub {
		return 0
	}
	return total * sign
}
