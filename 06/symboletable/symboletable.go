package symboletable

import (
	"fmt"
)

type SymboleTable struct {
	table map[string]int
}

func New(commands []string) *SymboleTable {
	table := map[string]int{}
	table["SP"] = 0
	table["LCL"] = 1
	table["ARG"] = 2
	table["THIS"] = 3
	table["THAT"] = 4
	for i := 0; i < 16; i++ {
		table[fmt.Sprintf("R%d", i)] = i
	}
	table["SCREEN"] = 16384
	table["KBD"] = 24567
	st := &SymboleTable{table}
	st.scan(commands)
	return st
}

func (st *SymboleTable) scan(commands []string) {
	i := 0
	for _, command := range commands {
		lastIndex := len(command) - 1
		if command[0:1] == "(" && command[lastIndex:] == ")" {
			label := command[1:lastIndex]
			st.AddEntry(label, i)
			continue
		}
		i++
	}
}

func (st *SymboleTable) AddEntry(symbol string, address int) {
	st.table[symbol] = address
}

func (st *SymboleTable) Contains(symbol string) bool {
	_, ok := st.table[symbol]
	return ok
}

func (st *SymboleTable) GetAddress(symbol string) int {
	return st.table[symbol]
}
