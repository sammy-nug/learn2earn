package piscine

func Join(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}
	final := strs[0]
	for i := 1; i < len(strs); i++ {
		final += sep + strs[i]
	}
	return final
}
