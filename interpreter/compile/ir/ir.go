package ir

import (
	"fmt"

	"github.com/ycl2018/pie-go/interpreter/vm"
)

type StackInstr struct {
	OpCode   int
	Operands []int32
}

func (s StackInstr) Dump() string {
	if len(s.Operands) > 0 {
		switch s.OpCode {
		case vm.InstrBR, vm.InstrBRT, vm.InstrBRF:
			return fmt.Sprintf("%-11s\ttag#%d\n", vm.Instr(s.OpCode), s.Operands[0])
		}
		return fmt.Sprintf("%-11s\t%v\n", vm.Instr(s.OpCode), s.Operands[0])
	} else {
		return fmt.Sprintf("%-11s\t\n", vm.Instr(s.OpCode))
	}

}

// ConstKind 常量类型，存储于全局常量池
type ConstKind uint8

const (
	ConstFloat32 ConstKind = iota
	ConstString
	ConstStruct
	ConstFunc
)

func (c ConstKind) String() string {
	switch c {
	case ConstFloat32:
		return "float32"
	case ConstString:
		return "string"
	case ConstStruct:
		return "struct"
	case ConstFunc:
		return "func"
	default:
		panic(fmt.Sprintf("unknown const kind %d", c))
	}
}

type StructDecl struct {
	Name   string
	Fields []string
}
type FuncDecl struct {
	FuncName string
	NArgs    int32
	NLocals  int32
	Structs  []StructDecl
	BodyCode []StackInstr
}
