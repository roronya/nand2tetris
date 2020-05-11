package symboletable

import (
	"fmt"
	"strings"
)

type SymboleTable struct {
	table map[string]int
}

func New(input string) *SymboleTable {
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
	st.scan(input)
	return st
}

func (st *SymboleTable) scan(input string) {
	commands := strings.Split(input, "\n")
	for i, command := range commands {
		trimed := strings.TrimSpace(command)
		lastIndex := len(trimed) - 1
		if trimed[0:1] == "(" && trimed[lastIndex:] == ")" {
			label := trimed[1:lastIndex]
			st.addEntry(label, i)
		}
		i++
	}
}

func (st *SymboleTable) addEntry(symbol string, address int) {
	st.table[symbol] = address
}

func (st *SymboleTable) contains(symbol string) bool {
	_, ok := st.table[symbol]
	return ok
}

func (st *SymboleTable) getAddress(symbol string) int {
	return st.table[symbol]
}
