package compile

import (
	"bytes"
	"errors"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	gen2 "github.com/ycl2018/pie-go/gen"
)

const (
	Version     uint16 = 1
	MagicNumber int32  = 0x25504446 // %PIE
)

type PieCompiler struct {
	InterpreterListener *DefaultInterpreterListener
	compileVisitor      *StackCompileVisitor
	ErrWriter           *bytes.Buffer
}

func NewPieCompiler() *PieCompiler {
	errWriter := &bytes.Buffer{}
	return &PieCompiler{
		ErrWriter: errWriter,
		InterpreterListener: &DefaultInterpreterListener{
			Writer:      os.Stdout,
			ErrorWriter: errWriter,
		},
	}
}

func (p *PieCompiler) Compile(input antlr.CharStream) ([]byte, error) {
	lexer := gen2.NewPieLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, 0)
	ps := gen2.NewPieParser(tokens)
	programContext := ps.Program()
	if ps.GetError() != nil {
		return nil, errors.New(ps.GetError().GetMessage())
	}
	// 符号&作用域解析
	log := p.InterpreterListener
	defVisitor := NewPieDefineVisitor(log)
	defVisitor.Visit(programContext)
	// 执行
	p.compileVisitor = NewStackCompileVisitor(defVisitor.Scopes, defVisitor.GlobalScope, log)
	p.compileVisitor.Visit(programContext)
	if p.ErrWriter.String() != "" {
		return nil, errors.New(p.ErrWriter.String())
	}
	assembler := NewAssembler(defVisitor.GlobalScope, p.compileVisitor.MainFunc, p.compileVisitor.AllFuncs, p.compileVisitor.ToBeFilled)
	return assembler.Assemble(), nil
}

func (p *PieCompiler) Dump() string {
	var sb strings.Builder
	sb.WriteString("Dump:\n")
	// 常量池
	sb.WriteString(p.compileVisitor.GlobalScope.Dump())
	sb.WriteString("\nCode:\n")
	sb.WriteString(p.compileVisitor.MainFunc.Dump())
	for _, f := range p.compileVisitor.AllFuncs {
		sb.WriteString(f.Dump())
	}
	return sb.String()
}
