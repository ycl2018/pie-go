// Code generated from /Users/bytedance/go/src/github.com/ycl2018/asm/gen/Assembler.g4 by ANTLR 4.13.1. DO NOT EDIT.

package gen // Assembler
import "github.com/antlr4-go/antlr/v4"


// A complete Visitor for a parse tree produced by AssemblerParser.
type AssemblerVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by AssemblerParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by AssemblerParser#globals.
	VisitGlobals(ctx *GlobalsContext) interface{}

	// Visit a parse tree produced by AssemblerParser#functionDeclaration.
	VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{}

	// Visit a parse tree produced by AssemblerParser#instr.
	VisitInstr(ctx *InstrContext) interface{}

	// Visit a parse tree produced by AssemblerParser#operand.
	VisitOperand(ctx *OperandContext) interface{}

	// Visit a parse tree produced by AssemblerParser#label.
	VisitLabel(ctx *LabelContext) interface{}

}