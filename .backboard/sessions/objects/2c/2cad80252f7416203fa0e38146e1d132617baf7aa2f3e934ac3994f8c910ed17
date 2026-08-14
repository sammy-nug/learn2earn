package main

func RenderToString(lines []string, banner []string) string {

	result := ""

	for _, text := range lines {

		if text == "" {
			result += "\n"
			continue
		}

		for row := 0; row < 8; row++ {

			for _, char := range text {

				if !IsPrintable(char) {
					continue
				}

				index := GetCharIndex(char)
				result += banner[index+row]
			}

			result += "\n"
		}
	}

	return result
}
