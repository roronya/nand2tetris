package parser

import (
	"testing"
)

func TestNew(t *testing.T) {
	commands := []string{"@Xxx", "M=1"}
	expected := []string{"@Xxx", "M=1"}
	p := New(commands)
	if p.commands[0] != expected[0] {
		t.Fatalf("p.commands[0] is exptected %s. got=%s", expected[0], p.commands[0])
	}
	if p.commands[1] != expected[1] {
		t.Fatalf("p.commands[1] is exptected %s. got=%s", expected[1], p.commands[1])
	}
	if p.position != -1 {
		t.Fatalf("p.position is exptected 0. got=%d", p.position)
	}
	if p.peekPosition != 0 {
		t.Fatalf("p.peekPosition is exptected 0. got=%d", p.peekPosition)
	}
}

func TestHasMoreCommands(t *testing.T) {
	commands := []string{"@Xxx", "M=1"}
	p := New(commands)
	if !p.HasMoreCommands() {
		t.Fatalf("p.HasMoreCommands is expected false. got=true.")
	}
	p.position = 1
	if p.HasMoreCommands() {
		t.Fatalf("p.HasMoreCommands is expected true. got=false.")
	}
}

func TestParse(t *testing.T) {
	commands := []string{"@1", "@16383", "@0", "@-1"}
	expecteds := []string{
		"1",
		"16383",
		"0",
		"-1",
	}
	p := New(commands)
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

func TestDest(t *testing.T) {
	commands := []string{"M=1"}
	expectedDest := "M"
	expectedComp := "1"
	p := New(commands)
	p.Advance()
	actualDest := p.Dest()
	actualComp := p.Comp()
	if actualDest != expectedDest {
		t.Fatalf("p.Dest() is expected %s. got=%s", expectedDest, actualDest)
	}
	if actualComp != expectedComp {
		t.Fatalf("p.Comp() is expected %s. got=%s", expectedComp, actualComp)
	}
}
func TestJump(t *testing.T) {
	commands := []string{"D;JGT"}
	expectedJump := "JGT"
	expectedComp := "D"
	p := New(commands)
	p.Advance()
	actualJump := p.Jump()
	actualComp := p.Comp()
	if actualJump != expectedJump {
		t.Fatalf("p.Dest() is expected %s. got=%s", expectedJump, actualJump)
	}
	if actualComp != expectedComp {
		t.Fatalf("p.Dest() is expected %s. got=%s", expectedComp, actualComp)
	}
}

func TestC(t *testing.T) {
	commands := []string{"D=1"}
	p := New(commands)
	if p.HasMoreCommands() {
		p.Advance()
	}
	p.Dest()

}

func TestAdd(t *testing.T) {
	commands := []string{"@2", "D=A", "@3", "D=D+A", "@0", "M=D"}
	p := New(commands)
	if p.HasMoreCommands() {
		p.Advance()
	}
}
