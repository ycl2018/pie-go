package tree

import (
	"github.com/ycl2018/pie-go/gen"
)

type Symbol interface {
	GetName() string
	SetScope(scope Scope)
}

type Scope interface {
	Resolve(string) Symbol
	GetName() string
	Define(symbol Symbol)
}

type GlobalScope struct {
	Symbols map[string]Symbol
}

func (g *GlobalScope) Define(symbol Symbol) {
	if g.Symbols[symbol.GetName()] != nil {
		return
	}
	g.Symbols[symbol.GetName()] = symbol
	symbol.SetScope(g)
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
}

func (l *LocalScope) Define(symbol Symbol) {
	if l.Symbols[symbol.GetName()] != nil {
		return
	}
	l.Symbols[symbol.GetName()] = symbol
	symbol.SetScope(l)
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
	Name       string
	FormalArgs []Symbol
	BlockAst   gen.ISlistContext

	EnclosingScope Scope
}

func (f *FunctionSymbol) SetScope(scope Scope) {
	f.EnclosingScope = scope
}

func (f *FunctionSymbol) Define(symbol Symbol) {
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
	Fields []Symbol
	Name   string

	EnclosingScope Scope
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
	Name  string
	Scope Scope
}

func (v *VariableSymbol) SetScope(scope Scope) {
	v.Scope = scope
}

func (v *VariableSymbol) GetName() string {
	return v.Name
}
