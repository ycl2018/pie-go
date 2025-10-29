package compile

import (
	"github.com/antlr4-go/antlr/v4"
	gen2 "github.com/ycl2018/pie-go/gen"
	"github.com/ycl2018/pie-go/interpreter/tree"
)

type PieCompiler struct {
	InterpreterListener tree.InterpreterListener
}

func NewPieCompiler() *PieCompiler {
	return &PieCompiler{
		InterpreterListener: tree.DefaultInterpreterListener{},
	}
}

func (p *PieCompiler) Execute(input antlr.CharStream) {
	lexer := gen2.NewPieLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, 0)
	ps := gen2.NewPieParser(tokens)
	programContext := ps.Program()
	if ps.GetError() == nil {
		// 符号&作用域解析
		defVisitor := tree.NewPieDefineVisitor()
		defVisitor.Visit(programContext)
		// 执行
		runVisitor := tree.NewRuntimeVisitor(p.InterpreterListener, defVisitor.Scopes)
		runVisitor.Visit(programContext)
	}
}