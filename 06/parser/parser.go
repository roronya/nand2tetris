package parser

import (
	"strings"
)

type Parser struct {
	commands []string
	position int
}

const (
	A_COMMAND = iota
	C_COMMAND
	L_COMMAND
)

func New(input string) *Parser {
	splited := strings.Split(strings.TrimSpace(input), "\n")
	return &Parser{splited, -1}
}

func (p *Parser) curCommand() string {
	return p.commands[p.position]
}

func (p *Parser) HasMoreCommands() bool {
	return p.position+1 < len(p.commands)
}

func (p *Parser) Advance() {
	p.position++
}

func (p *Parser) CommandType() int {
	command := p.curCommand()
	if command[0] == '@' {
		return A_COMMAND
	}
	// TODO: CとLの判定
	return C_COMMAND
}

func (p *Parser) Symbol() string {
	command := p.curCommand()
	return command[1:]
}
