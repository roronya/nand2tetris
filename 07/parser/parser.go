package parser

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
	commands     [][]string
	position     int
	nextPosition int
	CommandType  VMCommandType
}

func New(commands [][]string) *Parser {
	return &Parser{commands: commands}
}

func (p *Parser) HasMoreCommands() bool {
	return p.commands[p.nextPosition][0] != "EOF"
}

func (p *Parser) Advance() {
	p.position = p.nextPosition
	p.nextPosition = p.nextPosition + 1
	command := p.commands[p.position][0]
	p.CommandType = VMCommandTypeMap[command]
}
