package compile

import (
	"fmt"
	"strings"

	"github.com/ycl2018/pie-go/gen"
	"github.com/ycl2018/pie-go/interpreter/compile/ir"
)

const GlobalScopeName = "global"

type Symbol interface {
	GetName() string
	SetScope(scope Scope)
	Scope() Scope
	GetAddress() int32
	SetAddress(address int32)
}

type Scope interface {
	Resolve(string) Symbol
	GetName() string
	Define(symbol Symbol)
	ParentScope() Scope
}

type GlobalScope struct {
	Symbols map[string]Symbol
	Consts  map[string]*ConstSymbol
}

func NewGlobalScope() *GlobalScope {
	return &GlobalScope{
		Symbols: map[string]Symbol{},
		Consts:  map[string]*ConstSymbol{},
	}
}

func (g *GlobalScope) Define(symbol Symbol) {
	switch s := symbol.(type) {
	case *ConstSymbol:
		s.SetScope(g)
		name := s.Name
		if g.Consts[name] != nil {
			return
		}
		s.SetAddress(int32(len(g.Consts)))
		g.Consts[s.GetName()] = s
	case *StructSymbol:
		if g.Symbols[symbol.GetName()] != nil {
			return
		}
		fields := make([]int32, len(s.Fields))
		for _, field := range s.Fields {
			sym := getStringConst(field.GetName(), g)
			fields = append(fields, sym.GetAddress())
		}
		constSym := &ConstSymbol{
			Name:   s.Name,
			Kind:   ir.ConstStruct,
			Fields: fields,
		}
		g.Define(constSym) // 存储结构体常量到全局常量池
		g.Symbols[s.Name] = s
		s.SetScope(g)
		symbol.SetAddress(constSym.GetAddress())
	default:
		if g.Symbols[symbol.GetName()] != nil {
			return
		}
		symbol.SetAddress(int32(len(g.Symbols)))
		g.Symbols[symbol.GetName()] = symbol
		symbol.SetScope(g)
	}
}

func (g *GlobalScope) DefineOrGetConst(symbol *ConstSymbol) (Symbol, bool) {
	name := symbol.Name
	if g.Consts[name] != nil {
		return g.Consts[name], false
	}
	g.Define(symbol)
	return symbol, true
}

func (g *GlobalScope) Resolve(name string) Symbol {
	if v, ok := g.Symbols[name]; ok {
		return v
	}
	return nil
}

func (g *GlobalScope) GetName() string {
	return GlobalScopeName
}
func (g *GlobalScope) ParentScope() Scope {
	return nil
}

type LocalScope struct {
	ID             int32
	Symbols        map[string]Symbol
	EnclosingScope Scope
	StructConsts   map[string]*ConstSymbol
	BaseAllocAddr  int32
}

func (l *LocalScope) Define(symbol Symbol) {
	switch s := symbol.(type) {
	case *ConstSymbol:
		panic("const symbol define in local scope")
	case *StructSymbol:
		if l.Symbols[symbol.GetName()] != nil {
			panic(fmt.Sprintf("duplicate struct %s in local scope", symbol.GetName()))
			return
		}
		// 同时存储在全局 scope 中
		globalScope := l.EnclosingScope
		for globalScope != nil && globalScope.GetName() != GlobalScopeName {
			globalScope = globalScope.ParentScope()
		}
		gName := fmt.Sprintf("%d::%s", l.ID, symbol.GetName())
		globalScope.Define(&StructSymbol{
			Fields: s.Fields,
			// 通过 scope ID 来避免命名冲突
			Name: gName,
		})
		sym := globalScope.Resolve(gName)
		symbol.SetAddress(sym.GetAddress()) // 指向全局 scope 中的结构体
		l.Symbols[symbol.GetName()] = symbol
		symbol.SetScope(l)
	default:
		if l.Symbols[symbol.GetName()] != nil {
			panic(fmt.Sprintf("duplicate %s in local scope", symbol.GetName()))
			return
		}
		symbol.SetAddress(l.BaseAllocAddr + int32(len(l.Symbols)))
		l.Symbols[symbol.GetName()] = symbol
		symbol.SetScope(l)
	}
}

func (l *LocalScope) Resolve(name string) Symbol {
	if v, ok := l.Symbols[name]; ok {
		return v
	}
	if l.EnclosingScope != nil {
		return l.EnclosingScope.Resolve(name)
	}
	return nil
}

func (l *LocalScope) GetName() string {
	return "local"
}
func (l *LocalScope) ParentScope() Scope {
	return l.EnclosingScope
}

// FunctionSymbol 函数符号
type FunctionSymbol struct {
	Name           string
	ID             int32
	FormalArgs     []Symbol
	BlockAst       gen.ISlistContext
	Address        int32
	EnclosingScope Scope
	BodyScope      *LocalScope
	Line           int32
	Code           []*ir.StackInstr
}

