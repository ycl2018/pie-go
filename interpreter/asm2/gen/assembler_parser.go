// Code generated from /Users/bytedance/go/src/github.com/ycl2018/pie-go/interpreter/asm2/Assembler.g4 by ANTLR 4.13.2. DO NOT EDIT.

package gen // Assembler
import (
	"fmt"
	"strconv"
  	"sync"

	"github.com/antlr4-go/antlr/v4"
)



type AsmGenerator interface {
    Gen(instrs ...antlr.Token)
    DefineConst(tag string, value string)
    DefineStruct(name string, fields []antlr.Token)
    CheckForUnresolvedReferences()
    DefineFunction(id antlr.Token, nargs, nlocals int)
    DefineDataSize(n int)
    DefineLabel(id antlr.Token)
    DefineFuncStruct(funcName antlr.Token, structName string, fields []antlr.Token)
}


// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}


type AssemblerParser struct {
	*antlr.BaseParser
	// Grammar author supplied members of the instance struct
	AsmGenerator
}

var AssemblerParserStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  LiteralNames           []string
  SymbolicNames          []string
  RuleNames              []string
  PredictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func assemblerParserInit() {
  staticData := &AssemblerParserStaticData
  staticData.LiteralNames = []string{
    "", "'.globals'", "'.consts'", "'{'", "','", "'}'", "'.def'", "':'", 
    "'args'", "'='", "'locals'", "'f'", "'s'", "'struct'",
  }
  staticData.SymbolicNames = []string{
    "", "", "", "", "", "", "", "", "", "", "", "TAG_FLOAT", "TAG_STRING", 
    "TAG_STUCT", "REG", "ID", "FUNC", "INT", "CHAR", "STRING", "FLOAT", 
    "WS", "NEWLINE",
  }
  staticData.RuleNames = []string{
    "program", "globals", "constantPool", "constantPoolEntry", "structDecl", 
    "functionDeclaration", "instr", "operand", "label",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 22, 151, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 3, 0, 20, 8, 0, 
	1, 0, 3, 0, 23, 8, 0, 1, 0, 1, 0, 1, 0, 1, 0, 4, 0, 29, 8, 0, 11, 0, 12, 
	0, 30, 1, 0, 1, 0, 1, 1, 5, 1, 36, 8, 1, 10, 1, 12, 1, 39, 9, 1, 1, 1, 
	1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 5, 2, 47, 8, 2, 10, 2, 12, 2, 50, 9, 2, 1, 
	2, 1, 2, 1, 2, 1, 2, 5, 2, 56, 8, 2, 10, 2, 12, 2, 59, 9, 2, 1, 3, 1, 3, 
	1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 73, 8, 
	3, 10, 3, 12, 3, 76, 9, 3, 3, 3, 78, 8, 3, 1, 3, 1, 3, 3, 3, 82, 8, 3, 
	1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 90, 8, 4, 10, 4, 12, 4, 93, 9, 
	4, 3, 4, 95, 8, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 
	1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 4, 5, 111, 8, 5, 11, 5, 12, 5, 112, 3, 5, 
	115, 8, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 
	1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 
	1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 143, 8, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 
	1, 8, 1, 8, 0, 0, 9, 0, 2, 4, 6, 8, 10, 12, 14, 16, 0, 1, 1, 0, 14, 20, 
	161, 0, 19, 1, 0, 0, 0, 2, 37, 1, 0, 0, 0, 4, 48, 1, 0, 0, 0, 6, 81, 1, 
	0, 0, 0, 8, 83, 1, 0, 0, 0, 10, 98, 1, 0, 0, 0, 12, 142, 1, 0, 0, 0, 14, 
	144, 1, 0, 0, 0, 16, 146, 1, 0, 0, 0, 18, 20, 3, 2, 1, 0, 19, 18, 1, 0, 
	0, 0, 19, 20, 1, 0, 0, 0, 20, 22, 1, 0, 0, 0, 21, 23, 3, 4, 2, 0, 22, 21, 
	1, 0, 0, 0, 22, 23, 1, 0, 0, 0, 23, 28, 1, 0, 0, 0, 24, 29, 3, 10, 5, 0, 
	25, 29, 3, 12, 6, 0, 26, 29, 3, 16, 8, 0, 27, 29, 5, 22, 0, 0, 28, 24, 
	1, 0, 0, 0, 28, 25, 1, 0, 0, 0, 28, 26, 1, 0, 0, 0, 28, 27, 1, 0, 0, 0, 
	29, 30, 1, 0, 0, 0, 30, 28, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 32, 1, 
	0, 0, 0, 32, 33, 6, 0, -1, 0, 33, 1, 1, 0, 0, 0, 34, 36, 5, 22, 0, 0, 35, 
	34, 1, 0, 0, 0, 36, 39, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 37, 38, 1, 0, 0, 
	0, 38, 40, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 40, 41, 5, 1, 0, 0, 41, 42, 
	5, 17, 0, 0, 42, 43, 5, 22, 0, 0, 43, 44, 6, 1, -1, 0, 44, 3, 1, 0, 0, 
	0, 45, 47, 5, 22, 0, 0, 46, 45, 1, 0, 0, 0, 47, 50, 1, 0, 0, 0, 48, 46, 
	1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 51, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 
	51, 52, 5, 2, 0, 0, 52, 53, 5, 17, 0, 0, 53, 57, 5, 22, 0, 0, 54, 56, 3, 
	6, 3, 0, 55, 54, 1, 0, 0, 0, 56, 59, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 57, 
	58, 1, 0, 0, 0, 58, 5, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0, 60, 61, 5, 11, 0, 
	0, 61, 62, 5, 20, 0, 0, 62, 82, 6, 3, -1, 0, 63, 64, 5, 12, 0, 0, 64, 65, 
	5, 19, 0, 0, 65, 82, 6, 3, -1, 0, 66, 67, 5, 13, 0, 0, 67, 68, 5, 15, 0, 
	0, 68, 77, 5, 3, 0, 0, 69, 74, 5, 15, 0, 0, 70, 71, 5, 4, 0, 0, 71, 73, 
	5, 15, 0, 0, 72, 70, 1, 0, 0, 0, 73, 76, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 
	74, 75, 1, 0, 0, 0, 75, 78, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0, 77, 69, 1, 
	0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 80, 5, 5, 0, 0, 80, 
	82, 6, 3, -1, 0, 81, 60, 1, 0, 0, 0, 81, 63, 1, 0, 0, 0, 81, 66, 1, 0, 
	0, 0, 82, 7, 1, 0, 0, 0, 83, 84, 5, 13, 0, 0, 84, 85, 5, 15, 0, 0, 85, 
	94, 5, 3, 0, 0, 86, 91, 5, 15, 0, 0, 87, 88, 5, 4, 0, 0, 88, 90, 5, 15, 
	0, 0, 89, 87, 1, 0, 0, 0, 90, 93, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 91, 92, 
	1, 0, 0, 0, 92, 95, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 94, 86, 1, 0, 0, 0, 
	94, 95, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 97, 5, 5, 0, 0, 97, 9, 1, 0, 
	0, 0, 98, 99, 5, 6, 0, 0, 99, 100, 5, 15, 0, 0, 100, 101, 5, 7, 0, 0, 101, 
	102, 5, 8, 0, 0, 102, 103, 5, 9, 0, 0, 103, 104, 5, 17, 0, 0, 104, 105, 
	5, 4, 0, 0, 105, 106, 5, 10, 0, 0, 106, 107, 5, 9, 0, 0, 107, 108, 5, 17, 
	0, 0, 108, 114, 5, 22, 0, 0, 109, 111, 3, 8, 4, 0, 110, 109, 1, 0, 0, 0, 
	111, 112, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 
	115, 1, 0, 0, 0, 114, 110, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 116, 
	1, 0, 0, 0, 116, 117, 6, 5, -1, 0, 117, 11, 1, 0, 0, 0, 118, 119, 5, 15, 
	0, 0, 119, 120, 5, 22, 0, 0, 120, 143, 6, 6, -1, 0, 121, 122, 5, 15, 0, 
	0, 122, 123, 3, 14, 7, 0, 123, 124, 5, 22, 0, 0, 124, 125, 6, 6, -1, 0, 
	125, 143, 1, 0, 0, 0, 126, 127, 5, 15, 0, 0, 127, 128, 3, 14, 7, 0, 128, 
	129, 5, 4, 0, 0, 129, 130, 3, 14, 7, 0, 130, 131, 5, 22, 0, 0, 131, 132, 
	6, 6, -1, 0, 132, 143, 1, 0, 0, 0, 133, 134, 5, 15, 0, 0, 134, 135, 3, 
	14, 7, 0, 135, 136, 5, 4, 0, 0, 136, 137, 3, 14, 7, 0, 137, 138, 5, 4, 
	0, 0, 138, 139, 3, 14, 7, 0, 139, 140, 5, 22, 0, 0, 140, 141, 6, 6, -1, 
	0, 141, 143, 1, 0, 0, 0, 142, 118, 1, 0, 0, 0, 142, 121, 1, 0, 0, 0, 142, 
	126, 1, 0, 0, 0, 142, 133, 1, 0, 0, 0, 143, 13, 1, 0, 0, 0, 144, 145, 7, 
	0, 0, 0, 145, 15, 1, 0, 0, 0, 146, 147, 5, 15, 0, 0, 147, 148, 5, 7, 0, 
	0, 148, 149, 6, 8, -1, 0, 149, 17, 1, 0, 0, 0, 15, 19, 22, 28, 30, 37, 
	48, 57, 74, 77, 81, 91, 94, 112, 114, 142,
}
  deserializer := antlr.NewATNDeserializer(nil)
  staticData.atn = deserializer.Deserialize(staticData.serializedATN)
  atn := staticData.atn
  staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
  decisionToDFA := staticData.decisionToDFA
  for index, state := range atn.DecisionToState {
    decisionToDFA[index] = antlr.NewDFA(state, index)
  }
}

