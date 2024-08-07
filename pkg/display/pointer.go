package display

import controls2 "ads-cw/pkg/controls"

const (
	reset        = "\033[0m"
	whiteBGBlack = "\033[47m\033[30m"
)

type Pointer struct {
	X                  int
	Y                  int
	controls           controls2.Controls
	GridX              int
	GridY              int
	SelectedTileColour string
}

func NewPointer(x int, y int, controls controls2.Controls, gridX int, gridY int) *Pointer {
	return &Pointer{
		X:                  x,
		Y:                  y,
		controls:           controls,
		GridX:              gridX,
		GridY:              gridY,
		SelectedTileColour: whiteBGBlack,
	}
}

func (p *Pointer) Up() {
	p.Y--
}

func (p *Pointer) Down() {
	p.Y++
}

func (p *Pointer) Right() {
	p.X++
}

func (p *Pointer) Left() {
	p.X--
}
