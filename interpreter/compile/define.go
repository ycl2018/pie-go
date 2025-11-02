package compile

import (
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	gen2 "github.com/ycl2018/pie-go/gen"
	"github.com/ycl2018/pie-go/interpreter/compile/ir"
)

var _ gen2.PieVisitor = (*PieDefineVisitor)(nil)

type PieDefineVisitor struct {
	*gen2.BasePieVisitor
	Log            InterpreterListener
	CurScope       Scope
	Scopes         map[antlr.ParserRuleContext]Scope
	GlobalScope    *GlobalScope
	ScopeAllocator int32
}

func (p *PieDefineVisitor) VisitFloatAtom(ctx *gen2.FloatAtomContext) interface{} {
	value, _ := strconv.ParseFloat(ctx.FLOAT().GetText(), 32)
	p.GlobalScope.Define(&ConstSymbol{
		Name:  fmt.Sprintf("%s_%s", ir.ConstFloat32, ctx.FLOAT().GetText()),
		Line:  int32(ctx.GetStart().GetLine()),
		Kind:  ir.ConstFloat32,
		Value: float32(value),
	})
	return nil
}

func (p *PieDefineVisitor) VisitStringAtom(ctx *gen2.StringAtomContext) interface{} {
	value := ctx.STRING().GetText()[1 : len(ctx.STRING().GetText())-1]
	p.GlobalScope.Define(&ConstSymbol{
		Name:  fmt.Sprintf("%s_%s", ir.ConstString, value),
		Line:  int32(ctx.GetStart().GetLine()),
		Kind:  ir.ConstString,
		Value: value,
	})
	return nil
}

func (p *PieDefineVisitor) VisitQid(ctx *gen2.QidContext) interface{} {
	p.SaveScope(ctx, p.CurScope)
	firstID := ctx.ID(0)
	// 检查符号是否定义
	if p.CurScope.Resolve(firstID.GetText()) == nil {
		p.Log.ErrorToken(firstID.GetSymbol(), "undefined symbol: %s", firstID.GetText())
		return nil
	}
	return p.VisitChildren(ctx)
}

func NewPieDefineVisitor(log InterpreterListener) *PieDefineVisitor {
	gScope := NewGlobalScope()
	ret := &PieDefineVisitor{
		CurScope:    gScope,
		GlobalScope: gScope,
		Scopes:      make(map[antlr.ParserRuleContext]Scope),
		Log:         log,
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

func (p *PieDefineVisitor) VisitStructDefinition(ctx *gen2.StructDefinitionContext) interface{} {
	// 'struct' name=ID '{' vardef (',' vardef)* '}' NL
	structName := ctx.ID().GetText()
	structSymbol := NewStructSymbol(structName, p.CurScope)
	p.CurScope = structSymbol
	for _, vardef := range ctx.AllVardef() {
		p.Visit(vardef)
	}
	p.CurScope = structSymbol.EnclosingScope
	p.CurScope.Define(structSymbol)
	return nil
}

func (p *PieDefineVisitor) VisitFunctionDefinition(ctx *gen2.FunctionDefinitionContext) interface{} {
	// 'def' ID '(' (vardef (',' vardef)* )? ')' slist
	funcName := ctx.ID().GetText()
	funcSymbol := &FunctionSymbol{
		Name:           funcName,
		EnclosingScope: p.CurScope,
	}

	p.SaveScope(ctx, funcSymbol)
	p.CurScope = funcSymbol
	// vardefs
	for _, varDef := range ctx.AllVardef() {
		p.Visit(varDef)
	}
	// slist
	p.CurScope = &LocalScope{
		Symbols:        make(map[string]Symbol),
		EnclosingScope: p.CurScope,
		BaseAllocAddr:  int32(len(funcSymbol.FormalArgs)),
		ID:             p.ScopeAllocator,
	}
	p.ScopeAllocator++
	// 保存函数的local scope
	funcSymbol.BodyScope = p.CurScope.(*LocalScope)
	p.Visit(ctx.Slist())
	// 这里直接回退到funcSymbol的scope
	p.CurScope = funcSymbol.EnclosingScope
	p.CurScope.Define(funcSymbol)
	return nil
}

func (p *PieDefineVisitor) VisitAssignmentStatement(ctx *gen2.AssignmentStatementContext) interface{} {
	// qid '=' expr NL
	p.SaveScope(ctx, p.CurScope)
	qidCtx := ctx.Qid()
	p.SaveScope(qidCtx.(*gen2.QidContext), p.CurScope)
	vName := qidCtx.ID(0).GetText()
	if len(ctx.Qid().AllID()) == 1 && p.CurScope.Resolve(vName) == nil {
		vs := &VariableSymbol{
			// We can have : u.name.y => only take the left part "u"
			// u = new User
			// u.name = "parrt"    # make u.name a string
			// u.name.y = "parrt"  # u.name is a string not a struct
			Name: vName,
			Line: int32(ctx.GetStart().GetLine()),
		}
		p.CurScope.Define(vs)
	}
	if len(ctx.Qid().AllID()) > 1 {
		if p.CurScope.Resolve(vName) == nil {
			p.Log.ErrorToken(ctx.GetStart(), "undefined symbol: %s", vName)
			return nil
		}
	}
	p.Visit(ctx.Expr())
	return nil
}

func (p *PieDefineVisitor) VisitCall(ctx *gen2.CallContext) interface{} {
	// name=ID '(' (expr (',' expr )*)? ')'
	p.SaveScope(ctx, p.CurScope)
	for _, context := range ctx.AllExpr() {
		context.Accept(p)
	}
	return nil
}

func (p *PieDefineVisitor) VisitInstance(ctx *gen2.InstanceContext) interface{} {
	// 'new' sname=ID NL
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
