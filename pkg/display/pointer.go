package display

const (
	reset        = "\033[0m"
	whiteBGBlack = "\033[47m\033[30m"
)

type Pointer struct {
	x                  int
	y                  int
	controls           Controls
	selectedComponent  *ComponentNode
	selectedTileColour string
}

func NewPointer(x int, y int, controls Controls, selectedComponent *ComponentNode) *Pointer {
	return &Pointer{
		x:                  x,
		y:                  y,
		controls:           controls,
		selectedComponent:  selectedComponent,
		selectedTileColour: whiteBGBlack,
	}
}

func (p *Pointer) Up() {
	p.y--
}

func (p *Pointer) Down() {
	p.y++
}

func (p *Pointer) Right() {
	p.x++
}

func (p *Pointer) Left() {
	p.x--
}
