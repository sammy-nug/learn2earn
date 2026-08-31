package piscine

func Capitalize(s string) string {
	y := []rune(s)
	Isletter := true

	for i := 0; i < len(y); i++ {
		x := y[i]
		if (x >= 'A' && x <= 'Z') || (x >= 'a' && x <= 'z') || (x >= '0' && x <= '9') {
			if Isletter { // if we find any new alphabet
				if x >= 'a' && x <= 'z' { // if the first alphabet encounterd is small
					y[i] = x - 32
				}
				Isletter = false
			} else {
				if x >= 'A' && x <= 'Z' { // if the second alphabet we encounterd is big we reduce it
					y[i] = x + 32
				}
			}
		} else {
			Isletter = true
		}
	}
	return string(y)
}
