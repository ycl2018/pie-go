package reg

import (
	"testing"

	"github.com/antlr4-go/antlr/v4"
)

func TestInterpreter_Run(t *testing.T) {
	tests := []struct {
		Name string
		Code string
	}{
		{Name: "loop.pcode", Code: `
.def main: args=0, locals=4
; n = 100
	iconst r1, 100
; i = 0
	iconst r2, 0
; get 1 into a register
        iconst r4, 1
; while i<n:
start:
	ilt r2, r1, r3
	brf r3, done
;         i = i + 1
	iadd r4, r2, r2
	br start
done:
; print "looped "+n+" times."
	sconst r4, "done"
	print r4
	halt
`},
		{Name: "struct.pcode", Code: `
; T t
.globals 1

.def main: args=0, locals=2
; t = null
        null r1
	gstore r1, 0     ; t=null
; t = new T()
	struct r1, 2    ; hold t in r1
	gstore r1, 0     ; t points at struct
; t.x = 1
	iconst r2, 1
	fstore r2, r1, 0 ; field 0
; t.y = "foo"
	sconst r2, "foo"
	fstore r2, r1, 1 ; field 1
; print t.x
	fload r2, r1, 0
	print r2
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
