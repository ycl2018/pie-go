package compile

import (
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	gen2 "github.com/ycl2018/pie-go/gen"
	"github.com/ycl2018/pie-go/interpreter/compile/ir"
	"github.com/ycl2018/pie-go/interpreter/vm"
)

var _ gen2.PieVisitor = (*StackCompileVisitor)(nil)

type StackCompileVisitor struct {
	gen2.BasePieVisitor
	Log         InterpreterListener
	Scopes      map[antlr.ParserRuleContext]Scope // 作用域树&符号表存储
	GlobalScope *GlobalScope
	AllFuncs    []*FunctionSymbol // 所有函数,不包含主函数
	MainFunc    *FunctionSymbol
	CurFunc     *FunctionSymbol
	ToBeFilled  map[*ir.StackInstr][]*ir.StackInstr // 目标分支跳转Tag -> 需要回填的分支跳转指令列表
	TagAlloc    int32
}

func NewStackCompileVisitor(scopes map[antlr.ParserRuleContext]Scope, globalScope *GlobalScope, log InterpreterListener) *StackCompileVisitor {
	mainFunc := &FunctionSymbol{
		Name: "main",
	}
	s := &StackCompileVisitor{
		Log:         log,
		Scopes:      scopes,
		GlobalScope: globalScope,
		ToBeFilled:  map[*ir.StackInstr][]*ir.StackInstr{},
		MainFunc:    mainFunc,
		CurFunc:     mainFunc,
	}
	s.BasePieVisitor = gen2.BasePieVisitor{ParseTreeVisitor: &BaseVisitor{realVisitor: s}}
	return s
}

func (s *StackCompileVisitor) WriteInstr(instr *ir.StackInstr) {
	s.CurFunc.Code = append(s.CurFunc.Code, instr)
}

func (s *StackCompileVisitor) AllocTag() int {
	s.TagAlloc--
	return int(s.TagAlloc)
}

func (s *StackCompileVisitor) fillTargetTag(instr *ir.StackInstr, targetTag *ir.StackInstr) {
	s.ToBeFilled[targetTag] = append(s.ToBeFilled[targetTag], instr)
	instr.Operands[0] = int32(targetTag.OpCode)
}

func (s *StackCompileVisitor) VisitProgram(ctx *gen2.ProgramContext) interface{} {
	s.VisitChildren(ctx)
	// 补充主函数Halt指令IR
	if len(s.MainFunc.Code) == 0 || s.MainFunc.Code[len(s.MainFunc.Code)-1].OpCode != vm.InstrHalt {
		s.MainFunc.Code = append(s.MainFunc.Code, &ir.StackInstr{
			OpCode:   vm.InstrHalt,
			Operands: []int32{},
		})
	}
	return nil
}

func (s *StackCompileVisitor) VisitFunctionDefinition(ctx *gen2.FunctionDefinitionContext) interface{} {
	// 'def' ID '(' (vardef (',' vardef)* )? ')' slist
	// 生成函数定义 IR
	funcName := ctx.ID().GetText()
	scope := s.Scopes[ctx]
	funcSymbol := scope.Resolve(funcName).(*FunctionSymbol)
	s.CurFunc = funcSymbol
	// 参数列表入栈由解释器负责自动执行
	// 生成函数体IR
	ctx.Slist().Accept(s)
	// 补充返回指令IR
	if len(s.CurFunc.Code) == 0 || s.CurFunc.Code[len(s.CurFunc.Code)-1].OpCode != vm.InstrReturn {
		s.WriteInstr(&ir.StackInstr{
			OpCode:   vm.InstrReturn,
			Operands: []int32{},
		})
	}
	s.AllFuncs = append(s.AllFuncs, funcSymbol)
	s.CurFunc = s.MainFunc
	return nil
}

