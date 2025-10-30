package stack2

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/ycl2018/pie-go/interpreter/asm"
	"github.com/ycl2018/pie-go/interpreter/asm/gen"
)

type Option func(interpreter *Interpreter)

func WithEnableTrace() Option {
	return func(i *Interpreter) {
		i.enableTrace = true
	}
}

func WithEnableDump() Option {
	return func(i *Interpreter) {
		i.dump = true
	}
}
func WithDisassemble() Option {
	return func(i *Interpreter) {
		i.disassemble = true
	}
}

type Interpreter struct {
	disAssembler  *asm.DisAssembler
	IP            int32  // 指令地址
	Code          []byte // 代码
	ConstPool     []any
	FuncConstPool []*asm.FunctionSymbol
	MainFunc      *asm.FunctionSymbol // Main函数入口地址

	// 函数调用栈
	Calls []*StackFrame
	FP    int32 // 栈桢计数器

	Operands []any // 操作数栈，全局复用
	SP       int32 // 操作数栈计数器

	// 虚拟全局内存
	Globals  []any
	DataSize int32

	// ops
	enableTrace bool
	disassemble bool
	dump        bool
}

const DefaultOperandStackSize = 100

type StackFrame struct {
	ReturnAddr int32 // 返回值
	FuncSymbol *asm.FunctionSymbol
	Locals     []any // 参数和本地变量
}

func NewStackFrame(f *asm.FunctionSymbol, returnAddr int32) *StackFrame {
	return &StackFrame{
		ReturnAddr: returnAddr,
		FuncSymbol: f,
		Locals:     make([]any, f.NLocals+f.NArgs),
	}
}

func NewInterpreter(input antlr.CharStream, ops ...Option) *Interpreter {
	// 编译
	assembler := asm.NewByteCodeAssembler(Instructions)
	lexer := gen.NewAssemblerLexer(input)
	parser := gen.NewAssemblerParser(antlr.NewCommonTokenStream(lexer, 0))
	parser.AsmGenerator = assembler
	parser.Program()

	mainFunc := assembler.MainFunc

	if mainFunc == nil {
		mainFunc = &asm.FunctionSymbol{
			Name:      "main",
			IsDefined: true,
		}
	}
	i := &Interpreter{
		IP:            -1,
		Code:          assembler.Code,
		ConstPool:     assembler.ConstPool,
		FuncConstPool: assembler.FuncConsPool,
		MainFunc:      mainFunc,
		Calls:         nil,
		FP:            -1,
		Operands:      make([]any, DefaultOperandStackSize),
		SP:            -1,
		Globals:       make([]any, assembler.DataSize),
		DataSize:      assembler.DataSize,
		disAssembler:  asm.NewDisAssembler(Instructions, assembler.Code, assembler.ConstPool, assembler.FuncConsPool),
	}
	for _, op := range ops {
		op(i)
	}
	return i
}

func (i *Interpreter) Run() {
	sf := NewStackFrame(i.MainFunc, i.IP)
	i.Calls = append(i.Calls, sf)
	i.FP++

	i.IP = i.MainFunc.Address
	i.cpu()

	if i.disassemble {
		i.Disassemble()
	}
	if i.dump {
		i.Dump()
	}
}

func (i *Interpreter) PopOpStack() any {
	v := i.Operands[i.SP]
	i.SP--
	return v
}

func (i *Interpreter) PushOpStack(v any) {
	i.SP++
	i.Operands[i.SP] = v
}

func (i *Interpreter) NextOperand() int32 {
	val := asm.GetInt(int(i.IP), i.Code)
	i.IP += 4
	return val
}

