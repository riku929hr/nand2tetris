// Code package provides functions to convert Hack assembly instructions
// into binary code according to the Hack computer architecture specification.
// It includes functions for A-instructions, C-instructions, and jump/dest/comp codes.
package code

func Dest(dest string) string {
	switch dest {
	case "null":
		return "000"
	case "M":
		return "001"
	case "D":
		return "010"
	case "MD":
		return "011"
	case "A":
		return "100"
	case "AM":
		return "101"
	case "AD":
		return "110"
	case "ADM":
		return "111"
	}
	return "000" // invalid dest, return null
}

func Comp(comp string) string {
	switch comp {
	case "0":
		return "101010"
	case "1":
		return "111111"
	case "-1":
		return "111010"
	case "D":
		return "001100"
	case "A":
		return "110000"
	case "!D":
		return "001101"
	case "!A":
		return "110001"
	case "-D":
		return "001111"
	case "-A":
		return "110011"
	case "D+1":
		return "011111"
	case "A+1":
		return "110111"
	case "D-1":
		return "001110"
	case "A-1":
		return "110010"
	case "D+A":
		return "000010"
	case "D-A":
		return "010011"
	case "A-D":
		return "000111"
	case "D&A":
		return "000000"
	case "D|A":
		return "010101"
	default:
		return ""
	}
}

func Jump(jump string) string {
	switch jump {
	case "null":
		return "000"
	case "JGT":
		return "001"
	case "JEQ":
		return "010"
	case "JGE":
		return "011"
	case "JLT":
		return "100"
	case "JNE":
		return "101"
	case "JLE":
		return "110"
	case "JMP":
		return "111"
	}
	return "000"
}
