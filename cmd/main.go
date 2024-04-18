package main

import (
	"ads-cw/pkg/game"
	"ads-cw/pkg/term"
)

func main() {
	terminal := term.NewTerminal()
	err := terminal.EnableRawMode()
	if err != nil {
		panic(err)
	}
	defer terminal.Restore()

	sudoku := game.NewSudoku()
	sudoku.Play()
}
