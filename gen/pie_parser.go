// Code generated from /Users/bytedance/go/src/github.com/ycl2018/pie-go/gen/Pie.g4 by ANTLR 4.13.2. DO NOT EDIT.

package gen // Pie
import (
	"fmt"
	"strconv"
  	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}


type PieParser struct {
	*antlr.BaseParser
}

var PieParserStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  LiteralNames           []string
  SymbolicNames          []string
  RuleNames              []string
  PredictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func pieParserInit() {
  staticData := &PieParserStaticData
  staticData.LiteralNames = []string{
    "", "'{'", "','", "'}'", "'('", "')'", "':'", "'else'", "'if'", "'='", 
    "'print'", "'while'", "'return'", "'def'", "'+'", "'-'", "'*'", "'/'", 
    "'=='", "'<'", "'>'", "'>='", "'<='", "'!='", "'struct'", "'.'", "'new'",
  }
  staticData.SymbolicNames = []string{
    "", "", "", "", "", "", "", "", "IF", "ASSIGN", "PRINT", "WHILE", "RETURN", 
    "DEF", "ADD", "SUB", "MUL", "DIV", "EQ", "LT", "GT", "GEQ", "LEQ", "NEQ", 
    "STRUCT", "DOT", "NEW", "NL", "ID", "CHAR", "STRING", "INT", "FLOAT", 
    "WS", "SL_COMMENT",
  }
  staticData.RuleNames = []string{
    "program", "structDefinition", "functionDefinition", "slist", "statement", 
    "call", "expr", "addexpr", "mulexpr", "atom", "instance", "qid", "vardef", 
    "multOp", "addOp", "compOp",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 34, 186, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7, 
	10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15, 
	1, 0, 1, 0, 4, 0, 35, 8, 0, 11, 0, 12, 0, 36, 1, 0, 1, 0, 1, 1, 1, 1, 1, 
	1, 1, 1, 1, 1, 1, 1, 5, 1, 47, 8, 1, 10, 1, 12, 1, 50, 9, 1, 1, 1, 1, 1, 
	1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 61, 8, 2, 10, 2, 12, 2, 
	64, 9, 2, 3, 2, 66, 8, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 4, 3, 74, 
	8, 3, 11, 3, 12, 3, 75, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 82, 8, 3, 1, 4, 1, 
	4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 92, 8, 4, 1, 4, 1, 4, 1, 4, 
	1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 104, 8, 4, 1, 4, 1, 4, 
	1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 114, 8, 4, 1, 5, 1, 5, 1, 5, 
	1, 5, 1, 5, 5, 5, 121, 8, 5, 10, 5, 12, 5, 124, 9, 5, 3, 5, 126, 8, 5, 
	1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 134, 8, 6, 1, 7, 1, 7, 1, 7, 
	1, 7, 5, 7, 140, 8, 7, 10, 7, 12, 7, 143, 9, 7, 1, 8, 1, 8, 1, 8, 1, 8, 
	5, 8, 149, 8, 8, 10, 8, 12, 8, 152, 9, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 
	1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 165, 8, 9, 1, 10, 1, 10, 1, 10, 
	1, 11, 1, 11, 1, 11, 5, 11, 173, 8, 11, 10, 11, 12, 11, 176, 9, 11, 1, 
	12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 0, 0, 16, 0, 
	2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 0, 3, 1, 0, 16, 
	17, 1, 0, 14, 15, 1, 0, 18, 23, 198, 0, 34, 1, 0, 0, 0, 2, 40, 1, 0, 0, 
	0, 4, 54, 1, 0, 0, 0, 6, 81, 1, 0, 0, 0, 8, 113, 1, 0, 0, 0, 10, 115, 1, 
	0, 0, 0, 12, 129, 1, 0, 0, 0, 14, 135, 1, 0, 0, 0, 16, 144, 1, 0, 0, 0, 
	18, 164, 1, 0, 0, 0, 20, 166, 1, 0, 0, 0, 22, 169, 1, 0, 0, 0, 24, 177, 
	1, 0, 0, 0, 26, 179, 1, 0, 0, 0, 28, 181, 1, 0, 0, 0, 30, 183, 1, 0, 0, 
	0, 32, 35, 3, 4, 2, 0, 33, 35, 3, 8, 4, 0, 34, 32, 1, 0, 0, 0, 34, 33, 
	1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 34, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 
	37, 38, 1, 0, 0, 0, 38, 39, 5, 0, 0, 1, 39, 1, 1, 0, 0, 0, 40, 41, 5, 24, 
	0, 0, 41, 42, 5, 28, 0, 0, 42, 43, 5, 1, 0, 0, 43, 48, 3, 24, 12, 0, 44, 
	45, 5, 2, 0, 0, 45, 47, 3, 24, 12, 0, 46, 44, 1, 0, 0, 0, 47, 50, 1, 0, 
	0, 0, 48, 46, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 51, 1, 0, 0, 0, 50, 48, 
	1, 0, 0, 0, 51, 52, 5, 3, 0, 0, 52, 53, 5, 27, 0, 0, 53, 3, 1, 0, 0, 0, 
	54, 55, 5, 13, 0, 0, 55, 56, 5, 28, 0, 0, 56, 65, 5, 4, 0, 0, 57, 62, 3, 
	24, 12, 0, 58, 59, 5, 2, 0, 0, 59, 61, 3, 24, 12, 0, 60, 58, 1, 0, 0, 0, 
	61, 64, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 66, 1, 
	0, 0, 0, 64, 62, 1, 0, 0, 0, 65, 57, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 
	67, 1, 0, 0, 0, 67, 68, 5, 5, 0, 0, 68, 69, 3, 6, 3, 0, 69, 5, 1, 0, 0, 
	0, 70, 71, 5, 6, 0, 0, 71, 73, 5, 27, 0, 0, 72, 74, 3, 8, 4, 0, 73, 72, 
	1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 
	76, 77, 1, 0, 0, 0, 77, 78, 5, 25, 0, 0, 78, 79, 5, 27, 0, 0, 79, 82, 1, 
	0, 0, 0, 80, 82, 3, 8, 4, 0, 81, 70, 1, 0, 0, 0, 81, 80, 1, 0, 0, 0, 82, 
	7, 1, 0, 0, 0, 83, 114, 3, 2, 1, 0, 84, 85, 3, 22, 11, 0, 85, 86, 5, 9, 
	0, 0, 86, 87, 3, 12, 6, 0, 87, 88, 5, 27, 0, 0, 88, 114, 1, 0, 0, 0, 89, 
	91, 5, 12, 0, 0, 90, 92, 3, 12, 6, 0, 91, 90, 1, 0, 0, 0, 91, 92, 1, 0, 
	0, 0, 92, 93, 1, 0, 0, 0, 93, 114, 5, 27, 0, 0, 94, 95, 5, 10, 0, 0, 95, 
	96, 3, 12, 6, 0, 96, 97, 5, 27, 0, 0, 97, 114, 1, 0, 0, 0, 98, 99, 5, 8, 
	0, 0, 99, 100, 3, 12, 6, 0, 100, 103, 3, 6, 3, 0, 101, 102, 5, 7, 0, 0, 
	102, 104, 3, 6, 3, 0, 103, 101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 
	114, 1, 0, 0, 0, 105, 106, 5, 11, 0, 0, 106, 107, 3, 12, 6, 0, 107, 108, 
	3, 6, 3, 0, 108, 114, 1, 0, 0, 0, 109, 110, 3, 10, 5, 0, 110, 111, 5, 27, 
	0, 0, 111, 114, 1, 0, 0, 0, 112, 114, 5, 27, 0, 0, 113, 83, 1, 0, 0, 0, 
	113, 84, 1, 0, 0, 0, 113, 89, 1, 0, 0, 0, 113, 94, 1, 0, 0, 0, 113, 98, 
	1, 0, 0, 0, 113, 105, 1, 0, 0, 0, 113, 109, 1, 0, 0, 0, 113, 112, 1, 0, 
	0, 0, 114, 9, 1, 0, 0, 0, 115, 116, 5, 28, 0, 0, 116, 125, 5, 4, 0, 0, 
	117, 122, 3, 12, 6, 0, 118, 119, 5, 2, 0, 0, 119, 121, 3, 12, 6, 0, 120, 
	118, 1, 0, 0, 0, 121, 124, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 122, 123, 
	1, 0, 0, 0, 123, 126, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 125, 117, 1, 0, 
	0, 0, 125, 126, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 128, 5, 5, 0, 0, 
	128, 11, 1, 0, 0, 0, 129, 133, 3, 14, 7, 0, 130, 131, 3, 30, 15, 0, 131, 
	132, 3, 14, 7, 0, 132, 134, 1, 0, 0, 0, 133, 130, 1, 0, 0, 0, 133, 134, 
	1, 0, 0, 0, 134, 13, 1, 0, 0, 0, 135, 141, 3, 16, 8, 0, 136, 137, 3, 28, 
	14, 0, 137, 138, 3, 16, 8, 0, 138, 140, 1, 0, 0, 0, 139, 136, 1, 0, 0, 
	0, 140, 143, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 
	15, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 144, 150, 3, 18, 9, 0, 145, 146, 
	3, 26, 13, 0, 146, 147, 3, 18, 9, 0, 147, 149, 1, 0, 0, 0, 148, 145, 1, 
	0, 0, 0, 149, 152, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 150, 151, 1, 0, 0, 
	0, 151, 17, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 153, 165, 5, 31, 0, 0, 154, 
	165, 5, 29, 0, 0, 155, 165, 5, 32, 0, 0, 156, 165, 5, 30, 0, 0, 157, 165, 
	3, 22, 11, 0, 158, 165, 3, 10, 5, 0, 159, 165, 3, 20, 10, 0, 160, 161, 
	5, 4, 0, 0, 161, 162, 3, 12, 6, 0, 162, 163, 5, 5, 0, 0, 163, 165, 1, 0, 
	0, 0, 164, 153, 1, 0, 0, 0, 164, 154, 1, 0, 0, 0, 164, 155, 1, 0, 0, 0, 
	164, 156, 1, 0, 0, 0, 164, 157, 1, 0, 0, 0, 164, 158, 1, 0, 0, 0, 164, 
	159, 1, 0, 0, 0, 164, 160, 1, 0, 0, 0, 165, 19, 1, 0, 0, 0, 166, 167, 5, 
	26, 0, 0, 167, 168, 5, 28, 0, 0, 168, 21, 1, 0, 0, 0, 169, 174, 5, 28, 
	0, 0, 170, 171, 5, 25, 0, 0, 171, 173, 5, 28, 0, 0, 172, 170, 1, 0, 0, 
	0, 173, 176, 1, 0, 0, 0, 174, 172, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 
	23, 1, 0, 0, 0, 176, 174, 1, 0, 0, 0, 177, 178, 5, 28, 0, 0, 178, 25, 1, 
	0, 0, 0, 179, 180, 7, 0, 0, 0, 180, 27, 1, 0, 0, 0, 181, 182, 7, 1, 0, 
	0, 182, 29, 1, 0, 0, 0, 183, 184, 7, 2, 0, 0, 184, 31, 1, 0, 0, 0, 17, 
	34, 36, 48, 62, 65, 75, 81, 91, 103, 113, 122, 125, 133, 141, 150, 164, 
	174,
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

// PieParserInit initializes any static state used to implement PieParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPieParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PieParserInit() {
  staticData := &PieParserStaticData
  staticData.once.Do(pieParserInit)
}

// NewPieParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPieParser(input antlr.TokenStream) *PieParser {
	PieParserInit()
	this := new(PieParser)
	this.BaseParser = antlr.NewBaseParser(input)
  staticData := &PieParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Pie.g4"

	return this
}


