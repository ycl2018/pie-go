package asm

import "github.com/ycl2018/pie-go/interpreter/asm/gen"

// 指令集定义

const (
	REG  = gen.AssemblerParserREG
	FUNC = gen.AssemblerParserFUNC
	INT  = gen.AssemblerParserINT
	POLL = 1000
)

type Instruction struct {
	Name       string
	OpRandType []int32 // 最长3个操作数
}
