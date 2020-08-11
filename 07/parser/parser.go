package parser

import (
	"bufio"
)

type VMCommandType int

const (
	C_ARITHMETIC VMCommandType = iota
	C_PUSH
	C_POP
	C_LABEL
	C_GOTO
	C_IF
	C_FUNCTION
	C_RETURN
	C_CALL
)

var VMCommandTypeMap = map[string]VMCommandType{
	"add":  C_ARITHMETIC,
	"sub":  C_ARITHMETIC,
	"neg":  C_ARITHMETIC,
	"eq":   C_ARITHMETIC,
	"gt":   C_ARITHMETIC,
	"lt":   C_ARITHMETIC,
	"and":  C_ARITHMETIC,
	"or":   C_ARITHMETIC,
	"not":  C_ARITHMETIC,
	"push": C_PUSH,
}

type Parser struct {
	scanner     *bufio.Scanner
	CommandType VMCommandType
	command     []string
	Command     string
	Arg1        string
	Arg2        int
}

func New(scanner *bufio.Scanner) *Parser {
	return &Parser{scanner: scanner}
}

func (p *Parser) HasMoreCommands() bool {
	return p.scanner.Scan()
}

/**
func (p *Parser) Advance() error {
	p.position = p.nextPosition
	p.nextPosition = p.nextPosition + 1
	p.command = p.commands[p.position]
	p.Command = p.command[0]
	p.CommandType = VMCommandTypeMap[p.Command]
	switch p.CommandType {
	case C_PUSH:
		p.Arg1 = p.command[1]
		i, err := strconv.Atoi(p.command[2])
		if err != nil {
			return err
		}
		p.Arg2 = i
	}
	return nil
}
**/
