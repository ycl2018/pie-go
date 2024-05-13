package asm

import "fmt"

type FunctionSymbol struct {
	Name      string
	NArgs     int32
	NLocals   int32
	Address   int32
	IsDefined bool
}

func (f *FunctionSymbol) String() string {
	return fmt.Sprintf("func:<%s()@%d nargs:%d nlocals:%d>", f.Name, f.Address, f.NArgs, f.NLocals)
}

type LabelSymbol struct {
	Name string
	/** Address in code memory */
	Address int32
	/** Is this ref'd before def'd. */
	IsDefined bool
	/** Set when we see actual ID: definition */
	IsForwardRef bool
	/** List of operands in memory we need to update after seeing def */
	ForwardRefs []int32
}
