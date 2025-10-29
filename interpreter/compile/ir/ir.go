package ir

import "fmt"

type Instr int
type StackInstr struct {
	OpCode   Instr
	Operands []int32
}

// ConstKind 常量类型，存储于全局常量池
type ConstKind int

const (
	ConstFloat32 = iota
	ConstString
	ConstStruct
)

func (c ConstKind) String() string {
	switch c {
	case ConstFloat32:
		return "float32"
	case ConstString:
		return "string"
	case ConstStruct:
		return "struct"
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
