package hack_assembler

import (
	"fmt"
	"strconv"

	"github.com/riku929hr/nand2tetris/assembler/hack_assembler/code"
	"github.com/riku929hr/nand2tetris/assembler/hack_assembler/parser"
	symboltable "github.com/riku929hr/nand2tetris/assembler/hack_assembler/symbol_table"
)

type Assembler struct {
	Parser      *parser.Parser
	SymbolTable *symboltable.SymbolTable
}

func NewAssembler(p *parser.Parser, st *symboltable.SymbolTable) *Assembler {
	return &Assembler{
		Parser:      p,
		SymbolTable: st,
	}
}

func (a *Assembler) Assemble() (string, error) {
	p := a.Parser
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