// AssemblerParserInit initializes any static state used to implement AssemblerParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAssemblerParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AssemblerParserInit() {
  staticData := &AssemblerParserStaticData
  staticData.once.Do(assemblerParserInit)
}

// NewAssemblerParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAssemblerParser(input antlr.TokenStream) *AssemblerParser {
	AssemblerParserInit()
	this := new(AssemblerParser)
	this.BaseParser = antlr.NewBaseParser(input)
  staticData := &AssemblerParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Assembler.g4"

	return this
}


// AssemblerParser tokens.
const (
	AssemblerParserEOF = antlr.TokenEOF
	AssemblerParserT__0 = 1
	AssemblerParserT__1 = 2
	AssemblerParserT__2 = 3
	AssemblerParserT__3 = 4
	AssemblerParserT__4 = 5
	AssemblerParserT__5 = 6
	AssemblerParserT__6 = 7
	AssemblerParserT__7 = 8
	AssemblerParserT__8 = 9
	AssemblerParserT__9 = 10
	AssemblerParserTAG_FLOAT = 11
	AssemblerParserTAG_STRING = 12
	AssemblerParserTAG_STUCT = 13
	AssemblerParserREG = 14
	AssemblerParserID = 15
	AssemblerParserFUNC = 16
	AssemblerParserINT = 17
	AssemblerParserCHAR = 18
	AssemblerParserSTRING = 19
	AssemblerParserFLOAT = 20
	AssemblerParserWS = 21
	AssemblerParserNEWLINE = 22
)

