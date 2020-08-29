package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/roronya/nand2tetris/08/codewriter"
	"github.com/roronya/nand2tetris/08/parser"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("invalid arguments. usage: ./VMtranslator source")
	}
	cw := codewriter.New("main.asm")
	cw.WriteBootstrap()

	for _, arg := range os.Args[1:] {
		source, err := ioutil.ReadFile(arg)
		if err != nil {
			log.Fatal(err)
		}
		commands := string(source)
		buf := strings.NewReader(commands)
		scanner := bufio.NewScanner(buf)
		p := parser.New(scanner)
		for p.HasMoreCommands() {
			p.Advance()
			switch p.CommandType {
			case parser.C_PUSH:
				fallthrough
			case parser.C_POP:
				cw.WritePushPop(p.CommandType, p.Arg1, p.Arg2)
			case parser.C_ARITHMETIC:
				cw.WriteArithmetic(p.Command)
			case parser.C_LABEL:
				cw.WriteLabel(p.Arg1)
			case parser.C_IF:
				cw.WriteIfGoto(p.Arg1)
			case parser.C_GOTO:
				cw.WriteGoto(p.Arg1)
			case parser.C_FUNCTION:
				cw.WriteFunction(p.Arg1, p.Arg2)
			case parser.C_RETURN:
				cw.WriteReturn()
			case parser.C_CALL:
				cw.WriteCall(p.Arg1, p.Arg2)
			}
		}
	}
	cw.Close()
}