// PieParser tokens.
const (
	PieParserEOF = antlr.TokenEOF
	PieParserT__0 = 1
	PieParserT__1 = 2
	PieParserT__2 = 3
	PieParserT__3 = 4
	PieParserT__4 = 5
	PieParserT__5 = 6
	PieParserT__6 = 7
	PieParserIF = 8
	PieParserASSIGN = 9
	PieParserPRINT = 10
	PieParserWHILE = 11
	PieParserRETURN = 12
	PieParserDEF = 13
	PieParserADD = 14
	PieParserSUB = 15
	PieParserMUL = 16
	PieParserDIV = 17
	PieParserEQ = 18
	PieParserLT = 19
	PieParserGT = 20
	PieParserGEQ = 21
	PieParserLEQ = 22
	PieParserNEQ = 23
	PieParserSTRUCT = 24
	PieParserDOT = 25
	PieParserNEW = 26
	PieParserNL = 27
	PieParserID = 28
	PieParserCHAR = 29
	PieParserSTRING = 30
	PieParserINT = 31
	PieParserFLOAT = 32
	PieParserWS = 33
	PieParserSL_COMMENT = 34
)

// PieParser rules.
const (
	PieParserRULE_program = 0
	PieParserRULE_structDefinition = 1
	PieParserRULE_functionDefinition = 2
	PieParserRULE_slist = 3
	PieParserRULE_statement = 4
	PieParserRULE_call = 5
	PieParserRULE_expr = 6
	PieParserRULE_addexpr = 7
	PieParserRULE_mulexpr = 8
	PieParserRULE_atom = 9
	PieParserRULE_instance = 10
	PieParserRULE_qid = 11
	PieParserRULE_vardef = 12
	PieParserRULE_multOp = 13
	PieParserRULE_addOp = 14
	PieParserRULE_compOp = 15
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllFunctionDefinition() []IFunctionDefinitionContext
	FunctionDefinition(i int) IFunctionDefinitionContext
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

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
	p.RuleIndex = PieParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PieParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PieParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(PieParserEOF, 0)
}

