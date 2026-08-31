package piscine

func BasicJoin(elems []string) string {
	x := ""
	for i := 0; i < len(elems); i++ {
		x += elems[i]
	}
	return x
}