// AssemblerParser rules.
const (
	AssemblerParserRULE_program = 0
	AssemblerParserRULE_globals = 1
	AssemblerParserRULE_constantPool = 2
	AssemblerParserRULE_constantPoolEntry = 3
	AssemblerParserRULE_structDecl = 4
	AssemblerParserRULE_functionDeclaration = 5
	AssemblerParserRULE_instr = 6
	AssemblerParserRULE_operand = 7
	AssemblerParserRULE_label = 8
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Globals() IGlobalsContext
	ConstantPool() IConstantPoolContext
	AllFunctionDeclaration() []IFunctionDeclarationContext
	FunctionDeclaration(i int) IFunctionDeclarationContext
	AllInstr() []IInstrContext
	Instr(i int) IInstrContext
	AllLabel() []ILabelContext
	Label(i int) ILabelContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblerParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Globals() IGlobalsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGlobalsContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGlobalsContext)
}

func (s *ProgramContext) ConstantPool() IConstantPoolContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantPoolContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantPoolContext)
}

func (s *ProgramContext) AllFunctionDeclaration() []IFunctionDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IFunctionDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionDeclarationContext); ok {
			tst[i] = t.(IFunctionDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) FunctionDeclaration(i int) IFunctionDeclarationContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *ProgramContext) AllInstr() []IInstrContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstrContext); ok {
			len++
		}
	}

	tst := make([]IInstrContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstrContext); ok {
			tst[i] = t.(IInstrContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Instr(i int) IInstrContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstrContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstrContext)
}

func (s *ProgramContext) AllLabel() []ILabelContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILabelContext); ok {
			len++
		}
	}

	tst := make([]ILabelContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILabelContext); ok {
			tst[i] = t.(ILabelContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Label(i int) ILabelContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *ProgramContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(AssemblerParserNEWLINE)
}

func (s *ProgramContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(AssemblerParserNEWLINE, i)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AssemblerVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *AssemblerParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AssemblerParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(19)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(18)
			p.Globals()
		}

		} else if p.HasError() { // JIM
			goto errorExit
	}
	p.SetState(22)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(21)
			p.ConstantPool()
		}

		} else if p.HasError() { // JIM
			goto errorExit
	}
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 4227136) != 0) {
		p.SetState(28)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(24)
				p.FunctionDeclaration()
			}


		case 2:
			{
				p.SetState(25)
				p.Instr()
			}


		case 3:
			{
				p.SetState(26)
				p.Label()
			}


		case 4:
			{
				p.SetState(27)
				p.Match(AssemblerParserNEWLINE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	p.CheckForUnresolvedReferences()



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IGlobalsContext is an interface to support dynamic dispatch.
type IGlobalsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_INT returns the _INT token.
	Get_INT() antlr.Token 


	// Set_INT sets the _INT token.
	Set_INT(antlr.Token) 


	// Getter signatures
	INT() antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsGlobalsContext differentiates from other interfaces.
	IsGlobalsContext()
}

type GlobalsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	_INT antlr.Token
}

