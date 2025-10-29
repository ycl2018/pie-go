package compile

import (
	"fmt"

	"github.com/ycl2018/pie-go/gen"
	"github.com/ycl2018/pie-go/interpreter/compile/ir"
)

type Symbol interface {
	GetName() string
	SetScope(scope Scope)
	GetAddress() int32
	SetAddress(address int32)
}

type Scope interface {
	Resolve(string) Symbol
	GetName() string
	Define(symbol Symbol)
}

type GlobalScope struct {
	Symbols map[string]Symbol
	Consts  map[string]*ConstSymbol
}

func (g *GlobalScope) Define(symbol Symbol) {
	switch symbol.(type) {
	case *ConstSymbol:
		s := symbol.(*ConstSymbol)
		s.SetScope(g)
		name := s.Name
		if g.Consts[name] != nil {
			return
		}
		s.SetAddress(int32(len(g.Consts)))
		g.Consts[s.GetName()] = s
	default:
		if g.Symbols[symbol.GetName()] != nil {
			return
		}
		symbol.SetAddress(int32(len(g.Symbols)))
		g.Symbols[symbol.GetName()] = symbol
		symbol.SetScope(g)
	}
}

func (g *GlobalScope) DefineOrGetConst(symbol *ConstSymbol) (Symbol,bool) {
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
	return "global"
}

type LocalScope struct {
	Symbols        map[string]Symbol
	EnclosingScope Scope
	StructConsts   map[string]*ConstSymbol
}

func (l *LocalScope) Define(symbol Symbol) {
	switch symbol.(type) {
	case *ConstSymbol:
		if s := symbol.(*ConstSymbol); s.Kind == ir.ConstStruct {
			panic(fmt.Sprintf("local scope can only define struct const, but got %s", s.Kind))
		}
		if l.StructConsts[symbol.GetName()] != nil {
			panic(fmt.Sprintf("duplicate struct const %s in local scope", symbol.GetName()))
		}
		symbol.SetAddress(int32(len(l.StructConsts)))
		symbol.SetScope(l)
		l.StructConsts[symbol.GetName()] = symbol.(*ConstSymbol)
	default:
		if l.Symbols[symbol.GetName()] != nil {
			return
		}
		symbol.SetAddress(int32(len(l.Symbols)))
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

// FunctionSymbol 函数符号
type FunctionSymbol struct {
	Name           string
	FormalArgs     []Symbol
	BlockAst       gen.ISlistContext
	Address        int32
	EnclosingScope Scope
	BodyScope      *LocalScope
	Line           int32
	Code           []*ir.StackInstr
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

// VariableSymbol 变量符号
type VariableSymbol struct {
	Name    string
	Address int32
	Scope   Scope
	Line    int32
}

func (v *VariableSymbol) GetAddress() int32 {
	return v.Address
}

func (v *VariableSymbol) SetAddress(address int32) {
	v.Address = address
}

func (v *VariableSymbol) SetScope(scope Scope) {
	v.Scope = scope
}

func (v *VariableSymbol) GetName() string {
	return v.Name
}

type ConstSymbol struct {
	Name    string
	Address int32
	Scope   Scope
	Line    int32
	Kind    ir.ConstKind
	Fields  []*VariableSymbol // 结构体常量的字段
}

func (c *ConstSymbol) GetName() string {
	return c.Name
}

func (c *ConstSymbol) SetScope(scope Scope) {
	c.Scope = scope
}

func (c *ConstSymbol) GetAddress() int32 {
	return c.Address
}

func (c *ConstSymbol) SetAddress(address int32) {
	c.Address = address
}
