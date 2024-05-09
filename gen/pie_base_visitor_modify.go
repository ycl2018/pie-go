// fixed: CodeGen
package gen // Pie
import "github.com/antlr4-go/antlr/v4"

type BasePieVisitor struct {
	antlr.ParseTreeVisitor
}

func (v *BasePieVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePieVisitor) VisitStructDefinition(ctx *StructDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePieVisitor) VisitFunctionDefinition(ctx *FunctionDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePieVisitor) VisitSlist(ctx *SlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePieVisitor) VisitStructDefinitionStatement(ctx *StructDefinitionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePieVisitor) VisitAssignementStatement(ctx *AssignementStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePieVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePieVisitor) VisitPrintStatement(ctx *PrintStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePieVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePieVisitor) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePieVisitor) VisitCallStatement(ctx *CallStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePieVisitor) VisitNlStatement(ctx *NlStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePieVisitor) VisitCall(ctx *CallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePieVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePieVisitor) VisitAddexpr(ctx *AddexprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePieVisitor) VisitMulexpr(ctx *MulexprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePieVisitor) VisitIntAtom(ctx *IntAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePieVisitor) VisitCharAtom(ctx *CharAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePieVisitor) VisitFloatAtom(ctx *FloatAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePieVisitor) VisitStringAtom(ctx *StringAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePieVisitor) VisitQidAtom(ctx *QidAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePieVisitor) VisitCallAtom(ctx *CallAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePieVisitor) VisitInstanceAtom(ctx *InstanceAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePieVisitor) VisitParenthesizedAtom(ctx *ParenthesizedAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePieVisitor) VisitInstance(ctx *InstanceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePieVisitor) VisitQid(ctx *QidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePieVisitor) VisitVardef(ctx *VardefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePieVisitor) VisitMultOp(ctx *MultOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePieVisitor) VisitAddOp(ctx *AddOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePieVisitor) VisitCompOp(ctx *CompOpContext) interface{} {
	return v.VisitChildren(ctx)
}