func (s *ProgramContext) AllFunctionDefinition() []IFunctionDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionDefinitionContext); ok {
			len++
		}
	}

	tst := make([]IFunctionDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionDefinitionContext); ok {
			tst[i] = t.(IFunctionDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) FunctionDefinition(i int) IFunctionDefinitionContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDefinitionContext); ok {
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

	return t.(IFunctionDefinitionContext)
}

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PieVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *PieParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PieParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 419446016) != 0) {
		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case PieParserDEF:
			{
				p.SetState(32)
				p.FunctionDefinition()
			}


		case PieParserIF, PieParserPRINT, PieParserWHILE, PieParserRETURN, PieParserSTRUCT, PieParserNL, PieParserID:
			{
				p.SetState(33)
				p.Statement()
			}



		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(38)
		p.Match(PieParserEOF)
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


// IStructDefinitionContext is an interface to support dynamic dispatch.
type IStructDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token 


	// SetName sets the name token.
	SetName(antlr.Token) 


	// Getter signatures
	STRUCT() antlr.TerminalNode
	AllVardef() []IVardefContext
	Vardef(i int) IVardefContext
	NL() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsStructDefinitionContext differentiates from other interfaces.
	IsStructDefinitionContext()
}

type StructDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name antlr.Token
}

func NewEmptyStructDefinitionContext() *StructDefinitionContext {
	var p = new(StructDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PieParserRULE_structDefinition
	return p
}

func InitEmptyStructDefinitionContext(p *StructDefinitionContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PieParserRULE_structDefinition
}

func (*StructDefinitionContext) IsStructDefinitionContext() {}

func NewStructDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDefinitionContext {
	var p = new(StructDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PieParserRULE_structDefinition

	return p
}

func (s *StructDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDefinitionContext) GetName() antlr.Token { return s.name }


func (s *StructDefinitionContext) SetName(v antlr.Token) { s.name = v }


func (s *StructDefinitionContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(PieParserSTRUCT, 0)
}

func (s *StructDefinitionContext) AllVardef() []IVardefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVardefContext); ok {
			len++
		}
	}

	tst := make([]IVardefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVardefContext); ok {
			tst[i] = t.(IVardefContext)
			i++
		}
	}

	return tst
}

func (s *StructDefinitionContext) Vardef(i int) IVardefContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVardefContext); ok {
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

	return t.(IVardefContext)
}

func (s *StructDefinitionContext) NL() antlr.TerminalNode {
	return s.GetToken(PieParserNL, 0)
}

func (s *StructDefinitionContext) ID() antlr.TerminalNode {
	return s.GetToken(PieParserID, 0)
}

