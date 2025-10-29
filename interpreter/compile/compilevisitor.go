package compile

import (
	"github.com/antlr4-go/antlr/v4"
	gen2 "github.com/ycl2018/pie-go/gen"
	"github.com/ycl2018/pie-go/interpreter/compile/ir"
)

type StackCompileVisitor struct {
	gen2.BasePieVisitor
	InterpreterListener InterpreterListener
	Scopes              map[antlr.ParserRuleContext]Scope // 作用域树&符号表存储
	GlobalScope         *GlobalScope
	AllFuncCodes        map[string][]ir.StackInstr
	MainFuncCodes       []ir.StackInstr
	CurFuncCodes        []ir.StackInstr
}

func (s *StackCompileVisitor) VisitStructDefinition(ctx *gen2.StructDefinitionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *StackCompileVisitor) VisitSlist(ctx *gen2.SlistContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *StackCompileVisitor) VisitStructDefinitionStatement(ctx *gen2.StructDefinitionStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *StackCompileVisitor) VisitAssignementStatement(ctx *gen2.AssignementStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *StackCompileVisitor) VisitReturnStatement(ctx *gen2.ReturnStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *StackCompileVisitor) VisitPrintStatement(ctx *gen2.PrintStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *StackCompileVisitor) VisitIfStatement(ctx *gen2.IfStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *StackCompileVisitor) VisitWhileStatement(ctx *gen2.WhileStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *StackCompileVisitor) VisitCallStatement(ctx *gen2.CallStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *StackCompileVisitor) VisitNlStatement(ctx *gen2.NlStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *StackCompileVisitor) VisitCall(ctx *gen2.CallContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *StackCompileVisitor) VisitExpr(ctx *gen2.ExprContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *StackCompileVisitor) VisitAddexpr(ctx *gen2.AddexprContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *StackCompileVisitor) VisitMulexpr(ctx *gen2.MulexprContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *StackCompileVisitor) VisitIntAtom(ctx *gen2.IntAtomContext) interface{} {

}

func (s *StackCompileVisitor) VisitCharAtom(ctx *gen2.CharAtomContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *StackCompileVisitor) VisitFloatAtom(ctx *gen2.FloatAtomContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *StackCompileVisitor) VisitStringAtom(ctx *gen2.StringAtomContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *StackCompileVisitor) VisitQidAtom(ctx *gen2.QidAtomContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *StackCompileVisitor) VisitCallAtom(ctx *gen2.CallAtomContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *StackCompileVisitor) VisitInstanceAtom(ctx *gen2.InstanceAtomContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *StackCompileVisitor) VisitParenthesizedAtom(ctx *gen2.ParenthesizedAtomContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *StackCompileVisitor) VisitInstance(ctx *gen2.InstanceContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *StackCompileVisitor) VisitQid(ctx *gen2.QidContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *StackCompileVisitor) VisitVardef(ctx *gen2.VardefContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *StackCompileVisitor) VisitMultOp(ctx *gen2.MultOpContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *StackCompileVisitor) VisitAddOp(ctx *gen2.AddOpContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *StackCompileVisitor) VisitCompOp(ctx *gen2.CompOpContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *StackCompileVisitor) VisitFunctionDefinition(ctx *gen2.FunctionDefinitionContext) interface{} {
	// 'def' ID '(' (vardef (',' vardef)* )? ')' slist
	funcName := ctx.ID().GetText()
	scope := s.Scopes[ctx]
	funcSymbol := scope.Resolve(funcName)
	// 生成函数定义 IR

	return nil
}

func (s *StackCompileVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(s)
}

func (s *StackCompileVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	ctx := node.(antlr.ParserRuleContext)
	for _, child := range ctx.GetChildren() {
		child.(antlr.ParseTree).Accept(s)
	}
	return nil
}

func (s *StackCompileVisitor) VisitAtom(ctx *gen2.AtomContext) interface{} {
	return nil
}
