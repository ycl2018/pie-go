// Code generated from github.com/ycl2018/asm/gen/Assembler.g4 by ANTLR 4.13.1. DO NOT EDIT.

package gen // Assembler
import (
	"fmt"
	"strconv"
  	"sync"

	"github.com/antlr4-go/antlr/v4"
)



type AsmGenerator interface {
    Gen(instrs ...antlr.Token)
    CheckForUnresolvedReferences()
    DefineFunction(id antlr.Token, nargs, nlocals int)
    DefineDataSize(n int)
    DefineLabel(id antlr.Token)
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
    "", "'.globals'", "'.def'", "':'", "'args'", "'='", "','", "'locals'",
  }
  staticData.SymbolicNames = []string{
    "", "", "", "", "", "", "", "", "REG", "ID", "FUNC", "INT", "CHAR", 
    "STRING", "FLOAT", "WS", "NEWLINE",
  }
  staticData.RuleNames = []string{
    "program", "globals", "functionDeclaration", "instr", "operand", "label",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 16, 82, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 2, 5, 7, 5, 1, 0, 3, 0, 14, 8, 0, 1, 0, 1, 0, 1, 0, 1, 0, 4, 0, 20, 
	8, 0, 11, 0, 12, 0, 21, 1, 0, 1, 0, 1, 1, 5, 1, 27, 8, 1, 10, 1, 12, 1, 
	30, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 
	2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 
	3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 
	3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 74, 8, 3, 1, 4, 1, 4, 
	1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 0, 0, 6, 0, 2, 4, 6, 8, 10, 0, 1, 1, 0, 8, 
	14, 84, 0, 13, 1, 0, 0, 0, 2, 28, 1, 0, 0, 0, 4, 36, 1, 0, 0, 0, 6, 73, 
	1, 0, 0, 0, 8, 75, 1, 0, 0, 0, 10, 77, 1, 0, 0, 0, 12, 14, 3, 2, 1, 0, 
	13, 12, 1, 0, 0, 0, 13, 14, 1, 0, 0, 0, 14, 19, 1, 0, 0, 0, 15, 20, 3, 
	4, 2, 0, 16, 20, 3, 6, 3, 0, 17, 20, 3, 10, 5, 0, 18, 20, 5, 16, 0, 0, 
	19, 15, 1, 0, 0, 0, 19, 16, 1, 0, 0, 0, 19, 17, 1, 0, 0, 0, 19, 18, 1, 
	0, 0, 0, 20, 21, 1, 0, 0, 0, 21, 19, 1, 0, 0, 0, 21, 22, 1, 0, 0, 0, 22, 
	23, 1, 0, 0, 0, 23, 24, 6, 0, -1, 0, 24, 1, 1, 0, 0, 0, 25, 27, 5, 16, 
	0, 0, 26, 25, 1, 0, 0, 0, 27, 30, 1, 0, 0, 0, 28, 26, 1, 0, 0, 0, 28, 29, 
	1, 0, 0, 0, 29, 31, 1, 0, 0, 0, 30, 28, 1, 0, 0, 0, 31, 32, 5, 1, 0, 0, 
	32, 33, 5, 11, 0, 0, 33, 34, 5, 16, 0, 0, 34, 35, 6, 1, -1, 0, 35, 3, 1, 
	0, 0, 0, 36, 37, 5, 2, 0, 0, 37, 38, 5, 9, 0, 0, 38, 39, 5, 3, 0, 0, 39, 
	40, 5, 4, 0, 0, 40, 41, 5, 5, 0, 0, 41, 42, 5, 11, 0, 0, 42, 43, 5, 6, 
	0, 0, 43, 44, 5, 7, 0, 0, 44, 45, 5, 5, 0, 0, 45, 46, 5, 11, 0, 0, 46, 
	47, 5, 16, 0, 0, 47, 48, 6, 2, -1, 0, 48, 5, 1, 0, 0, 0, 49, 50, 5, 9, 
	0, 0, 50, 51, 5, 16, 0, 0, 51, 74, 6, 3, -1, 0, 52, 53, 5, 9, 0, 0, 53, 
	54, 3, 8, 4, 0, 54, 55, 5, 16, 0, 0, 55, 56, 6, 3, -1, 0, 56, 74, 1, 0, 
	0, 0, 57, 58, 5, 9, 0, 0, 58, 59, 3, 8, 4, 0, 59, 60, 5, 6, 0, 0, 60, 61, 
	3, 8, 4, 0, 61, 62, 5, 16, 0, 0, 62, 63, 6, 3, -1, 0, 63, 74, 1, 0, 0, 
	0, 64, 65, 5, 9, 0, 0, 65, 66, 3, 8, 4, 0, 66, 67, 5, 6, 0, 0, 67, 68, 
	3, 8, 4, 0, 68, 69, 5, 6, 0, 0, 69, 70, 3, 8, 4, 0, 70, 71, 5, 16, 0, 0, 
	71, 72, 6, 3, -1, 0, 72, 74, 1, 0, 0, 0, 73, 49, 1, 0, 0, 0, 73, 52, 1, 
	0, 0, 0, 73, 57, 1, 0, 0, 0, 73, 64, 1, 0, 0, 0, 74, 7, 1, 0, 0, 0, 75, 
	76, 7, 0, 0, 0, 76, 9, 1, 0, 0, 0, 77, 78, 5, 9, 0, 0, 78, 79, 5, 3, 0, 
	0, 79, 80, 6, 5, -1, 0, 80, 11, 1, 0, 0, 0, 5, 13, 19, 21, 28, 73,
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
	AssemblerParserREG = 8
	AssemblerParserID = 9
	AssemblerParserFUNC = 10
	AssemblerParserINT = 11
	AssemblerParserCHAR = 12
	AssemblerParserSTRING = 13
	AssemblerParserFLOAT = 14
	AssemblerParserWS = 15
	AssemblerParserNEWLINE = 16
)

