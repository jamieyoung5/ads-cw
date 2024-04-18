package game

import (
	"ads-cw/internal/pkg/components/menu"
	"ads-cw/internal/pkg/gamemodes/classic"
	"ads-cw/internal/pkg/gamemodes/time_trials"
	"ads-cw/internal/pkg/gamemodes/two_player"
	"ads-cw/pkg/display"
)

type Sudoku struct {
}

func NewSudoku() *Sudoku {
	return &Sudoku{}
}

func (s *Sudoku) Play() {
	oldState, err := display.TerminalRawMode()
	if err != nil {
		panic(err)
	}
	defer display.RestoreTerminal(oldState)

	for {
		grid := [][]*display.ComponentNode{
			{
				&display.ComponentNode{Component: menu.Menu{
					{
						Name:    "Classic Sudoku",
						Summary: "classic sudoku! fill in a board with numbers without repeating any in rows, columns, or regions",
						Runner:  classic.DifficultySelect,
					},
					{
						Name:    "Time Trials",
						Summary: "play sudoku against a timer!",
						Runner:  time_trials.DifficultySelect,
					},
					{
						Name:    "Two Player Mode",
						Summary: "play against friends to see who can solve their board first",
						Runner:  two_player.DifficultySelect,
					},
				}},
			},
		}

		pointer := display.NewPointer(0, 0, display.MenuControls, 0, 0)
		grid[0][0].Pointer = pointer

		canvas := display.NewCanvas(grid, []*display.Pointer{pointer}, true)
		canvas.Render()
	}
}
