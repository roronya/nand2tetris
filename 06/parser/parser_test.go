package parser

import (
	"testing"
)

func TestNew(t *testing.T) {
	input := `@Xxx
M=1`
	expected := []string{"@Xxx", "M=1"}
	p := New(input)
	if p.commands[0] != expected[0] {
		t.Fatalf("p.commands[0] is exptected %s. got=%s", expected[0], p.commands[0])
	}
	if p.commands[1] != expected[1] {
		t.Fatalf("p.commands[1] is exptected %s. got=%s", expected[1], p.commands[1])
	}
	if p.position != -1 {
		t.Fatalf("p.position is exptected 0. got=%d", p.position)
	}
}

func TestHasMoreCommands(t *testing.T) {
	input := `@Xxx
M=1`
	p := New(input)
	if !p.HasMoreCommands() {
		t.Fatalf("p.HasMoreCommands is expected false. got=true.")
	}
	p.position = 1
	if p.HasMoreCommands() {
		t.Fatalf("p.HasMoreCommands is expected true. got=false.")
	}
}

func TestParse(t *testing.T) {
	input := `@1
@16383
@0
@-1`
	expecteds := []string{
		"1",
		"16383",
		"0",
		"-1",
	}
	p := New(input)
	for _, expected := range expecteds {
		if !p.HasMoreCommands() {
			t.Fatalf("p.HadMoreCommands() is expected true. got=false")
		}
		p.Advance()
		if p.Symbol() != expected {
			t.Fatalf("p.Symbol() is expected %s. got=%s", expected, p.Symbol())
		}
	}
}
