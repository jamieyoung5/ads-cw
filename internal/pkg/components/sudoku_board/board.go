package sudoku_board

import (
	"ads-cw/pkg/display"
	"errors"
	"fmt"
)

type Board [][]Tile

func (b Board) Print(pointer *display.Pointer) {
	fmt.Printf(b.Serialize(pointer))
}

func (b Board) Serialize(pointer *display.Pointer) string {
	return SerializeBoard(b, pointer)
}

func (b Board) GetDimensions() (height, width int) {
	return len(b), len(b[0])
}

func (b Board) Select(pointer *display.Pointer, keycode byte) (exit bool, err error) {
	if keycode > '1' && keycode <= '9' {
		b[pointer.Y][pointer.X].Value = int(keycode - '0')
		return false, nil
	}

	return false, errors.New("you can only change a tile to a number between 1 and 9")
}
