package hack_assembler_test

import (
	"strings"
	"testing"

	"github.com/riku929hr/nand2tetris/assembler/hack_assembler"
	symboltable "github.com/riku929hr/nand2tetris/assembler/hack_assembler/symbol_table"
)

func Test_FirstPass(t *testing.T) {
	t.Run("it should parse the file and add symbols to the symbol table", func(t *testing.T) {
		input := `
			(LOOP)
			@R0
			D=M
			(END)
			@END
			0;JMP`
		st := symboltable.NewSymbolTable()
		a := hack_assembler.NewAssembler(strings.NewReader(input), st)

		err := a.FirstPass()
		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}

		if !st.Contains("LOOP") || !st.Contains("END") {
			t.Error("Expected symbols 'LOOP' and 'END' to be added to the symbol table.")
		}

		if address, err := st.GetAddress("LOOP"); err != nil || address != 0 {
			t.Errorf("Expected address for 'LOOP' to be 0, but got: %d, error: %v", address, err)
		}
		if address, err := st.GetAddress("END"); err != nil || address != 2 {
			t.Errorf("Expected address for 'END' to be 2, but got: %d, error: %v", address, err)
		}
	})
}
