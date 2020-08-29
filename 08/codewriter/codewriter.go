package codewriter

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/roronya/nand2tetris/08/parser"
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

func (cw *CodeWriter) WriteLabel(name string) {
	cw.writeCode(fmt.Sprintf("(%s)", name))
}

func (cw *CodeWriter) WriteIfGoto(label string) {
	cw.writePop()
	cw.writeCodes([]string{
		"D=M",
		fmt.Sprintf("@%s", label),
		"D;JGT",
	})
}

func (cw *CodeWriter) WriteGoto(label string) {
	cw.writeCodes([]string{
		fmt.Sprintf("@%s", label),
		"D;JMP",
	})
}

func (cw *CodeWriter) WriteFunction(name string, localVariableCount int) {
	cw.WriteLabel(name)
	/**
	callの処理っぽいのでやらないでみる
	// 呼び出す前のLCL,ARG,THIS,THATをpushしておき、return時に巻き戻せるようにしておく
	cw.writeCodes([]string{"@LCL", "D=M"})
	cw.writePush()
	cw.writeCodes([]string{"@ARG", "D=M"})
	cw.writePush()
	cw.writeCodes([]string{"@THIS", "D=M"})
	cw.writePush()
	cw.writeCodes([]string{"@THAT", "D=M"})
	cw.writePush()
	**/
	// 呼び出す前のLCL,ARG,THIS,THATをpushした直後のSPのアドレスをLCLのポインタとする
	cw.writeCodes([]string{
		"@SP",
		"D=M",
		"@LCL",
		"M=D",
	})
	// localVariableCount回初期化しておく
	for i := 0; i < localVariableCount; i++ {
		cw.writeCodes([]string{
			"@0",
			"D=A",
		})
		cw.writePush()
	}
}

func (cw *CodeWriter) WriteReturn() {
	// 演算結果であるスタックのヘッドを、取り出してR13に入れておく
	cw.writePop()
	cw.writeCodes([]string{"D=M", "@R13", "M=D"})
	// このあとの処理でARGを上書きしてしまうが、この時点でのARGの値を後で使いたいのでR14に入れておく
	cw.writeCodes([]string{"@ARG", "D=M", "@R14", "M=D"})
	// SPをこの時点でのLCLの位置まで戻す
	cw.writeCodes([]string{"@LCL", "D=M", "@SP", "M=D"})
	// このSPからpopしていくとTHAT, THIS, ARG, LCLの順でcallerの状態が手に入るのでセットしていく
	cw.writePop()
	cw.writeCodes([]string{"D=M", "@THAT", "M=D"})
	cw.writePop()
	cw.writeCodes([]string{"D=M", "@THIS", "M=D"})
	cw.writePop()
	cw.writeCodes([]string{"D=M", "@ARG", "M=D"})
	cw.writePop()
	cw.writeCodes([]string{"D=M", "@LCL", "M=D"})
	// R14にいれたARGの位置までSPを戻す
	cw.writeCodes([]string{"@R14", "D=M", "@SP", "M=D"})
	// R13にいれたスタックのヘッドをpushする
	cw.writeCodes([]string{"@R13", "D=M"})
	cw.writePush()
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
	case "pointer":
		cw.writeCodes([]string{
			fmt.Sprintf("@R%d", 3+index),
			"D=M",
		})
	case "static":
		cw.writeCodes([]string{
			fmt.Sprintf("@%s.%d", cw.filename, index),
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
	} else if segment == "pointer" {
		cw.writeCode(fmt.Sprintf("@R%d", 3+index))
	} else if segment == "static" {
		cw.writeCode(fmt.Sprintf("@%s.%d", cw.filename, index))
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
