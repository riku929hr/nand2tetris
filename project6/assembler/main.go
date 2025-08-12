package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/riku929hr/nand2tetris/assembler/code"
	"github.com/riku929hr/nand2tetris/assembler/parser"
)

func main() {
	// read from file
	data, err := os.Open("sample.asm")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	defer data.Close()
	// print

	scanner := bufio.NewScanner(data)

	p := parser.NewParser(scanner)

	for p.HasMoreLines() {
		p.Advance()
		instructionType, err := p.InstructionType()
		if err != nil {
			fmt.Println(err)
			return
		}

		if instructionType == parser.AInstruction {
			symbol, err := p.Symbol()
			if err != nil {
				fmt.Println(err)
				return
			}
			decimal, _ := strconv.ParseInt(symbol, 10, 64)
			binaryString := fmt.Sprintf("%016b", decimal)
			fmt.Println(binaryString)
		}

		if instructionType == parser.CInstruction {
			dest, err := p.Dest()
			if err != nil {
				fmt.Println(err)
				return
			}

			comp, err := p.Comp()
			if err != nil {
				fmt.Println(err)
				return
			}

			jump, err := p.Jump()
			if err != nil {
				fmt.Println(err)
				return
			}
			// Convert comp, dest, and jump to binary
			compBinary := code.Comp(comp)
			destBinary := code.Dest(dest)
			jumpBinary := code.Jump(jump)
			fmt.Println("1110" + compBinary + destBinary + jumpBinary)
		}
	}
}
