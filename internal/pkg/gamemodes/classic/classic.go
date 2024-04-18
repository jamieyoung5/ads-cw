package classic

import (
	"ads-cw/internal/pkg/components/menu"
	"ads-cw/internal/pkg/components/sudoku_board"
	"ads-cw/pkg/controls"
	"ads-cw/pkg/display"
)

func Easy() *display.State {
	boardContent := sudoku_board.GenerateBoard(4)
	board, _ := sudoku_board.NewBoard(boardContent)
	return play(board)
}

func Normal() *display.State {
	boardContent := sudoku_board.GenerateBoard(9)
	board, _ := sudoku_board.NewBoard(boardContent)
	return play(board)
}

func Hard() *display.State {
	boardContent := sudoku_board.GenerateBoard(16)
	board, _ := sudoku_board.NewBoard(boardContent)
	return play(board)
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

func play(board *sudoku_board.Board) *display.State {
	gridMap := [][]*display.ComponentNode{
		{

			&display.ComponentNode{Component: board},
		},
	}

	pointer := display.NewPointer(0, 0, controls.SudokuControls, 0, 0)
	gridMap[0][0].Pointer = pointer

	return &display.State{
		Components: gridMap,
		Pointers:   []*display.Pointer{pointer},
		Persist:    false,
	}
}
