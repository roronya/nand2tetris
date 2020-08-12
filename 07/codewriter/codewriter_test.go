package codewriter

import (
	"testing"
)

func TestNew(t *testing.T) {
	cw := New("filename")
	if cw.filename != "filename" {
		t.Fatalf("cw.filename is expected \"filename\". got=%#v", cw.filename)
	}
	if cw.buffer.String() != "" {
		t.Fatalf("cw.buffer is expected empty string. got=%#v", cw.buffer.String())
	}
}

func TestWritePush(t *testing.T) {
	cw := New("filename")
	cw.writePush("constant", 0)
	expected := "@0\nD=A\n@SP\nM=D\nA=A+1\n"
	actual := cw.buffer.String()
	if actual != expected {
		t.Fatalf("cw.buffer.String() is %#v . got=%#v", expected, actual)
	}
}

func TestWriteAdd(t *testing.T) {
	cw := New("filename")
	cw.writeAdd()
	expected := "@SP\nD=M\nA=A-1\nM=M+D\n"
	actual := cw.buffer.String()
	if actual != expected {
		t.Fatalf("cw.buffer.String() is %#v . got=%#v", expected, actual)
	}
}