func (i *Interpreter) cpu() {
	// 取指令，并执行
	instr := i.Code[i.IP]
	for i.IP < int32(len(i.Code)) && instr != InstrHalt {
		if i.enableTrace {
			i.trace()
		}
		i.IP++ // next instruction or first operand
		switch instr {
		case InstrAdd:
			// support auto type convert
			// int32 + float32 = int32
			// int32 + string = string
			// float32 + string = string
			op2, op1 := i.PopOpStack(), i.PopOpStack()
			op1Type, op2Type := reflect.TypeOf(op1), reflect.TypeOf(op2)
			if op1Type.Kind() == reflect.String || op2Type.Kind() == reflect.String {
				i.PushOpStack(toString(op1) + toString(op2))
			} else if op1Type.Kind() == reflect.Int32 && op2Type.Kind() == reflect.Int32 {
				i.PushOpStack(toInt32(op2) + toInt32(op1))
			} else  {
				i.PushOpStack(toFloat32(op2) + toFloat32(op1))
			}
		case InstrSub, InstrMul, InstrDiv, InstrLT, InstrEQ, InstrLEQ,InstrNEQ,InstrGEQ, InstrGT:
			// 弹出两个操作数，相加，push
			op2, op1 := i.PopOpStack(), i.PopOpStack()
			t2, t1 := reflect.TypeOf(op2), reflect.TypeOf(op1)
			if t2.Kind() != t1.Kind() {
				panic(fmt.Sprintf("type not match:%v %v", t1, t2))
			}
			switch instr {
			case InstrSub:
				if t1.Kind() == reflect.Int32 {
					i.PushOpStack(toInt32(op1) - toInt32(op2))
				} else {
					i.PushOpStack(toFloat32(op1) - toFloat32(op2))
				}
			case InstrMul:
				if t1.Kind() == reflect.Int32 {
					i.PushOpStack(toInt32(op1) * toInt32(op2))
				} else {
					i.PushOpStack(toFloat32(op1) * toFloat32(op2))
				}
			case InstrDiv:
				if t1.Kind() == reflect.Int32 {
					i.PushOpStack(toInt32(op1) / toInt32(op2))
				} else {
					i.PushOpStack(toFloat32(op1) / toFloat32(op2))
				}
			case InstrLT:
				if t1.Kind() == reflect.Int32 {
					i.PushOpStack(toInt32(op1) < toInt32(op2))
				} else {
					i.PushOpStack(toFloat32(op1) < toFloat32(op2))
				}
			case InstrEQ:
				if t1.Kind() == reflect.Int32 {
					i.PushOpStack(toInt32(op1) == toInt32(op2))
				} else {
					i.PushOpStack(toFloat32(op1) == toFloat32(op2))
				}
			case InstrLEQ:
				if t1.Kind() == reflect.Int32 {
					i.PushOpStack(toInt32(op1) <= toInt32(op2))
				} else {
					i.PushOpStack(toFloat32(op1) <= toFloat32(op2))
				}
			case InstrNEQ:
				if t1.Kind() == reflect.Int32 {
					i.PushOpStack(toInt32(op1) != toInt32(op2))
				} else {
					i.PushOpStack(toFloat32(op1) != toFloat32(op2))
				}
			case InstrGEQ:
				if t1.Kind() == reflect.Int32 {
					i.PushOpStack(toInt32(op1) >= toInt32(op2))
				} else {
					i.PushOpStack(toFloat32(op1) >= toFloat32(op2))
				}
			case InstrGT:
				if t1.Kind() == reflect.Int32 {
					i.PushOpStack(toInt32(op1) > toInt32(op2))
				} else {
					i.PushOpStack(toFloat32(op1) > toFloat32(op2))
				}
			}

		case InstrCall:
			// 函数调用
			funcIndex := i.NextOperand()
			fs := i.FuncConstPool[funcIndex]
			funcStack := NewStackFrame(fs, i.IP)
			i.FP++
			i.Calls = append(i.Calls, funcStack)
			// 拷贝操作数到参数中
			// move args from operand stack to top frame on call stack
			for a := fs.NArgs - 1; a >= 0; a-- {
				funcStack.Locals[a] = i.PopOpStack()
			}
			i.IP = fs.Address
		case InstrReturn:
			curFs := i.Calls[i.FP]
			i.FP--
			i.IP = curFs.ReturnAddr
		case InstrBR:
			toAddr := i.NextOperand()
			i.IP = toAddr
		case InstrBRT:
			toAddr := i.NextOperand()
			if i.PopOpStack().(bool) {
				i.IP = toAddr
			}
		case InstrBRF:
			toAddr := i.NextOperand()
			if !i.PopOpStack().(bool) {
				i.IP = toAddr
			}
		case InstrCConst, InstrIConst:
			val := i.NextOperand()
			i.PushOpStack(val)
		case InstrFConst, InstrSConst:
			poolIndex := i.NextOperand()
			fConst := i.ConstPool[poolIndex]
			i.PushOpStack(fConst)
		case InstrLoad:
			argIndex := i.NextOperand()
			curStack := i.Calls[i.FP]
			i.PushOpStack(curStack.Locals[argIndex])
		case InstrGLoad:
			argIndex := i.NextOperand()
			gVal := i.Globals[argIndex]
			i.PushOpStack(gVal)
		case InstrFLoad:
			// 字段加载,字段地址在操作数栈中，字段index为下个操作数
			s := i.PopOpStack().(*StructSpace)
			index := i.NextOperand()
			i.PushOpStack(s.Field(int(index)))
		case InstrStore:
			addr := i.NextOperand()
			i.Calls[i.FP].Locals[addr] = i.PopOpStack()
		case InstrGStore:
			addr := i.NextOperand()
			i.Globals[addr] = i.PopOpStack()
		case InstrFStore:
			s := i.PopOpStack().(*StructSpace)
			index := i.NextOperand()
			s.Set(int(index), i.PopOpStack())
		case InstrPrint:
			fmt.Println(i.PopOpStack())
		case InstrStruct:
			// push struct
			s := NewStructSpace(int(i.NextOperand()))
			i.PushOpStack(s)
		case InstrNull:
			i.PushOpStack(nil)
		case InstrPop:
			i.PopOpStack()
		case InstrHalt:
			return
		default:
			panic(fmt.Sprintf("unknown opcode:%d", instr))
		}
		instr = i.Code[i.IP]
	}
}

