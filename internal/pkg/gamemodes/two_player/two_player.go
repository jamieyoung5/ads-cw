package two_player

import (
	"ads-cw/internal/pkg/components/menu"
	"ads-cw/internal/pkg/components/sudoku_board"
	"ads-cw/pkg/controls"
	"ads-cw/pkg/display"
)

func Easy() *display.State {
	board1, _ := sudoku_board.NewBoard(sudoku_board.GenerateBoard(4))
	board2, _ := sudoku_board.NewBoard(sudoku_board.GenerateBoard(4))
	return run(board1, board2)
}

func Normal() *display.State {
	board1, _ := sudoku_board.NewBoard(sudoku_board.GenerateBoard(9))
	board2, _ := sudoku_board.NewBoard(sudoku_board.GenerateBoard(9))
	return run(board1, board2)
}

func Hard() *display.State {
	board1, _ := sudoku_board.NewBoard(sudoku_board.GenerateBoard(16))
	board2, _ := sudoku_board.NewBoard(sudoku_board.GenerateBoard(16))
	return run(board1, board2)
}

func DifficultySelect() *display.State {
	gridMap := [][]*display.ComponentNode{
		{
			&display.ComponentNode{Component: menu.Menu{
				{
					Name:    "Easy",
					Summary: "A small 4x4 board",
					Runner:  Easy,
				},
				{
					Name:    "Normal",
					Summary: "Normal sized 9x9 board",
					Runner:  Normal,
				},
				{
					Name:    "Hard",
					Summary: "Large 16x16 board",
					Runner:  Hard,
				},
			}},
		},
	}

	pointer := display.NewPointer(0, 0, controls.MenuControls, 0, 0)
	gridMap[0][0].Pointer = pointer

	return &display.State{
		Components: gridMap,
		Pointers:   []*display.Pointer{pointer},
		Persist:    false,
	}
}

func run(board1 *sudoku_board.Board, board2 *sudoku_board.Board) *display.State {
	gridMap := [][]*display.ComponentNode{
		{
			&display.ComponentNode{Component: board1},
			&display.ComponentNode{Component: board2},
		},
	}

	pointer1 := display.NewPointer(0, 0, controls.SudokuControls, 0, 0)
	gridMap[0][0].Pointer = pointer1

	pointer2 := display.NewPointer(0, 0, controls.SudokuControlsAlternate, 1, 0)
	gridMap[0][1].Pointer = pointer2

	return &display.State{
		Components: gridMap,
		Pointers:   []*display.Pointer{pointer1, pointer2},
		Persist:    false,
	}
}
