package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/roronya/nand2tetris/06/parser"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("invalid arguments. usage: assembler input_file outpu_file")
	}

	asm, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	input := string(asm)

	p := parser.New(input)
	evaluated := []string{}
	for p.HasMoreCommands() {
		// A命令かC命令かの分岐
		p.Advance()
		symbol, _ := strconv.Atoi(p.Symbol())
		binary := toBynary(symbol)
		evaluated = append(evaluated, binary)
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
