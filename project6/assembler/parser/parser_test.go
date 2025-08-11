package parser_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/riku929hr/nand2tetris/assembler/parser" // Adjust the import path as necessary
)

func Test_InstructionType(t *testing.T) {
	tests := []struct {
		line     string
		expected parser.InstructionType
	}{
		{"@R0", parser.AInstruction},
		{"(LOOP)", parser.LInstruction},
		{"D=M", parser.CInstruction},
	}

	for _, test := range tests {
		scanner := bufio.NewScanner(strings.NewReader(test.line))
		p := parser.NewParser(scanner)
		p.Advance()
		result, _ := p.InstructionType()
		if result != test.expected {
			t.Errorf("Expected %s, got %s for line '%s'", test.expected, result, test.line)
		}
	}
}

func Test_InstructionType_it_fails_for_unknown_instruction(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader("INVALID"))
	p := parser.NewParser(scanner)
	p.Advance()
	_, err := p.InstructionType()
	if err == nil {
		t.Error("Expected error for unknown instruction, but got none")
	}
}

func Test_Symbol_it_succeeds_if_A_or_L_instruction(t *testing.T) {
	tests := []struct {
		line     string
		expected string
	}{
		{"@R0", "R0"},
		{"(LOOP)", "LOOP"},
	}

	for _, test := range tests {
		scanner := bufio.NewScanner(strings.NewReader(test.line))
		p := parser.NewParser(scanner)
		p.Advance()
		result, err := p.Symbol()
		if err != nil && test.expected != "" {
			t.Errorf("Unexpected error for line '%s': %v", test.line, err)
			continue
		}
		if result != test.expected {
			t.Errorf("Expected symbol '%s', got '%s' for line '%s'", test.expected, result, test.line)
		}
	}
}

func Test_Symbol_it_fails_if_not_A_or_L_instruction(t *testing.T) {
	tests := []struct {
		line string
	}{
		{"D=M"},
		{"0;JMP"},
	}

	for _, test := range tests {
		scanner := bufio.NewScanner(strings.NewReader(test.line))
		p := parser.NewParser(scanner)
		p.Advance()
		_, err := p.Symbol()
		if err == nil {
			t.Errorf("Expected error for line '%s', but got none", test.line)
		}
	}
}

func Test_Dest_it_succeeds_for_C_instruction(t *testing.T) {
	tests := []struct {
		line     string
		expected string
	}{
		{"D=M", "D"},
		{"A=1", "A"},
		{"0;JMP", ""},
	}

	for _, test := range tests {
		scanner := bufio.NewScanner(strings.NewReader(test.line))
		p := parser.NewParser(scanner)
		p.Advance()
		result, err := p.Dest()
		if err != nil && test.expected != "" {
			t.Errorf("Unexpected error for line '%s': %v", test.line, err)
			continue
		}
		if result != test.expected {
			t.Errorf("Expected dest '%s', got '%s' for line '%s'", test.expected, result, test.line)
		}
	}
}

func Test_Dest_it_fails_for_non_C_instruction(t *testing.T) {
	tests := []struct {
		line string
	}{
		{"@R0"},
		{"(LOOP)"},
	}

	for _, test := range tests {
		scanner := bufio.NewScanner(strings.NewReader(test.line))
		p := parser.NewParser(scanner)
		p.Advance()
		_, err := p.Dest()
		if err == nil {
			t.Errorf("Expected error for line '%s', but got none", test.line)
		}
	}
}

func Test_Comp_it_succeeds_for_C_instruction(t *testing.T) {
	tests := []struct {
		line     string
		expected string
	}{
		{"D=M;JMP", "M"},
		{"A=1;JGT", "1"},
	}

	for _, test := range tests {
		scanner := bufio.NewScanner(strings.NewReader(test.line))
		p := parser.NewParser(scanner)
		p.Advance()
		result, err := p.Comp()
		if err != nil && test.expected != "" {
			t.Errorf("Unexpected error for line '%s': %v", test.line, err)
			continue
		}
		if result != test.expected {
			t.Errorf("Expected comp '%s', got '%s' for line '%s'", test.expected, result, test.line)
		}
	}
}

func Test_Jump_it_succeeds_for_C_instruction(t *testing.T) {
	tests := []struct {
		line     string
		expected string
	}{
		{"D=M;JGT", "JGT"},
		{"A=1;JMP", "JMP"},
		{"D;JEQ", "JEQ"},
		{"0;JMP", "JMP"},
	}

	for _, test := range tests {
		scanner := bufio.NewScanner(strings.NewReader(test.line))
		p := parser.NewParser(scanner)
		p.Advance()
		result, err := p.Jump()
		if err != nil && test.expected != "" {
			t.Errorf("Unexpected error for line '%s': %v", test.line, err)
			continue
		}
		if result != test.expected {
			t.Errorf("Expected jump '%s', got '%s' for line '%s'", test.expected, result, test.line)
		}
	}
}

func Test_Jump_it_fails_for_non_C_instruction(t *testing.T) {
	tests := []struct {
		line string
	}{
		{"@R0"},
		{"(LOOP)"},
		{"D=M"},
	}

	for _, test := range tests {
		scanner := bufio.NewScanner(strings.NewReader(test.line))
		p := parser.NewParser(scanner)
		p.Advance()
		_, err := p.Jump()
		if err == nil {
			t.Errorf("Expected error for line '%s', but got none", test.line)
		}
	}
}
