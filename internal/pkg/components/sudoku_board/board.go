package sudoku_board

import (
	"ads-cw/pkg/display"
	"errors"
	"fmt"
	"math"
)

type Board [][]int

func NewBoard(board [][]int) (Board, error) {
	verticalSize := len(board)
	if verticalSize != len(board[0]) || !isPerfectSquare(verticalSize) {
		return nil, fmt.Errorf("the provided board is not a perfect square")
	}

	return board, nil
}

func (b Board) Print(pointer *display.Pointer) {
	fmt.Printf(b.Serialize(pointer))
}

func (b Board) Serialize(pointer *display.Pointer) string {
	gridSize := b.GetGridSize()
	return SerializeBoard(b, gridSize, pointer)
}

func (b Board) GetDimensions() (height, width int) {
	return len(b), len(b[0])
}

func (b Board) Select(pointer *display.Pointer, keycode byte) (exit bool, err error) {
	if keycode > '1' && keycode <= '9' {
		b[pointer.Y][pointer.X] = int(keycode - '0')
		return false, nil
	}

	return false, errors.New("you can only change a tile to a number between 1 and 9")
}

func (b Board) GetGridSize() int {
	boardLength := len(b)
	return int(math.Sqrt(float64(boardLength)))
}

func isPerfectSquare(n int) bool {
	sqrt := int(math.Sqrt(float64(n)))
	return sqrt*sqrt == n
}
