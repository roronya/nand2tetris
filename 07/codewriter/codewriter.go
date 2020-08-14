package codewriter

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/roronya/nand2tetris/07/parser"
)

var segmentMap = map[string]string{}

type CodeWriter struct {
	filename    string
	buffer      *bytes.Buffer
	labelNumber int
}

func New(filename string) *CodeWriter {
	cw := &CodeWriter{}
	cw.setFileName(filename)
	return cw
}

func (cw *CodeWriter) setFileName(filename string) {
	cw.filename = filename
	cw.buffer = bytes.NewBuffer([]byte(""))
}

func (cw *CodeWriter) WritePushPop(command parser.VMCommandType, segment string, index int) {
	switch command {
	case parser.C_PUSH:
		cw.writePush(segment, index)
	}
}

func (cw *CodeWriter) writePush(segment string, index int) {
	switch segment {
	case "constant":
		cw.writeCodes([]string{
			fmt.Sprintf("@%d", index),
			"D=A",
			"@SP",
			"M=D",
			"A=A+1",
		})
	}
}

func (cw *CodeWriter) WriteArithmetic(command string) {
	switch command {
	case "add":
		fallthrough
	case "sub":
		fallthrough
	case "and":
		fallthrough
	case "or":
		cw.writeCalc(command)
	case "eq":
		fallthrough
	case "gt":
		fallthrough
	case "lt":
		cw.writeCmp(command)
	case "not":
		cw.writeNot()
	case "neg":
		cw.writeNeg()
	}
}

func (cw *CodeWriter) writeCalc(command string) {
	var commandType string
	if command == "add" {
		commandType = "M+D"
	} else if command == "sub" {
		commandType = "M-D"
	} else if command == "and" {
		commandType = "M&D"
	} else {
		commandType = "M|D"
	}
	cw.writeCodes([]string{
		"@SP",
		"D=M",
		"A=A-1",
		fmt.Sprintf("M=%s", commandType),
	})
}

func (cw *CodeWriter) writeNeg() {
	cw.writeCodes([]string{
		"@SP",
		"M=-M",
	})
}

func (cw *CodeWriter) writeNot() {
	cw.writeCodes([]string{
		"@SP",
		"M=!M",
	})
}

func (cw *CodeWriter) writeCmp(command string) {
	var commandType string
	if command == "eq" {
		commandType = "JEQ"
	} else if command == "gt" {
		commandType = "JGT"
	} else {
		commandType = "JLT"
	}
	trueLabel := cw.generateNewLabel()
	endLabel := cw.generateNewLabel()
	cw.writeCodes([]string{
		"@SP",
		"D=M",
		"A=A-1",
		"D=M-D",
		trueLabel,
		fmt.Sprintf("D;%s", commandType),
		"M=0",
		endLabel,
		"0;JMP",
		fmt.Sprintf("(%s)", trueLabel),
		"M=-1",
		endLabel,
		"0;JMP",
		fmt.Sprintf("(%s)", endLabel),
	})
}

func (cw *CodeWriter) Close() {
	ioutil.WriteFile(cw.filename, cw.buffer.Bytes(), 0664)
}

func (cw *CodeWriter) writeCode(code string) {
	cw.buffer.WriteString(code)
	cw.buffer.WriteString("\n")
}

func (cw *CodeWriter) writeCodes(codes []string) {
	for _, code := range codes {
		cw.writeCode(code)
	}
}

func (cw *CodeWriter) generateNewLabel() string {
	newLabel := fmt.Sprintf("LABEL%d", cw.labelNumber)
	cw.labelNumber++
	return newLabel
}
