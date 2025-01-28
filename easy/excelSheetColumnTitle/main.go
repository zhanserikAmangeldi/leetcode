package main

func convertToTitle(columnNumber int) string {
	str := ""
	var char rune

	for columnNumber > 0 {
		columnNumber--
		char = 'A' + rune(columnNumber%26)
		str = string(char) + str
		columnNumber /= 26
	}

	return str
}
