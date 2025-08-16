package hack_assembler

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/riku929hr/nand2tetris/assembler/hack_assembler/code"
	"github.com/riku929hr/nand2tetris/assembler/hack_assembler/parser"
	symboltable "github.com/riku929hr/nand2tetris/assembler/hack_assembler/symbol_table"
)

type Assembler struct {
	content     string
	SymbolTable *symboltable.SymbolTable
}

func NewAssembler(reader io.Reader, st *symboltable.SymbolTable) *Assembler {
	content, _ := io.ReadAll(reader)
	return &Assembler{
		content:     string(content),
		SymbolTable: st,
	}
}

func (a *Assembler) FirstPass() error {
	p := parser.NewParser(strings.NewReader(a.content))
	st := a.SymbolTable

	currentLineNumber := 0

	for p.HasMoreLines() {
		p.Advance()
		instructionType, err := p.InstructionType()
		if err != nil {
			return err
		}

		if instructionType == parser.LInstruction {
			symbol, err := p.Symbol()
			if err != nil {
				return err
			}
			st.AddEntry(symbol, currentLineNumber+1)
		}
		currentLineNumber++
	}
	return nil
}

func (a *Assembler) SecondPass() error {
	p := parser.NewParser(strings.NewReader(a.content))
	st := a.SymbolTable

	currentVariableAddress := 16

	for p.HasMoreLines() {
		p.Advance()
		instructionType, err := p.InstructionType()
		if err != nil {
			return err
		}

		if instructionType == parser.AInstruction {
			symbol, err := p.Symbol()
			if err != nil {
				return err
			}

			if _, err := strconv.Atoi(symbol); err != nil { // If symbol is not a number
				if !st.Contains(symbol) {
					st.AddEntry(symbol, currentVariableAddress)
					currentVariableAddress++
				}
			}
		}
	}
	return nil
}

func (a *Assembler) Assemble() (string, error) {
	if err := a.FirstPass(); err != nil {
		return "", err
	}
	if err := a.SecondPass(); err != nil {
		return "", err
	}
	p := parser.NewParser(strings.NewReader(a.content))
	for p.HasMoreLines() {
		p.Advance()
		instructionType, err := p.InstructionType()
		if err != nil {
			return "", err
		}

		if instructionType == parser.AInstruction {
			symbol, err := p.Symbol()
			if err != nil {
				return "", err
			}

			decimal, _ := strconv.ParseInt(symbol, 10, 64)
			binaryString := fmt.Sprintf("%016b", decimal)
			fmt.Println(binaryString)
		}

		if instructionType == parser.CInstruction {
			dest, err := p.Dest()
			if err != nil {
				return "", err
			}

			comp, err := p.Comp()
			if err != nil {
				return "", err
			}

			jump, err := p.Jump()
			if err != nil {
				return "", err
			}
			// Convert comp, dest, and jump to binary
			compBinary := code.Comp(comp)
			destBinary := code.Dest(dest)
			jumpBinary := code.Jump(jump)
			fmt.Println("1110" + compBinary + destBinary + jumpBinary)
		}
	}
	return "success", nil
}
