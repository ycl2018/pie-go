package compile

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	gen2 "github.com/ycl2018/pie-go/gen"
)

var _ gen2.PieVisitor = (*RuntimeVisitor)(nil)

type RuntimeVisitor struct {
	gen2.BasePieVisitor
	InterpreterListener InterpreterListener
	Scopes              map[antlr.ParserRuleContext]Scope // 作用域树&符号表存储
	GlobalMemory        *MemorySpace                      // global memory
	CurMemory           *MemorySpace
	FunctionStack       []*FunctionMemorySpace // call stack
}

func NewRuntimeVisitor(iListener InterpreterListener, scopes map[antlr.ParserRuleContext]Scope) *RuntimeVisitor {
	globalMem := NewMemorySpace("global")
	ret := &RuntimeVisitor{
		InterpreterListener: iListener,
		Scopes:              scopes,
		GlobalMemory:        globalMem,
		CurMemory:           globalMem,
	}
	ret.BasePieVisitor = gen2.BasePieVisitor{ParseTreeVisitor: BaseVisitor{ret}}
	return ret
}

func (v *RuntimeVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *RuntimeVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	for _, tree := range node.GetChildren() {
		v.Visit(tree.(antlr.ParseTree))
	}
	return nil
}

func (v *RuntimeVisitor) VisitProgram(ctx *gen2.ProgramContext) interface{} {
	for _, stmtCtx := range ctx.AllStatement() {
		v.Visit(stmtCtx)
	}
	return nil
}

func (v *RuntimeVisitor) VisitAssignementStatement(ctx *gen2.AssignementStatementContext) interface{} {
	// qid '=' expr NL
	qid := ctx.Qid()
	expr := ctx.Expr()
	value := v.Visit(expr)
	if qid.GetChildCount() > 1 {
		v.FieldAssign(qid, value)
		return nil
	}
	// variableAssign
	variableName := qid.ID(0).GetText()
	memSpace := v.GetSpaceWithSymbol(variableName)
	if memSpace == nil {
		// create in current memory
		memSpace = v.CurMemory
	}
	memSpace.Members[variableName] = value
	return nil
}

func (v *RuntimeVisitor) FieldAssign(ctx gen2.IQidContext, value any) {
	structInstance, err := v.GetStructInstanceWithQID(ctx)
	if err != nil {
		v.InterpreterListener.ErrorToken(ctx.GetStart(), "can't assign; %s", err)
		return
	}
	filedName := ctx.ID(len(ctx.AllID()) - 1).GetText()
	if structInstance.StructDef.Resolve(filedName) == nil {
		v.InterpreterListener.ErrorToken(ctx.GetStart(), "can't assign; no such field %s in struct %s", filedName, structInstance.StructDef.Name)
		return
	}
	structInstance.Members[filedName] = value
}

func (v *RuntimeVisitor) VisitReturnStatement(ctx *gen2.ReturnStatementContext) interface{} {
	// 'return' expr NL
	returnVal := v.Visit(ctx.Expr())
	panic(returnVal)
}

func (v *RuntimeVisitor) VisitPrintStatement(ctx *gen2.PrintStatementContext) interface{} {
	// 'print' expr NL
	value := v.Visit(ctx.Expr())
	if value == nil {
		value = "null"
	}
	fmt.Println(value)
	return nil
}

func (v *RuntimeVisitor) VisitIfStatement(ctx *gen2.IfStatementContext) interface{} {
	// 'if' expr c=slist ('else' el=slist)?
	valid := v.Visit(ctx.Expr()).(bool)
	if valid && ctx.GetC() != nil {
		v.Visit(ctx.GetC())
	} else if ctx.GetEl() != nil {
		v.Visit(ctx.GetEl())
	}
	return nil
}

func (v *RuntimeVisitor) VisitWhileStatement(ctx *gen2.WhileStatementContext) interface{} {
	// 'while' expr slist
	value := v.Visit(ctx.Expr()).(bool)
	for value {
		v.Visit(ctx.Slist())
		value = v.Visit(ctx.Expr()).(bool)
	}
	return nil
}

