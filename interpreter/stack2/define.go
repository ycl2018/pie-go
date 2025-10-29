package stack

import "github.com/ycl2018/pie-go/interpreter/asm"

const (
	FUNC = asm.FUNC
	INT  = asm.INT
	POLL = asm.POLL
)

const (
	InstrAdd = iota + 1
	InstrSub
	InstrMul
	InstrDiv
	InstrLT
	InstrEQ

	InstrCall
	InstrReturn
	InstrBR     // branch
	InstrBRT    // branch if true
	InstrBRF    // branch if false
	InstrCConst // push constant
	InstrIConst
	InstrFConst
	InstrSConst

	InstrLoad
	InstrGLoad  // global load
	InstrFLoad  // filed load
	InstrStore  // local store
	InstrGStore // global store
	InstrFStore // field store
	InstrPrint
	InstrStruct // push struct on stack
	InstrNull   // push null on stack
	InstrPop    // pop stack

	InstrHalt
)

// 基于栈的指令集
var Instructions = []*asm.Instruction{
	nil,
	{"add", []int32{}},
	{"sub", []int32{}},
	{"mul", []int32{}},
	{"div", []int32{}},
	{"lt", []int32{}},
	{"eq", []int32{}},
	{"call", []int32{FUNC}},
	{"ret", []int32{}},
	{"br", []int32{INT}},
	{"brt", []int32{INT}},
	{"brf", []int32{INT}},
	{"cconst", []int32{INT}},
	{"iconst", []int32{INT}},
	{"fconst", []int32{POLL}},
	{"sconst", []int32{POLL}},
	{"load", []int32{INT}},
	{"gload", []int32{INT}},
	{"fload", []int32{INT}},
	{"store", []int32{INT}},
	{"gstore", []int32{INT}},
	{"fstore", []int32{INT}},
	{"print", []int32{}},
	{"struct", []int32{INT}},
	{"null", []int32{}},
	{"pop", []int32{}},
	{"halt", []int32{}},
}
