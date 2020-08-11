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
