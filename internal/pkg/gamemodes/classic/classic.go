package classic

import (
	"ads-cw/internal/pkg/components/sudoku_board"
	"ads-cw/pkg/display"
)

func Play() *display.State {
	size := config()
	return run(size)
}

func config() (size int) {
	return 9
}

func run(size int) *display.State {
	board := sudoku_board.CreateNearlyComplete9x9TestBoard()
	gridMap := [][]*display.ComponentNode{
		{

			&display.ComponentNode{Component: board},
		},
	}

	pointer := display.NewPointer(0, 0, display.SudokuControls, 0, 0)
	gridMap[0][0].Pointer = pointer

	return &display.State{
		Components: gridMap,
		Pointers:   []*display.Pointer{pointer},
		Persist:    false,
	}
}
