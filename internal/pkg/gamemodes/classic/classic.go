package classic

import (
	"ads-cw/internal/pkg/components/sudoku_board"
	"ads-cw/pkg/display"
)

func Play() {
	size := config()
	run(size)
}

func config() (size int) {
	return 9
}

func run(size int) {
	board := sudoku_board.CreateNearlyComplete9x9TestBoard()
	gridMap := [][]*display.ComponentNode{
		{

			&display.ComponentNode{Component: board},
		},
	}

	pointer := display.NewPointer(0, 0, display.SudokuControls, 0, 0)
	gridMap[0][0].Pointer = pointer

	canvas := display.NewCanvas(gridMap, []*display.Pointer{pointer})
	canvas.Render()
}
