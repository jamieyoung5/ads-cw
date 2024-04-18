package menu

import (
	"ads-cw/pkg/controls"
	"ads-cw/pkg/display"
	"fmt"
	"strings"
)

type Action func() *display.State

type Item struct {
	Name    string
	Summary string
	Runner  Action
}

type Menu []*Item

func (m Menu) Print(pointer *display.Pointer) {
	fmt.Printf(m.Serialize(pointer))
}

func (m Menu) GetDimensions() (height int, width int) {
	return 4, 1
}

func (m Menu) Select(pointer *display.Pointer, macro string) (*display.State, bool) {
	if macro == controls.Selected {
		if len(m) < pointer.Y {
			pointer.Y = len(m) - 1
		}
		return m[pointer.Y].Runner(), true

	}

	return nil, false
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
