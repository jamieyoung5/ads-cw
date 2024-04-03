package menu

import (
	"fmt"
)

type GamemodeRunner func()

type Gamemode struct {
	Name    string
	Summary string
	Runner  GamemodeRunner
}

type Menu [4]*Gamemode

func (m Menu) Print(pointerX int, pointerY int) {
	for index, gamemode := range m {
		if index == pointerY {
			fmt.Printf("> %d: %s <\n%s\n\n", index+1, gamemode.Name, gamemode.Summary)
		} else {
			fmt.Printf("%d: %s\n%s\n\n", index+1, gamemode.Name, gamemode.Summary)
		}
	}
}

func (m Menu) GetInstructions() string {
	return "Use arrow keys to navigate the menu and press ENTER to select an item."
}

func (m Menu) GetDimensions() (height int, width int) {
	return 4, 1
}

func (m Menu) Select(pointerX int, pointerY int, keyCode byte) {
	if keyCode == 10 {
		m[pointerY].Runner()
	}
}
