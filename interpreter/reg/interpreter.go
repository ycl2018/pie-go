package reg

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
	return func(t *Interpreter) {
		t.enableTrace = true
	}
}

func WithEnableDump() Option {
	return func(t *Interpreter) {
		t.dump = true
	}
}
func WithDisassemble() Option {
	return func(t *Interpreter) {
		t.disassemble = true
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

	Regs []any // 操作数栈，全局复用

	// 虚拟全局内存
	Globals  []any
	DataSize int32

	// ops
	enableTrace bool
	disassemble bool
	dump        bool
}

type StackFrame struct {
	ReturnAddr int32
	FuncSymbol *asm.FunctionSymbol
	Regs       []any // 参数和本地变量
}

func NewStackFrame(f *asm.FunctionSymbol, returnAddr int32) *StackFrame {
	return &StackFrame{
		ReturnAddr: returnAddr,
		FuncSymbol: f,
		Regs:       make([]any, f.NLocals+f.NArgs+1), // plus r0
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
	t := &Interpreter{
		IP:            -1,
		Code:          assembler.Code,
		ConstPool:     assembler.ConstPool,
		FuncConstPool: assembler.FuncConsPool,
		MainFunc:      mainFunc,
		Calls:         nil,
		FP:            -1,
		Globals:       make([]any, assembler.DataSize),
		DataSize:      assembler.DataSize,
		disAssembler:  asm.NewDisAssembler(Instructions, assembler.Code, assembler.ConstPool, assembler.FuncConsPool),
	}
	for _, op := range ops {
		op(t)
	}
	return t
}

func (t *Interpreter) Run() {
	sf := NewStackFrame(t.MainFunc, t.IP)
	t.Calls = append(t.Calls, sf)
	t.FP++

	t.IP = t.MainFunc.Address
	t.cpu()

	if t.disassemble {
		t.Disassemble()
	}
	if t.dump {
		t.Dump()
	}
}

func (t *Interpreter) NextOperand() int32 {
	val := asm.GetInt(int(t.IP), t.Code)
	t.IP += 4
	return val
}

func (t *Interpreter) cpu() {
	// 取指令，并执行
	instr := t.Code[t.IP]
	for t.IP < int32(len(t.Code)) && instr != InstrHalt {
		if t.enableTrace {
			t.trace()
		}
		t.IP++ // next instruction or first operand
		r := t.Calls[t.FP].Regs
		switch instr {
		case InstrIAdd:
			j, k, i := t.NextOperand(), t.NextOperand(), t.NextOperand()
			r[i] = r[j].(int32) + r[k].(int32)
		case InstrFAdd:
			j, k, i := t.NextOperand(), t.NextOperand(), t.NextOperand()
			r[i] = r[j].(float32) + r[k].(float32)
		case InstrISub:
			j, k, i := t.NextOperand(), t.NextOperand(), t.NextOperand()
			r[i] = r[j].(int32) - r[k].(int32)
		case InstrFSub:
			j, k, i := t.NextOperand(), t.NextOperand(), t.NextOperand()
			r[i] = r[j].(float32) - r[k].(float32)
		case InstrIMul:
			j, k, i := t.NextOperand(), t.NextOperand(), t.NextOperand()
			r[i] = r[j].(int32) * r[k].(int32)
		case InstrFMul:
			j, k, i := t.NextOperand(), t.NextOperand(), t.NextOperand()
			r[i] = r[j].(float32) * r[k].(float32)
		case InstrILT:
			j, k, i := t.NextOperand(), t.NextOperand(), t.NextOperand()
			r[i] = r[j].(int32) < r[k].(int32)
		case InstrFLT:
			j, k, i := t.NextOperand(), t.NextOperand(), t.NextOperand()
			r[i] = r[j].(float32) < r[k].(float32)
		case InstrIEQ:
			j, k, i := t.NextOperand(), t.NextOperand(), t.NextOperand()
			r[i] = r[j].(int32) == r[k].(int32)
		case InstrFEQ:
			j, k, i := t.NextOperand(), t.NextOperand(), t.NextOperand()
			r[i] = r[j].(float32) == r[k].(float32)
		case InstrI2F:
			j, i := t.NextOperand(), t.NextOperand()
			r[i] = int32(r[j].(float32))
		case InstrCall:
			// 函数调用
			funcIndex, baseRegIndex := t.NextOperand(), t.NextOperand()
			fs := t.FuncConstPool[funcIndex]
			funcStack := NewStackFrame(fs, t.IP)
			t.FP++
			t.Calls = append(t.Calls, funcStack)
			// 拷贝操作数到参数中
			// move args from operand stack to top frame on call stack
			for a := int32(0); a < fs.NArgs; a++ {
				funcStack.Regs[a+1] = r[baseRegIndex+a]
			}
			t.IP = fs.Address
		case InstrReturn:
			curFs := t.Calls[t.FP]
			t.FP--
			t.IP = curFs.ReturnAddr
		case InstrBR:
			toAddr := t.NextOperand()
			t.IP = toAddr
		case InstrBRT:
			i, toAddr := t.NextOperand(), t.NextOperand()
			if r[i].(bool) {
				t.IP = toAddr
			}
		case InstrBRF:
			i, toAddr := t.NextOperand(), t.NextOperand()
			if !r[i].(bool) {
				t.IP = toAddr
			}
		case InstrCConst, InstrIConst:
			i, val := t.NextOperand(), t.NextOperand()
			r[i] = val
		case InstrFConst, InstrSConst:
			i, poolIndex := t.NextOperand(), t.NextOperand()
			fConst := t.ConstPool[poolIndex]
			r[i] = fConst
		case InstrGLoad:
			i, argIndex := t.NextOperand(), t.NextOperand()
			gVal := t.Globals[argIndex]
			r[i] = gVal
		case InstrFLoad:
			// 字段加载,字段地址在操作数栈中，字段index为下个操作数
			i, j, k := t.NextOperand(), t.NextOperand(), t.NextOperand()
			s := r[j].(*StructSpace)
			r[i] = (s.Field(int(k)))
		case InstrGStore:
			i, addr := t.NextOperand(), t.NextOperand()
			t.Globals[addr] = r[i]
		case InstrFStore:
			i, j, k := t.NextOperand(), t.NextOperand(), t.NextOperand()
			s := r[j].(*StructSpace)
			s.Set(int(k), r[i])
		case InstrMove:
			i, j := t.NextOperand(), t.NextOperand()
			r[j] = r[i]
		case InstrPrint:
			i := t.NextOperand()
			fmt.Println(r[i])
		case InstrStruct:
			// push struct
			i, s := t.NextOperand(), NewStructSpace(int(t.NextOperand()))
			r[i] = s
		case InstrNull:
			i := t.NextOperand()
			r[i] = nil
		case InstrHalt:
			return
		default:
			panic(fmt.Sprintf("unknown opcode:%d", instr))
		}
		instr = t.Code[t.IP]
	}
}

func (t *Interpreter) trace() {
	// asm code
	_, str, err := t.disAssembler.DisAssembleInstruction(int(t.IP))
	if err != nil {
		fmt.Printf("disAssemble instruction err:%v on ip:%d", err, t.IP)
		return
	}
	fmt.Println(str)
	// call stack registers
	r := t.Calls[t.FP].Regs
	if len(r) > 0 {
		fmt.Printf("\t%s.registers=[", t.Calls[t.FP].FuncSymbol.Name)
		for j := (0); j < len(r); j++ {
			if j == 1 {
				fmt.Print(" |")
			}
			if j == int(t.Calls[t.FP].FuncSymbol.NArgs+1) && j == 1 {
				fmt.Print("|")
			} else if j == int(t.Calls[t.FP].FuncSymbol.NArgs+1) {
				fmt.Print(" |")
			}
			fmt.Print(" ")
			if r[j] == nil {
				fmt.Print("?")
			} else {
				fmt.Printf("%v", r[j])
			}
		}
		fmt.Print(" ]")
	}

	// call stack
	if t.FP >= 0 {
		fmt.Printf(", calls=[")
		for j := int32(0); j <= t.FP; j++ {
			fmt.Printf(" " + t.Calls[j].FuncSymbol.Name)
		}
		fmt.Print(" ]")
	}
	fmt.Println()
}

func (t *Interpreter) Disassemble() {
	str, err := t.disAssembler.DisAssemble()
	if err != nil {
		fmt.Printf("disAssemble err:%v\n", err)
		return
	}
	fmt.Println(str)
}

func (t *Interpreter) Dump() {
	if len(t.ConstPool) > 0 {
		t.dumpConstPool()
	}
	if len(t.Globals) > 0 {
		t.dumpDataMemory()
	}
	t.dumpCodeMemory()
}

func (t *Interpreter) dumpConstPool() {
	fmt.Println("Constant Pool:")
	for j, a := range append(t.ConstPool, t.FuncConstPool) {
		switch a.(type) {
		case string:
			fmt.Printf("%04d: \"%s\"\n", j, a)
		default:
			fmt.Printf("%04d: %s\n", j, a)
		}
	}
	fmt.Println()
}

func (t *Interpreter) dumpCodeMemory() {
	fmt.Println("Code memory:")
	for j := 0; j < len(t.Code); j++ {
		if j%8 == 0 && j != 0 {
			fmt.Println()
		}
		if j%8 == 0 {
			fmt.Printf("%04d:", j)
		}
		fmt.Printf(" %3d", (int)(t.Code[j]))
	}
	fmt.Println()
}

func (t *Interpreter) dumpDataMemory() {
	fmt.Println("Data memory:")
	for j, a := range t.Globals {
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
