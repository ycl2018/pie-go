// Code generated from /Users/bytedance/go/src/github.com/ycl2018/pie-go/interpreter/asm2/Assembler.g4 by ANTLR 4.13.2. DO NOT EDIT.

package gen // Assembler
import "github.com/antlr4-go/antlr/v4"


// A complete Visitor for a parse tree produced by AssemblerParser.
type AssemblerVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by AssemblerParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by AssemblerParser#globals.
	VisitGlobals(ctx *GlobalsContext) interface{}

	// Visit a parse tree produced by AssemblerParser#constantPool.
	VisitConstantPool(ctx *ConstantPoolContext) interface{}

	// Visit a parse tree produced by AssemblerParser#constantPoolEntry.
	VisitConstantPoolEntry(ctx *ConstantPoolEntryContext) interface{}

	// Visit a parse tree produced by AssemblerParser#structDecl.
	VisitStructDecl(ctx *StructDeclContext) interface{}

	// Visit a parse tree produced by AssemblerParser#functionDeclaration.
	VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{}

	// Visit a parse tree produced by AssemblerParser#instr.
	VisitInstr(ctx *InstrContext) interface{}

	// Visit a parse tree produced by AssemblerParser#operand.
	VisitOperand(ctx *OperandContext) interface{}

	// Visit a parse tree produced by AssemblerParser#label.
	VisitLabel(ctx *LabelContext) interface{}

}