func NewEmptyGlobalsContext() *GlobalsContext {
	var p = new(GlobalsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_globals
	return p
}

func InitEmptyGlobalsContext(p *GlobalsContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_globals
}

func (*GlobalsContext) IsGlobalsContext() {}

func NewGlobalsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GlobalsContext {
	var p = new(GlobalsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblerParserRULE_globals

	return p
}

func (s *GlobalsContext) GetParser() antlr.Parser { return s.parser }

func (s *GlobalsContext) Get_INT() antlr.Token { return s._INT }


func (s *GlobalsContext) Set_INT(v antlr.Token) { s._INT = v }


func (s *GlobalsContext) INT() antlr.TerminalNode {
	return s.GetToken(AssemblerParserINT, 0)
}

func (s *GlobalsContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(AssemblerParserNEWLINE)
}

func (s *GlobalsContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(AssemblerParserNEWLINE, i)
}

func (s *GlobalsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GlobalsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *GlobalsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AssemblerVisitor:
		return t.VisitGlobals(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *AssemblerParser) Globals() (localctx IGlobalsContext) {
	localctx = NewGlobalsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AssemblerParserRULE_globals)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == AssemblerParserNEWLINE {
		{
			p.SetState(34)
			p.Match(AssemblerParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(40)
		p.Match(AssemblerParserT__0)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(41)

		var _m = p.Match(AssemblerParserINT)

		localctx.(*GlobalsContext)._INT = _m
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(42)
		p.Match(AssemblerParserNEWLINE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.DefineDataSize((func() int { if localctx.(*GlobalsContext).Get_INT() == nil { return 0 } else { i, _ := strconv.Atoi(localctx.(*GlobalsContext).Get_INT().GetText()); return i }}()))



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IConstantPoolContext is an interface to support dynamic dispatch.
type IConstantPoolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCount returns the count token.
	GetCount() antlr.Token 


	// SetCount sets the count token.
	SetCount(antlr.Token) 


	// Get_constantPoolEntry returns the _constantPoolEntry rule contexts.
	Get_constantPoolEntry() IConstantPoolEntryContext


	// Set_constantPoolEntry sets the _constantPoolEntry rule contexts.
	Set_constantPoolEntry(IConstantPoolEntryContext)


	// GetEntries returns the entries rule context list.
	GetEntries() []IConstantPoolEntryContext


	// SetEntries sets the entries rule context list.
	SetEntries([]IConstantPoolEntryContext) 


	// Getter signatures
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	INT() antlr.TerminalNode
	AllConstantPoolEntry() []IConstantPoolEntryContext
	ConstantPoolEntry(i int) IConstantPoolEntryContext

	// IsConstantPoolContext differentiates from other interfaces.
	IsConstantPoolContext()
}

type ConstantPoolContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	count antlr.Token
	_constantPoolEntry IConstantPoolEntryContext 
	entries []IConstantPoolEntryContext
}

func NewEmptyConstantPoolContext() *ConstantPoolContext {
	var p = new(ConstantPoolContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_constantPool
	return p
}

func InitEmptyConstantPoolContext(p *ConstantPoolContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_constantPool
}

func (*ConstantPoolContext) IsConstantPoolContext() {}

func NewConstantPoolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantPoolContext {
	var p = new(ConstantPoolContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblerParserRULE_constantPool

	return p
}

func (s *ConstantPoolContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantPoolContext) GetCount() antlr.Token { return s.count }


func (s *ConstantPoolContext) SetCount(v antlr.Token) { s.count = v }


func (s *ConstantPoolContext) Get_constantPoolEntry() IConstantPoolEntryContext { return s._constantPoolEntry }


func (s *ConstantPoolContext) Set_constantPoolEntry(v IConstantPoolEntryContext) { s._constantPoolEntry = v }


func (s *ConstantPoolContext) GetEntries() []IConstantPoolEntryContext { return s.entries }


func (s *ConstantPoolContext) SetEntries(v []IConstantPoolEntryContext) { s.entries = v }


func (s *ConstantPoolContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(AssemblerParserNEWLINE)
}

func (s *ConstantPoolContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(AssemblerParserNEWLINE, i)
}

func (s *ConstantPoolContext) INT() antlr.TerminalNode {
	return s.GetToken(AssemblerParserINT, 0)
}

func (s *ConstantPoolContext) AllConstantPoolEntry() []IConstantPoolEntryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstantPoolEntryContext); ok {
			len++
		}
	}

	tst := make([]IConstantPoolEntryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstantPoolEntryContext); ok {
			tst[i] = t.(IConstantPoolEntryContext)
			i++
		}
	}

	return tst
}

func (s *ConstantPoolContext) ConstantPoolEntry(i int) IConstantPoolEntryContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantPoolEntryContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantPoolEntryContext)
}

func (s *ConstantPoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantPoolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ConstantPoolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AssemblerVisitor:
		return t.VisitConstantPool(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *AssemblerParser) ConstantPool() (localctx IConstantPoolContext) {
	localctx = NewConstantPoolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AssemblerParserRULE_constantPool)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == AssemblerParserNEWLINE {
		{
			p.SetState(45)
			p.Match(AssemblerParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(51)
		p.Match(AssemblerParserT__1)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(52)

		var _m = p.Match(AssemblerParserINT)

		localctx.(*ConstantPoolContext).count = _m
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(53)
		p.Match(AssemblerParserNEWLINE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 14336) != 0) {
		{
			p.SetState(54)

			var _x = p.ConstantPoolEntry()


			localctx.(*ConstantPoolContext)._constantPoolEntry = _x
		}
		localctx.(*ConstantPoolContext).entries = append(localctx.(*ConstantPoolContext).entries, localctx.(*ConstantPoolContext)._constantPoolEntry)


		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IConstantPoolEntryContext is an interface to support dynamic dispatch.
type IConstantPoolEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTag returns the tag token.
	GetTag() antlr.Token 

	// GetValue returns the value token.
	GetValue() antlr.Token 

	// GetName returns the name token.
	GetName() antlr.Token 

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token 


	// SetTag sets the tag token.
	SetTag(antlr.Token) 

	// SetValue sets the value token.
	SetValue(antlr.Token) 

	// SetName sets the name token.
	SetName(antlr.Token) 

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token) 


	// GetFields returns the fields token list.
	GetFields() []antlr.Token


	// SetFields sets the fields token list.
	SetFields([]antlr.Token)


	// Getter signatures
	TAG_FLOAT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	TAG_STRING() antlr.TerminalNode
	STRING() antlr.TerminalNode
	TAG_STUCT() antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode

	// IsConstantPoolEntryContext differentiates from other interfaces.
	IsConstantPoolEntryContext()
}

type ConstantPoolEntryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	tag antlr.Token
	value antlr.Token
	name antlr.Token
	_ID antlr.Token
	fields []antlr.Token
}

func NewEmptyConstantPoolEntryContext() *ConstantPoolEntryContext {
	var p = new(ConstantPoolEntryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_constantPoolEntry
	return p
}

func InitEmptyConstantPoolEntryContext(p *ConstantPoolEntryContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_constantPoolEntry
}

func (*ConstantPoolEntryContext) IsConstantPoolEntryContext() {}

func NewConstantPoolEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantPoolEntryContext {
	var p = new(ConstantPoolEntryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblerParserRULE_constantPoolEntry

	return p
}

func (s *ConstantPoolEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantPoolEntryContext) GetTag() antlr.Token { return s.tag }

func (s *ConstantPoolEntryContext) GetValue() antlr.Token { return s.value }

func (s *ConstantPoolEntryContext) GetName() antlr.Token { return s.name }

func (s *ConstantPoolEntryContext) Get_ID() antlr.Token { return s._ID }


func (s *ConstantPoolEntryContext) SetTag(v antlr.Token) { s.tag = v }

func (s *ConstantPoolEntryContext) SetValue(v antlr.Token) { s.value = v }

func (s *ConstantPoolEntryContext) SetName(v antlr.Token) { s.name = v }

func (s *ConstantPoolEntryContext) Set_ID(v antlr.Token) { s._ID = v }


func (s *ConstantPoolEntryContext) GetFields() []antlr.Token { return s.fields }


func (s *ConstantPoolEntryContext) SetFields(v []antlr.Token) { s.fields = v }


func (s *ConstantPoolEntryContext) TAG_FLOAT() antlr.TerminalNode {
	return s.GetToken(AssemblerParserTAG_FLOAT, 0)
}

func (s *ConstantPoolEntryContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(AssemblerParserFLOAT, 0)
}

func (s *ConstantPoolEntryContext) TAG_STRING() antlr.TerminalNode {
	return s.GetToken(AssemblerParserTAG_STRING, 0)
}

func (s *ConstantPoolEntryContext) STRING() antlr.TerminalNode {
	return s.GetToken(AssemblerParserSTRING, 0)
}

func (s *ConstantPoolEntryContext) TAG_STUCT() antlr.TerminalNode {
	return s.GetToken(AssemblerParserTAG_STUCT, 0)
}

func (s *ConstantPoolEntryContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(AssemblerParserID)
}

func (s *ConstantPoolEntryContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(AssemblerParserID, i)
}

func (s *ConstantPoolEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantPoolEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ConstantPoolEntryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AssemblerVisitor:
		return t.VisitConstantPoolEntry(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *AssemblerParser) ConstantPoolEntry() (localctx IConstantPoolEntryContext) {
	localctx = NewConstantPoolEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AssemblerParserRULE_constantPoolEntry)
	var _la int

	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AssemblerParserTAG_FLOAT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(60)

			var _m = p.Match(AssemblerParserTAG_FLOAT)

			localctx.(*ConstantPoolEntryContext).tag = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(61)

			var _m = p.Match(AssemblerParserFLOAT)

			localctx.(*ConstantPoolEntryContext).value = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.DefineConst((func() string { if localctx.(*ConstantPoolEntryContext).GetTag() == nil { return "" } else { return localctx.(*ConstantPoolEntryContext).GetTag().GetText() }}()), (func() string { if localctx.(*ConstantPoolEntryContext).GetValue() == nil { return "" } else { return localctx.(*ConstantPoolEntryContext).GetValue().GetText() }}()))


	case AssemblerParserTAG_STRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(63)

			var _m = p.Match(AssemblerParserTAG_STRING)

			localctx.(*ConstantPoolEntryContext).tag = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(64)

			var _m = p.Match(AssemblerParserSTRING)

			localctx.(*ConstantPoolEntryContext).value = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.DefineConst((func() string { if localctx.(*ConstantPoolEntryContext).GetTag() == nil { return "" } else { return localctx.(*ConstantPoolEntryContext).GetTag().GetText() }}()), (func() string { if localctx.(*ConstantPoolEntryContext).GetValue() == nil { return "" } else { return localctx.(*ConstantPoolEntryContext).GetValue().GetText() }}()))


	case AssemblerParserTAG_STUCT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(66)

			var _m = p.Match(AssemblerParserTAG_STUCT)

			localctx.(*ConstantPoolEntryContext).tag = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(67)

			var _m = p.Match(AssemblerParserID)

			localctx.(*ConstantPoolEntryContext).name = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(68)
			p.Match(AssemblerParserT__2)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if _la == AssemblerParserID {
			{
				p.SetState(69)

				var _m = p.Match(AssemblerParserID)

				localctx.(*ConstantPoolEntryContext)._ID = _m
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			localctx.(*ConstantPoolEntryContext).fields = append(localctx.(*ConstantPoolEntryContext).fields, localctx.(*ConstantPoolEntryContext)._ID)
			p.SetState(74)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)


			for _la == AssemblerParserT__3 {
				{
					p.SetState(70)
					p.Match(AssemblerParserT__3)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(71)

					var _m = p.Match(AssemblerParserID)

					localctx.(*ConstantPoolEntryContext)._ID = _m
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				localctx.(*ConstantPoolEntryContext).fields = append(localctx.(*ConstantPoolEntryContext).fields, localctx.(*ConstantPoolEntryContext)._ID)


				p.SetState(76)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
			    	goto errorExit
			    }
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(79)
			p.Match(AssemblerParserT__4)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		 p.DefineStruct((func() string { if localctx.(*ConstantPoolEntryContext).GetName() == nil { return "" } else { return localctx.(*ConstantPoolEntryContext).GetName().GetText() }}()), localctx.(*ConstantPoolEntryContext).GetFields())



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}


errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IStructDeclContext is an interface to support dynamic dispatch.
type IStructDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token 

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token 


	// SetName sets the name token.
	SetName(antlr.Token) 

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token) 


	// GetFields returns the fields token list.
	GetFields() []antlr.Token


	// SetFields sets the fields token list.
	SetFields([]antlr.Token)


	// Getter signatures
	TAG_STUCT() antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode

	// IsStructDeclContext differentiates from other interfaces.
	IsStructDeclContext()
}

type StructDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name antlr.Token
	_ID antlr.Token
	fields []antlr.Token
}

func NewEmptyStructDeclContext() *StructDeclContext {
	var p = new(StructDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_structDecl
	return p
}

func InitEmptyStructDeclContext(p *StructDeclContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_structDecl
}

func (*StructDeclContext) IsStructDeclContext() {}

func NewStructDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDeclContext {
	var p = new(StructDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblerParserRULE_structDecl

	return p
}

func (s *StructDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDeclContext) GetName() antlr.Token { return s.name }

func (s *StructDeclContext) Get_ID() antlr.Token { return s._ID }


func (s *StructDeclContext) SetName(v antlr.Token) { s.name = v }

func (s *StructDeclContext) Set_ID(v antlr.Token) { s._ID = v }


func (s *StructDeclContext) GetFields() []antlr.Token { return s.fields }


func (s *StructDeclContext) SetFields(v []antlr.Token) { s.fields = v }


func (s *StructDeclContext) TAG_STUCT() antlr.TerminalNode {
	return s.GetToken(AssemblerParserTAG_STUCT, 0)
}

func (s *StructDeclContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(AssemblerParserID)
}

func (s *StructDeclContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(AssemblerParserID, i)
}

func (s *StructDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StructDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AssemblerVisitor:
		return t.VisitStructDecl(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *AssemblerParser) StructDecl() (localctx IStructDeclContext) {
	localctx = NewStructDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AssemblerParserRULE_structDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(83)
		p.Match(AssemblerParserTAG_STUCT)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(84)

		var _m = p.Match(AssemblerParserID)

		localctx.(*StructDeclContext).name = _m
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(85)
		p.Match(AssemblerParserT__2)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == AssemblerParserID {
		{
			p.SetState(86)

			var _m = p.Match(AssemblerParserID)

			localctx.(*StructDeclContext)._ID = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		localctx.(*StructDeclContext).fields = append(localctx.(*StructDeclContext).fields, localctx.(*StructDeclContext)._ID)
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == AssemblerParserT__3 {
			{
				p.SetState(87)
				p.Match(AssemblerParserT__3)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(88)

				var _m = p.Match(AssemblerParserID)

				localctx.(*StructDeclContext)._ID = _m
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			localctx.(*StructDeclContext).fields = append(localctx.(*StructDeclContext).fields, localctx.(*StructDeclContext)._ID)


			p.SetState(93)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(96)
		p.Match(AssemblerParserT__4)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IFunctionDeclarationContext is an interface to support dynamic dispatch.
type IFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token 

	// GetA returns the a token.
	GetA() antlr.Token 

	// GetN returns the n token.
	GetN() antlr.Token 


	// SetName sets the name token.
	SetName(antlr.Token) 

	// SetA sets the a token.
	SetA(antlr.Token) 

	// SetN sets the n token.
	SetN(antlr.Token) 


	// Get_structDecl returns the _structDecl rule contexts.
	Get_structDecl() IStructDeclContext


	// Set_structDecl sets the _structDecl rule contexts.
	Set_structDecl(IStructDeclContext)


	// GetStructs returns the structs rule context list.
	GetStructs() []IStructDeclContext


	// SetStructs sets the structs rule context list.
	SetStructs([]IStructDeclContext) 


	// Getter signatures
	NEWLINE() antlr.TerminalNode
	ID() antlr.TerminalNode
	AllINT() []antlr.TerminalNode
	INT(i int) antlr.TerminalNode
	AllStructDecl() []IStructDeclContext
	StructDecl(i int) IStructDeclContext

	// IsFunctionDeclarationContext differentiates from other interfaces.
	IsFunctionDeclarationContext()
}

type FunctionDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name antlr.Token
	a antlr.Token
	n antlr.Token
	_structDecl IStructDeclContext 
	structs []IStructDeclContext
}

func NewEmptyFunctionDeclarationContext() *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_functionDeclaration
	return p
}

func InitEmptyFunctionDeclarationContext(p *FunctionDeclarationContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_functionDeclaration
}

func (*FunctionDeclarationContext) IsFunctionDeclarationContext() {}

func NewFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblerParserRULE_functionDeclaration

	return p
}

func (s *FunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationContext) GetName() antlr.Token { return s.name }

func (s *FunctionDeclarationContext) GetA() antlr.Token { return s.a }

func (s *FunctionDeclarationContext) GetN() antlr.Token { return s.n }


func (s *FunctionDeclarationContext) SetName(v antlr.Token) { s.name = v }

func (s *FunctionDeclarationContext) SetA(v antlr.Token) { s.a = v }

func (s *FunctionDeclarationContext) SetN(v antlr.Token) { s.n = v }


func (s *FunctionDeclarationContext) Get_structDecl() IStructDeclContext { return s._structDecl }


func (s *FunctionDeclarationContext) Set_structDecl(v IStructDeclContext) { s._structDecl = v }


func (s *FunctionDeclarationContext) GetStructs() []IStructDeclContext { return s.structs }


func (s *FunctionDeclarationContext) SetStructs(v []IStructDeclContext) { s.structs = v }


func (s *FunctionDeclarationContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(AssemblerParserNEWLINE, 0)
}

func (s *FunctionDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(AssemblerParserID, 0)
}

func (s *FunctionDeclarationContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(AssemblerParserINT)
}

func (s *FunctionDeclarationContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(AssemblerParserINT, i)
}

func (s *FunctionDeclarationContext) AllStructDecl() []IStructDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructDeclContext); ok {
			len++
		}
	}

	tst := make([]IStructDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructDeclContext); ok {
			tst[i] = t.(IStructDeclContext)
			i++
		}
	}

	return tst
}

func (s *FunctionDeclarationContext) StructDecl(i int) IStructDeclContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructDeclContext)
}

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FunctionDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AssemblerVisitor:
		return t.VisitFunctionDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *AssemblerParser) FunctionDeclaration() (localctx IFunctionDeclarationContext) {
	localctx = NewFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AssemblerParserRULE_functionDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Match(AssemblerParserT__5)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(99)

		var _m = p.Match(AssemblerParserID)

		localctx.(*FunctionDeclarationContext).name = _m
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(100)
		p.Match(AssemblerParserT__6)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(101)
		p.Match(AssemblerParserT__7)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(102)
		p.Match(AssemblerParserT__8)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(103)

		var _m = p.Match(AssemblerParserINT)

		localctx.(*FunctionDeclarationContext).a = _m
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(104)
		p.Match(AssemblerParserT__3)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(105)
		p.Match(AssemblerParserT__9)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(106)
		p.Match(AssemblerParserT__8)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(107)

		var _m = p.Match(AssemblerParserINT)

		localctx.(*FunctionDeclarationContext).n = _m
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(108)
		p.Match(AssemblerParserNEWLINE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == AssemblerParserTAG_STUCT {
		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = _la == AssemblerParserTAG_STUCT {
			{
				p.SetState(109)

				var _x = p.StructDecl()


				localctx.(*FunctionDeclarationContext)._structDecl = _x
			}
			localctx.(*FunctionDeclarationContext).structs = append(localctx.(*FunctionDeclarationContext).structs, localctx.(*FunctionDeclarationContext)._structDecl)


			p.SetState(112)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}

	}

	        p.DefineFunction(localctx.(*FunctionDeclarationContext).GetName(), (func() int { if localctx.(*FunctionDeclarationContext).GetA() == nil { return 0 } else { i, _ := strconv.Atoi(localctx.(*FunctionDeclarationContext).GetA().GetText()); return i }}()), (func() int { if localctx.(*FunctionDeclarationContext).GetN() == nil { return 0 } else { i, _ := strconv.Atoi(localctx.(*FunctionDeclarationContext).GetN().GetText()); return i }}()))
	        for _, v := range localctx.(*FunctionDeclarationContext).GetStructs() {
	          p.DefineFuncStruct(localctx.(*FunctionDeclarationContext).GetName(), v.GetName().GetText(), v.GetFields())
	        }
	      



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IInstrContext is an interface to support dynamic dispatch.
type IInstrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token 


	// Set_ID sets the _ID token.
	Set_ID(antlr.Token) 


	// Get_operand returns the _operand rule contexts.
	Get_operand() IOperandContext

	// GetA returns the a rule contexts.
	GetA() IOperandContext

	// GetB returns the b rule contexts.
	GetB() IOperandContext

	// GetC returns the c rule contexts.
	GetC() IOperandContext


	// Set_operand sets the _operand rule contexts.
	Set_operand(IOperandContext)

	// SetA sets the a rule contexts.
	SetA(IOperandContext)

	// SetB sets the b rule contexts.
	SetB(IOperandContext)

	// SetC sets the c rule contexts.
	SetC(IOperandContext)


	// Getter signatures
	ID() antlr.TerminalNode
	NEWLINE() antlr.TerminalNode
	AllOperand() []IOperandContext
	Operand(i int) IOperandContext

	// IsInstrContext differentiates from other interfaces.
	IsInstrContext()
}

type InstrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	_ID antlr.Token
	_operand IOperandContext 
	a IOperandContext 
	b IOperandContext 
	c IOperandContext 
}

