package compile

type MemorySpace struct {
	Name    string
	Members map[string]any
}

func NewMemorySpace(name string) *MemorySpace {
	return &MemorySpace{
		Name:    name,
		Members: make(map[string]any),
	}
}

func (m MemorySpace) Get(name string) any {
	return m.Members[name]
}

func (m MemorySpace) Contains(name string) bool {
	_, ok := m.Members[name]
	return ok
}

type FunctionMemorySpace struct {
	MemorySpace
	FunctionSymbol *FunctionSymbol
}
