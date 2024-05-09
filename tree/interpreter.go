package tree

import (
	"github.com/antlr4-go/antlr/v4"
	gen2 "github.com/ycl2018/pie-go/gen"
)

type PieInterpreter struct {
	InterpreterListener InterpreterListener
}

func NewPieInterpreter() *PieInterpreter {
	return &PieInterpreter{
		InterpreterListener: DefaultInterpreterListener{},
	}
}

func (p *PieInterpreter) Execute(input antlr.CharStream) {
	lexer := gen2.NewPieLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, 0)
	ps := gen2.NewPieParser(tokens)
	programContext := ps.Program()
	if ps.GetError() == nil {
		// 符号&作用域解析
		defVisitor := NewPieDefineVisitor(p)
		defVisitor.Visit(programContext)
		// 执行
		runVisitor := NewRuntimeVisitor(p, defVisitor.Scopes)
		runVisitor.Visit(programContext)
	}
}
