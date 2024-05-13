package asm

import (
	"fmt"
	"strconv"
	"strings"
)

// 反汇编： 将编译的二进制文件反编译为字符串代码，只有指令

type DisAssembler struct {
	Code          []byte
	ConstPool     []any
	FuncConstPool []*FunctionSymbol
	Instructions  []*Instruction
	sb            strings.Builder
}

func NewDisAssembler(instrs []*Instruction, code []byte, constPool []any, funcConstPool []*FunctionSymbol) *DisAssembler {
	return &DisAssembler{
		Code:          code,
		ConstPool:     constPool,
		FuncConstPool: funcConstPool,
		Instructions:  instrs,
	}
}

func (d *DisAssembler) DisAssemble() (string, error) {
	d.sb.WriteString("Disassembly:")

	var i int
	for i < len(d.Code) {
		d.sb.WriteString("\n")
		next, str, err := d.DisAssembleInstruction(i)
		if err != nil {
			return "", err
		}
		d.sb.WriteString(str)
		i = next
	}
	return d.sb.String(), nil
}

func (d *DisAssembler) DisAssembleInstruction(i int) (int, string, error) {
	var sb strings.Builder
	instr := d.Instructions[int(d.Code[i])]
	if instr == nil {
		return 0, "", fmt.Errorf("unknown instruction:%d in code[%d]", d.Code[i], i)
	}
	instrName := instr.Name
	sb.WriteString(fmt.Sprintf("%04d:\t%-11s", i, instrName))
	i++
	if len(instr.OpRandType) == 0 {
		sb.WriteString("    ")
		return i, sb.String(), nil
	}
	var opRand []string
	for _, t := range instr.OpRandType {
		opRandNum := GetInt(i, d.Code)
		i += 4
		switch t {
		case INT:
			opStr := strconv.Itoa(int(opRandNum))
			opRand = append(opRand, opStr)
		case FUNC:
			funcSym := d.FuncConstPool[opRandNum]
			opRand = append(opRand, fmt.Sprintf("#%d:%s()@%d", opRandNum, funcSym.Name, funcSym.Address))
		case POLL:
			opRand = append(opRand, d.showConstPoolOperand(opRandNum))
		case REG:
			opRand = append(opRand, fmt.Sprintf("r%d", opRandNum))
		}
	}

	for i, s := range opRand {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(s)
	}
	return i, sb.String(), nil
}

func GetInt(ip int, code []byte) int32 {
	var num = int32(0)
	for i := 0; i < 4; i++ {
		b := int32(code[ip+i] << (8 * (3 - i)))
		num = num | b
	}
	return num
}

func (d *DisAssembler) showConstPoolOperand(index int32) string {
	var sb strings.Builder
	sb.WriteString("#")
	sb.WriteString(strconv.Itoa(int(index)) + ":")
	cst := d.ConstPool[index]
	switch cst.(type) {
	case float32:
		sb.WriteString(strconv.FormatFloat(float64(cst.(float32)), 'g', 10, 32))
	case string:
		sb.WriteString("\"" + cst.(string) + "\"")
	}
	return sb.String()
}