func (s *StackCompileVisitor) VisitAssignmentStatement(ctx *gen2.AssignmentStatementContext) interface{} {
	ctx.Expr().Accept(s)
	qid := ctx.Qid()
	qidSymbol := s.Scopes[qid.(*gen2.QidContext)].Resolve(qid.ID(0).GetText())
	symbolScope := qidSymbol.Scope()
	isGlobal := symbolScope.GetName() == GlobalScopeName
	loadInstr := vm.InstrLoad
	storeInstr := vm.InstrStore
	if isGlobal {
		loadInstr = vm.InstrGLoad
		storeInstr = vm.InstrGStore
	}
	if len(qid.AllID()) == 1 {
		s.WriteInstr(&ir.StackInstr{
			OpCode:   storeInstr,
			Operands: []int32{qidSymbol.GetAddress()},
		})
	} else {
		// 加载全局 Struct
		s.WriteInstr(&ir.StackInstr{
			OpCode:   loadInstr,
			Operands: []int32{qidSymbol.GetAddress()},
		})
		for i := 1; i < len(qid.AllID())-1; i++ {
			// 加载结构体成员
			memberName := qid.ID(i).GetText()
			cSymbol := getStringConst(memberName, s.GlobalScope)
			opCode := cSymbol.GetAddress() // 成员名称常量地址
			s.WriteInstr(&ir.StackInstr{
				OpCode:   vm.InstrFLoad,
				Operands: []int32{opCode},
			})
		}
		lastFieldName := qid.AllID()[len(qid.AllID())-1].GetText()
		cSymbol := getStringConst(lastFieldName, s.GlobalScope)
		opCode := cSymbol.GetAddress() // 成员名称常量地址
		// 生成赋值语句IR
		s.WriteInstr(&ir.StackInstr{
			OpCode:   vm.InstrFStore,
			Operands: []int32{opCode}, // 成员名称常量地址
		})
	}
	return nil
}

func getStringConst(memberName string, scope *GlobalScope) Symbol {
	constSymbol := &ConstSymbol{
		Name:  fmt.Sprintf("%s_%s", ir.ConstString, memberName),
		Kind:  ir.ConstString,
		Value: memberName,
	}
	cSymbol, _ := scope.DefineOrGetConst(constSymbol)
	return cSymbol
}

func (s *StackCompileVisitor) VisitReturnStatement(ctx *gen2.ReturnStatementContext) interface{} {
	if ctx.Expr() != nil {
		ctx.Expr().Accept(s)
	}
	s.WriteInstr(&ir.StackInstr{
		OpCode:   vm.InstrReturn,
		Operands: []int32{},
	})
	return nil
}

func (s *StackCompileVisitor) VisitPrintStatement(ctx *gen2.PrintStatementContext) interface{} {
	ctx.Expr().Accept(s)
	s.WriteInstr(&ir.StackInstr{
		OpCode:   vm.InstrPrint,
		Operands: []int32{},
	})
	return nil
}

func (s *StackCompileVisitor) VisitIfStatement(ctx *gen2.IfStatementContext) interface{} {
	ctx.Expr().Accept(s)
	// 生成分支跳转IR
	brfInstr := &ir.StackInstr{
		OpCode:   vm.InstrBRF,
		Operands: []int32{-1}, // 占位，等编译为二进制时再回填
	}
	s.WriteInstr(brfInstr)
	// 生成分支语句IR
	ctx.GetC().Accept(s)
	var joinTag *ir.StackInstr
	if ctx.GetEl() != nil {
		var brInstr *ir.StackInstr
		brInstr = &ir.StackInstr{
			OpCode:   vm.InstrBR,
			Operands: []int32{-1}, // 占位，等编译为二进制时再回填
		}
		s.WriteInstr(brInstr)
		// 回填分支跳转地址
		elTag := &ir.StackInstr{
			OpCode:   s.AllocTag(),
			Operands: []int32{},
		}
		s.fillTargetTag(brfInstr, elTag)
		s.WriteInstr(elTag)
		if el := ctx.GetEl(); el != nil {
			el.Accept(s)
		}
		joinTag = &ir.StackInstr{
			OpCode:   s.AllocTag(),
			Operands: []int32{},
		}
		s.fillTargetTag(brInstr, joinTag)
		s.WriteInstr(joinTag)
	} else {
		joinTag = &ir.StackInstr{
			OpCode:   s.AllocTag(),
			Operands: []int32{},
		}
		s.fillTargetTag(brfInstr, joinTag)
		s.WriteInstr(joinTag)
	}

	return nil
}

