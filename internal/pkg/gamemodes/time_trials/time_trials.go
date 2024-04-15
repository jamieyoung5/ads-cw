package time_trials

import (
	"ads-cw/internal/pkg/components/sudoku_board"
	"ads-cw/internal/pkg/components/timer"
	"ads-cw/pkg/display"
	"time"
)

func Play() *display.State {
	duration, size := config()
	return run(size, duration)
}

func config() (duration time.Duration, size int) {
	return 10 * time.Second, 9
}

func run(size int, duration time.Duration) *display.State {
	board := sudoku_board.CreateNearlyComplete9x9TestBoard()
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
