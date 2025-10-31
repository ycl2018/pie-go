package compile

import (
	"fmt"
	"strings"

	"github.com/ycl2018/pie-go/interpreter/compile/ir"
)

type Symbol interface {
	GetName() string
	SetScope(scope Scope)
	Scope() Scope
	GetAddress() int32
	SetAddress(address int32)
}

// FunctionSymbol 函数符号
type FunctionSymbol struct {
	Name           string
	FormalArgs     []Symbol
	Address        int32 // 函数在全局符号表中的地址
	EnclosingScope Scope
	BodyScope      *LocalScope
	Line           int32
	Code           []*ir.StackInstr
	CodeAddr       int32 // 函数入口
}

func (f *FunctionSymbol) Dump() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf(".def %s args=%d locals=%d\n", f.Name, len(f.FormalArgs), f.LocalNums()))
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
			panic(fmt.Sprintf("duplicate formal arg %s in function %s line %d",
				symbol.GetName(), f.Name, f.Line))
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

func (f *FunctionSymbol) LocalNums() int {
	locals := 0
	if f.BodyScope != nil {
		// 过滤 structSymbol
		locals = int(f.BodyScope.LocalVarAllocator)
	}
	return locals
}

// StructSymbol 结构体符号
type StructSymbol struct {
	Fields         []Symbol
	Name           string
	Address        int32 // 结构体在全局符号表中的地址
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
	Address        int32 // 常量在全局符号表中的地址
	EnclosingScope Scope
	Line           int32
	Kind           ir.ConstKind
	Value          any
	// Kind==ConstStruct:structName+存储字段名常量的地址
	// Kind==ConstFunc:存储args,locals的数量，以及函数体代码地址
	Fields []int32
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

func DumpSymbol(s Symbol) string {
	switch s := s.(type) {
	case *ConstSymbol:
		switch s.Kind {
		case ir.ConstFunc:
			return fmt.Sprintf("%d:%s\t%s\targs=%d locals=%d addr=%d\n",
				s.Address, s.Kind, s.Name, s.Fields[1], s.Fields[2], s.Fields[3])
		case ir.ConstStruct:
			// 结构体常量
			var sb strings.Builder
			sb.WriteString(fmt.Sprintf("%d:%s\t%s\n", s.Address, s.Kind, s.Name))
			for i, field := range s.Fields {
				if i == 0 {
					continue
				}
				sb.WriteString(fmt.Sprintf("\tfield%d:\t.const->%d\n", i-1, field))
			}
			return sb.String()
		case ir.ConstString:
			return fmt.Sprintf("%d:%s\t%q\n", s.Address, s.Kind, s.Value)
		default:
			return fmt.Sprintf("%d:%s\t%s\n", s.Address, s.Kind, s.Name)
		}
	case *StructSymbol, *FunctionSymbol:
		return ""
	default:
		return fmt.Sprintf("%d:\t%s\n", s.GetAddress(), s.GetName())
	}
}
