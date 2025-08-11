package parser

import "errors"

func isCommentLine(lineStr string) bool {
	return lineStr[0:2] == "//"
}

func getInstructionType(lineStr string) (InstructionType, error) {
	if lineStr[0] == '@' {
		return AInstruction, nil
	}

	if lineStr[0] == '(' && lineStr[len(lineStr)-1] == ')' {
		return LInstruction, nil
	}
	if stringIndex(lineStr, '=') != -1 || stringIndex(lineStr, ';') != -1 {
		return CInstruction, nil
	}

	return "", errors.New("parse error: unknown instruction type: " + lineStr)
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
