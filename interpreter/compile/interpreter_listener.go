package compile

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type InterpreterListener interface {
	Infof(string, ...any)
	Errorf(string, ...any)
	ErrorToken(antlr.Token, string, ...any)
}

type DefaultInterpreterListener struct {
}

func (d DefaultInterpreterListener) Infof(s string, a ...any) {
	fmt.Printf(s+"\n", a...)

}

func (d DefaultInterpreterListener) Errorf(s string, a ...any) {
	fmt.Printf(s+"\n", a...)
}

func (d DefaultInterpreterListener) ErrorToken(token antlr.Token, s string, a ...any) {
	var sb strings.Builder
	lineStr := fmt.Sprintf("<line %d>: ", token.GetLine())
	sb.WriteString(lineStr)
	sb.WriteString(fmt.Sprintf(s, a...))
	fmt.Println(sb.String())
}