func (s *StructDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StructDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PieVisitor:
		return t.VisitStructDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *PieParser) StructDefinition() (localctx IStructDefinitionContext) {
	localctx = NewStructDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PieParserRULE_structDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
		p.Match(PieParserSTRUCT)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(41)

		var _m = p.Match(PieParserID)

		localctx.(*StructDefinitionContext).name = _m
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(42)
		p.Match(PieParserT__0)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(43)
		p.Vardef()
	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == PieParserT__1 {
		{
			p.SetState(44)
			p.Match(PieParserT__1)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(45)
			p.Vardef()
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
		p.Match(PieParserT__2)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(52)
		p.Match(PieParserNL)
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


// IFunctionDefinitionContext is an interface to support dynamic dispatch.
type IFunctionDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEF() antlr.TerminalNode
	ID() antlr.TerminalNode
	Slist() ISlistContext
	AllVardef() []IVardefContext
	Vardef(i int) IVardefContext

	// IsFunctionDefinitionContext differentiates from other interfaces.
	IsFunctionDefinitionContext()
}

type FunctionDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDefinitionContext() *FunctionDefinitionContext {
	var p = new(FunctionDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PieParserRULE_functionDefinition
	return p
}

func InitEmptyFunctionDefinitionContext(p *FunctionDefinitionContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PieParserRULE_functionDefinition
}

func (*FunctionDefinitionContext) IsFunctionDefinitionContext() {}

func NewFunctionDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDefinitionContext {
	var p = new(FunctionDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PieParserRULE_functionDefinition

	return p
}

func (s *FunctionDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDefinitionContext) DEF() antlr.TerminalNode {
	return s.GetToken(PieParserDEF, 0)
}

func (s *FunctionDefinitionContext) ID() antlr.TerminalNode {
	return s.GetToken(PieParserID, 0)
}

func (s *FunctionDefinitionContext) Slist() ISlistContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISlistContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISlistContext)
}

func (s *FunctionDefinitionContext) AllVardef() []IVardefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVardefContext); ok {
			len++
		}
	}

	tst := make([]IVardefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVardefContext); ok {
			tst[i] = t.(IVardefContext)
			i++
		}
	}

	return tst
}

func (s *FunctionDefinitionContext) Vardef(i int) IVardefContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVardefContext); ok {
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

	return t.(IVardefContext)
}

func (s *FunctionDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FunctionDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PieVisitor:
		return t.VisitFunctionDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *PieParser) FunctionDefinition() (localctx IFunctionDefinitionContext) {
	localctx = NewFunctionDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PieParserRULE_functionDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)
		p.Match(PieParserDEF)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(55)
		p.Match(PieParserID)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(56)
		p.Match(PieParserT__3)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == PieParserID {
		{
			p.SetState(57)
			p.Vardef()
		}
		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == PieParserT__1 {
			{
				p.SetState(58)
				p.Match(PieParserT__1)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(59)
				p.Vardef()
			}


			p.SetState(64)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(67)
		p.Match(PieParserT__4)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(68)
		p.Slist()
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


// ISlistContext is an interface to support dynamic dispatch.
type ISlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	DOT() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsSlistContext differentiates from other interfaces.
	IsSlistContext()
}

type SlistContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySlistContext() *SlistContext {
	var p = new(SlistContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PieParserRULE_slist
	return p
}

func InitEmptySlistContext(p *SlistContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PieParserRULE_slist
}

func (*SlistContext) IsSlistContext() {}

func NewSlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SlistContext {
	var p = new(SlistContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PieParserRULE_slist

	return p
}

func (s *SlistContext) GetParser() antlr.Parser { return s.parser }

func (s *SlistContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(PieParserNL)
}

func (s *SlistContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(PieParserNL, i)
}

func (s *SlistContext) DOT() antlr.TerminalNode {
	return s.GetToken(PieParserDOT, 0)
}

func (s *SlistContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *SlistContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *SlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SlistContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PieVisitor:
		return t.VisitSlist(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *PieParser) Slist() (localctx ISlistContext) {
	localctx = NewSlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PieParserRULE_slist)
	var _la int

	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PieParserT__5:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(70)
			p.Match(PieParserT__5)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(71)
			p.Match(PieParserNL)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 419437824) != 0) {
			{
				p.SetState(72)
				p.Statement()
			}


			p.SetState(75)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(77)
			p.Match(PieParserDOT)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(78)
			p.Match(PieParserNL)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case PieParserIF, PieParserPRINT, PieParserWHILE, PieParserRETURN, PieParserSTRUCT, PieParserNL, PieParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(80)
			p.Statement()
		}



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


// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PieParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PieParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PieParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyAll(ctx *StatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type StructDefinitionStatementContext struct {
	StatementContext
}

func NewStructDefinitionStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructDefinitionStatementContext {
	var p = new(StructDefinitionStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *StructDefinitionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDefinitionStatementContext) StructDefinition() IStructDefinitionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructDefinitionContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructDefinitionContext)
}


func (s *StructDefinitionStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PieVisitor:
		return t.VisitStructDefinitionStatement(s)

	default:
		return t.VisitChildren(s)
	}
}


type WhileStatementContext struct {
	StatementContext
}

func NewWhileStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileStatementContext {
	var p = new(WhileStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *WhileStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatementContext) WHILE() antlr.TerminalNode {
	return s.GetToken(PieParserWHILE, 0)
}

func (s *WhileStatementContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhileStatementContext) Slist() ISlistContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISlistContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISlistContext)
}


func (s *WhileStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PieVisitor:
		return t.VisitWhileStatement(s)

	default:
		return t.VisitChildren(s)
	}
}


type PrintStatementContext struct {
	StatementContext
}

func NewPrintStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintStatementContext {
	var p = new(PrintStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *PrintStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintStatementContext) PRINT() antlr.TerminalNode {
	return s.GetToken(PieParserPRINT, 0)
}

func (s *PrintStatementContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrintStatementContext) NL() antlr.TerminalNode {
	return s.GetToken(PieParserNL, 0)
}


func (s *PrintStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PieVisitor:
		return t.VisitPrintStatement(s)

	default:
		return t.VisitChildren(s)
	}
}


type AssignmentStatementContext struct {
	StatementContext
}

func NewAssignmentStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *AssignmentStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementContext) Qid() IQidContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQidContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQidContext)
}

func (s *AssignmentStatementContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(PieParserASSIGN, 0)
}

func (s *AssignmentStatementContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignmentStatementContext) NL() antlr.TerminalNode {
	return s.GetToken(PieParserNL, 0)
}


func (s *AssignmentStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PieVisitor:
		return t.VisitAssignmentStatement(s)

	default:
		return t.VisitChildren(s)
	}
}


type NlStatementContext struct {
	StatementContext
}

func NewNlStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NlStatementContext {
	var p = new(NlStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *NlStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NlStatementContext) NL() antlr.TerminalNode {
	return s.GetToken(PieParserNL, 0)
}


func (s *NlStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PieVisitor:
		return t.VisitNlStatement(s)

	default:
		return t.VisitChildren(s)
	}
}


type CallStatementContext struct {
	StatementContext
}

func NewCallStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallStatementContext {
	var p = new(CallStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *CallStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallStatementContext) Call() ICallContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallContext)
}

func (s *CallStatementContext) NL() antlr.TerminalNode {
	return s.GetToken(PieParserNL, 0)
}


func (s *CallStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PieVisitor:
		return t.VisitCallStatement(s)

	default:
		return t.VisitChildren(s)
	}
}


type ReturnStatementContext struct {
	StatementContext
}

func NewReturnStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) RETURN() antlr.TerminalNode {
	return s.GetToken(PieParserRETURN, 0)
}

func (s *ReturnStatementContext) NL() antlr.TerminalNode {
	return s.GetToken(PieParserNL, 0)
}

func (s *ReturnStatementContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *ReturnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PieVisitor:
		return t.VisitReturnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}


type IfStatementContext struct {
	StatementContext
	c ISlistContext 
	el ISlistContext 
}

func NewIfStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfStatementContext {
	var p = new(IfStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}


func (s *IfStatementContext) GetC() ISlistContext { return s.c }

func (s *IfStatementContext) GetEl() ISlistContext { return s.el }


func (s *IfStatementContext) SetC(v ISlistContext) { s.c = v }

func (s *IfStatementContext) SetEl(v ISlistContext) { s.el = v }

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(PieParserIF, 0)
}

func (s *IfStatementContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfStatementContext) AllSlist() []ISlistContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISlistContext); ok {
			len++
		}
	}

	tst := make([]ISlistContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISlistContext); ok {
			tst[i] = t.(ISlistContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) Slist(i int) ISlistContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISlistContext); ok {
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

	return t.(ISlistContext)
}


func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PieVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *PieParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PieParserRULE_statement)
	var _la int

	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStructDefinitionStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(83)
			p.StructDefinition()
		}


	case 2:
		localctx = NewAssignmentStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(84)
			p.Qid()
		}
		{
			p.SetState(85)
			p.Match(PieParserASSIGN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(86)
			p.Expr()
		}
		{
			p.SetState(87)
			p.Match(PieParserNL)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 3:
		localctx = NewReturnStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(89)
			p.Match(PieParserRETURN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 8388608016) != 0) {
			{
				p.SetState(90)
				p.Expr()
			}

		}
		{
			p.SetState(93)
			p.Match(PieParserNL)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 4:
		localctx = NewPrintStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(94)
			p.Match(PieParserPRINT)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(95)
			p.Expr()
		}
		{
			p.SetState(96)
			p.Match(PieParserNL)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 5:
		localctx = NewIfStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(98)
			p.Match(PieParserIF)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(99)
			p.Expr()
		}
		{
			p.SetState(100)

			var _x = p.Slist()


			localctx.(*IfStatementContext).c = _x
		}
		p.SetState(103)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(101)
				p.Match(PieParserT__6)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(102)

				var _x = p.Slist()


				localctx.(*IfStatementContext).el = _x
			}

			} else if p.HasError() { // JIM
				goto errorExit
		}


	case 6:
		localctx = NewWhileStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(105)
			p.Match(PieParserWHILE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(106)
			p.Expr()
		}
		{
			p.SetState(107)
			p.Slist()
		}


	case 7:
		localctx = NewCallStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(109)
			p.Call()
		}
		{
			p.SetState(110)
			p.Match(PieParserNL)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 8:
		localctx = NewNlStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(112)
			p.Match(PieParserNL)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

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


// ICallContext is an interface to support dynamic dispatch.
type ICallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token 


	// SetName sets the name token.
	SetName(antlr.Token) 


	// Getter signatures
	ID() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsCallContext differentiates from other interfaces.
	IsCallContext()
}

type CallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name antlr.Token
}

