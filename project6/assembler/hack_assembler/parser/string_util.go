package parser

func isCommentLine(lineStr string) bool {
	if lineStr == "" {
		return false
	}
	return lineStr[0:2] == "//"
}

func removeSpaces(lineStr string) string {
	// remove all spaces from the line
	result := ""
	for _, char := range lineStr {
		if char != ' ' && char != '\t' {
			result += string(char)
		}
	}

	return result
}

func stringIndex(s string, char rune) int {
	for i, c := range s {
		if c == char {
			return i
		}
	}
	return -1 // not found
}
