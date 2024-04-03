package classic

import (
	"ads-cw/pkg/display"
	"ads-cw/pkg/sudoku_board"
)

func Play() {
	size := config()
	run(size)
}

func config() (size int) {
	return 9
}

func run(size int) {
	board := sudoku_board.GenerateBoard(size)
	var x, y int
	for {
		x, y = display.Display(board, x, y)
	}
}
