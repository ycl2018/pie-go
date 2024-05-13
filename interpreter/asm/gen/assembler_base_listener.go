// Code generated from /Users/bytedance/go/src/github.com/ycl2018/asm/gen/Assembler.g4 by ANTLR 4.13.1. DO NOT EDIT.

package gen // Assembler
import "github.com/antlr4-go/antlr/v4"

// BaseAssemblerListener is a complete listener for a parse tree produced by AssemblerParser.
type BaseAssemblerListener struct{}

var _ AssemblerListener = &BaseAssemblerListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAssemblerListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAssemblerListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAssemblerListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAssemblerListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseAssemblerListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseAssemblerListener) ExitProgram(ctx *ProgramContext) {}

// EnterGlobals is called when production globals is entered.
func (s *BaseAssemblerListener) EnterGlobals(ctx *GlobalsContext) {}

// ExitGlobals is called when production globals is exited.
func (s *BaseAssemblerListener) ExitGlobals(ctx *GlobalsContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BaseAssemblerListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BaseAssemblerListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterInstr is called when production instr is entered.
func (s *BaseAssemblerListener) EnterInstr(ctx *InstrContext) {}

// ExitInstr is called when production instr is exited.
func (s *BaseAssemblerListener) ExitInstr(ctx *InstrContext) {}

// EnterOperand is called when production operand is entered.
func (s *BaseAssemblerListener) EnterOperand(ctx *OperandContext) {}

// ExitOperand is called when production operand is exited.
func (s *BaseAssemblerListener) ExitOperand(ctx *OperandContext) {}

// EnterLabel is called when production label is entered.
func (s *BaseAssemblerListener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BaseAssemblerListener) ExitLabel(ctx *LabelContext) {}
