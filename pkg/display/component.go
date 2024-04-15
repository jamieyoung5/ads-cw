package display

type ComponentNode struct {
	Pointer   *Pointer
	Component Component
}

type Component interface {
	GetDimensions() (height int, width int)
	Print(pointer *Pointer)
	Serialize(pointer *Pointer) string
	Select(pointer *Pointer, macro string) (state *State, exit bool)
}