func NewEmptyInstrContext() *InstrContext {
	var p = new(InstrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_instr
	return p
}

func InitEmptyInstrContext(p *InstrContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_instr
}

func (*InstrContext) IsInstrContext() {}

func NewInstrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstrContext {
	var p = new(InstrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblerParserRULE_instr

	return p
}

func (s *InstrContext) GetParser() antlr.Parser { return s.parser }

func (s *InstrContext) Get_ID() antlr.Token { return s._ID }


func (s *InstrContext) Set_ID(v antlr.Token) { s._ID = v }


func (s *InstrContext) Get_operand() IOperandContext { return s._operand }

func (s *InstrContext) GetA() IOperandContext { return s.a }

func (s *InstrContext) GetB() IOperandContext { return s.b }

func (s *InstrContext) GetC() IOperandContext { return s.c }


func (s *InstrContext) Set_operand(v IOperandContext) { s._operand = v }

func (s *InstrContext) SetA(v IOperandContext) { s.a = v }

func (s *InstrContext) SetB(v IOperandContext) { s.b = v }

func (s *InstrContext) SetC(v IOperandContext) { s.c = v }


func (s *InstrContext) ID() antlr.TerminalNode {
	return s.GetToken(AssemblerParserID, 0)
}

func (s *InstrContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(AssemblerParserNEWLINE, 0)
}

func (s *InstrContext) AllOperand() []IOperandContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOperandContext); ok {
			len++
		}
	}

	tst := make([]IOperandContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOperandContext); ok {
			tst[i] = t.(IOperandContext)
			i++
		}
	}

	return tst
}

