package parser

import (
	"io/ioutil"
	"log"
	"strings"
)

type Parser struct {
	commands [][]string
}

func New(fileName string) *Parser {
	vmFunc, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	input := string(vmFunc)
	rows := strings.Split(input, "\n")
	commands := [][]string{}
	for _, row := range rows {
		command := []string{}
		elements := strings.Split(row, " ")
		for _, element := range elements {
			command = append(command, element)
		}
		commands = append(commands, command)
	}
	return &Parser{commands: commands}
}
