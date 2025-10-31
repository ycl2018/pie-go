package vm

import (
	"fmt"
)

type FunctionSymbol struct {
	Name	   string
	Args       uint8
	Locals     uint16
	Addr       int32
}

type Instruction struct {
	Name       string
	OpRandType []int32 // 最长3个操作数
}

type Const struct {
	Value any
	Kind  ConstKind
}

// ConstKind 常量类型，存储于全局常量池
// 注意和 interpreter/compile/ir/ir.go 中的 ConstKind 保持一致
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


type ConstStructDef struct {
	Name        string
	MemberNames []string
}