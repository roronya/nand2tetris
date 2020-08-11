package codewriter

import (
	"bufio"
	"bytes"

	"github.com/roronya/nand2tetris/07/parser"
)

type CodeWriter struct {
	filename string
	writer   *bufio.Writer
	buffer   *bytes.Buffer
}

func New(filename string) *CodeWriter {
	cw := &CodeWriter{}
	cw.setFileName(filename)
	return cw
}

func (cw *CodeWriter) setFileName(filename string) {
	cw.filename = filename
	cw.buffer = new(bytes.Buffer)
	cw.writer = bufio.NewWriter(cw.buffer)
}

func (cw *CodeWriter) writePushPop(command parser.VMCommandType, segment string, index int) {
	switch command {
	case parser.C_PUSH:
		// アセンブリでsegmentに書き込むようなコードをwriterで書き込む
		// M=index
		// D=M
		// M=segment
		// M=D
	}
}
