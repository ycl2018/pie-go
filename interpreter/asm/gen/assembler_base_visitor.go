// Code generated from /Users/bytedance/go/src/github.com/ycl2018/asm/gen/Assembler.g4 by ANTLR 4.13.1. DO NOT EDIT.

package gen // Assembler
import "github.com/antlr4-go/antlr/v4"


type BaseAssemblerVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseAssemblerVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAssemblerVisitor) VisitGlobals(ctx *GlobalsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAssemblerVisitor) VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAssemblerVisitor) VisitInstr(ctx *InstrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAssemblerVisitor) VisitOperand(ctx *OperandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAssemblerVisitor) VisitLabel(ctx *LabelContext) interface{} {
	return v.VisitChildren(ctx)
}
