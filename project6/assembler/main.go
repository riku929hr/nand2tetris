package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/riku929hr/nand2tetris/assembler/hack_assembler"
	symboltable "github.com/riku929hr/nand2tetris/assembler/hack_assembler/symbol_table"
)

func main() {
	// read from file
	content, err := os.ReadFile("sample.asm")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	assembler := hack_assembler.NewAssembler(
		strings.NewReader(string(content)),
		symboltable.NewSymbolTable(),
	)

	assembler.Assemble()
}
