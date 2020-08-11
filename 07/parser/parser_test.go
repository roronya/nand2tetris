package parser

import (
	"testing"
)

func TestHasMoreCommands(t *testing.T) {
	commands := [][]string{
		[]string{"push", "constant", "7"},
		[]string{"push", "constant", "8"},
		[]string{"add"},
		[]string{"EOF"},
	}
	p := New(commands)
	if !p.HasMoreCommands() {
		t.Fatalf("p.HasMoreCommands is expected true. got=false.")
	}
	p.nextPosition = 3
	if p.HasMoreCommands() {
		t.Fatalf("p.HasMoreCommands is expected false. got=true.")
	}
}

func TestHasMoreCommands_空ファイルのとき(t *testing.T) {
	commands := [][]string{
		[]string{"EOF"},
	}
	p := New(commands)
	if p.HasMoreCommands() {
		t.Fatalf("p.HasMoreCommands is expected false. got=true.")
	}
}

func TestAdvance(t *testing.T) {
	commands := [][]string{
		[]string{"push", "constant", "7"},
		[]string{"push", "constant", "8"},
		[]string{"add"},
		[]string{"EOF"},
	}
	p := New(commands)
	p.Advance()
	if !(p.position == 0 && p.nextPosition == 1) {
		t.Fatalf("p.Advance() is expected setting position = 0, nextPosition = 1. got=position=%d, nextPosition=%d.", p.position, p.nextPosition)
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
