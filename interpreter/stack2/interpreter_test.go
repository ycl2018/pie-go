package stack

import (
	"testing"

	"github.com/antlr4-go/antlr/v4"
)

func TestInterpreter_Run(t *testing.T) {
	tests := []struct {
		Name string
		Code string
	}{
		{Name: "factorial.pcode", Code: `
.def fact: args=1, locals=0
;	if n < 2 return 1
	load 0
	iconst 2
	ilt
	brf cont
	iconst 1
	ret
cont:
;	return n * fact(n-1)
	load 0
	load 0
	iconst 1
	isub
	call fact()
	imul
	ret

.def main: args=0, locals=0
; print fact(10)
	iconst 10
	call fact()
	print
	halt
`},
		{Name: "loop.pcode", Code: `
.globals 2; n, i
; n = 10000
	iconst 10000
	gstore 0
; i = 0
	iconst 0
	gstore 1
; while i<n:
start:
	gload 1
	gload 0
	ilt
	brf done
;         i = i + 1
	gload 1
	iconst 1
	iadd
	gstore 1
	br start
done:
; print "looped "+n+" times."
	sconst "done"
	print
	halt
`},
		{Name: "struct.pcode", Code: `
; T t
.globals 1

.def main: args=0, locals=0
; t = new T()
	struct 2
	gstore 0
; t.x = 1
	iconst 1
	gload 0
	fstore 0
; t.y = "foo"
	sconst "foo"
	gload 0
	fstore 1
; print t.x
	gload 0
	fload 0
	print
	halt
`},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			i := NewInterpreter(antlr.NewInputStream(test.Code),
				WithEnableTrace(), WithEnableDump(), WithDisassemble())
			i.Run()
		})
	}
}