func NewEmptyCallContext() *CallContext {
	var p = new(CallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PieParserRULE_call
	return p
}

func InitEmptyCallContext(p *CallContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PieParserRULE_call
}

func (*CallContext) IsCallContext() {}

func NewCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallContext {
	var p = new(CallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PieParserRULE_call

	return p
}

func (s *CallContext) GetParser() antlr.Parser { return s.parser }

func (s *CallContext) GetName() antlr.Token { return s.name }


func (s *CallContext) SetName(v antlr.Token) { s.name = v }


func (s *CallContext) ID() antlr.TerminalNode {
	return s.GetToken(PieParserID, 0)
}

func (s *CallContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *CallContext) Expr(i int) IExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *CallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *CallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PieVisitor:
		return t.VisitCall(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *PieParser) Call() (localctx ICallContext) {
	localctx = NewCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PieParserRULE_call)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)

		var _m = p.Match(PieParserID)

		localctx.(*CallContext).name = _m
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(116)
		p.Match(PieParserT__3)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 8388608016) != 0) {
		{
			p.SetState(117)
			p.Expr()
		}
		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == PieParserT__1 {
			{
				p.SetState(118)
				p.Match(PieParserT__1)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(119)
				p.Expr()
			}


			p.SetState(124)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(127)
		p.Match(PieParserT__4)
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


// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAddexpr() []IAddexprContext
	Addexpr(i int) IAddexprContext
	CompOp() ICompOpContext

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PieParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PieParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PieParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) AllAddexpr() []IAddexprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAddexprContext); ok {
			len++
		}
	}

	tst := make([]IAddexprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAddexprContext); ok {
			tst[i] = t.(IAddexprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Addexpr(i int) IAddexprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddexprContext); ok {
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

	return t.(IAddexprContext)
}

func (s *ExprContext) CompOp() ICompOpContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompOpContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompOpContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PieVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *PieParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PieParserRULE_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(129)
		p.Addexpr()
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 16515072) != 0) {
		{
			p.SetState(130)
			p.CompOp()
		}
		{
			p.SetState(131)
			p.Addexpr()
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


// IAddexprContext is an interface to support dynamic dispatch.
type IAddexprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMulexpr() []IMulexprContext
	Mulexpr(i int) IMulexprContext
	AllAddOp() []IAddOpContext
	AddOp(i int) IAddOpContext

	// IsAddexprContext differentiates from other interfaces.
	IsAddexprContext()
}

type AddexprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddexprContext() *AddexprContext {
	var p = new(AddexprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PieParserRULE_addexpr
	return p
}

func InitEmptyAddexprContext(p *AddexprContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PieParserRULE_addexpr
}

func (*AddexprContext) IsAddexprContext() {}

func NewAddexprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddexprContext {
	var p = new(AddexprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PieParserRULE_addexpr

	return p
}

func (s *AddexprContext) GetParser() antlr.Parser { return s.parser }

func (s *AddexprContext) AllMulexpr() []IMulexprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMulexprContext); ok {
			len++
		}
	}

	tst := make([]IMulexprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMulexprContext); ok {
			tst[i] = t.(IMulexprContext)
			i++
		}
	}

	return tst
}

func (s *AddexprContext) Mulexpr(i int) IMulexprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMulexprContext); ok {
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

	return t.(IMulexprContext)
}

func (s *AddexprContext) AllAddOp() []IAddOpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAddOpContext); ok {
			len++
		}
	}

	tst := make([]IAddOpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAddOpContext); ok {
			tst[i] = t.(IAddOpContext)
			i++
		}
	}

	return tst
}

func (s *AddexprContext) AddOp(i int) IAddOpContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddOpContext); ok {
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

	return t.(IAddOpContext)
}

func (s *AddexprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddexprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AddexprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PieVisitor:
		return t.VisitAddexpr(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *PieParser) Addexpr() (localctx IAddexprContext) {
	localctx = NewAddexprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PieParserRULE_addexpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.Mulexpr()
	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == PieParserADD || _la == PieParserSUB {
		{
			p.SetState(136)
			p.AddOp()
		}
		{
			p.SetState(137)
			p.Mulexpr()
		}


		p.SetState(143)
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


// IMulexprContext is an interface to support dynamic dispatch.
type IMulexprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAtom() []IAtomContext
	Atom(i int) IAtomContext
	AllMultOp() []IMultOpContext
	MultOp(i int) IMultOpContext

	// IsMulexprContext differentiates from other interfaces.
	IsMulexprContext()
}

type MulexprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMulexprContext() *MulexprContext {
	var p = new(MulexprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PieParserRULE_mulexpr
	return p
}

func InitEmptyMulexprContext(p *MulexprContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PieParserRULE_mulexpr
}

func (*MulexprContext) IsMulexprContext() {}

func NewMulexprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MulexprContext {
	var p = new(MulexprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PieParserRULE_mulexpr

	return p
}

func (s *MulexprContext) GetParser() antlr.Parser { return s.parser }

func (s *MulexprContext) AllAtom() []IAtomContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAtomContext); ok {
			len++
		}
	}

	tst := make([]IAtomContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAtomContext); ok {
			tst[i] = t.(IAtomContext)
			i++
		}
	}

	return tst
}

func (s *MulexprContext) Atom(i int) IAtomContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomContext); ok {
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

	return t.(IAtomContext)
}

func (s *MulexprContext) AllMultOp() []IMultOpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMultOpContext); ok {
			len++
		}
	}

	tst := make([]IMultOpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMultOpContext); ok {
			tst[i] = t.(IMultOpContext)
			i++
		}
	}

	return tst
}

func (s *MulexprContext) MultOp(i int) IMultOpContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultOpContext); ok {
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

	return t.(IMultOpContext)
}

func (s *MulexprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulexprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MulexprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PieVisitor:
		return t.VisitMulexpr(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *PieParser) Mulexpr() (localctx IMulexprContext) {
	localctx = NewMulexprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PieParserRULE_mulexpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Atom()
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == PieParserMUL || _la == PieParserDIV {
		{
			p.SetState(145)
			p.MultOp()
		}
		{
			p.SetState(146)
			p.Atom()
		}


		p.SetState(152)
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


// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PieParserRULE_atom
	return p
}

func InitEmptyAtomContext(p *AtomContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PieParserRULE_atom
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PieParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) CopyAll(ctx *AtomContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type QidAtomContext struct {
	AtomContext
}

func NewQidAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *QidAtomContext {
	var p = new(QidAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *QidAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QidAtomContext) Qid() IQidContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQidContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQidContext)
}


func (s *QidAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PieVisitor:
		return t.VisitQidAtom(s)

	default:
		return t.VisitChildren(s)
	}
}


type CallAtomContext struct {
	AtomContext
}

func NewCallAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallAtomContext {
	var p = new(CallAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *CallAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallAtomContext) Call() ICallContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallContext)
}


