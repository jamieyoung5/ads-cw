package menu

import (
	"ads-cw/internal/pkg/classic"
	"ads-cw/internal/pkg/game_history"
	"ads-cw/internal/pkg/time_trials"
	"ads-cw/internal/pkg/two_player"
)

var Content = Menu{
	{
		Name:    "Classic Sudoku",
		Summary: "classic sudoku! fill in a board with numbers without repeating any in rows, columns, or regions",
		Runner:  classic.Play,
	},
	{
		Name:    "Time Trials",
		Summary: "play sudoku against a timer!",
		Runner:  time_trials.Play,
	},
	{
		Name:    "Two Player Mode",
		Summary: "play against friends to see who can solve their board first",
		Runner:  two_player.Play,
	},
	{
		Name:    "Game History",
		Summary: "view and replay previous games",
		Runner:  game_history.Open,
	},
}
