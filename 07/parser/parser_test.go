package parser

import (
	"bufio"
	"strings"
	"testing"
)

func newParser(commands string) *Parser {
	buf := strings.NewReader(commands)
	scanner := bufio.NewScanner(buf)
	return New(scanner)
}
func TestNew(t *testing.T) {
	commands := "push constant 7\npush constant 8\nadd"
	buf := strings.NewReader(commands)
	scanner := bufio.NewScanner(buf)
	p := New(scanner)
	p.scanner.Scan()
	scaned := p.scanner.Text()
	if scaned != "push constant 7" {
		t.Fatalf("scaned is expected \"push constant 7\". got=%v", scaned)
	}
}

/**
func TestAdvance(t *testing.T) {
	commands := "push constant 7\npush constant 8\nadd"
	p := newParser(commands)
	p.Advance()
	if !(p.position == 0 && p.nextPosition == 1) {
		t.Fatalf("p.Advance() is expected setting position = 0, nextPosition = 1. got=position=%d, nextPosition=%d.", p.position, p.nextPosition)
	}
}
**/

func TestHasMoreCommands(t *testing.T) {
	commands := "push constant 7\npush constant 8\nadd"
	p := newParser(commands)
	if !p.HasMoreCommands() {
		t.Fatalf("p.HasMoreCommands is expected true. got=false.")
	}
	if !p.HasMoreCommands() {
		t.Fatalf("p.HasMoreCommands is expected true. got=false.")
	}
	if !p.HasMoreCommands() {
		t.Fatalf("p.HasMoreCommands is expected true. got=false.")
	}
	if p.HasMoreCommands() {
		t.Fatalf("p.HasMoreCommands is expected false. got=true.")
	}
}

/**
func TestHasMoreCommands_空ファイルのとき(t *testing.T) {
	commands := [][]string{
		[]string{"EOF"},
	}
	p := New(commands)
	if p.HasMoreCommands() {
		t.Fatalf("p.HasMoreCommands is expected false. got=true.")
	}
}

func TestCommandType(t *testing.T) {
	commands := [][]string{
		[]string{"add"},
		[]string{"sub"},
		[]string{"neg"},
		[]string{"eq"},
		[]string{"gt"},
		[]string{"lt"},
		[]string{"and"},
		[]string{"or"},
		[]string{"not"},
		[]string{"push"},
		[]string{"EOF"},
	}
	expecteds := []VMCommandType{
		C_ARITHMETIC,
		C_ARITHMETIC,
		C_ARITHMETIC,
		C_ARITHMETIC,
		C_ARITHMETIC,
		C_ARITHMETIC,
		C_ARITHMETIC,
		C_ARITHMETIC,
		C_ARITHMETIC,
		C_PUSH,
	}
	p := New(commands)
	for _, expected := range expecteds {
		p.Advance()
		if p.CommandType != expected {
			t.Fatalf("p.CommandType of %v is expected %v. got=%v", p.commands[p.position][0], expected, p.CommandType)
		}
	}
}

func TestArg(t *testing.T) {
	commands := [][]string{
		[]string{"push", "constant", "10"},
	}
	p := New(commands)
	p.Advance()
	if !(p.Arg1 == "constant" && p.Arg2 == 10) {
		t.Fatalf("p.Arg1 and p.Arg2 are expected Arg1=constant, Arg2=10. got=Arg1=%v, Arg2=%v", p.Arg1, p.Arg2)
	}
}
**/