// AssemblerParser rules.
const (
	AssemblerParserRULE_program = 0
	AssemblerParserRULE_globals = 1
	AssemblerParserRULE_functionDeclaration = 2
	AssemblerParserRULE_instr = 3
	AssemblerParserRULE_operand = 4
	AssemblerParserRULE_label = 5
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Globals() IGlobalsContext
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


func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.ExitProgram(s)
	}
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
	p.SetState(13)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(12)
			p.Globals()
		}

		} else if p.HasError() { // JIM
			goto errorExit
	}
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 66052) != 0) {
		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(15)
				p.FunctionDeclaration()
			}


		case 2:
			{
				p.SetState(16)
				p.Instr()
			}


		case 3:
			{
				p.SetState(17)
				p.Label()
			}


		case 4:
			{
				p.SetState(18)
				p.Match(AssemblerParserNEWLINE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(21)
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


func (s *GlobalsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.EnterGlobals(s)
	}
}

func (s *GlobalsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.ExitGlobals(s)
	}
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
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == AssemblerParserNEWLINE {
		{
			p.SetState(25)
			p.Match(AssemblerParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(31)
		p.Match(AssemblerParserT__0)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(32)

		var _m = p.Match(AssemblerParserINT)

		localctx.(*GlobalsContext)._INT = _m
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(33)
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


	// Getter signatures
	NEWLINE() antlr.TerminalNode
	ID() antlr.TerminalNode
	AllINT() []antlr.TerminalNode
	INT(i int) antlr.TerminalNode

	// IsFunctionDeclarationContext differentiates from other interfaces.
	IsFunctionDeclarationContext()
}

type FunctionDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name antlr.Token
	a antlr.Token
	n antlr.Token
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

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FunctionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.EnterFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.ExitFunctionDeclaration(s)
	}
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
	p.EnterRule(localctx, 4, AssemblerParserRULE_functionDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(36)
		p.Match(AssemblerParserT__1)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(37)

		var _m = p.Match(AssemblerParserID)

		localctx.(*FunctionDeclarationContext).name = _m
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(38)
		p.Match(AssemblerParserT__2)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(39)
		p.Match(AssemblerParserT__3)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(40)
		p.Match(AssemblerParserT__4)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(41)

		var _m = p.Match(AssemblerParserINT)

		localctx.(*FunctionDeclarationContext).a = _m
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(42)
		p.Match(AssemblerParserT__5)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(43)
		p.Match(AssemblerParserT__6)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(44)
		p.Match(AssemblerParserT__4)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(45)

		var _m = p.Match(AssemblerParserINT)

		localctx.(*FunctionDeclarationContext).n = _m
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(46)
		p.Match(AssemblerParserNEWLINE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.DefineFunction(localctx.(*FunctionDeclarationContext).GetName(), (func() int { if localctx.(*FunctionDeclarationContext).GetA() == nil { return 0 } else { i, _ := strconv.Atoi(localctx.(*FunctionDeclarationContext).GetA().GetText()); return i }}()), (func() int { if localctx.(*FunctionDeclarationContext).GetN() == nil { return 0 } else { i, _ := strconv.Atoi(localctx.(*FunctionDeclarationContext).GetN().GetText()); return i }}()))



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
	parser   antlr.Parser
	_ID      antlr.Token
	_operand IOperandContext
	a        IOperandContext
	b        IOperandContext
	c        IOperandContext
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


func (s *InstrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.EnterInstr(s)
	}
}

func (s *InstrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.ExitInstr(s)
	}
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
	p.EnterRule(localctx, 6, AssemblerParserRULE_instr)
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(49)

			var _m = p.Match(AssemblerParserID)

			localctx.(*InstrContext)._ID = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(50)
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
			p.SetState(52)

			var _m = p.Match(AssemblerParserID)

			localctx.(*InstrContext)._ID = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(53)

			var _x = p.Operand()


			localctx.(*InstrContext)._operand = _x
		}
		{
			p.SetState(54)
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
			p.SetState(57)

			var _m = p.Match(AssemblerParserID)

			localctx.(*InstrContext)._ID = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(58)

			var _x = p.Operand()


			localctx.(*InstrContext).a = _x
		}
		{
			p.SetState(59)
			p.Match(AssemblerParserT__5)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(60)

			var _x = p.Operand()


			localctx.(*InstrContext).b = _x
		}
		{
			p.SetState(61)
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
			p.SetState(64)

			var _m = p.Match(AssemblerParserID)

			localctx.(*InstrContext)._ID = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(65)

			var _x = p.Operand()


			localctx.(*InstrContext).a = _x
		}
		{
			p.SetState(66)
			p.Match(AssemblerParserT__5)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(67)

			var _x = p.Operand()


			localctx.(*InstrContext).b = _x
		}
		{
			p.SetState(68)
			p.Match(AssemblerParserT__5)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(69)

			var _x = p.Operand()


			localctx.(*InstrContext).c = _x
		}
		{
			p.SetState(70)
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


func (s *OperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.EnterOperand(s)
	}
}

func (s *OperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.ExitOperand(s)
	}
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
	p.EnterRule(localctx, 8, AssemblerParserRULE_operand)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		_la = p.GetTokenStream().LA(1)

		if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 32512) != 0)) {
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


func (s *LabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.EnterLabel(s)
	}
}

func (s *LabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.ExitLabel(s)
	}
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
	p.EnterRule(localctx, 10, AssemblerParserRULE_label)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)

		var _m = p.Match(AssemblerParserID)

		localctx.(*LabelContext)._ID = _m
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(78)
		p.Match(AssemblerParserT__2)
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


