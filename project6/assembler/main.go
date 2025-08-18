package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/riku929hr/nand2tetris/assembler/hack_assembler"
	symboltable "github.com/riku929hr/nand2tetris/assembler/hack_assembler/symbol_table"
)

func main() {
	// filepath from args
	if len(os.Args) < 2 {
		fmt.Println("Usage: assembler [path_to_asm_file]")
		return
	}
	filePath := os.Args[1]
	if filePath == "" {
		fmt.Println("Please provide a valid file path.")
		return
	}

	// check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Printf("File does not exist: %s\n", filePath)
		return
	}

	// read from file
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	assembler := hack_assembler.NewAssembler(
		strings.NewReader(string(content)),
		symboltable.NewSymbolTable(),
	)

	_, err = assembler.Assemble()
	if err != nil {
		fmt.Printf("Assembly error: %v\n", err)
	}
}
