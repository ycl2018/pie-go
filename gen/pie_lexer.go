// Code generated from github.com/ycl2018/pie-go/gen/Pie.g4 by ANTLR 4.13.1. DO NOT EDIT.

package gen
import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr4-go/antlr/v4"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter


type PieLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var PieLexerLexerStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  ChannelNames           []string
  ModeNames              []string
  LiteralNames           []string
  SymbolicNames          []string
  RuleNames              []string
  PredictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func pielexerLexerInit() {
  staticData := &PieLexerLexerStaticData
  staticData.ChannelNames = []string{
    "DEFAULT_TOKEN_CHANNEL", "HIDDEN",
  }
  staticData.ModeNames = []string{
    "DEFAULT_MODE",
  }
  staticData.LiteralNames = []string{
    "", "'{'", "','", "'}'", "'('", "')'", "':'", "'else'", "'if'", "'='", 
    "'print'", "'while'", "'return'", "'def'", "'+'", "'-'", "'*'", "'/'", 
    "'=='", "'<'", "'struct'", "'.'", "'new'",
  }
  staticData.SymbolicNames = []string{
    "", "", "", "", "", "", "", "", "IF", "ASSIGN", "PRINT", "WHILE", "RETURN", 
    "DEF", "ADD", "SUB", "MUL", "DIV", "EQ", "LT", "STRUCT", "DOT", "NEW", 
    "NL", "ID", "CHAR", "STRING", "INT", "FLOAT", "WS", "SL_COMMENT",
  }
  staticData.RuleNames = []string{
    "T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "IF", "ASSIGN", 
    "PRINT", "WHILE", "RETURN", "DEF", "ADD", "SUB", "MUL", "DIV", "EQ", 
    "LT", "STRUCT", "DOT", "NEW", "NL", "ID", "LETTER", "CHAR", "STRING", 
    "INT", "FLOAT", "WS", "SL_COMMENT",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 0, 30, 196, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 
	4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 
	10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 
	7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 
	20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 
	2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 1, 
	0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 
	6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 
	9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 
	11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 
	1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 
	18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 
	1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 3, 22, 136, 8, 22, 1, 22, 1, 22, 1, 
	23, 1, 23, 1, 23, 5, 23, 143, 8, 23, 10, 23, 12, 23, 146, 9, 23, 1, 24, 
	1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 5, 26, 156, 8, 26, 10, 
	26, 12, 26, 159, 9, 26, 1, 26, 1, 26, 1, 27, 4, 27, 164, 8, 27, 11, 27, 
	12, 27, 165, 1, 28, 1, 28, 1, 28, 5, 28, 171, 8, 28, 10, 28, 12, 28, 174, 
	9, 28, 1, 28, 1, 28, 4, 28, 178, 8, 28, 11, 28, 12, 28, 179, 3, 28, 182, 
	8, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 5, 30, 190, 8, 30, 10, 
	30, 12, 30, 193, 9, 30, 1, 30, 1, 30, 1, 157, 0, 31, 1, 1, 3, 2, 5, 3, 
	7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 
	27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 
	45, 23, 47, 24, 49, 0, 51, 25, 53, 26, 55, 27, 57, 28, 59, 29, 61, 30, 
	1, 0, 3, 2, 0, 65, 90, 97, 122, 2, 0, 9, 9, 32, 32, 2, 0, 10, 10, 13, 13, 
	203, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 
	0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 
	0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 
	1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 
	31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 
	0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 
	0, 0, 47, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 
	0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 1, 63, 1, 
	0, 0, 0, 3, 65, 1, 0, 0, 0, 5, 67, 1, 0, 0, 0, 7, 69, 1, 0, 0, 0, 9, 71, 
	1, 0, 0, 0, 11, 73, 1, 0, 0, 0, 13, 75, 1, 0, 0, 0, 15, 80, 1, 0, 0, 0, 
	17, 83, 1, 0, 0, 0, 19, 85, 1, 0, 0, 0, 21, 91, 1, 0, 0, 0, 23, 97, 1, 
	0, 0, 0, 25, 104, 1, 0, 0, 0, 27, 108, 1, 0, 0, 0, 29, 110, 1, 0, 0, 0, 
	31, 112, 1, 0, 0, 0, 33, 114, 1, 0, 0, 0, 35, 116, 1, 0, 0, 0, 37, 119, 
	1, 0, 0, 0, 39, 121, 1, 0, 0, 0, 41, 128, 1, 0, 0, 0, 43, 130, 1, 0, 0, 
	0, 45, 135, 1, 0, 0, 0, 47, 139, 1, 0, 0, 0, 49, 147, 1, 0, 0, 0, 51, 149, 
	1, 0, 0, 0, 53, 153, 1, 0, 0, 0, 55, 163, 1, 0, 0, 0, 57, 181, 1, 0, 0, 
	0, 59, 183, 1, 0, 0, 0, 61, 187, 1, 0, 0, 0, 63, 64, 5, 123, 0, 0, 64, 
	2, 1, 0, 0, 0, 65, 66, 5, 44, 0, 0, 66, 4, 1, 0, 0, 0, 67, 68, 5, 125, 
	0, 0, 68, 6, 1, 0, 0, 0, 69, 70, 5, 40, 0, 0, 70, 8, 1, 0, 0, 0, 71, 72, 
	5, 41, 0, 0, 72, 10, 1, 0, 0, 0, 73, 74, 5, 58, 0, 0, 74, 12, 1, 0, 0, 
	0, 75, 76, 5, 101, 0, 0, 76, 77, 5, 108, 0, 0, 77, 78, 5, 115, 0, 0, 78, 
	79, 5, 101, 0, 0, 79, 14, 1, 0, 0, 0, 80, 81, 5, 105, 0, 0, 81, 82, 5, 
	102, 0, 0, 82, 16, 1, 0, 0, 0, 83, 84, 5, 61, 0, 0, 84, 18, 1, 0, 0, 0, 
	85, 86, 5, 112, 0, 0, 86, 87, 5, 114, 0, 0, 87, 88, 5, 105, 0, 0, 88, 89, 
	5, 110, 0, 0, 89, 90, 5, 116, 0, 0, 90, 20, 1, 0, 0, 0, 91, 92, 5, 119, 
	0, 0, 92, 93, 5, 104, 0, 0, 93, 94, 5, 105, 0, 0, 94, 95, 5, 108, 0, 0, 
	95, 96, 5, 101, 0, 0, 96, 22, 1, 0, 0, 0, 97, 98, 5, 114, 0, 0, 98, 99, 
	5, 101, 0, 0, 99, 100, 5, 116, 0, 0, 100, 101, 5, 117, 0, 0, 101, 102, 
	5, 114, 0, 0, 102, 103, 5, 110, 0, 0, 103, 24, 1, 0, 0, 0, 104, 105, 5, 
	100, 0, 0, 105, 106, 5, 101, 0, 0, 106, 107, 5, 102, 0, 0, 107, 26, 1, 
	0, 0, 0, 108, 109, 5, 43, 0, 0, 109, 28, 1, 0, 0, 0, 110, 111, 5, 45, 0, 
	0, 111, 30, 1, 0, 0, 0, 112, 113, 5, 42, 0, 0, 113, 32, 1, 0, 0, 0, 114, 
	115, 5, 47, 0, 0, 115, 34, 1, 0, 0, 0, 116, 117, 5, 61, 0, 0, 117, 118, 
	5, 61, 0, 0, 118, 36, 1, 0, 0, 0, 119, 120, 5, 60, 0, 0, 120, 38, 1, 0, 
	0, 0, 121, 122, 5, 115, 0, 0, 122, 123, 5, 116, 0, 0, 123, 124, 5, 114, 
	0, 0, 124, 125, 5, 117, 0, 0, 125, 126, 5, 99, 0, 0, 126, 127, 5, 116, 
	0, 0, 127, 40, 1, 0, 0, 0, 128, 129, 5, 46, 0, 0, 129, 42, 1, 0, 0, 0, 
	130, 131, 5, 110, 0, 0, 131, 132, 5, 101, 0, 0, 132, 133, 5, 119, 0, 0, 
	133, 44, 1, 0, 0, 0, 134, 136, 5, 13, 0, 0, 135, 134, 1, 0, 0, 0, 135, 
	136, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 138, 5, 10, 0, 0, 138, 46, 
	1, 0, 0, 0, 139, 144, 3, 49, 24, 0, 140, 143, 3, 49, 24, 0, 141, 143, 2, 
	48, 57, 0, 142, 140, 1, 0, 0, 0, 142, 141, 1, 0, 0, 0, 143, 146, 1, 0, 
	0, 0, 144, 142, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 48, 1, 0, 0, 0, 
	146, 144, 1, 0, 0, 0, 147, 148, 7, 0, 0, 0, 148, 50, 1, 0, 0, 0, 149, 150, 
	5, 39, 0, 0, 150, 151, 9, 0, 0, 0, 151, 152, 5, 39, 0, 0, 152, 52, 1, 0, 
	0, 0, 153, 157, 5, 34, 0, 0, 154, 156, 9, 0, 0, 0, 155, 154, 1, 0, 0, 0, 
	156, 159, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 158, 
	160, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 160, 161, 5, 34, 0, 0, 161, 54, 
	1, 0, 0, 0, 162, 164, 2, 48, 57, 0, 163, 162, 1, 0, 0, 0, 164, 165, 1, 
	0, 0, 0, 165, 163, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 56, 1, 0, 0, 
	0, 167, 168, 3, 55, 27, 0, 168, 172, 5, 46, 0, 0, 169, 171, 3, 55, 27, 
	0, 170, 169, 1, 0, 0, 0, 171, 174, 1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 172, 
	173, 1, 0, 0, 0, 173, 182, 1, 0, 0, 0, 174, 172, 1, 0, 0, 0, 175, 177, 
	5, 46, 0, 0, 176, 178, 3, 55, 27, 0, 177, 176, 1, 0, 0, 0, 178, 179, 1, 
	0, 0, 0, 179, 177, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 182, 1, 0, 0, 
	0, 181, 167, 1, 0, 0, 0, 181, 175, 1, 0, 0, 0, 182, 58, 1, 0, 0, 0, 183, 
	184, 7, 1, 0, 0, 184, 185, 1, 0, 0, 0, 185, 186, 6, 29, 0, 0, 186, 60, 
	1, 0, 0, 0, 187, 191, 5, 35, 0, 0, 188, 190, 8, 2, 0, 0, 189, 188, 1, 0, 
	0, 0, 190, 193, 1, 0, 0, 0, 191, 189, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 
	192, 194, 1, 0, 0, 0, 193, 191, 1, 0, 0, 0, 194, 195, 6, 30, 0, 0, 195, 
	62, 1, 0, 0, 0, 10, 0, 135, 142, 144, 157, 165, 172, 179, 181, 191, 1, 
	0, 1, 0,
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

// PieLexerInit initializes any static state used to implement PieLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewPieLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func PieLexerInit() {
  staticData := &PieLexerLexerStaticData
  staticData.once.Do(pielexerLexerInit)
}

// NewPieLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewPieLexer(input antlr.CharStream) *PieLexer {
  PieLexerInit()
	l := new(PieLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
  staticData := &PieLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Pie.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// PieLexer tokens.
const (
	PieLexerT__0 = 1
	PieLexerT__1 = 2
	PieLexerT__2 = 3
	PieLexerT__3 = 4
	PieLexerT__4 = 5
	PieLexerT__5 = 6
	PieLexerT__6 = 7
	PieLexerIF = 8
	PieLexerASSIGN = 9
	PieLexerPRINT = 10
	PieLexerWHILE = 11
	PieLexerRETURN = 12
	PieLexerDEF = 13
	PieLexerADD = 14
	PieLexerSUB = 15
	PieLexerMUL = 16
	PieLexerDIV = 17
	PieLexerEQ = 18
	PieLexerLT = 19
	PieLexerSTRUCT = 20
	PieLexerDOT = 21
	PieLexerNEW = 22
	PieLexerNL = 23
	PieLexerID = 24
	PieLexerCHAR = 25
	PieLexerSTRING = 26
	PieLexerINT = 27
	PieLexerFLOAT = 28
	PieLexerWS = 29
	PieLexerSL_COMMENT = 30
)

