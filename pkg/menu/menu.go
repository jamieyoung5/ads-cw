package menu

import (
	"errors"
	"fmt"
	"strings"
)

type GamemodeRunner func()

type Gamemode struct {
	Name    string
	Summary string
	Runner  GamemodeRunner
}

type Menu [4]*Gamemode

func (m Menu) Print(pointerX, pointerY int, selectedColour string) {
	fmt.Printf(m.Serialize(pointerX, pointerY, selectedColour))
}

func (m Menu) GetDimensions() (height int, width int) {
	return 4, 1
}

func (m Menu) Select(pointerX int, pointerY int, keyCode byte) (exit bool, err error) {
	if keyCode == 10 {
		if len(m) > pointerY {
			m[pointerY].Runner()
			return true, nil
		} else {
			return false, errors.New("invalid menu item selected")
		}
	}

	return false, errors.New("invalid input, navigate to menu item and press enter to select")
}

func (m Menu) Serialize(pointerX, pointerY int, selectedColour string) string {
	var builder strings.Builder

	for index, gamemode := range m {
		if index == pointerY {
			item := fmt.Sprintf("%s %d: %s %s\n%s \n\n", selectedColour, index+1, gamemode.Name, "\u001B[0m", gamemode.Summary)
			builder.WriteString(item)
		} else {
			item := fmt.Sprintf("%d: %s\n%s\n\n", index+1, gamemode.Name, gamemode.Summary)
			builder.WriteString(item)
		}
	}

	return builder.String()
}
