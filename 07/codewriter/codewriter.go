package codewriter

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/roronya/nand2tetris/07/parser"
)

var registerMap = map[string]string{
	"local":    "LCL",
	"argument": "ARG",
	"this":     "THIS",
	"that":     "THAT",
}

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
		cw.push(segment, index)
	case parser.C_POP:
		cw.pop(segment, index)
	}
}

func (cw *CodeWriter) push(segment string, index int) {
	switch segment {
	case "constant":
		cw.writeCodes([]string{
			fmt.Sprintf("@%d", index),
			"D=A",
		})
	case "temp":
		cw.writeCodes([]string{
			fmt.Sprintf("@R%d", 5+index),
			"D=M",
		})
	default:
		cw.writeCodes([]string{
			fmt.Sprintf("@%s", registerMap[segment]),
			"A=M",
		})
		for i := 0; i < index; i++ {
			cw.writeCode("A=A+1")
		}
		cw.writeCode("D=M")
	}
	cw.writePush()
}

func (cw *CodeWriter) pop(segment string, index int) {
	cw.writePop()
	cw.writeCode("D=M")
	if segment == "temp" {
		cw.writeCode(fmt.Sprintf("@R%d", 5+index))
	} else {
		cw.writeCodes([]string{
			fmt.Sprintf("@%s", registerMap[segment]),
			"A=M",
		})
		for i := 0; i < index; i++ {
			cw.writeCode("A=A+1")
		}
	}
	cw.writeCodes([]string{
		"M=D",
	})
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
	// M=x, D=y
	var calc string
	if command == "add" {
		calc = "M+D"
	} else if command == "sub" {
		calc = "M-D"
	} else if command == "and" {
		calc = "M&D"
	} else {
		calc = "M|D"
	}
	cw.writePoppop()
	cw.writeCode(fmt.Sprintf("D=%s", calc))
	cw.writePush()
}

func (cw *CodeWriter) writeNeg() {
	cw.writePop()
	cw.writeCode("D=-M")
	cw.writePush()
}

func (cw *CodeWriter) writeNot() {
	cw.writePop()
	cw.writeCode("D=!M")
	cw.writePush()
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
	cw.writePoppop()
	cw.writeCodes([]string{
		"D=M-D", // D=x-y
		fmt.Sprintf("@%s", trueLabel),
		fmt.Sprintf("D;%s", commandType), // x == y
		// falseのとき
		"D=0",
	})
	cw.writePush()
	cw.writeCodes([]string{
		fmt.Sprintf("@%s", endLabel),
		"0;JMP",
		fmt.Sprintf("(%s)", trueLabel),
		// trueのとき
		"D=-1",
	})
	cw.writePush()
	cw.writeCode(fmt.Sprintf("(%s)", endLabel))
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

// Dの値をスタックの一番上に積んでSPをインクリメントする
// 事後条件
// A: SPのアドレス = 0
// M: スタックの一番上のアドレス
// D: 呼び出す前と同じ
func (cw *CodeWriter) writePush() {
	cw.writeCodes([]string{
		"@SP",
		"A=M",
		"M=D",
		"@SP",
		"M=M+1",
	})
}

// スタックの一番上をDにセットしてSPをデクリメントする
// 事後条件
// A: スタックの一番上のアドレス
// M: スタックの一番上の値
// D: 呼び出す前と同じ
func (cw *CodeWriter) writePop() {
	cw.writeCodes([]string{
		"@SP",
		"M=M-1",
		"A=M",
	})
}

// 2回popする
// 事後条件
// A: 呼び出し後のスタックの一番上のアドレス
// M: 呼び出し後のスタックの一番上の値
// D: 呼び出す前のスタックの一番上だった値
func (cw *CodeWriter) writePoppop() {
	cw.writePop()
	cw.writeCode("D=M")
	cw.writePop()
}
