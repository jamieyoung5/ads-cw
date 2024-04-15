package sudoku_board

import (
	"ads-cw/pkg/display"
	"ads-cw/pkg/dlx"
	"errors"
	"fmt"
	"math"
)

type Board struct {
	Content      [][]int
	initialBoard [][]int
	emptyCells   int
	size         int
	subGridSize  int
}

func NewBoard(board [][]int) (*Board, error) {
	verticalSize := len(board)
	if verticalSize != len(board[0]) || !isPerfectSquare(verticalSize) {
		return nil, fmt.Errorf("the provided board is not a perfect square")
	}

	emptyCells := 0
	for _, row := range board {
		for _, cell := range row {
			if cell == 0 {
				emptyCells++
			}
		}
	}

	subGridSize := int(math.Sqrt(float64(verticalSize)))
	return &Board{Content: board, initialBoard: board, emptyCells: emptyCells, size: verticalSize, subGridSize: subGridSize}, nil
}

func (b *Board) Print(pointer *display.Pointer) {
	fmt.Printf(b.Serialize(pointer))
}

func (b *Board) Serialize(pointer *display.Pointer) string {
	return SerializeBoard(b.Content, b.subGridSize, pointer)
}

func (b *Board) GetDimensions() (height, width int) {
	return len(b.Content), len(b.Content[0])
}

func (b *Board) Select(pointer *display.Pointer, keyCode []byte) (exit bool, err error) {
	if b.initialBoard[pointer.Y][pointer.X] != 0 {
		return false, errors.New("you cannot edit a pre-set cell")
	}

	if keyCode[0] >= '1' && keyCode[0] <= '9' {
		if b.Content[pointer.Y][pointer.X] == 0 {
			b.emptyCells--
		}
		b.Content[pointer.Y][pointer.X] = int(keyCode[0] - '0')
		valid, end := b.Validate()
		if end {
			return true, nil
		}

		if !valid {
			err = errors.New("invalid placement")
		}

		return false, err
	}

	return false, errors.New("you can only change a tile to a number between 1 and 9")
}

func (b *Board) Validate() (validBoard bool, boardSolved bool) {
	if b.emptyCells == 0 {
		valid := b.IsValidSolution()
		return valid, valid
	}

	matrix := generateDancingLinksMatrix(b.size, b.subGridSize, b.Content)

	results := make([][]*dlx.Node, 0)
	matrix.Search([]*dlx.Node{}, &results)
	s, _ := solutionToBoard(results[0], b.size)
	fmt.Println(SerializeBoard(s, b.subGridSize, &display.Pointer{X: 1, Y: 1, SelectedTileColour: ""}))
	if len(results) == 0 {
		return false, false
	}
	return true, false
}

func (b *Board) IsValidSolution() bool {
	seen := make(map[string]bool)

	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			num := b.Content[i][j]
			if num < 1 || num > b.size {
				return false
			}
			rowKey := fmt.Sprintf("row_%d_%d", i, num)
			colKey := fmt.Sprintf("col_%d_%d", j, num)
			sqKey := fmt.Sprintf("sq_%d_%d_%d", i/b.subGridSize, j/b.subGridSize, num)

			if seen[rowKey] || seen[colKey] || seen[sqKey] {
				return false
			}
			seen[rowKey], seen[colKey], seen[sqKey] = true, true, true
		}
	}
	return true
}

func isPerfectSquare(n int) bool {
	sqrt := int(math.Sqrt(float64(n)))
	return sqrt*sqrt == n
}
