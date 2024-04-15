package two_player

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
	board1 := sudoku_board.CreateNearlyComplete9x9TestBoard()
	board2 := sudoku_board.CreateNearlyComplete9x9TestBoard()
	gridMap := [][]*display.ComponentNode{
		{
			&display.ComponentNode{Component: board1},
			&display.ComponentNode{Component: board2},
		},
	}

	pointer1 := display.NewPointer(0, 0, display.SudokuControls, 0, 0)
	gridMap[0][0].Pointer = pointer1

	pointer2 := display.NewPointer(0, 0, display.SudokuControlsAlternate, 0, 0)
	gridMap[0][1].Pointer = pointer2

	return &display.State{
		Components: gridMap,
		Pointers:   []*display.Pointer{pointer1, pointer2},
		Persist:    false,
	}
}
