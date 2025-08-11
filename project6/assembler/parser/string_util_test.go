package parser

import "testing"

func Test_isCommentLine_returns_true_when_starting_with_double_slash(t *testing.T) {
	result := isCommentLine("// this is a comment")

	if result != true {
		t.Errorf("Expected true, got %v", result)
	}
}

func Test_isCommentLine_returns_false_when_not_starting_with_double_slash(t *testing.T) {
	result := isCommentLine("this is not a comment")

	if result != false {
		t.Errorf("Expected false, got %v", result)
	}
}

func Test_getInstructionType(t *testing.T) {
	tests := []struct {
		lineStr          string
		expectedType     InstructionType
		expectedErrorMsg string
	}{
		{"@R0", AInstruction, ""},
		{"(LOOP)", LInstruction, ""},
		{"D=M", CInstruction, ""},
		{"0;JMP", CInstruction, ""},
		{"INVALID", "", "parse error: unknown instruction type: INVALID"},
	}

	for _, test := range tests {
		result, err := getInstructionType(test.lineStr)
		if err != nil && err.Error() != test.expectedErrorMsg {
			t.Errorf("Expected error '%s', got '%s'", test.expectedErrorMsg, err.Error())
			continue
		}
		if result != test.expectedType {
			t.Errorf("Expected type '%s', got '%s' for line '%s'", test.expectedType, result, test.lineStr)
		}
	}
}

func Test_removeSpaces(t *testing.T) {
	tests := []struct {
		lineStr          string
		expectedResult   string
		expectedErrorMsg string
	}{
		{"   D=M   ", "D=M", ""},
		{"A=1", "A=1", ""},
		{"   ", "", ""},
		{"( LOOP) ", "(LOOP)", ""},
	}

	for _, test := range tests {
		result := removeSpaces(test.lineStr)
		if result != test.expectedResult {
			t.Errorf("Expected '%s', got '%s'", test.expectedResult, result)
		}
	}
}

func Test_stringIndex(t *testing.T) {
	tests := []struct {
		s                string
		char             rune
		expectedIndex    int
		expectedErrorMsg string
	}{
		{"hello", 'e', 1, ""},
		{"world", 'd', 4, ""},
		{"test", 'x', -1, ""},
		{"", 'a', -1, ""},
		{"abc", 'c', 2, ""},
	}
	for _, test := range tests {

		result := stringIndex(test.s, test.char)
		if result != test.expectedIndex {
			t.Errorf("Expected index %d for '%s' with char '%c', got %d", test.expectedIndex, test.s, test.char, result)
		}
	}
}
