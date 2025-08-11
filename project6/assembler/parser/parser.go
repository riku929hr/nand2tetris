package parser

import (
	"bufio"
	"errors"
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
}

func NewParser(scanner *bufio.Scanner) *Parser {
	return &Parser{
		scanner: scanner,
	}
}

func (p *Parser) Advance() {
	// if line starts with "//", read next line
	p.scanner.Scan()
	p.currentLine = p.scanner.Text()
	if isCommentLine(p.currentLine) {
		p.Advance() // recursively call Advance to skip comments
		return
	}

	// if includes space, remove them all
	p.currentLine = removeSpaces(p.currentLine)
	if p.currentLine == "" {
		p.Advance() // skip empty lines
		return
	}
}

func (p *Parser) InstructionType() (InstructionType, error) {
	return getInstructionType(p.currentLine)
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

	// dest=comp;jumpの=の前までを取得して返す
	equalIndex := stringIndex(p.currentLine, '=')
	if equalIndex == -1 {
		return "", errors.New("no dest found in C instruction")
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

	// dest=comp;jumpの=の後から;までを取得して返す
	semicolonIndex := stringIndex(p.currentLine, ';')
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

	// dest=comp;jumpの;の後を取得して返す
	semicolonIndex := stringIndex(p.currentLine, ';')
	if semicolonIndex == -1 {
		return "", errors.New("no jump found in C instruction")
	}
	jump := p.currentLine[semicolonIndex+1:]
	if jump == "" {
		return "", errors.New("jump is empty in C instruction")
	}
	return jump, nil
}
