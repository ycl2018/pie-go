package compile

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	gen2 "github.com/ycl2018/pie-go/gen"
	"github.com/ycl2018/pie-go/interpreter/compile/ir"
	"github.com/ycl2018/pie-go/interpreter/stack2"
)

type StackCompileVisitor struct {
	gen2.BasePieVisitor
	InterpreterListener InterpreterListener
	Scopes              map[antlr.ParserRuleContext]Scope // 作用域树&符号表存储
	GlobalScope         *GlobalScope
	AllFuncCodes        map[string][]*ir.StackInstr
	MainFuncCodes       []*ir.StackInstr
	CurFuncCodes        []*ir.StackInstr
	ToBeFilledBrfAddrs  map[*ir.StackInstr][]*ir.StackInstr // 目标分支跳转指令 -> 需要回填的分支跳转指令列表
	Code []byte
}

func (s *StackCompileVisitor) VisitAssignementStatement(ctx *gen2.AssignementStatementContext) interface{} {
	ctx.Expr().Accept(s)
	scope := s.Scopes[ctx]
	qidCtx := ctx.Qid()
	var vName = qidCtx.ID(0).GetText()
	varSymbol := scope.Resolve(vName).(*VariableSymbol)
	// 普通变量赋值
	varAddr := varSymbol.Address
	// 生成赋值语句IR
	if len(qidCtx.AllID()) == 1 {
		s.CurFuncCodes = append(s.CurFuncCodes, &ir.StackInstr{
			OpCode:   stack2.InstrStore,
			Operands: []int32{varAddr},
		})
		return nil
	} else {
		// 结构体成员赋值
		// 生成加载结构体变量地址IR
		ss := scope.Resolve(qidCtx.ID(0).GetText())
		opCode := ss.GetAddress()
		s.CurFuncCodes = append(s.CurFuncCodes, &ir.StackInstr{
			OpCode:   stack2.InstrLoad,
			Operands: []int32{opCode},
		})
		for i := 1; i < len(qidCtx.AllID())-1; i++ {
			// 生成加载结构体成员地址IR
			memberName := qidCtx.ID(i).GetText()
			constSymbol := &ConstSymbol{
				Name: fmt.Sprintf("%s_%s", ir.ConstString, memberName),
				Line: int32(qidCtx.GetStart().GetLine()),
				Kind: ir.ConstString,
			}
			cSymbol, _ := s.GlobalScope.DefineOrGetConst(constSymbol)
			opCode := cSymbol.GetAddress()
			s.CurFuncCodes = append(s.CurFuncCodes, &ir.StackInstr{
				OpCode:   stack2.InstrFLoad,
				Operands: []int32{opCode}, // 成员名称常量地址
			})
		}
		// 生成赋值语句IR
		memberName := qidCtx.AllID()[len(qidCtx.AllID())-1].GetText()
		constSymbol := &ConstSymbol{
			Name: fmt.Sprintf("%s_%s", ir.ConstString, memberName),
			Line: int32(qidCtx.GetStart().GetLine()),
			Kind: ir.ConstString,
		}
		cSymbol, _ := s.GlobalScope.DefineOrGetConst(constSymbol)
		opCode = cSymbol.GetAddress()
		s.CurFuncCodes = append(s.CurFuncCodes, &ir.StackInstr{
			OpCode:   stack2.InstrFStore,
			Operands: []int32{opCode}, // 成员名称常量地址
		})
	}
	return nil
}

func (s *StackCompileVisitor) VisitReturnStatement(ctx *gen2.ReturnStatementContext) interface{} {
	ctx.Expr().Accept(s)
	s.CurFuncCodes = append(s.CurFuncCodes, &ir.StackInstr{
		OpCode:   stack2.InstrReturn,
		Operands: []int32{},
	})
	return nil
}

func (s *StackCompileVisitor) VisitPrintStatement(ctx *gen2.PrintStatementContext) interface{} {
	ctx.Expr().Accept(s)
	s.CurFuncCodes = append(s.CurFuncCodes, &ir.StackInstr{
		OpCode:   stack2.InstrPrint,
		Operands: []int32{},
	})
	return nil
}

func (s *StackCompileVisitor) VisitIfStatement(ctx *gen2.IfStatementContext) interface{} {
	ctx.Expr().Accept(s)
	// 生成分支跳转IR
	brfInstr := &ir.StackInstr{
		OpCode:   stack2.InstrBRF,
		Operands: []int32{-1}, // 占位，等编译为二进制时再回填
	}
	s.CurFuncCodes = append(s.CurFuncCodes, brfInstr)
	// 生成分支语句IR
	ctx.GetC().Accept(s)
	brInstr := &ir.StackInstr{
		OpCode:   stack2.InstrBR,
		Operands: []int32{-1}, // 占位，等编译为二进制时再回填
	}
	s.CurFuncCodes = append(s.CurFuncCodes, brInstr)
	// 回填分支跳转地址
	brfInstr.Operands[0] = int32(len(s.Code))
	if el := ctx.GetEl(); el != nil {
		el.Accept(s)
	}
	brInstr.Operands[0] = int32(len(s.Code))
	return nil
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
	s.CurFuncCodes = funcSymbol.(*FunctionSymbol).Code
	// 生成参数列表IR
	paramLen := len(ctx.AllVardef())
	for i := paramLen - 1; i >= 0; i-- {
		s.CurFuncCodes = append(s.CurFuncCodes, &ir.StackInstr{
			OpCode:   stack2.InstrLoad,
			Operands: []int32{int32(i)},
		})
	}
	// 生成函数体IR
	ctx.Slist().Accept(s)
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
