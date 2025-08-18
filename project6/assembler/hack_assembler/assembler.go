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

func isNumber(s string) bool {
	for _, char := range s {
		if char < '0' || char > '9' {
			return false
		}
	}
	return len(s) > 0 // ensure it's not an empty string
}

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
			// Label points to the next instruction address
			st.AddEntry(symbol, currentLineNumber)
		} else {
			// Only count A and C instructions as actual code lines
			currentLineNumber++
		}
	}
	return nil
}

func (a *Assembler) Assemble() (string, error) {
	if err := a.FirstPass(); err != nil {
		return "", err
	}
	currentVariableAddress := 16

	p := parser.NewParser(strings.NewReader(a.content))
	for {
		p.Advance()

		if !p.HasMoreLines() {
			break
		}

		instructionType, err := p.InstructionType()
		if err != nil {
			return "", err
		}

		// Skip L instructions (labels) as they don't generate machine code
		if instructionType == parser.LInstruction {
			continue
		}

		if instructionType == parser.AInstruction {
			symbol, err := p.Symbol()
			if err != nil {
				return "", err
			}

			if isNumber(symbol) && !a.SymbolTable.Contains(symbol) {
				a.SymbolTable.AddEntry(symbol, currentVariableAddress)
				currentVariableAddress++
			}

			var decimal int64
			if isNumber(symbol) {
				// Symbol is a number, convert to int64
				num, _ := strconv.ParseInt(symbol, 10, 64)
				decimal = int64(num)
			} else {
				// Symbol is a variable or label, get from symbol table
				if addr, err := a.SymbolTable.GetAddress(symbol); err == nil {
					decimal = int64(addr)
				} else {
					return "", fmt.Errorf("undefined symbol: %s", symbol)
				}
			}

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
			compBinary, err := code.Comp(comp)
			if err != nil {
				return "", err
			}

			destBinary, err := code.Dest(dest)
			if err != nil {
				return "", err
			}

			jumpBinary, err := code.Jump(jump)
			if err != nil {
				return "", err
			}
			fmt.Println("111" + compBinary + destBinary + jumpBinary)
		}
	}
	return "success", nil
}
