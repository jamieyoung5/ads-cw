package display

type ComponentNode struct {
	Pointer   *Pointer
	Component Component
}

type Component interface {
	GetDimensions() (height int, width int)
	Print(pointer *Pointer)
	Serialize(pointer *Pointer) string
	Select(pointer *Pointer, keyCode []byte) (state *State, exit bool)
}
