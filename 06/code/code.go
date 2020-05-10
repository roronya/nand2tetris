package code

import (
	"fmt"
	"strings"
)

func Dest(dest string) string {
	a := 0
	d := 0
	m := 0
	if strings.Contains(dest, "A") {
		a = 1
	}
	if strings.Contains(dest, "D") {
		d = 1
	}
	if strings.Contains(dest, "M") {
		m = 1
	}
	return fmt.Sprintf("%d%d%d", a, d, m)
}

func Comp(comp string) string {
	switch comp {
	case "0":
		return "0101010"
	case "1":
		return "0111111"
	case "-1":
		return "0111010"
	case "D":
		return "0001100"
	case "A":
		return "0110000"
	case "!D":
		return "0001101"
	case "!A":
		return "0110001"
	case "-D":
		return "0001111"
	case "-A":
		return "0110011"
	case "D+1":
		return "0011111"
	case "A+1":
		return "0110111"
	case "D-1":
		return "0001110"
	case "A-1":
		return "0110010"
	case "D+A":
		return "0000010"
	case "D-A":
		return "0010011"
	case "A-D":
		return "0000111"
	case "D&A":
		return "0000000"
	case "D|A":
		return "0010101"
	case "M":
		return "1110000"
	case "!M":
		return "1110001"
	case "-M":
		return "1110111"
	case "M+1":
		return "1110111"
	case "M-1":
		return "1110010"
	case "D+M":
		return "1000010"
	case "D-M":
		return "1010011"
	case "M-D":
		return "1000111"
	case "D&M":
		return "1000000"
	case "D|M":
		return "1010101"
	}
	panic(fmt.Sprintf("unknown comp %s", comp))
}

func Jump(jump string) string {
	switch jump {
	case "":
		return "000"
	case "JGT":
		return "001"
	case "JEQ":
		return "010"
	case "JGE":
		return "011"
	case "JLT":
		return "100"
	case "JNE":
		return "101"
	case "JLE":
		return "110"
	case "JMP":
		return "111"
	}
	panic(fmt.Sprintf("unknown jump %s", jump))
}
