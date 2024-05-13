package stack

import "github.com/ycl2018/pie-go/interpreter/asm"

const (
	InstrIAdd = iota + 1
	InstrISub = iota + 1
	InstrIMul = iota + 1
	InstrILT  = iota + 1
	InstrIEQ  = iota + 1
	InstrFAdd = iota + 1
	InstrFSub = iota + 1
	InstrFMul = iota + 1
	InstrFLT  = iota + 1
	InstrFEQ  = iota + 1

	InstrI2F    = iota + 1 // int to float
	InstrCall   = iota + 1
	InstrReturn = iota + 1
	InstrBR     = iota + 1 // branch
	InstrBRT    = iota + 1 // branch if true
	InstrBRF    = iota + 1 // branch if false
	InstrCConst = iota + 1 // push constant
	InstrIConst = iota + 1
	InstrFConst = iota + 1
	InstrSConst = iota + 1

	InstrLoad   = iota + 1
	InstrGLoad  = iota + 1 // global load
	InstrFLoad  = iota + 1 // filed load
	InstrStore  = iota + 1 // local store
	InstrGStore = iota + 1 // global store
	InstrFStore = iota + 1 // field store
	InstrPrint  = iota + 1
	InstrStruct = iota + 1 // push struct on stack
	InstrNull   = iota + 1 // push null on stack
	InstrPop    = iota + 1 // pop stack

	InstrHalt = iota + 1
)

// 基于栈的指令集
var Instructions = []*asm.Instruction{
	nil,
	{"iadd", []int32{}},
	{"isub", []int32{}},
	{"imul", []int32{}},
	{"ilt", []int32{}},
	{"ieq", []int32{}},
	{"fadd", []int32{}},
	{"fsub", []int32{}},
	{"fmul", []int32{}},
	{"flt", []int32{}},
	{"feq", []int32{}},
	{"itof", []int32{}},
	{"call", []int32{asm.FUNC}},
	{"ret", []int32{}},
	{"br", []int32{asm.INT}},
	{"brt", []int32{asm.INT}},
	{"brf", []int32{asm.INT}},
	{"cconst", []int32{asm.INT}},
	{"iconst", []int32{asm.INT}},
	{"fconst", []int32{asm.POLL}},
	{"sconst", []int32{asm.POLL}},
	{"load", []int32{asm.INT}},
	{"gload", []int32{asm.INT}},
	{"fload", []int32{asm.INT}},
	{"store", []int32{asm.INT}},
	{"gstore", []int32{asm.INT}},
	{"fstore", []int32{asm.INT}},
	{"print", []int32{}},
	{"struct", []int32{asm.INT}},
	{"null", []int32{}},
	{"pop", []int32{}},
	{"halt", []int32{}},
}
