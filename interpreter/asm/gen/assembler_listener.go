// Code generated from /Users/bytedance/go/src/github.com/ycl2018/asm/gen/Assembler.g4 by ANTLR 4.13.1. DO NOT EDIT.

package gen // Assembler
import "github.com/antlr4-go/antlr/v4"


// AssemblerListener is a complete listener for a parse tree produced by AssemblerParser.
type AssemblerListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterGlobals is called when entering the globals production.
	EnterGlobals(c *GlobalsContext)

	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// EnterInstr is called when entering the instr production.
	EnterInstr(c *InstrContext)

	// EnterOperand is called when entering the operand production.
	EnterOperand(c *OperandContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitGlobals is called when exiting the globals production.
	ExitGlobals(c *GlobalsContext)

	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitInstr is called when exiting the instr production.
	ExitInstr(c *InstrContext)

	// ExitOperand is called when exiting the operand production.
	ExitOperand(c *OperandContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)
}
