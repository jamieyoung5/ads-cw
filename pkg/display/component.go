package display

type ComponentNode struct {
	Left  *ComponentNode
	Right *ComponentNode
	Up    *ComponentNode
	Down  *ComponentNode

	Pointer   *Pointer
	Component Component
}

type Component interface {
	GetDimensions() (height int, width int)
	Print(pointerX, pointerY int, selectedColour string)
	Serialize(pointerX, pointerY int, selectedColour string) string
	Select(pointerX, pointerY int, keyCode byte) (exit bool, err error)
}
