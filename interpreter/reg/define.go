package reg

import "github.com/ycl2018/pie-go/interpreter/asm"

const (
	REG  = asm.REG
	FUNC = asm.FUNC
	INT  = asm.INT
	POLL = asm.POLL
)

const (
	InstrIAdd = iota + 1
	InstrISub
	InstrIMul
	InstrILT
	InstrIEQ
	InstrFAdd
	InstrFSub
	InstrFMul
	InstrFLT
	InstrFEQ

	InstrI2F // int to float
	InstrCall
	InstrReturn
	InstrBR     // branch
	InstrBRT    // branch if true
	InstrBRF    // branch if false
	InstrCConst // push constant
	InstrIConst
	InstrFConst
	InstrSConst

	InstrGLoad  // global load
	InstrGStore // global store
	InstrFLoad  // filed load
	InstrFStore // field store
	InstrMove   // reg to reg move
	InstrPrint
	InstrStruct // push struct on stack
	InstrNull   // push null on stack

	InstrHalt
)

// 基于栈的指令集
var Instructions = []*asm.Instruction{
	nil,
	{"iadd", []int32{REG, REG, REG}},
	{"isub", []int32{REG, REG, REG}},
	{"imul", []int32{REG, REG, REG}},
	{"ilt", []int32{REG, REG, REG}},
	{"ieq", []int32{REG, REG, REG}},
	{"fadd", []int32{REG, REG, REG}},
	{"fsub", []int32{REG, REG, REG}},
	{"fmul", []int32{REG, REG, REG}},
	{"flt", []int32{REG, REG, REG}},
	{"feq", []int32{REG, REG, REG}},
	{"itof", []int32{REG, REG}},
	{"call", []int32{FUNC, REG}},
	{"ret", []int32{}},
	{"br", []int32{INT}},
	{"brt", []int32{REG, INT}},
	{"brf", []int32{REG, INT}},
	{"cconst", []int32{REG, INT}},
	{"iconst", []int32{REG, INT}},
	{"fconst", []int32{REG, POLL}},
	{"sconst", []int32{REG, POLL}},
	{"gload", []int32{REG, INT}},
	{"gstore", []int32{REG, INT}},
	{"fload", []int32{REG, REG, INT}}, // 地址，寄存器，field
	{"fstore", []int32{REG, REG, INT}},
	{"move", []int32{REG, REG}},
	{"print", []int32{REG}},
	{"struct", []int32{REG, INT}},
	{"null", []int32{REG}},
	{"halt", []int32{}},
}
