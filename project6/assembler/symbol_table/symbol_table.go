package symboltable

import "errors"

type SymbolTable struct {
	symbols map[string]int
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		symbols: make(map[string]int),
	}
}

func (st *SymbolTable) AddEntry(symbol string, address int) {
	st.symbols[symbol] = address
}

func (st *SymbolTable) Contains(symbol string) bool {
	_, exists := st.symbols[symbol]
	return exists
}

func (st *SymbolTable) GetAddress(symbol string) (int, error) {
	if address, exists := st.symbols[symbol]; exists {
		return address, nil
	}
	return -1, errors.New("symbol not found: " + symbol)
}
