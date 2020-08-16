package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/roronya/nand2tetris/08/codewriter"
	"github.com/roronya/nand2tetris/08/parser"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("invalid arguments. usage: ./VMtranslator source")
	}
	path := os.Args[1]
	e := filepath.Ext(path)
	if e != ".vm" {
		log.Fatalf("invalid file format. expected .vm, got=%s", e)
	}
	rep := regexp.MustCompile(`.vm$`)
	name := filepath.Base(rep.ReplaceAllString(path, ""))

	source, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	commands := string(source)
	buf := strings.NewReader(commands)
	scanner := bufio.NewScanner(buf)
	p := parser.New(scanner)
	cw := codewriter.New(fmt.Sprintf("%s.asm", name))
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
		}
	}
	cw.Close()
}
