// Code generated from /Users/bytedance/go/src/github.com/ycl2018/pie-go/parser/Pie.g4 by ANTLR 4.13.1. DO NOT EDIT.

package gen // Pie
import "github.com/antlr4-go/antlr/v4"


// A complete Visitor for a parse tree produced by PieParser.
type PieVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by PieParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by PieParser#structDefinition.
	VisitStructDefinition(ctx *StructDefinitionContext) interface{}

	// Visit a parse tree produced by PieParser#functionDefinition.
	VisitFunctionDefinition(ctx *FunctionDefinitionContext) interface{}

	// Visit a parse tree produced by PieParser#slist.
	VisitSlist(ctx *SlistContext) interface{}

	// Visit a parse tree produced by PieParser#structDefinitionStatement.
	VisitStructDefinitionStatement(ctx *StructDefinitionStatementContext) interface{}

	// Visit a parse tree produced by PieParser#assignementStatement.
	VisitAssignementStatement(ctx *AssignementStatementContext) interface{}

	// Visit a parse tree produced by PieParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by PieParser#printStatement.
	VisitPrintStatement(ctx *PrintStatementContext) interface{}

	// Visit a parse tree produced by PieParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by PieParser#whileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by PieParser#callStatement.
	VisitCallStatement(ctx *CallStatementContext) interface{}

	// Visit a parse tree produced by PieParser#nlStatement.
	VisitNlStatement(ctx *NlStatementContext) interface{}

	// Visit a parse tree produced by PieParser#call.
	VisitCall(ctx *CallContext) interface{}

	// Visit a parse tree produced by PieParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by PieParser#addexpr.
	VisitAddexpr(ctx *AddexprContext) interface{}

	// Visit a parse tree produced by PieParser#mulexpr.
	VisitMulexpr(ctx *MulexprContext) interface{}

	// Visit a parse tree produced by PieParser#intAtom.
	VisitIntAtom(ctx *IntAtomContext) interface{}

	// Visit a parse tree produced by PieParser#charAtom.
	VisitCharAtom(ctx *CharAtomContext) interface{}

	// Visit a parse tree produced by PieParser#floatAtom.
	VisitFloatAtom(ctx *FloatAtomContext) interface{}

	// Visit a parse tree produced by PieParser#stringAtom.
	VisitStringAtom(ctx *StringAtomContext) interface{}

	// Visit a parse tree produced by PieParser#qidAtom.
	VisitQidAtom(ctx *QidAtomContext) interface{}

	// Visit a parse tree produced by PieParser#callAtom.
	VisitCallAtom(ctx *CallAtomContext) interface{}

	// Visit a parse tree produced by PieParser#instanceAtom.
	VisitInstanceAtom(ctx *InstanceAtomContext) interface{}

	// Visit a parse tree produced by PieParser#parenthesizedAtom.
	VisitParenthesizedAtom(ctx *ParenthesizedAtomContext) interface{}

	// Visit a parse tree produced by PieParser#instance.
	VisitInstance(ctx *InstanceContext) interface{}

	// Visit a parse tree produced by PieParser#qid.
	VisitQid(ctx *QidContext) interface{}

	// Visit a parse tree produced by PieParser#vardef.
	VisitVardef(ctx *VardefContext) interface{}

	// Visit a parse tree produced by PieParser#multOp.
	VisitMultOp(ctx *MultOpContext) interface{}

	// Visit a parse tree produced by PieParser#addOp.
	VisitAddOp(ctx *AddOpContext) interface{}

	// Visit a parse tree produced by PieParser#compOp.
	VisitCompOp(ctx *CompOpContext) interface{}

}