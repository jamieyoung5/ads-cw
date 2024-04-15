package game

import (
	"ads-cw/internal/pkg/components/menu"
	"ads-cw/pkg/display"
)

type Sudoku struct {
}

func NewSudoku() *Sudoku {
	return &Sudoku{}
}

func (s *Sudoku) Play() {
	oldState, err := display.TerminalRawMode()
	if err != nil {
		panic(err)
	}
	defer display.RestoreTerminal(oldState)

	for {
		grid := [][]*display.ComponentNode{
			{
				&display.ComponentNode{Component: menu.MainMenu},
			},
		}

		pointer := display.NewPointer(0, 0, display.MenuControls, 0, 0)
		grid[0][0].Pointer = pointer

		canvas := display.NewCanvas(grid, []*display.Pointer{pointer})
		canvas.Render()
	}
}