func (f *FunctionSymbol) Dump() string {
	var sb strings.Builder
	var locals int
	// 计算 locals 数量
	if f.BodyScope != nil {
		locals = len(f.BodyScope.Symbols)
	}
	fmt.Fprintf(&sb, ".def %s args=%d locals=%d\n", f.Name, len(f.FormalArgs), locals)
	for _, instr := range f.Code {
		sb.WriteString(instr.Dump())
	}
	return sb.String()
}

func (f *FunctionSymbol) GetAddress() int32 {
	return f.Address
}

func (f *FunctionSymbol) SetAddress(address int32) {
	f.Address = address
}

func (f *FunctionSymbol) SetScope(scope Scope) {
	f.EnclosingScope = scope
}

func (f *FunctionSymbol) Scope() Scope {
	return f.EnclosingScope
}

func (f *FunctionSymbol) Define(symbol Symbol) {
	// 检查是否重复定义
	for _, arg := range f.FormalArgs {
		if arg.GetName() == symbol.GetName() {
			panic(fmt.Sprintf("duplicate formal arg %s in function %s line %d", symbol.GetName(), f.Name, f.Line))
			return
		}
	}
	// 分配地址
	symbol.SetAddress(int32(len(f.FormalArgs)))
	f.FormalArgs = append(f.FormalArgs, symbol)
	symbol.SetScope(f)
}

func (f *FunctionSymbol) Resolve(s string) Symbol {
	for _, arg := range f.FormalArgs {
		if arg.GetName() == s {
			return arg
		}
	}
	if f.EnclosingScope != nil {
		return f.EnclosingScope.Resolve(s)
	}
	return nil
}

func (f *FunctionSymbol) GetName() string {
	return f.Name
}
func (f *FunctionSymbol) ParentScope() Scope {
	return f.EnclosingScope
}

// StructSymbol 结构体符号
type StructSymbol struct {
	Fields         []Symbol
	Name           string
	Address        int32
	Line           int32
	EnclosingScope Scope
}

func (s *StructSymbol) GetAddress() int32 {
	return s.Address
}

func (s *StructSymbol) SetAddress(address int32) {
	s.Address = address
}

func NewStructSymbol(name string, parentScope Scope) *StructSymbol {
	return &StructSymbol{
		Name:           name,
		EnclosingScope: parentScope,
	}
}

func (s *StructSymbol) SetScope(scope Scope) {
	s.EnclosingScope = scope
}

func (s *StructSymbol) Define(symbol Symbol) {
	var hasField bool
	for i := 0; i < len(s.Fields); i++ {
		if s.Fields[i].GetName() == symbol.GetName() {
			hasField = true
			s.Fields[i] = symbol
		}
	}
	if !hasField {
		s.Fields = append(s.Fields, symbol)
	}
	symbol.SetScope(s)
}

func (s *StructSymbol) GetName() string {
	return s.Name
}

func (s *StructSymbol) Resolve(name string) Symbol {
	for i := 0; i < len(s.Fields); i++ {
		if s.Fields[i].GetName() == name {
			return s.Fields[i]
		}
	}
	if s.EnclosingScope != nil {
		return s.EnclosingScope.Resolve(name)
	}
	return nil
}
func (s *StructSymbol) Scope() Scope {
	return s.EnclosingScope
}

func (s *StructSymbol) ParentScope() Scope {
	return s.EnclosingScope
}

// VariableSymbol 变量符号
type VariableSymbol struct {
	Name           string
	Address        int32
	EnclosingScope Scope
	Line           int32
}

func (v *VariableSymbol) GetAddress() int32 {
	return v.Address
}

func (v *VariableSymbol) SetAddress(address int32) {
	v.Address = address
}

func (v *VariableSymbol) SetScope(scope Scope) {
	v.EnclosingScope = scope
}

func (v *VariableSymbol) GetName() string {
	return v.Name
}

func (v *VariableSymbol) Scope() Scope {
	return v.EnclosingScope
}

type ConstSymbol struct {
	Name           string
	Address        int32
	EnclosingScope Scope
	Line           int32
	Kind           ir.ConstKind
	Fields         []int32 // 字段名常量的地址
}

func (c *ConstSymbol) GetName() string {
	return c.Name
}

func (c *ConstSymbol) SetScope(scope Scope) {
	c.EnclosingScope = scope
}

func (c *ConstSymbol) GetAddress() int32 {
	return c.Address
}

func (c *ConstSymbol) SetAddress(address int32) {
	c.Address = address
}

func (c *ConstSymbol) Scope() Scope {
	return c.EnclosingScope
}
