package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/riku929hr/nand2tetris/assembler/hack_assembler"
	"github.com/riku929hr/nand2tetris/assembler/hack_assembler/parser"
	symboltable "github.com/riku929hr/nand2tetris/assembler/hack_assembler/symbol_table"
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

	assembler := hack_assembler.NewAssembler(
		parser.NewParser(scanner),
		symboltable.NewSymbolTable(),
	)

	assembler.Assemble()
}
