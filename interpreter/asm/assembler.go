package asm

import (
	"fmt"
	"os"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/ycl2018/pie-go/interpreter/asm/gen"
)

// ByteCodeAssembler 将字符串汇编文本编译为二进制代码
type ByteCodeAssembler struct {
	Code         []byte                  // 字节码
	Instructions map[string]int32        // 指令映射
	Labels       map[string]*LabelSymbol // 跳转
	ConstPool    []any                   // 常量池，包含浮点数、字符串等
	FuncConsPool []*FunctionSymbol       // 函数常量池
	DataSize     int32                   // 数据大小
	MainFunc     *FunctionSymbol         // Main函数
	IP           int32
}

func NewByteCodeAssembler(instructions []*Instruction) *ByteCodeAssembler {
	bca := &ByteCodeAssembler{
		Instructions: make(map[string]int32),
		Labels:       make(map[string]*LabelSymbol),
	}
	for i, ins := range instructions {
		if ins == nil {
			continue
		}
		bca.Instructions[ins.Name] = int32(i)
	}
	return bca
}

func (b *ByteCodeAssembler) Gen(tt ...antlr.Token) {
	instr := tt[0]
	instrName := instr.GetText()
	opCode, ok := b.Instructions[instrName]
	if !ok {
		fmt.Fprintf(os.Stderr, "line %d: unknown instruction:%s", instr.GetLine(), instrName)
		return
	}
	b.appendByte(byte(opCode & 0xFF)) // 指令占一个字节
	if len(tt) > 1 {
		for i := 1; i < len(tt); i++ {
			opRand := tt[i]
			text := opRand.GetText()
			var v int32 // 操作数都是32位
			switch opRand.GetTokenType() {
			case gen.AssemblerParserINT, gen.AssemblerParserCHAR:
				val, _ := strconv.ParseInt(text, 10, 32)
				v = int32(val)
			case gen.AssemblerParserFLOAT:
				val, _ := strconv.ParseFloat(text, 32)
				v = b.getConstPoolIndex(float32(val))
			case gen.AssemblerParserSTRING:
				v = b.getConstPoolIndex(text)
			case gen.AssemblerParserID:
				v = b.getLabelAddress(text)
			case gen.AssemblerParserFUNC:
				text = text[:len(text)-2]
				v = b.getFuncIndex(text)
			case gen.AssemblerParserREG:
				v = b.getRegisterNumber(text)
			default:
				fmt.Fprintf(os.Stderr, "line:%d unknown oprand:%s", opRand.GetLine(), text)
				return
			}
			b.writeInt32(v)
		}
	}
}

func (b *ByteCodeAssembler) appendByte(bt byte) {
	b.Code = append(b.Code, bt)
	b.IP++
}

func (b *ByteCodeAssembler) CheckForUnresolvedReferences() {
	for name, symbol := range b.Labels {
		if !symbol.IsDefined {
			fmt.Fprintf(os.Stderr, "unresolved referance: %s", name)
		}
	}
	for _, symbol := range b.FuncConsPool {
		if !symbol.IsDefined {
			fmt.Fprintf(os.Stderr, "unresolved func: %s", symbol.Name)
		}
	}
}

func (b *ByteCodeAssembler) DefineFunction(id antlr.Token, nargs, nlocals int) {
	funcName := id.GetText()
	for _, fs := range b.FuncConsPool {
		if fs.Name == funcName {
			if fs.IsDefined {
				fmt.Fprintf(os.Stderr, "duplicated define func:%s", funcName)
				return
			} else {
				fs.IsDefined = true
				fs.NArgs = int32(nargs)
				fs.NLocals = int32(nlocals)
				return
			}
		}
	}
	f := &FunctionSymbol{
		Name:      funcName,
		NArgs:     int32(nargs),
		NLocals:   int32(nlocals),
		Address:   b.IP,
		IsDefined: true,
	}
	if funcName == "main" {
		b.MainFunc = f
	}
	b.FuncConsPool = append(b.FuncConsPool, f)
}

func (b *ByteCodeAssembler) DefineDataSize(n int) {
	b.DataSize = int32(n)
}

func (b *ByteCodeAssembler) DefineLabel(id antlr.Token) {
	if l := b.Labels[id.GetText()]; l == nil {
		b.Labels[id.GetText()] = &LabelSymbol{
			Name:      id.GetText(),
			Address:   b.IP,
			IsDefined: true,
		}
		return
	} else {
		if l.IsDefined {
			fmt.Fprintf(os.Stderr, "duplicated define label:%s", l.Name)
			return
		}
		l.IsDefined = true
		l.Address = b.IP
		// 更新前向引用
		if l.IsForwardRef {
			for _, ref := range l.ForwardRefs {
				for i := 0; i < 4; i++ {
					b.Code[int(ref)+i] = (byte)((l.Address >> (8 * (3 - i))) & 0xFF)
				}
			}
		}
	}
}

/** Write value at index into a byte array highest to lowest byte,
 *  left to right.
 */
func (b *ByteCodeAssembler) writeInt32(v int32) {
	b.appendByte(byte((v >> (8 * 3)) & 0xFF))
	b.appendByte(byte((v >> (8 * 2)) & 0xFF))
	b.appendByte(byte((v >> (8 * 1)) & 0xFF))
	b.appendByte(byte(v & 0xFF))
}

func (b *ByteCodeAssembler) getConstPoolIndex(val any) int32 {
	for i, a := range b.ConstPool {
		if a == val {
			return int32(i)
		}
	}
	b.ConstPool = append(b.ConstPool, val)
	return int32(len(b.ConstPool) - 1)
}

func (b *ByteCodeAssembler) getLabelAddress(text string) int32 {
	v, ok := b.Labels[text]
	if !ok {
		// 前向引用的标签
		label := &LabelSymbol{
			Name:         text,
			Address:      0,
			IsDefined:    false,
			IsForwardRef: true,
			ForwardRefs:  []int32{b.IP}, // 记录当前引用该标签的IP，后续以便写入真实地址
		}
		b.Labels[text] = label
		return -1
	}
	if !v.IsDefined {
		v.ForwardRefs = append(v.ForwardRefs, b.IP)
		return -1
	} else {
		return v.Address
	}
}

func (b *ByteCodeAssembler) getFuncIndex(funcName string) int32 {
	for i, fs := range b.FuncConsPool {
		if fs.Name == funcName {
			return int32(i)
		}
	}
	f := &FunctionSymbol{
		Name: funcName,
	}
	b.FuncConsPool = append(b.FuncConsPool, f)
	return int32(len(b.FuncConsPool) - 1)
}

func (b *ByteCodeAssembler) getRegisterNumber(text string) int32 {
	i, _ := strconv.ParseInt(text[1:], 10, 32)
	return int32(i)
}
