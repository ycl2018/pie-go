package tree

import (
	"fmt"
	"strings"
)

type StructInstance struct {
	MemorySpace
	StructDef *StructSymbol
}

func NewStructInstance(structDef *StructSymbol) *StructInstance {
	return &StructInstance{
		MemorySpace: MemorySpace{
			Members: make(map[string]any),
			Name:    structDef.Name,
		},
		StructDef: structDef,
	}
}

func (s *StructInstance) String() string {
	var sb strings.Builder
	sb.WriteString("{")
	for i, field := range s.StructDef.Fields {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(field.GetName())
		sb.WriteString("=")
		value := s.Get(field.GetName())
		if value == nil {
			sb.WriteString("null")
		} else {
			sb.WriteString(fmt.Sprintf("%s", value))
		}
	}
	sb.WriteString("}")
	return sb.String()
}
