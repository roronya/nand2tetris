package parser

import (
	"regexp"
	"strings"
)

type Parser struct {
	commands     []string
	position     int
	peekPosition int
}

const (
	A_COMMAND = iota
	C_COMMAND
	L_COMMAND
)

var EQUAL_REGEXP = regexp.MustCompile("=")
var SEMICOLON_REGEXP = regexp.MustCompile(";")

func New(input string) *Parser {
	return &Parser{commands: strings.Split(strings.TrimSpace(input), "\n"), position: -1}
}

func (p *Parser) skipComment() {
	p.peekPosition = p.position + 1
	for p.peekPosition < len(p.commands) && (len(p.peekCommand()) == 0 || p.peekCommand()[0:2] == "//") {
		p.peekPosition++
	}
}

func (p *Parser) trimComment(command string) string {
	return strings.Split(command, " ")[0]
}

func (p *Parser) curCommand() string {
	return strings.TrimSpace(p.commands[p.position])
}

func (p *Parser) peekCommand() string {
	return strings.TrimSpace(p.commands[p.peekPosition])
}

func (p *Parser) HasMoreCommands() bool {
	p.skipComment()
	return p.position+1 < len(p.commands)
}

func (p *Parser) Advance() {
	p.position = p.peekPosition
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
	command := p.trimComment(p.curCommand())
	return command[1:]
}

func (p *Parser) Dest() string {
	command := p.trimComment(p.curCommand())
	index := EQUAL_REGEXP.FindStringIndex(command)
	if len(index) == 0 {
		return ""
	}
	return command[0:index[0]]
}

func (p *Parser) Comp() string {
	command := p.trimComment(p.curCommand())
	index := EQUAL_REGEXP.FindStringIndex(command)
	if len(index) != 0 {
		return command[index[1]:]
	}
	index = SEMICOLON_REGEXP.FindStringIndex(command)
	return command[0:index[0]]
}

func (p *Parser) Jump() string {
	command := p.trimComment(p.curCommand())
	index := SEMICOLON_REGEXP.FindStringIndex(command)
	if len(index) == 0 {
		return ""
	}
	return command[index[1]:]
}
