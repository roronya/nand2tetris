package main

import (
	"bufio"
	"strings"

	"github.com/roronya/nand2tetris/07/codewriter"
	"github.com/roronya/nand2tetris/07/parser"
)

func main() {
	commands := "push constant 10\npush constant 10\nadd"
	buf := strings.NewReader(commands)
	scanner := bufio.NewScanner(buf)
	p := parser.New(scanner)
	cw := codewriter.New("result.asm")
	for p.HasMoreCommands() {
		p.Advance()
		switch p.CommandType {
		case parser.C_PUSH:
			cw.WritePushPop(p.CommandType, p.Arg1, p.Arg2)
		}
	}
	cw.Close()
}