func (v *RuntimeVisitor) VisitCall(ctx *gen2.CallContext) interface{} {
	// name=ID '(' (expr (',' expr)*)? ')'
	fname := ctx.ID().GetText()
	scope := v.Scopes[ctx]
	fnSymbol := scope.Resolve(fname)
	if fnSymbol == nil {
		v.InterpreterListener.ErrorToken(ctx.GetStart(), "function %s not found", fname)
		return nil
	}
	savedCurMemory := v.CurMemory
	funcSymbol := fnSymbol.(*FunctionSymbol)
	curFnStack := &FunctionMemorySpace{
		MemorySpace: MemorySpace{
			Name:    fmt.Sprintf("%s:%d", fname, len(v.FunctionStack)),
			Members: make(map[string]any),
		},
		FunctionSymbol: funcSymbol,
	}
	v.CurMemory = &curFnStack.MemorySpace
	if len(ctx.AllExpr()) != len(funcSymbol.FormalArgs) {
		v.InterpreterListener.ErrorToken(ctx.GetStart(), "function %s expects %d arguments but got %d", fname, len(funcSymbol.FormalArgs), len(ctx.AllExpr()))
		return nil
	}
	for i, arg := range ctx.AllExpr() {
		actualArg := v.Visit(arg)
		v.CurMemory.Members[funcSymbol.FormalArgs[i].GetName()] = actualArg
	}
	// 参数传递完之后，才能将函数栈push进去
	v.FunctionStack = append(v.FunctionStack, curFnStack)
	retVal := v.visitFuncBlock(funcSymbol)
	v.FunctionStack = v.FunctionStack[:len(v.FunctionStack)-1]
	v.CurMemory = savedCurMemory
	return retVal
}
func (v *RuntimeVisitor) visitFuncBlock(funcSymbol *FunctionSymbol) (retVal interface{}) {
	defer func() {
		retVal = recover()
	}()
	v.Visit(funcSymbol.BlockAst)
	return
}

func (v *RuntimeVisitor) VisitExpr(ctx *gen2.ExprContext) interface{} {
	// addexpr (compOp addexpr)?
	if len(ctx.AllAddexpr()) == 1 {
		return v.Visit(ctx.Addexpr(0))
	} else if len(ctx.AllAddexpr()) == 2 {
		left := v.Visit(ctx.Addexpr(0))
		right := v.Visit(ctx.Addexpr(1))
		switch ctx.CompOp().GetText() {
		case "==":
			return v.IsEquals(left, right)
		case "<":
			return v.LessThan(left, right)
		default:
			v.InterpreterListener.ErrorToken(ctx.GetStart(), "invalid comp op: %s", ctx.CompOp().GetText())
		}
	}
	panic("not implemented")
}

func (v *RuntimeVisitor) VisitAddexpr(ctx *gen2.AddexprContext) interface{} {
	// mulexpr (addOp mulexpr)*
	if len(ctx.AllMulexpr()) == 1 {
		return v.Visit(ctx.Mulexpr(0))
	}
	var accumulator = v.Visit(ctx.Mulexpr(0))
	for i := 0; i < len(ctx.AllAddOp()); i++ {
		op := ctx.AddOp(i)
		right := v.Visit(ctx.Mulexpr(i + 1))
		switch op.GetText() {
		case "+":
			accumulator = v.Add(accumulator, right)
		case "-":
			accumulator = v.Sub(accumulator, right)
		default:
			v.InterpreterListener.ErrorToken(ctx.AddOp(i+1).GetStart(), "invalid add op: %s", op.GetText())
		}
	}
	return accumulator
}

func (v *RuntimeVisitor) VisitMulexpr(ctx *gen2.MulexprContext) interface{} {
	// atom (multOp atom)*
	if len(ctx.AllAtom()) == 1 {
		return v.Visit(ctx.Atom(0))
	}
	var accumulator = v.Visit(ctx.Atom(0))
	for i := 0; i < len(ctx.AllMultOp()); i++ {
		op := ctx.MultOp(i)
		right := v.Visit(ctx.Atom(i + 1))
		switch op.GetText() {
		case "*":
			accumulator = v.Mul(accumulator, right)
		case "/":
			accumulator = v.Div(accumulator, right)
		default:
			v.InterpreterListener.ErrorToken(ctx.Atom(i+1).GetStart(), "invalid mult op: %s", op.GetText())
		}
	}
	return accumulator
}