func (s *CallAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PieVisitor:
		return t.VisitCallAtom(s)

	default:
		return t.VisitChildren(s)
	}
}


type FloatAtomContext struct {
	AtomContext
}

func NewFloatAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatAtomContext {
	var p = new(FloatAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *FloatAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatAtomContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(PieParserFLOAT, 0)
}


func (s *FloatAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PieVisitor:
		return t.VisitFloatAtom(s)

	default:
		return t.VisitChildren(s)
	}
}


type IntAtomContext struct {
	AtomContext
}

func NewIntAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntAtomContext {
	var p = new(IntAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *IntAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntAtomContext) INT() antlr.TerminalNode {
	return s.GetToken(PieParserINT, 0)
}


func (s *IntAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PieVisitor:
		return t.VisitIntAtom(s)

	default:
		return t.VisitChildren(s)
	}
}


type StringAtomContext struct {
	AtomContext
}

func NewStringAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringAtomContext {
	var p = new(StringAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *StringAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringAtomContext) STRING() antlr.TerminalNode {
	return s.GetToken(PieParserSTRING, 0)
}


func (s *StringAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PieVisitor:
		return t.VisitStringAtom(s)

	default:
		return t.VisitChildren(s)
	}
}


type ParenthesizedAtomContext struct {
	AtomContext
}

func NewParenthesizedAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesizedAtomContext {
	var p = new(ParenthesizedAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *ParenthesizedAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesizedAtomContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *ParenthesizedAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PieVisitor:
		return t.VisitParenthesizedAtom(s)

	default:
		return t.VisitChildren(s)
	}
}


type InstanceAtomContext struct {
	AtomContext
}

func NewInstanceAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InstanceAtomContext {
	var p = new(InstanceAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *InstanceAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstanceAtomContext) Instance() IInstanceContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstanceContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstanceContext)
}


func (s *InstanceAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PieVisitor:
		return t.VisitInstanceAtom(s)

	default:
		return t.VisitChildren(s)
	}
}


type CharAtomContext struct {
	AtomContext
}

func NewCharAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CharAtomContext {
	var p = new(CharAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *CharAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharAtomContext) CHAR() antlr.TerminalNode {
	return s.GetToken(PieParserCHAR, 0)
}


func (s *CharAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PieVisitor:
		return t.VisitCharAtom(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *PieParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PieParserRULE_atom)
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIntAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(153)
			p.Match(PieParserINT)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 2:
		localctx = NewCharAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(154)
			p.Match(PieParserCHAR)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 3:
		localctx = NewFloatAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(155)
			p.Match(PieParserFLOAT)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 4:
		localctx = NewStringAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(156)
			p.Match(PieParserSTRING)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 5:
		localctx = NewQidAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(157)
			p.Qid()
		}


	case 6:
		localctx = NewCallAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(158)
			p.Call()
		}


	case 7:
		localctx = NewInstanceAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(159)
			p.Instance()
		}


	case 8:
		localctx = NewParenthesizedAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(160)
			p.Match(PieParserT__3)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(161)
			p.Expr()
		}
		{
			p.SetState(162)
			p.Match(PieParserT__4)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

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


// IInstanceContext is an interface to support dynamic dispatch.
type IInstanceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSname returns the sname token.
	GetSname() antlr.Token 


	// SetSname sets the sname token.
	SetSname(antlr.Token) 


	// Getter signatures
	NEW() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsInstanceContext differentiates from other interfaces.
	IsInstanceContext()
}

type InstanceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	sname antlr.Token
}

