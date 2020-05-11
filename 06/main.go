package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/roronya/nand2tetris/06/code"
	"github.com/roronya/nand2tetris/06/parser"
	"github.com/roronya/nand2tetris/06/symboletable"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("invalid arguments. usage: assembler input_file output_file")
	}

	asm, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	input := string(asm)
	spliteds := strings.Split(strings.TrimSpace(input), "\n")
	commands := []string{}
	for _, splited := range spliteds {
		command := strings.TrimSpace(splited)
		if len(command) == 0 || command[0:2] == "//" {
			continue
		}
		commands = append(commands, command)
	}
	commands = append(commands, "EOF")

	address := 16
	st := symboletable.New(commands)
	p := parser.New(input)
	evaluated := []string{}
	for p.HasMoreCommands() {
		p.Advance()
		commandType := p.CommandType()
		switch commandType {
		case parser.A_COMMAND:
			symbol := p.Symbol()
			value, err := strconv.Atoi(symbol)
			if err != nil {
				if st.Contains(symbol) {
					value = st.GetAddress(symbol)
				} else {
					st.AddEntry(symbol, address)
					address++
					value = address
				}
			}
			binary := toBynary(value)
			evaluated = append(evaluated, binary)
		case parser.L_COMMAND:
			// do nothing
		default:
			binary := fmt.Sprintf("111%s%s%s", code.Comp(p.Comp()), code.Dest(p.Dest()), code.Jump(p.Jump()))
			evaluated = append(evaluated, binary)
		}
	}
	output := strings.Join(evaluated, "\n")

	err = ioutil.WriteFile(os.Args[2], []byte(output), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func toBynary(i int) string {
	if i < 0 {
		i = 32768 + i
	}
	return fmt.Sprintf("%016b", i)
}
