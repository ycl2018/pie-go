package tree

import (
	"github.com/antlr4-go/antlr/v4"
	gen2 "github.com/ycl2018/pie-go/gen"
)

var _ gen2.PieVisitor = (*PieDefineVisitor)(nil)

type PieDefineVisitor struct {
	*gen2.BasePieVisitor
	CurScope Scope
	Scopes   map[antlr.ParserRuleContext]Scope
}

func NewPieDefineVisitor() *PieDefineVisitor {
	ret := &PieDefineVisitor{
		CurScope: &GlobalScope{Symbols: make(map[string]Symbol)},
		Scopes:   make(map[antlr.ParserRuleContext]Scope),
	}
	ret.BasePieVisitor = &gen2.BasePieVisitor{ParseTreeVisitor: &BaseVisitor{realVisitor: ret}}
	return ret
}

func (p *PieDefineVisitor) SaveScope(ctx antlr.ParserRuleContext, scope Scope) {
	p.Scopes[ctx] = scope
}

func (p *PieDefineVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(p)
}

func (p *PieDefineVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	ctx := node.(antlr.ParserRuleContext)
	for _, child := range ctx.GetChildren() {
		child.(antlr.ParseTree).Accept(p)
	}
	return nil
}

func (p *PieDefineVisitor) VisitProgram(ctx *gen2.ProgramContext) interface{} {
	// ( functionDefinition | statement )+ EOF
	for _, f := range ctx.AllFunctionDefinition() {
		p.Visit(f)
	}
	for _, s := range ctx.AllStatement() {
		p.Visit(s)
	}
	return nil
}

func (p *PieDefineVisitor) VisitStructDefinition(ctx *gen2.StructDefinitionContext) interface{} {
	// 'struct' name=ID '{' vardef (',' vardef)* '}' NL
	structName := ctx.ID().GetText()
	structSymbol := NewStructSymbol(structName, p.CurScope)
	p.CurScope.Define(structSymbol)
	p.CurScope = structSymbol
	for _, vardef := range ctx.AllVardef() {
		p.Visit(vardef)
	}
	p.CurScope = structSymbol.EnclosingScope
	return nil
}

func (p *PieDefineVisitor) VisitFunctionDefinition(ctx *gen2.FunctionDefinitionContext) interface{} {
	// 'def' ID '(' (vardef (',' vardef)* )? ')' slist
	funcName := ctx.ID().GetText()
	funcSymbol := &FunctionSymbol{
		Name:           funcName,
		BlockAst:       ctx.Slist(),
		EnclosingScope: p.CurScope,
	}
	p.CurScope.Define(funcSymbol)
	p.CurScope = funcSymbol
	// vardefs
	for _, varDef := range ctx.AllVardef() {
		p.Visit(varDef)
	}
	// slist
	p.CurScope = &LocalScope{
		Symbols:        make(map[string]Symbol),
		EnclosingScope: p.CurScope,
	}
	p.Visit(ctx.Slist())
	// 这里直接回退到funcSymbol的scope
	p.CurScope = funcSymbol.EnclosingScope
	return nil
}

func (p *PieDefineVisitor) VisitSlist(ctx *gen2.SlistContext) interface{} {
	for _, stmtCtx := range ctx.AllStatement() {
		p.Visit(stmtCtx)
	}
	return nil
}

func (p *PieDefineVisitor) VisitAssignmentStatement(ctx *gen2.AssignmentStatementContext) interface{} {
	// qid '=' expr NL
	p.SaveScope(ctx, p.CurScope)
	qidCtx := ctx.Qid()
	vs := &VariableSymbol{
		// We can have : u.name.y => only take the left part "u"
		// u = new User
		// u.name = "parrt"    # make u.name a string
		// u.name.y = "parrt"  # u.name is a string not a struct
		Name: qidCtx.ID(0).GetText(),
	}
	p.CurScope.Define(vs)
	p.Visit(ctx.Expr())
	return nil
}

func (p *PieDefineVisitor) VisitCall(ctx *gen2.CallContext) interface{} {
	// name=ID '(' (expr (',' expr )*)? ')'
	p.SaveScope(ctx, p.CurScope)
	return nil
}

func (p *PieDefineVisitor) VisitInstance(ctx *gen2.InstanceContext) interface{} {
	// 'new' sname=ID
	p.SaveScope(ctx, p.CurScope)
	return nil
}

func (p *PieDefineVisitor) VisitVardef(ctx *gen2.VardefContext) interface{} {
	// ID
	p.SaveScope(ctx, p.CurScope)
	variableSymbol := &VariableSymbol{
		Name: ctx.ID().GetText(),
	}
	p.CurScope.Define(variableSymbol)
	return nil
}
