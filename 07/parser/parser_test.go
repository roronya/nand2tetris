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

func TestAdvance(t *testing.T) {
	commands := "push constant 7\npush constant 8\nadd"
	p := newParser(commands)
	p.HasMoreCommands()
	p.Advance()
	if p.command != "push constant 7" {
		t.Fatalf("p.scanner.Text() is expected \"push constant 7\". got=%v", p.scanner.Text())
	}
	p.HasMoreCommands()
	p.Advance()
	if p.command != "push constant 8" {
		t.Fatalf("p.scanner.Text() is expected \"push constant 8\". got=%v", p.scanner.Text())
	}
	p.HasMoreCommands()
	p.Advance()
	if p.command != "add" {
		t.Fatalf("p.scanner.Text() is expected \"add\". got=%v", p.scanner.Text())
	}
}

func TestCommandType(t *testing.T) {
	commands := `add
sub
net
eq
gt
lt
and
or
not
push`
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
	p := newParser(commands)
	for _, expected := range expecteds {
		p.HasMoreCommands()
		p.Advance()
		if p.CommandType != expected {
			t.Fatalf("p.CommandType of %v is expected %v. got=%v", p.Command, expected, p.CommandType)
		}
	}
}

func TestArg(t *testing.T) {
	commands := "push constant 10"
	p := newParser(commands)
	p.Advance()
	if !(p.Arg1 == "constant" && p.Arg2 == 10) {
		t.Fatalf("p.Arg1 and p.Arg2 are expected Arg1=constant, Arg2=10. got=Arg1=%v, Arg2=%v", p.Arg1, p.Arg2)
	}
}
