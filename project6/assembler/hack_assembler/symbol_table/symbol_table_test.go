package symboltable_test

import (
	"testing"

	symboltable "github.com/riku929hr/nand2tetris/assembler/hack_assembler/symbol_table"
)

func TestSymbolTable(t *testing.T) {
	t.Run("it returns the right address when add new entry and run GetAddress", func(t *testing.T) {
		// Example: Add a symbol to the table and verify it exists.
		symbolTable := symboltable.NewSymbolTable()
		symbolTable.AddEntry("TEST", 123)

		if !symbolTable.Contains("TEST") {
			t.Errorf("Expected symbol 'TEST' to be in the table, but it was not found.")
		}

		address, err := symbolTable.GetAddress("TEST")
		if err != nil || address != 123 {
			t.Errorf("Expected address for 'TEST' to be 123, but got: %d", address)
		}
	})

	t.Run("it returns false when symbol does not exist", func(t *testing.T) {
		// Example: Check for a symbol that has not been added.
		symbolTable := symboltable.NewSymbolTable()

		if symbolTable.Contains("NON_EXISTENT") {
			t.Errorf("Expected symbol 'NON_EXISTENT' to not be in the table, but it was found.")
		}

		_, err := symbolTable.GetAddress("NON_EXISTENT")
		if err == nil {
			t.Errorf("Expected error when getting address for 'NON_EXISTENT', but got none.")
		}
	})
}
