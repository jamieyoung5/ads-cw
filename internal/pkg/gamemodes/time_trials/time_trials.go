package time_trials

import (
	"ads-cw/internal/pkg/components/menu"
	"ads-cw/internal/pkg/components/sudoku_board"
	"ads-cw/internal/pkg/components/timer"
	"ads-cw/pkg/display"
	"time"
)

func Easy() *display.State {
	boardContent := sudoku_board.GenerateBoard(4)
	board, _ := sudoku_board.NewBoard(boardContent)
	return run(board, 15*time.Minute)
}

func Normal() *display.State {
	boardContent := sudoku_board.GenerateBoard(9)
	board, _ := sudoku_board.NewBoard(boardContent)
	return run(board, 15*time.Minute)
}

func Hard() *display.State {
	boardContent := sudoku_board.GenerateBoard(16)
	board, _ := sudoku_board.NewBoard(boardContent)
	return run(board, 20*time.Minute)
}

func DifficultySelect() *display.State {
	gridMap := [][]*display.ComponentNode{
		{
			&display.ComponentNode{Component: menu.Menu{
				{
					Name:    "Easy",
					Summary: "A small 4x4 board with a 15 minute timer",
					Runner:  Easy,
				},
				{
					Name:    "Normal",
					Summary: "Normal sized 9x9 board with a 15 minute timer",
					Runner:  Normal,
				},
				{
					Name:    "Hard",
					Summary: "Large 16x16 board with a 20 minute timer",
					Runner:  Hard,
				},
			}},
		},
	}

	pointer := display.NewPointer(0, 0, display.MenuControls, 0, 0)
	gridMap[0][0].Pointer = pointer

	return &display.State{
		Components: gridMap,
		Pointers:   []*display.Pointer{pointer},
		Persist:    false,
	}
}

func run(board *sudoku_board.Board, duration time.Duration) *display.State {
	timerComponent := timer.NewTimer(duration)
	gridMap := [][]*display.ComponentNode{
		{
			&display.ComponentNode{Component: board},
			&display.ComponentNode{Component: timerComponent},
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
