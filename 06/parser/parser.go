package parser

import (
	"regexp"
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

var EQUAL_REGEXP = regexp.MustCompile("=")
var SEMICOLON_REGEXP = regexp.MustCompile(";")

func New(commands []string) *Parser {
	return &Parser{commands: commands, position: -1}
}

func (p *Parser) curCommand() string {
	return p.commands[p.position]
}

func (p *Parser) HasMoreCommands() bool {
	if p.position > len(p.commands) {
		return false
	}
	peekCommand := p.commands[p.position+1]
	return peekCommand != "EOF"
}

func (p *Parser) Advance() {
	p.position++
}

func (p *Parser) CommandType() int {
	command := p.curCommand()
	if command[0] == '@' {
		return A_COMMAND
	}
	if command[0] == '(' {
		return L_COMMAND
	}
	return C_COMMAND
}

func (p *Parser) Symbol() string {
	return p.curCommand()[1:]
}

func (p *Parser) Dest() string {
	command := p.curCommand()
	index := EQUAL_REGEXP.FindStringIndex(command)
	if len(index) == 0 {
		return ""
	}
	return command[0:index[0]]
}

func (p *Parser) Comp() string {
	command := p.curCommand()
	index := EQUAL_REGEXP.FindStringIndex(command)
	if len(index) != 0 {
		return command[index[1]:]
	}
	index = SEMICOLON_REGEXP.FindStringIndex(command)
	return command[0:index[0]]
}

func (p *Parser) Jump() string {
	command := p.curCommand()
	index := SEMICOLON_REGEXP.FindStringIndex(command)
	if len(index) == 0 {
		return ""
	}
	return command[index[1]:]
}
