package compile

import (
	"strings"

	"github.com/antlr4-go/antlr/v4"
	gen2 "github.com/ycl2018/pie-go/gen"
)

type PieCompiler struct {
	InterpreterListener InterpreterListener
	compileVisitor      *StackCompileVisitor
}

func NewPieCompiler() *PieCompiler {
	return &PieCompiler{
		InterpreterListener: DefaultInterpreterListener{},
	}
}

func (p *PieCompiler) Execute(input antlr.CharStream) {
	lexer := gen2.NewPieLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, 0)
	ps := gen2.NewPieParser(tokens)
	programContext := ps.Program()
	if ps.GetError() == nil {
		// 符号&作用域解析
		defVisitor := NewPieDefineVisitor()
		defVisitor.Visit(programContext)
		// 执行
		p.compileVisitor = NewStackCompileVisitor(defVisitor.Scopes)
		p.compileVisitor.Visit(programContext)
	}
}

func (p *PieCompiler) Dump() string {
	var sb strings.Builder
	for _, f := range p.compileVisitor.AllFuncs {
		sb.WriteString(f.Dump())
	}
	sb.WriteString(p.compileVisitor.MainFunc.Dump())
	return sb.String()
}