func (v *RuntimeVisitor) IsEquals(left, right any) bool {
	return v.op(left, right, "==").(bool)
}
func (v *RuntimeVisitor) LessThan(left, right any) bool {
	return v.op(left, right, "<").(bool)
}

type unaryCalculator[T interface{ float64 | int64 | rune }] struct {
	l T
	r T
}

func (c unaryCalculator[T]) Cal(op string) any {
	switch op {
	case "+":
		return c.l + c.r
	case "-":
		return c.l - c.r
	case "*":
		return c.l * c.r
	case "/":
		return c.l / c.r
	case "<":
		return c.l < c.r
	case "==":
		return c.l == c.r
	}
	return nil
}

func (v *RuntimeVisitor) op(left, right any, op string) interface{} {
	lType, rType := reflect.TypeOf(left), reflect.TypeOf(right)
	if lType != rType {
		if lType.Kind() == reflect.Float64 || rType.Kind() == reflect.Float64 {
			left, right = cast2Float64(left), cast2Float64(right)

		} else if lType.Kind() == reflect.Int64 || rType.Kind() == reflect.Int64 {
			left, right = cast2Int64(left), cast2Int64(right)
		}
	}
	switch left.(type) {
	case int64:
		return unaryCalculator[int64]{l: left.(int64), r: right.(int64)}.Cal(op)
	case float64:
		return unaryCalculator[float64]{l: left.(float64), r: right.(float64)}.Cal(op)
	case rune:
		return unaryCalculator[rune]{l: left.(rune), r: right.(rune)}.Cal(op)
	}
	v.InterpreterListener.Errorf("invalid type for op: %s %s %s", reflect.TypeOf(left), op, reflect.TypeOf(right))
	return nil
}

func cast2Int64(v any) int64 {
	switch v.(type) {
	case int64:
		return v.(int64)
	case rune:
		return int64(v.(rune))
	}
	return 0
}

func cast2Float64(v any) float64 {
	switch v.(type) {
	case int64:
		return float64(v.(int64))
	case float64:
		return v.(float64)
	case rune:
		return float64(v.(rune))
	case string:
		ret, _ := strconv.ParseFloat(v.(string), 64)
		return ret
	default:
		return 0
	}
}

func cast2String(v any) string {
	switch v.(type) {
	case int64:
		return fmt.Sprintf("%d", v.(int64))
	case float64:
		return fmt.Sprintf("%f", v.(float64))
	case string:
		return v.(string)
	case rune:
		return fmt.Sprintf("%c", v.(rune))
	default:
		return fmt.Sprintf("%v", v)
	}
}

func (v *RuntimeVisitor) Add(left, right any) any {
	lType, rType := reflect.TypeOf(left), reflect.TypeOf(right)
	if lType.Kind() == reflect.String || rType.Kind() == reflect.String {
		return cast2String(left) + cast2String(right)
	}
	return v.op(left, right, "+")
}

func (v *RuntimeVisitor) Sub(left, right any) any {
	return v.op(left, right, "-")
}

func (v *RuntimeVisitor) Mul(left, right any) any {
	return v.op(left, right, "*")
}

func (v *RuntimeVisitor) Div(left, right any) any {
	return v.op(left, right, "/")
}

func (v *RuntimeVisitor) VisitIntAtom(ctx *gen2.IntAtomContext) interface{} {
	intStr := ctx.INT().GetText()
	if intStr != "" {
		intValue, err := strconv.ParseInt(intStr, 10, 64)
		if err != nil {
			v.InterpreterListener.ErrorToken(ctx.GetStart(), "parse int error: %s", err)
			return nil
		}
		return intValue
	}
	return nil
}

func (v *RuntimeVisitor) VisitCharAtom(ctx *gen2.CharAtomContext) interface{} {
	charStr := ctx.CHAR().GetText()
	if charStr != "" {
		charValue := rune(charStr[0])
		return charValue
	}
	return nil
}

