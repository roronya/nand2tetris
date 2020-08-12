package codewriter

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/roronya/nand2tetris/07/parser"
)

var segmentMap = map[string]string{}

type CodeWriter struct {
	filename string
	buffer   *bytes.Buffer
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
		cw.buffer.WriteString(fmt.Sprintf("@%d", index))
		cw.buffer.WriteString("\n")
		cw.buffer.WriteString("D=A")
		cw.buffer.WriteString("\n")
		cw.buffer.WriteString("@SP")
		cw.buffer.WriteString("\n")
		cw.buffer.WriteString("M=D")
		cw.buffer.WriteString("\n")
		cw.buffer.WriteString("A=A+1")
		cw.buffer.WriteString("\n")
	}
}

func (cw *CodeWriter) WriteArithmetic(command string) {
	switch command {
	case "add":
		cw.writeAdd()
	}
}

func (cw *CodeWriter) writeAdd() {
	cw.buffer.WriteString("@SP")
	cw.buffer.WriteString("\n")
	cw.buffer.WriteString("D=M")
	cw.buffer.WriteString("\n")
	cw.buffer.WriteString("A=A-1")
	cw.buffer.WriteString("\n")
	cw.buffer.WriteString("M=M+D")
	cw.buffer.WriteString("\n")
}

func (cw *CodeWriter) Close() {
	ioutil.WriteFile(cw.filename, cw.buffer.Bytes(), 0664)
}
