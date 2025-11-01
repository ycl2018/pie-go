package vm

import (
	"fmt"

	"github.com/ycl2018/pie-go/interpreter/asm"
)

const (
	INT  = asm.INT
	POLL = asm.POLL
)

type Instr int

func (i Instr) String() string {
	switch i {
	case InstrAdd:
		return "add"
	case InstrSub:
		return "sub"
	case InstrMul:
		return "mul"
	case InstrDiv:
		return "div"
	case InstrLT:
		return "lt"
	case InstrGT:
		return "gt"
	case InstrGEQ:
		return "geq"
	case InstrLEQ:
		return "leq"
	case InstrNEQ:
		return "neq"
	case InstrEQ:
		return "eq"
	case InstrCall:
		return "call"
	case InstrReturn:
		return "ret"
	case InstrBR:
		return "br"
	case InstrBRT:
		return "brt"
	case InstrBRF:
		return "brf"
	case InstrCConst:
		return "cconst"
	case InstrIConst:
		return "iconst"
	case InstrFConst:
		return "fconst"
	case InstrSConst:
		return "sconst"
	case InstrLoad:
		return "load"
	case InstrGLoad:
		return "gload"
	case InstrFLoad:
		return "fload"
	case InstrStore:
		return "store"
	case InstrGStore:
		return "gstore"
	case InstrFStore:
		return "fstore"
	case InstrPrint:
		return "print"
	case InstrStruct:
		return "struct"
	case InstrNull:
		return "null"
	case InstrPop:
		return "pop"
	case InstrHalt:
		return "halt"
	default:
		if i < 0 {
			return fmt.Sprintf("tag#%d:", i)
		}
		panic(fmt.Sprintf("unknown instr %d", i))
	}
}

const (
	InstrAdd = iota + 1
	InstrSub
	InstrMul
	InstrDiv
	InstrLT
	InstrGT
	InstrGEQ
	InstrLEQ
	InstrNEQ
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
var Instructions = []*Instruction{
	nil,
	{"add", []int32{}},
	{"sub", []int32{}},
	{"mul", []int32{}},
	{"div", []int32{}},
	{"lt", []int32{}},
	{"gt", []int32{}},
	{"geq", []int32{}},
	{"leq", []int32{}},
	{"neq", []int32{}},
	{"eq", []int32{}},
	{"call", []int32{POLL}},
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
	{"fload", []int32{POLL}},
	{"store", []int32{INT}},
	{"gstore", []int32{INT}},
	{"fstore", []int32{POLL}},
	{"print", []int32{}},
	{"struct", []int32{INT}},
	{"null", []int32{}},
	{"pop", []int32{}},
	{"halt", []int32{}},
}