func (v *RuntimeVisitor) VisitFloatAtom(ctx *gen2.FloatAtomContext) interface{} {
	intStr := ctx.FLOAT().GetText()
	if intStr != "" {
		intValue, err := strconv.ParseFloat(intStr, 64)
		if err != nil {
			v.InterpreterListener.ErrorToken(ctx.GetStart(), "parse float error: %s", err)
			return nil
		}
		return intValue
	}
	return nil
}

func (v *RuntimeVisitor) VisitStringAtom(ctx *gen2.StringAtomContext) interface{} {
	return ctx.STRING().GetText()[1 : len(ctx.STRING().GetText())-1]
}

func (v *RuntimeVisitor) VisitParenthesizedAtom(ctx *gen2.ParenthesizedAtomContext) interface{} {
	// '(' expr ')'
	return v.Visit(ctx.Expr())
}

func (v *RuntimeVisitor) VisitInstance(ctx *gen2.InstanceContext) interface{} {
	// 'new' sname=ID
	scope := v.Scopes[ctx]
	if scope == nil {
		return nil
	}
	// 解析结构体符号引用
	structSymbol := scope.Resolve(ctx.ID().GetText()).(*StructSymbol)
	if structSymbol != nil {
		return NewStructInstance(structSymbol)
	}
	return nil
}

func (v *RuntimeVisitor) VisitQid(ctx *gen2.QidContext) interface{} {
	// ID ('.' ID)*
	if len(ctx.AllID()) > 1 {
		return v.FiledLoad(ctx)
	}
	return v.VariableLoad(ctx)
}

func (v *RuntimeVisitor) VisitQidAtom(ctx *gen2.QidAtomContext) interface{} {
	// ID ('.' ID)*
	return v.Visit(ctx.Qid())
}

func (v *RuntimeVisitor) FiledLoad(ctx *gen2.QidContext) any {
	structInstance, err := v.GetStructInstanceWithQID(ctx)
	if err != nil || structInstance == nil {
		return nil
	}
	filedName := ctx.ID(len(ctx.AllID()) - 1).GetText()
	if structInstance.StructDef.Resolve(filedName) == nil {
		v.InterpreterListener.ErrorToken(ctx.GetStart(), "no such field %s in struct %s", filedName, structInstance.StructDef.Name)
	}
	return structInstance.Get(filedName)
}

func (v *RuntimeVisitor) VariableLoad(ctx gen2.IQidContext) any {
	variableName := ctx.ID(0).GetText()
	memSpace := v.GetSpaceWithSymbol(variableName)
	if memSpace != nil {
		return memSpace.Get(variableName)
	}
	v.InterpreterListener.ErrorToken(ctx.GetStart(), "no such variable %s", ctx.ID(0).GetSymbol())
	return nil
}

func (v *RuntimeVisitor) GetSpaceWithSymbol(name string) *MemorySpace {
	if len(v.FunctionStack) > 0 {
		funcSpace := v.FunctionStack[len(v.FunctionStack)-1]
		if funcSpace.Contains(name) {
			return &funcSpace.MemorySpace
		}
	}
	if v.GlobalMemory.Contains(name) {
		return v.GlobalMemory
	}
	return nil
}

func (v *RuntimeVisitor) GetStructInstanceWithQID(ctx gen2.IQidContext) (*StructInstance, error) {
	structInstance, ok := v.VariableLoad(ctx).(*StructInstance)
	if !ok {
		return nil, fmt.Errorf("%s is not a struct", ctx.ID(0).GetText())
	}
	var curFiledCompleteName = ctx.ID(0).GetText()
	for i := 1; i < len(ctx.AllID())-1; i++ {
		fieldName := ctx.ID(i).GetText()
		fieldSymbol := structInstance.StructDef.Resolve(fieldName)
		if fieldSymbol == nil {
			return nil, fmt.Errorf("%s has no %s filed", structInstance.StructDef.Name, fieldName)
		}
		curFiledCompleteName = curFiledCompleteName + "." + fieldName
		fieldValue := structInstance.Get(fieldName)
		if curStructInstance, ok := fieldValue.(*StructInstance); !ok || fieldValue == nil {
			return nil, fmt.Errorf("%s is not a struct", curFiledCompleteName)
		} else {
			structInstance = curStructInstance
		}
	}
	return structInstance, nil
}
