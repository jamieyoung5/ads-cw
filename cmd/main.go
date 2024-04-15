package main

import (
	"ads-cw/pkg/display"
	"ads-cw/pkg/game"
)

func main() {
	oldState, err := display.TerminalRawMode()
	if err != nil {
		panic(err)
	}
	defer display.RestoreTerminal(oldState)

	sudoku := game.NewSudoku()
	sudoku.Play()
}
