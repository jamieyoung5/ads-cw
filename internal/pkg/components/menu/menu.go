package menu

import (
	"ads-cw/pkg/display"
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

func (m Menu) Print(pointer *display.Pointer) {
	fmt.Printf(m.Serialize(pointer))
}

func (m Menu) GetDimensions() (height int, width int) {
	return 4, 1
}

func (m Menu) Select(pointer *display.Pointer, keyCode []byte) (exit bool, err error) {
	if keyCode[0] == 10 {
		if len(m) > pointer.Y {

			m[pointer.Y].Runner()
			return true, nil
		} else {
			return false, errors.New("invalid menu item selected")
		}
	}

	return false, errors.New("invalid input, navigate to menu item and press enter to select")
}

func (m Menu) Serialize(pointer *display.Pointer) string {
	var builder strings.Builder

	for index, gamemode := range m {
		if pointer != nil && index == pointer.Y {
			item := fmt.Sprintf("%s %d: %s %s\n%s \n\n", pointer.SelectedTileColour, index+1, gamemode.Name, "\u001B[0m", gamemode.Summary)
			builder.WriteString(item)
		} else {
			item := fmt.Sprintf("%d: %s\n%s\n\n", index+1, gamemode.Name, gamemode.Summary)
			builder.WriteString(item)
		}
	}

	return builder.String()
}
