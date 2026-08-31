package piscine

func StrLen(s string) int {
	y := 0
	for range s { // iterate over each rune of the string
		y++
	}
	return y
}
