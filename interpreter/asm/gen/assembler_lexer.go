// Code generated from /Users/bytedance/go/src/github.com/ycl2018/asm/gen/Assembler.g4 by ANTLR 4.13.1. DO NOT EDIT.

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


type AssemblerLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var AssemblerLexerLexerStaticData struct {
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

func assemblerlexerLexerInit() {
  staticData := &AssemblerLexerLexerStaticData
  staticData.ChannelNames = []string{
    "DEFAULT_TOKEN_CHANNEL", "HIDDEN",
  }
  staticData.ModeNames = []string{
    "DEFAULT_MODE",
  }
  staticData.LiteralNames = []string{
    "", "'.globals'", "'.def'", "':'", "'args'", "'='", "','", "'locals'",
  }
  staticData.SymbolicNames = []string{
    "", "", "", "", "", "", "", "", "REG", "ID", "FUNC", "INT", "CHAR", 
    "STRING", "FLOAT", "WS", "NEWLINE",
  }
  staticData.RuleNames = []string{
    "T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "REG", "ID", 
    "FUNC", "LETTER", "INT", "CHAR", "STRING", "STR_CHARS", "FLOAT", "WS", 
    "NEWLINE",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 0, 16, 145, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 
	4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 
	10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 
	7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 
	1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 
	1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 
	1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 5, 8, 76, 8, 8, 10, 8, 
	12, 8, 79, 9, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 3, 11, 88, 
	8, 11, 1, 11, 4, 11, 91, 8, 11, 11, 11, 12, 11, 92, 1, 12, 1, 12, 1, 12, 
	1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 5, 14, 104, 8, 14, 10, 14, 12, 
	14, 107, 9, 14, 1, 15, 1, 15, 1, 15, 5, 15, 112, 8, 15, 10, 15, 12, 15, 
	115, 9, 15, 1, 15, 1, 15, 4, 15, 119, 8, 15, 11, 15, 12, 15, 120, 3, 15, 
	123, 8, 15, 1, 16, 4, 16, 126, 8, 16, 11, 16, 12, 16, 127, 1, 16, 1, 16, 
	1, 17, 1, 17, 5, 17, 134, 8, 17, 10, 17, 12, 17, 137, 9, 17, 3, 17, 139, 
	8, 17, 1, 17, 3, 17, 142, 8, 17, 1, 17, 1, 17, 1, 135, 0, 18, 1, 1, 3, 
	2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 0, 23, 11, 
	25, 12, 27, 13, 29, 0, 31, 14, 33, 15, 35, 16, 1, 0, 3, 2, 0, 65, 90, 97, 
	122, 1, 0, 34, 34, 2, 0, 9, 9, 32, 32, 154, 0, 1, 1, 0, 0, 0, 0, 3, 1, 
	0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 
	0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 
	1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 
	31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 1, 37, 1, 0, 0, 0, 
	3, 46, 1, 0, 0, 0, 5, 51, 1, 0, 0, 0, 7, 53, 1, 0, 0, 0, 9, 58, 1, 0, 0, 
	0, 11, 60, 1, 0, 0, 0, 13, 62, 1, 0, 0, 0, 15, 69, 1, 0, 0, 0, 17, 72, 
	1, 0, 0, 0, 19, 80, 1, 0, 0, 0, 21, 84, 1, 0, 0, 0, 23, 87, 1, 0, 0, 0, 
	25, 94, 1, 0, 0, 0, 27, 98, 1, 0, 0, 0, 29, 105, 1, 0, 0, 0, 31, 122, 1, 
	0, 0, 0, 33, 125, 1, 0, 0, 0, 35, 138, 1, 0, 0, 0, 37, 38, 5, 46, 0, 0, 
	38, 39, 5, 103, 0, 0, 39, 40, 5, 108, 0, 0, 40, 41, 5, 111, 0, 0, 41, 42, 
	5, 98, 0, 0, 42, 43, 5, 97, 0, 0, 43, 44, 5, 108, 0, 0, 44, 45, 5, 115, 
	0, 0, 45, 2, 1, 0, 0, 0, 46, 47, 5, 46, 0, 0, 47, 48, 5, 100, 0, 0, 48, 
	49, 5, 101, 0, 0, 49, 50, 5, 102, 0, 0, 50, 4, 1, 0, 0, 0, 51, 52, 5, 58, 
	0, 0, 52, 6, 1, 0, 0, 0, 53, 54, 5, 97, 0, 0, 54, 55, 5, 114, 0, 0, 55, 
	56, 5, 103, 0, 0, 56, 57, 5, 115, 0, 0, 57, 8, 1, 0, 0, 0, 58, 59, 5, 61, 
	0, 0, 59, 10, 1, 0, 0, 0, 60, 61, 5, 44, 0, 0, 61, 12, 1, 0, 0, 0, 62, 
	63, 5, 108, 0, 0, 63, 64, 5, 111, 0, 0, 64, 65, 5, 99, 0, 0, 65, 66, 5, 
	97, 0, 0, 66, 67, 5, 108, 0, 0, 67, 68, 5, 115, 0, 0, 68, 14, 1, 0, 0, 
	0, 69, 70, 5, 114, 0, 0, 70, 71, 3, 23, 11, 0, 71, 16, 1, 0, 0, 0, 72, 
	77, 3, 21, 10, 0, 73, 76, 3, 21, 10, 0, 74, 76, 2, 48, 57, 0, 75, 73, 1, 
	0, 0, 0, 75, 74, 1, 0, 0, 0, 76, 79, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 77, 
	78, 1, 0, 0, 0, 78, 18, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 80, 81, 3, 17, 
	8, 0, 81, 82, 5, 40, 0, 0, 82, 83, 5, 41, 0, 0, 83, 20, 1, 0, 0, 0, 84, 
	85, 7, 0, 0, 0, 85, 22, 1, 0, 0, 0, 86, 88, 5, 45, 0, 0, 87, 86, 1, 0, 
	0, 0, 87, 88, 1, 0, 0, 0, 88, 90, 1, 0, 0, 0, 89, 91, 2, 48, 57, 0, 90, 
	89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 
	0, 93, 24, 1, 0, 0, 0, 94, 95, 5, 39, 0, 0, 95, 96, 9, 0, 0, 0, 96, 97, 
	5, 39, 0, 0, 97, 26, 1, 0, 0, 0, 98, 99, 5, 34, 0, 0, 99, 100, 3, 29, 14, 
	0, 100, 101, 5, 34, 0, 0, 101, 28, 1, 0, 0, 0, 102, 104, 8, 1, 0, 0, 103, 
	102, 1, 0, 0, 0, 104, 107, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 105, 106, 
	1, 0, 0, 0, 106, 30, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 108, 109, 3, 23, 
	11, 0, 109, 113, 5, 46, 0, 0, 110, 112, 3, 23, 11, 0, 111, 110, 1, 0, 0, 
	0, 112, 115, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 
	123, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 116, 118, 5, 46, 0, 0, 117, 119, 
	3, 23, 11, 0, 118, 117, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 118, 1, 
	0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 123, 1, 0, 0, 0, 122, 108, 1, 0, 0, 
	0, 122, 116, 1, 0, 0, 0, 123, 32, 1, 0, 0, 0, 124, 126, 7, 2, 0, 0, 125, 
	124, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 127, 128, 
	1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 130, 6, 16, 0, 0, 130, 34, 1, 0, 
	0, 0, 131, 135, 5, 59, 0, 0, 132, 134, 9, 0, 0, 0, 133, 132, 1, 0, 0, 0, 
	134, 137, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 136, 
	139, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 138, 131, 1, 0, 0, 0, 138, 139, 
	1, 0, 0, 0, 139, 141, 1, 0, 0, 0, 140, 142, 5, 13, 0, 0, 141, 140, 1, 0, 
	0, 0, 141, 142, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 144, 5, 10, 0, 0, 
	144, 36, 1, 0, 0, 0, 13, 0, 75, 77, 87, 92, 105, 113, 120, 122, 127, 135, 
	138, 141, 1, 6, 0, 0,
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

// AssemblerLexerInit initializes any static state used to implement AssemblerLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewAssemblerLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func AssemblerLexerInit() {
  staticData := &AssemblerLexerLexerStaticData
  staticData.once.Do(assemblerlexerLexerInit)
}

// NewAssemblerLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewAssemblerLexer(input antlr.CharStream) *AssemblerLexer {
  AssemblerLexerInit()
	l := new(AssemblerLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
  staticData := &AssemblerLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Assembler.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// AssemblerLexer tokens.
const (
	AssemblerLexerT__0 = 1
	AssemblerLexerT__1 = 2
	AssemblerLexerT__2 = 3
	AssemblerLexerT__3 = 4
	AssemblerLexerT__4 = 5
	AssemblerLexerT__5 = 6
	AssemblerLexerT__6 = 7
	AssemblerLexerREG = 8
	AssemblerLexerID = 9
	AssemblerLexerFUNC = 10
	AssemblerLexerINT = 11
	AssemblerLexerCHAR = 12
	AssemblerLexerSTRING = 13
	AssemblerLexerFLOAT = 14
	AssemblerLexerWS = 15
	AssemblerLexerNEWLINE = 16
)

