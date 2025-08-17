// Code package provides functions to convert Hack assembly instructions
// into binary code according to the Hack computer architecture specification.
// It includes functions for A-instructions, C-instructions, and jump/dest/comp codes.
package code

import (
	"errors"
)

var destMap = map[string]string{
	"":     "000", // when dest is omitted, it defaults to "000"
	"null": "000",
	"M":    "001",
	"D":    "010",
	"MD":   "011",
	"A":    "100",
	"AM":   "101",
	"AD":   "110",
	"ADM":  "111",
}

func Dest(dest string) (string, error) {
	if value, exists := destMap[dest]; exists {
		return value, nil
	}
	return "", errors.New("invalid dest: " + dest)
}

var compMap = map[string]string{
	"0":   "0101010",
	"1":   "0111111",
	"-1":  "0111010",
	"D":   "0001100",
	"A":   "0110000",
	"M":   "1110000",
	"!D":  "0001101",
	"!A":  "0110001",
	"-D":  "0001111",
	"-A":  "0110011",
	"D+1": "0011111",
	"A+1": "0110111",
	"M+1": "1110111",
	"D-1": "0001110",
	"A-1": "0110010",
	"M-1": "1110010",
	"D+A": "0000010",
	"D+M": "1000010",
	"D-A": "0010011",
	"D-M": "1010011",
	"A-D": "0000111",
	"M-D": "1000111",
	"D&A": "0000000",
	"D&M": "1000000",
	"D|A": "0010101",
	"D|M": "1010101",
}

func Comp(comp string) (string, error) {
	if value, exists := compMap[comp]; exists {
		return value, nil
	}
	return "", errors.New("invalid comp: " + comp)
}

var jumpMap = map[string]string{
	"":     "000", // when jump is omitted, it defaults to "000"
	"null": "000",
	"JGT":  "001",
	"JEQ":  "010",
	"JGE":  "011",
	"JLT":  "100",
	"JNE":  "101",
	"JLE":  "110",
	"JMP":  "111",
}

func Jump(jump string) (string, error) {
	if value, exists := jumpMap[jump]; exists {
		return value, nil
	}
	return "", errors.New("invalid jump: " + jump)
}
