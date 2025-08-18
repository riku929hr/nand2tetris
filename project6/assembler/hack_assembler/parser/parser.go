package parser

import (
	"bufio"
	"errors"
	"io"
)

type InstructionType string

const (
	AInstruction InstructionType = "A_INSTRUCTION"
	CInstruction InstructionType = "C_INSTRUCTION"
	LInstruction InstructionType = "L_INSTRUCTION"
)

type Parser struct {
	scanner     *bufio.Scanner
	currentLine string
	hasNext     bool
}

func NewParser(reader io.Reader) Parser {
	scanner := bufio.NewScanner(reader)
	hasNext := scanner.Scan() // Check if there's at least one line
	return Parser{
		scanner: scanner,
		hasNext: hasNext,
	}
}

func (p *Parser) HasMoreLines() bool {
	return p.hasNext
}

func (p *Parser) Advance() {
	if !p.hasNext {
		return
	}

	spaceRemoved := removeSpaces(p.scanner.Text())
	p.hasNext = p.scanner.Scan()

	if isCommentLine(spaceRemoved) {
		if p.hasNext {
			p.Advance() // recursively call Advance to skip comments
		}
		return
	}

	if spaceRemoved == "" {
		if p.hasNext {
			p.Advance() // skip empty lines
		}
		return
	}

	p.currentLine = spaceRemoved
}

func (p *Parser) InstructionType() (InstructionType, error) {
	if p.currentLine == "" {
		return "", errors.New("parse error: empty line")
	}

	if p.currentLine[0] == '@' {
		return AInstruction, nil
	}

	if p.currentLine[0] == '(' && p.currentLine[len(p.currentLine)-1] == ')' {
		return LInstruction, nil
	}

	if stringIndex(p.currentLine, '=') != -1 || stringIndex(p.currentLine, ';') != -1 {
		return CInstruction, nil
	}

	return "", errors.New("parse error: unknown instruction type: " + p.currentLine)
}

func (p *Parser) Symbol() (string, error) {
	instructionType, err := p.InstructionType()
	if err != nil {
		return "", err
	}
	if instructionType == CInstruction {
		return "", errors.New("not A or L instruction")
	}

	if instructionType == AInstruction {
		return p.currentLine[1:], nil
	}

	if instructionType == LInstruction {
		// remove the parentheses
		return p.currentLine[1 : len(p.currentLine)-1], nil
	}

	return "", nil
}

func (p *Parser) Dest() (string, error) {
	instructionType, err := p.InstructionType()
	if err != nil {
		return "", err
	}
	if instructionType != CInstruction {
		return "", errors.New("not C instruction")
	}

	equalIndex := stringIndex(p.currentLine, '=')

	// If there is no equal sign, it means there is no dest instruction
	if equalIndex == -1 {
		return "", nil
	}

	dest := p.currentLine[:equalIndex]
	if dest == "" {
		return "", errors.New("dest is empty in C instruction")
	}
	return dest, nil
}

func (p *Parser) Comp() (string, error) {
	instructionType, err := p.InstructionType()
	if err != nil {
		return "", err
	}
	if instructionType != CInstruction {
		return "", errors.New("not C instruction")
	}
	semicolonIndex := stringIndex(p.currentLine, ';')

	// If there is no semicolon, it means there is no jump instruction
	if semicolonIndex == -1 {
		return p.currentLine[stringIndex(p.currentLine, '=')+1:], nil
	}
	// If there is a semicolon, we take the part before it
	if stringIndex(p.currentLine, '=') == -1 {
		return p.currentLine[:semicolonIndex], nil
	}

	comp := p.currentLine[stringIndex(p.currentLine, '=')+1 : semicolonIndex]
	if comp == "" {
		return "", errors.New("comp is empty in C instruction")
	}
	return comp, nil
}

func (p *Parser) Jump() (string, error) {
	instructionType, err := p.InstructionType()
	if err != nil {
		return "", err
	}
	if instructionType != CInstruction {
		return "", errors.New("not C instruction")
	}

	semicolonIndex := stringIndex(p.currentLine, ';')

	// If there is no semicolon, it means there is no jump instruction
	if semicolonIndex == -1 {
		return "", nil
	}

	jump := p.currentLine[semicolonIndex+1:]
	if jump == "" {
		return "", errors.New("jump is empty in C instruction")
	}
	return jump, nil
}