func (s *StackCompileVisitor) VisitWhileStatement(ctx *gen2.WhileStatementContext) interface{} {
	condTag := &ir.StackInstr{
		OpCode:   s.AllocTag(),
		Operands: []int32{},
	}
	// 生成条件跳转Tag
	s.WriteInstr(condTag)
	ctx.Expr().Accept(s)
	// 生成分支跳转IR
	brfInstr := &ir.StackInstr{
		OpCode:   vm.InstrBRF,
		Operands: []int32{-1}, // 占位，等编译为二进制时再回填
	}
	s.WriteInstr(brfInstr)
	// 生成分支语句IR
	ctx.Slist().Accept(s)
	brInstr := &ir.StackInstr{
		OpCode:   vm.InstrBR,
		Operands: []int32{-1}, // 占位，等编译为二进制时再回填
	}
	s.WriteInstr(brInstr)
	s.fillTargetTag(brInstr, condTag)
	// 回填分支跳转地址
	endTag := &ir.StackInstr{
		OpCode:   s.AllocTag(),
		Operands: []int32{},
	}
	s.fillTargetTag(brfInstr, endTag)
	s.WriteInstr(endTag)
	return nil
}

func (s *StackCompileVisitor) VisitCall(ctx *gen2.CallContext) interface{} {
	for _, context := range ctx.AllExpr() {
		context.Accept(s)
	}
	funcName := ctx.GetName()
	if s.Scopes[ctx].Resolve(funcName.GetText()) == nil {
		s.Log.ErrorToken(ctx.GetStart(), fmt.Sprintf("undefined func: %s", funcName.GetText()))
		return nil
	}
	fSymbol := s.Scopes[ctx].Resolve(funcName.GetText()).(*FunctionSymbol)
	// 生成调用指令IR
	callInstr := &ir.StackInstr{
		OpCode:   vm.InstrCall,
		Operands: []int32{fSymbol.Address}, // 函数ID填充，编译为二进制时会替换为函数地址
	}
	s.WriteInstr(callInstr)
	return nil
}

func (s *StackCompileVisitor) VisitExpr(ctx *gen2.ExprContext) interface{} {
	// addexpr (compOp addexpr)?
	ctx.Addexpr(0).Accept(s)
	if ctx.CompOp() != nil {
		ctx.Addexpr(1).Accept(s)
		ctx.CompOp().Accept(s)
	}
	return nil
}

func (s *StackCompileVisitor) VisitAddexpr(ctx *gen2.AddexprContext) interface{} {
	// mulexpr (addOp mulexpr)*
	ctx.Mulexpr(0).Accept(s)
	for i := 0; i < len(ctx.AllAddOp()); i++ {
		ctx.Mulexpr(i + 1).Accept(s)
		ctx.AddOp(i).Accept(s)
	}
	return nil
}

func (s *StackCompileVisitor) VisitMulexpr(ctx *gen2.MulexprContext) interface{} {
	// atom (multOp atom)*
	ctx.Atom(0).Accept(s)
	for i := 0; i < len(ctx.AllMultOp()); i++ {
		ctx.Atom(i + 1).Accept(s)
		ctx.MultOp(i).Accept(s)
	}
	return nil
}

func (s *StackCompileVisitor) VisitIntAtom(ctx *gen2.IntAtomContext) interface{} {
	valStr := ctx.INT().GetText()
	ival, _ := strconv.ParseInt(valStr, 10, 32)
	s.WriteInstr(&ir.StackInstr{
		OpCode:   vm.InstrIConst,
		Operands: []int32{int32(ival)}, // 整数常量地址
	})
	return nil
}

func (s *StackCompileVisitor) VisitCharAtom(ctx *gen2.CharAtomContext) interface{} {
	valStr := ctx.CHAR().GetText()
	ival := rune(valStr[0])
	s.WriteInstr(&ir.StackInstr{
		OpCode:   vm.InstrCConst,
		Operands: []int32{ival}, // 字符常量地址
	})
	return nil
}

func (s *StackCompileVisitor) VisitFloatAtom(ctx *gen2.FloatAtomContext) interface{} {
	valStr := ctx.FLOAT().GetText()
	fval, _ := strconv.ParseFloat(valStr, 32)
	// 浮点数常量地址
	symbol := getFloatConst(float32(fval), s.GlobalScope)
	s.WriteInstr(&ir.StackInstr{
		OpCode:   vm.InstrFConst,
		Operands: []int32{symbol.GetAddress()}, // 浮点数常量地址
	})
	return nil
}