func (s *InstrContext) Operand(i int) IOperandContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *InstrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *InstrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AssemblerVisitor:
		return t.VisitInstr(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *AssemblerParser) Instr() (localctx IInstrContext) {
	localctx = NewInstrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AssemblerParserRULE_instr)
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(118)

			var _m = p.Match(AssemblerParserID)

			localctx.(*InstrContext)._ID = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(119)
			p.Match(AssemblerParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.Gen(localctx.(*InstrContext).Get_ID())


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(121)

			var _m = p.Match(AssemblerParserID)

			localctx.(*InstrContext)._ID = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(122)

			var _x = p.Operand()


			localctx.(*InstrContext)._operand = _x
		}
		{
			p.SetState(123)
			p.Match(AssemblerParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.Gen(localctx.(*InstrContext).Get_ID(),(func() antlr.Token { if localctx.(*InstrContext).Get_operand() == nil { return nil } else { return localctx.(*InstrContext).Get_operand().GetStart() }}()))


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(126)

			var _m = p.Match(AssemblerParserID)

			localctx.(*InstrContext)._ID = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(127)

			var _x = p.Operand()


			localctx.(*InstrContext).a = _x
		}
		{
			p.SetState(128)
			p.Match(AssemblerParserT__3)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(129)

			var _x = p.Operand()


			localctx.(*InstrContext).b = _x
		}
		{
			p.SetState(130)
			p.Match(AssemblerParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.Gen(localctx.(*InstrContext).Get_ID(),(func() antlr.Token { if localctx.(*InstrContext).GetA() == nil { return nil } else { return localctx.(*InstrContext).GetA().GetStart() }}()),(func() antlr.Token { if localctx.(*InstrContext).GetB() == nil { return nil } else { return localctx.(*InstrContext).GetB().GetStart() }}()))


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(133)

			var _m = p.Match(AssemblerParserID)

			localctx.(*InstrContext)._ID = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(134)

			var _x = p.Operand()


			localctx.(*InstrContext).a = _x
		}
		{
			p.SetState(135)
			p.Match(AssemblerParserT__3)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(136)

			var _x = p.Operand()


			localctx.(*InstrContext).b = _x
		}
		{
			p.SetState(137)
			p.Match(AssemblerParserT__3)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(138)

			var _x = p.Operand()


			localctx.(*InstrContext).c = _x
		}
		{
			p.SetState(139)
			p.Match(AssemblerParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.Gen(localctx.(*InstrContext).Get_ID(),(func() antlr.Token { if localctx.(*InstrContext).GetA() == nil { return nil } else { return localctx.(*InstrContext).GetA().GetStart() }}()),(func() antlr.Token { if localctx.(*InstrContext).GetB() == nil { return nil } else { return localctx.(*InstrContext).GetB().GetStart() }}()),(func() antlr.Token { if localctx.(*InstrContext).GetC() == nil { return nil } else { return localctx.(*InstrContext).GetC().GetStart() }}()))

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}


errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IOperandContext is an interface to support dynamic dispatch.
type IOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	REG() antlr.TerminalNode
	FUNC() antlr.TerminalNode
	INT() antlr.TerminalNode
	CHAR() antlr.TerminalNode
	STRING() antlr.TerminalNode
	FLOAT() antlr.TerminalNode

	// IsOperandContext differentiates from other interfaces.
	IsOperandContext()
}

type OperandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandContext() *OperandContext {
	var p = new(OperandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_operand
	return p
}

func InitEmptyOperandContext(p *OperandContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_operand
}

func (*OperandContext) IsOperandContext() {}

func NewOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandContext {
	var p = new(OperandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblerParserRULE_operand

	return p
}

func (s *OperandContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandContext) ID() antlr.TerminalNode {
	return s.GetToken(AssemblerParserID, 0)
}

func (s *OperandContext) REG() antlr.TerminalNode {
	return s.GetToken(AssemblerParserREG, 0)
}

func (s *OperandContext) FUNC() antlr.TerminalNode {
	return s.GetToken(AssemblerParserFUNC, 0)
}

func (s *OperandContext) INT() antlr.TerminalNode {
	return s.GetToken(AssemblerParserINT, 0)
}

func (s *OperandContext) CHAR() antlr.TerminalNode {
	return s.GetToken(AssemblerParserCHAR, 0)
}

func (s *OperandContext) STRING() antlr.TerminalNode {
	return s.GetToken(AssemblerParserSTRING, 0)
}

func (s *OperandContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(AssemblerParserFLOAT, 0)
}

func (s *OperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *OperandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AssemblerVisitor:
		return t.VisitOperand(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *AssemblerParser) Operand() (localctx IOperandContext) {
	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AssemblerParserRULE_operand)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		_la = p.GetTokenStream().LA(1)

		if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 2080768) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// ILabelContext is an interface to support dynamic dispatch.
type ILabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token 


	// Set_ID sets the _ID token.
	Set_ID(antlr.Token) 


	// Getter signatures
	ID() antlr.TerminalNode

	// IsLabelContext differentiates from other interfaces.
	IsLabelContext()
}

type LabelContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	_ID antlr.Token
}

func NewEmptyLabelContext() *LabelContext {
	var p = new(LabelContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_label
	return p
}

func InitEmptyLabelContext(p *LabelContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_label
}

func (*LabelContext) IsLabelContext() {}

func NewLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelContext {
	var p = new(LabelContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblerParserRULE_label

	return p
}

func (s *LabelContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelContext) Get_ID() antlr.Token { return s._ID }


func (s *LabelContext) Set_ID(v antlr.Token) { s._ID = v }


func (s *LabelContext) ID() antlr.TerminalNode {
	return s.GetToken(AssemblerParserID, 0)
}

func (s *LabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LabelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AssemblerVisitor:
		return t.VisitLabel(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *AssemblerParser) Label() (localctx ILabelContext) {
	localctx = NewLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, AssemblerParserRULE_label)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(146)

		var _m = p.Match(AssemblerParserID)

		localctx.(*LabelContext)._ID = _m
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(147)
		p.Match(AssemblerParserT__6)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.DefineLabel(localctx.(*LabelContext).Get_ID())



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


