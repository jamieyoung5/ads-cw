package sudoku_board

import (
	"errors"
	"fmt"
)

type Board [][]Tile

func (b Board) Print(pointerX, pointerY int, selectedColour string) {
	fmt.Printf(b.Serialize(pointerX, pointerY, selectedColour))
}

func (b Board) Serialize(pointerX, pointerY int, selectedColour string) string {
	return SerializeBoard(b, pointerX, pointerY, selectedColour)
}

func (b Board) GetDimensions() (height, width int) {
	return len(b), len(b[0])
}

func (b Board) Select(pointerX int, pointerY int, keycode byte) (exit bool, err error) {
	if keycode > '1' && keycode <= '9' {
		b[pointerY][pointerX].Value = int(keycode - '0')
		return false, nil
	}

	return false, errors.New("you can only change a tile to a number between 1 and 9")
}