func NewEmptyInstanceContext() *InstanceContext {
	var p = new(InstanceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PieParserRULE_instance
	return p
}

func InitEmptyInstanceContext(p *InstanceContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PieParserRULE_instance
}

func (*InstanceContext) IsInstanceContext() {}

func NewInstanceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstanceContext {
	var p = new(InstanceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PieParserRULE_instance

	return p
}

func (s *InstanceContext) GetParser() antlr.Parser { return s.parser }

func (s *InstanceContext) GetSname() antlr.Token { return s.sname }


func (s *InstanceContext) SetSname(v antlr.Token) { s.sname = v }


func (s *InstanceContext) NEW() antlr.TerminalNode {
	return s.GetToken(PieParserNEW, 0)
}

func (s *InstanceContext) ID() antlr.TerminalNode {
	return s.GetToken(PieParserID, 0)
}

func (s *InstanceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstanceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *InstanceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PieVisitor:
		return t.VisitInstance(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *PieParser) Instance() (localctx IInstanceContext) {
	localctx = NewInstanceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, PieParserRULE_instance)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(166)
		p.Match(PieParserNEW)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(167)

		var _m = p.Match(PieParserID)

		localctx.(*InstanceContext).sname = _m
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


// IQidContext is an interface to support dynamic dispatch.
type IQidContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsQidContext differentiates from other interfaces.
	IsQidContext()
}

type QidContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQidContext() *QidContext {
	var p = new(QidContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PieParserRULE_qid
	return p
}

func InitEmptyQidContext(p *QidContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PieParserRULE_qid
}

func (*QidContext) IsQidContext() {}

func NewQidContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QidContext {
	var p = new(QidContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PieParserRULE_qid

	return p
}

func (s *QidContext) GetParser() antlr.Parser { return s.parser }

func (s *QidContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(PieParserID)
}

func (s *QidContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(PieParserID, i)
}

func (s *QidContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(PieParserDOT)
}

func (s *QidContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(PieParserDOT, i)
}

func (s *QidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QidContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *QidContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PieVisitor:
		return t.VisitQid(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *PieParser) Qid() (localctx IQidContext) {
	localctx = NewQidContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, PieParserRULE_qid)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(PieParserID)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == PieParserDOT {
		{
			p.SetState(170)
			p.Match(PieParserDOT)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(171)
			p.Match(PieParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


		p.SetState(176)
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


// IVardefContext is an interface to support dynamic dispatch.
type IVardefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsVardefContext differentiates from other interfaces.
	IsVardefContext()
}

type VardefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVardefContext() *VardefContext {
	var p = new(VardefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PieParserRULE_vardef
	return p
}

func InitEmptyVardefContext(p *VardefContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PieParserRULE_vardef
}

func (*VardefContext) IsVardefContext() {}

func NewVardefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VardefContext {
	var p = new(VardefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PieParserRULE_vardef

	return p
}

func (s *VardefContext) GetParser() antlr.Parser { return s.parser }

func (s *VardefContext) ID() antlr.TerminalNode {
	return s.GetToken(PieParserID, 0)
}

func (s *VardefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VardefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *VardefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PieVisitor:
		return t.VisitVardef(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *PieParser) Vardef() (localctx IVardefContext) {
	localctx = NewVardefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, PieParserRULE_vardef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(177)
		p.Match(PieParserID)
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


// IMultOpContext is an interface to support dynamic dispatch.
type IMultOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MUL() antlr.TerminalNode
	DIV() antlr.TerminalNode

	// IsMultOpContext differentiates from other interfaces.
	IsMultOpContext()
}

type MultOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultOpContext() *MultOpContext {
	var p = new(MultOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PieParserRULE_multOp
	return p
}

func InitEmptyMultOpContext(p *MultOpContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PieParserRULE_multOp
}

func (*MultOpContext) IsMultOpContext() {}

func NewMultOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultOpContext {
	var p = new(MultOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PieParserRULE_multOp

	return p
}

func (s *MultOpContext) GetParser() antlr.Parser { return s.parser }

func (s *MultOpContext) MUL() antlr.TerminalNode {
	return s.GetToken(PieParserMUL, 0)
}

func (s *MultOpContext) DIV() antlr.TerminalNode {
	return s.GetToken(PieParserDIV, 0)
}

func (s *MultOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MultOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PieVisitor:
		return t.VisitMultOp(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *PieParser) MultOp() (localctx IMultOpContext) {
	localctx = NewMultOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, PieParserRULE_multOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PieParserMUL || _la == PieParserDIV) {
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


// IAddOpContext is an interface to support dynamic dispatch.
type IAddOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ADD() antlr.TerminalNode
	SUB() antlr.TerminalNode

	// IsAddOpContext differentiates from other interfaces.
	IsAddOpContext()
}

type AddOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddOpContext() *AddOpContext {
	var p = new(AddOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PieParserRULE_addOp
	return p
}

func InitEmptyAddOpContext(p *AddOpContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PieParserRULE_addOp
}

func (*AddOpContext) IsAddOpContext() {}

func NewAddOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddOpContext {
	var p = new(AddOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PieParserRULE_addOp

	return p
}

func (s *AddOpContext) GetParser() antlr.Parser { return s.parser }

func (s *AddOpContext) ADD() antlr.TerminalNode {
	return s.GetToken(PieParserADD, 0)
}

func (s *AddOpContext) SUB() antlr.TerminalNode {
	return s.GetToken(PieParserSUB, 0)
}

func (s *AddOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AddOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PieVisitor:
		return t.VisitAddOp(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *PieParser) AddOp() (localctx IAddOpContext) {
	localctx = NewAddOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, PieParserRULE_addOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PieParserADD || _la == PieParserSUB) {
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


// ICompOpContext is an interface to support dynamic dispatch.
type ICompOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQ() antlr.TerminalNode
	LT() antlr.TerminalNode
	GT() antlr.TerminalNode
	NEQ() antlr.TerminalNode
	GEQ() antlr.TerminalNode
	LEQ() antlr.TerminalNode

	// IsCompOpContext differentiates from other interfaces.
	IsCompOpContext()
}

type CompOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompOpContext() *CompOpContext {
	var p = new(CompOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PieParserRULE_compOp
	return p
}

func InitEmptyCompOpContext(p *CompOpContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PieParserRULE_compOp
}

func (*CompOpContext) IsCompOpContext() {}

func NewCompOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompOpContext {
	var p = new(CompOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PieParserRULE_compOp

	return p
}

func (s *CompOpContext) GetParser() antlr.Parser { return s.parser }

func (s *CompOpContext) EQ() antlr.TerminalNode {
	return s.GetToken(PieParserEQ, 0)
}

func (s *CompOpContext) LT() antlr.TerminalNode {
	return s.GetToken(PieParserLT, 0)
}

func (s *CompOpContext) GT() antlr.TerminalNode {
	return s.GetToken(PieParserGT, 0)
}

func (s *CompOpContext) NEQ() antlr.TerminalNode {
	return s.GetToken(PieParserNEQ, 0)
}

func (s *CompOpContext) GEQ() antlr.TerminalNode {
	return s.GetToken(PieParserGEQ, 0)
}

func (s *CompOpContext) LEQ() antlr.TerminalNode {
	return s.GetToken(PieParserLEQ, 0)
}

func (s *CompOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *CompOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PieVisitor:
		return t.VisitCompOp(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *PieParser) CompOp() (localctx ICompOpContext) {
	localctx = NewCompOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, PieParserRULE_compOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
		_la = p.GetTokenStream().LA(1)

		if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 16515072) != 0)) {
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