func getFloatConst(fval float32, scope *GlobalScope) Symbol {
	constSymbol := &ConstSymbol{
		Name:  fmt.Sprintf("%s_%f", ir.ConstFloat32, fval),
		Value: fval,
		Kind:  ir.ConstFloat32,
	}
	symbol, _ := scope.DefineOrGetConst(constSymbol)
	return symbol
}
func (s *StackCompileVisitor) VisitStringAtom(ctx *gen2.StringAtomContext) interface{} {
	valStr := ctx.STRING().GetText()
	// 字符串常量地址
	symbol := getStringConst(valStr[1:len(valStr)-1], s.GlobalScope)
	s.WriteInstr(&ir.StackInstr{
		OpCode:   vm.InstrSConst,
		Operands: []int32{symbol.GetAddress()}, // 字符串常量地址
	})
	return nil
}

func (s *StackCompileVisitor) VisitInstance(ctx *gen2.InstanceContext) interface{} {
	// 'new' sname=ID
	structName := ctx.GetSname().GetText()
	structSymbol := s.Scopes[ctx].Resolve(structName)
	if structSymbol == nil {
		s.Log.ErrorToken(ctx.GetStart(), fmt.Sprintf("undefined struct: %s", structName))
		return nil
	}
	s.WriteInstr(&ir.StackInstr{
		OpCode:   vm.InstrStruct,
		Operands: []int32{structSymbol.GetAddress()},
	})
	return nil
}

func (s *StackCompileVisitor) VisitQid(ctx *gen2.QidContext) interface{} {
	// 只处理右值
	vName := ctx.ID(0).GetText()
	scope := s.Scopes[ctx]
	loadInstr := vm.InstrLoad
	if scope.GetName() == GlobalScopeName {
		loadInstr = vm.InstrGLoad
	}
	sym := scope.Resolve(vName)
	s.WriteInstr(&ir.StackInstr{
		OpCode:   loadInstr,
		Operands: []int32{sym.GetAddress()}, // 变量地址
	})
	// 处理结构体字段访问
	for i := 1; i < len(ctx.AllID()); i++ {
		field := getStringConst(ctx.ID(i).GetText(), s.GlobalScope)
		s.WriteInstr(&ir.StackInstr{
			OpCode:   vm.InstrFLoad,
			Operands: []int32{field.GetAddress()}, // 变量地址
		})
	}
	return nil
}

func (s *StackCompileVisitor) VisitMultOp(ctx *gen2.MultOpContext) interface{} {
	// multOp	: MUL | DIV ;
	op := ctx.GetText()
	if op == "*" {
		s.WriteInstr(&ir.StackInstr{
			OpCode:   vm.InstrMul,
			Operands: []int32{},
		})
	} else if op == "/" {
		s.WriteInstr(&ir.StackInstr{
			OpCode:   vm.InstrDiv,
			Operands: []int32{},
		})
	}
	return nil
}

func (s *StackCompileVisitor) VisitAddOp(ctx *gen2.AddOpContext) interface{} {
	// addOp	: ADD | SUB ;
	op := ctx.GetText()
	if op == "+" {
		s.WriteInstr(&ir.StackInstr{
			OpCode:   vm.InstrAdd,
			Operands: []int32{},
		})
	} else if op == "-" {
		s.WriteInstr(&ir.StackInstr{
			OpCode:   vm.InstrSub,
			Operands: []int32{},
		})
	}
	return nil
}

func (s *StackCompileVisitor) VisitCompOp(ctx *gen2.CompOpContext) interface{} {
	// compOp	: EQ | LT | GT | GEQ | LEQ | NEQ ;
	op := ctx.GetText()
	switch op {
	case "==":
		s.WriteInstr(&ir.StackInstr{
			OpCode:   vm.InstrEQ,
			Operands: []int32{},
		})
	case "<":
		s.WriteInstr(&ir.StackInstr{
			OpCode:   vm.InstrLT,
			Operands: []int32{},
		})
	case ">":
		s.WriteInstr(&ir.StackInstr{
			OpCode:   vm.InstrGT,
			Operands: []int32{},
		})
	case ">=":
		s.WriteInstr(&ir.StackInstr{
			OpCode:   vm.InstrGEQ,
			Operands: []int32{},
		})
	case "<=":
		s.WriteInstr(&ir.StackInstr{
			OpCode:   vm.InstrLEQ,
			Operands: []int32{},
		})
	case "!=":
		s.WriteInstr(&ir.StackInstr{
			OpCode:   vm.InstrNEQ,
			Operands: []int32{},
		})
	default:
		s.Log.ErrorToken(ctx.GetStart(), fmt.Sprintf("unknown compOp %s", op))
	}
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