func toInt32(x any) (int32) {
	switch x.(type) {
	case int32:
		return x.(int32)
	case float32:
		return int32(x.(float32))
	default:
		return 0
	}
}

func toString(x any) string {
	switch x.(type) {
	case string:
		return x.(string)
	case int32:
		return fmt.Sprintf("%d", x.(int32))
	case float32:
		return fmt.Sprintf("%f", x.(float32))
	default:
		return ""
	}
}

func toFloat32(x any) (float32) {
	switch x.(type) {
	case float32:
		return x.(float32)
	case int32:
		return float32(x.(int32))
	default:
		return 0
	}
}

func (i *Interpreter) trace() {
	// asm code
	_, str, err := i.disAssembler.DisAssembleInstruction(int(i.IP))
	if err != nil {
		fmt.Printf("disAssemble instruction err:%v on ip:%d", err, i.IP)
		return
	}
	fmt.Println(str)
	// operand stack
	fmt.Printf("\tstack=[")
	for j := int32(0); j <= i.SP; j++ {
		fmt.Printf(" %v", i.Operands[j])
	}
	fmt.Print(" ]")
	// call stack
	if i.FP >= 0 {
		fmt.Printf(", calls=[")
		for j := int32(0); j <= i.FP; j++ {
			fmt.Printf(" " + i.Calls[j].FuncSymbol.Name)
		}
		fmt.Print(" ]")
	}
	fmt.Println()
}

func (i *Interpreter) Disassemble() {
	str, err := i.disAssembler.DisAssemble()
	if err != nil {
		fmt.Printf("disAssemble err:%v\n", err)
		return
	}
	fmt.Println(str)
}

func (i *Interpreter) Dump() {
	if len(i.ConstPool) > 0 {
		i.dumpConstPool()
	}
	if len(i.Globals) > 0 {
		i.dumpDataMemory()
	}
	i.dumpCodeMemory()
}

func (i *Interpreter) dumpConstPool() {
	fmt.Println("Constant Pool:")
	for j, a := range append(i.ConstPool, i.FuncConstPool) {
		switch a.(type) {
		case string:
			fmt.Printf("%04d: \"%s\"\n", j, a)
		default:
			fmt.Printf("%04d: %s\n", j, a)
		}
	}
	fmt.Println()
}

func (i *Interpreter) dumpCodeMemory() {
	fmt.Println("Code memory:")
	for j := 0; j < len(i.Code); j++ {
		if j%8 == 0 && j != 0 {
			fmt.Println()
		}
		if j%8 == 0 {
			fmt.Printf("%04d:", j)
		}
		fmt.Printf(" %3d", (int)(i.Code[j]))
	}
	fmt.Println()
}

func (i *Interpreter) dumpDataMemory() {
	fmt.Println("Data memory:")
	for j, a := range i.Globals {
		if a != nil {
			fmt.Printf("%04d: %v %s\n", j, a, reflect.TypeOf(a).Name())
		} else {
			fmt.Printf("%04d: null \n", j)
		}
	}
	fmt.Println()
}

type StructSpace struct {
	Fields []any
}

func (s *StructSpace) String() string {
	var sb strings.Builder
	sb.WriteString("struct{")
	for _, field := range s.Fields {
		if field == nil {
			sb.WriteString(" null")
		} else {
			sb.WriteString(fmt.Sprintf(" %v", field))
		}
	}
	sb.WriteString(" }")
	return sb.String()
}

func NewStructSpace(fieldsNum int) *StructSpace {
	return &StructSpace{Fields: make([]any, fieldsNum)}
}

func (s *StructSpace) Field(index int) any {
	return s.Fields[index]
}

func (s *StructSpace) Set(index int, val any) {
	s.Fields[index] = val
}
