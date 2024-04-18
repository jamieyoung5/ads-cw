package sudoku_board

import (
	"ads-cw/pkg/display"
	"ads-cw/pkg/dlx"
	"bytes"
	"encoding/gob"
	"fmt"
	"math"
	"strconv"
	"time"
)

type Board struct {
	Content           [][]int
	InitialBoard      [][]int
	EmptyCells        int
	Size              int
	SubGridSize       int
	FooterMessage     string
	InvalidPlacements map[[2]int]struct{}
	PreviousState     *Board
	FutureState       *Board
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
	initialBoardCopy := make([][]int, len(board))
	for i := range board {
		initialBoardCopy[i] = make([]int, len(board[i]))
		copy(initialBoardCopy[i], board[i])
	}

	gob.Register(&Board{})
	gob.Register(map[[2]int]struct{}{})
	return &Board{
		Content:           board,
		InitialBoard:      initialBoardCopy,
		EmptyCells:        emptyCells,
		Size:              verticalSize,
		SubGridSize:       subGridSize,
		InvalidPlacements: make(map[[2]int]struct{}),
	}, nil
}

func (b *Board) Print(pointer *display.Pointer) {
	fmt.Printf(b.Serialize(pointer))
}

func (b *Board) Serialize(pointer *display.Pointer) string {
	return b.serializeBoard(pointer)
}

func (b *Board) GetDimensions() (height, width int) {
	return len(b.Content), len(b.Content[0])
}

func (b *Board) Select(pointer *display.Pointer, macro string) (*display.State, bool) {
	b.FooterMessage = ""

	if b.InitialBoard[pointer.Y][pointer.X] != 0 {
		b.FooterMessage = "cannot edit pre-set static cells!"
		return nil, false
	}

	switch macro {
	case display.Clear:
		b.PreviousState = b.DeepCopy()
		b.EmptyCells++
		b.Content[pointer.Y][pointer.X] = 0
		delete(b.InvalidPlacements, [2]int{pointer.Y, pointer.X})
		return nil, false
	case display.Redo:
		if b.FutureState != nil {
			previousState := b.DeepCopy()
			b.MapNewState(b.FutureState)
			b.PreviousState = previousState
		}
		return nil, false
	case display.Undo:
		if b.PreviousState != nil {
			futureState := b.DeepCopy()
			b.MapNewState(b.PreviousState)
			b.FutureState = futureState
		}
		return nil, false
	}

	// check if there is space for the new value
	if b.Content[pointer.Y][pointer.X] != 0 {
		if len(strconv.Itoa(b.Size)) < len(strconv.Itoa(b.Content[pointer.Y][pointer.X]))+1 {
			return nil, false
		}
	}

	if isOneToNine(macro) {
		convertedNewValue, _ := strconv.Atoi(strconv.Itoa(b.Content[pointer.Y][pointer.X]) + macro)
		if convertedNewValue > b.Size {
			b.FooterMessage = fmt.Sprintf("max cell value is %d", convertedNewValue)
			return nil, false
		}

		b.EmptyCells--

		b.PreviousState = b.DeepCopy()
		b.Content[pointer.Y][pointer.X] = convertedNewValue
		valid, end := b.Validate()
		if end {
			b.FooterMessage = "Board complete!"
			time.Sleep(5 * time.Second)
			return nil, true
		}

		if !valid {
			b.InvalidPlacements[[2]int{pointer.Y, pointer.X}] = struct{}{}
			b.FooterMessage = "placement has led to no possible solutions"
			return nil, false
		} else {
			delete(b.InvalidPlacements, [2]int{pointer.Y, pointer.X})
			return nil, false
		}

	}

	b.FooterMessage = "cell inputs must be a number!"
	return nil, false
}

func (b *Board) DeepCopy() *Board {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	decoder := gob.NewDecoder(&buf)

	// Encode the current Board
	if err := encoder.Encode(b); err != nil {
		panic(err)
	}

	// Decode into a new Board object
	var deepCopy Board
	if err := decoder.Decode(&deepCopy); err != nil {
		panic(err)
	}

	return &deepCopy
}

func (b *Board) MapNewState(newState *Board) {
	b.Content = newState.Content
	b.InitialBoard = newState.InitialBoard
	b.EmptyCells = newState.EmptyCells
	b.Size = newState.Size
	b.SubGridSize = newState.SubGridSize
	b.FooterMessage = newState.FooterMessage
	b.InvalidPlacements = newState.InvalidPlacements
	b.PreviousState = newState.PreviousState
	b.FutureState = newState.FutureState
}

func (b *Board) Validate() (validBoard bool, boardSolved bool) {
	if b.EmptyCells == 0 {
		valid := b.IsValidSolution()
		return valid, valid
	}

	matrix := generateDancingLinksMatrix(b.Size, b.SubGridSize, b.Content)

	results := make([][]*dlx.Node, 0)
	matrix.Search([]*dlx.Node{}, &results)
	if len(results) == 0 {
		return false, false
	}
	return true, false
}

func (b *Board) IsValidSolution() bool {
	seen := make(map[string]bool)

	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			num := b.Content[i][j]
			if num < 1 || num > b.Size {
				return false
			}
			rowKey := fmt.Sprintf("row_%d_%d", i, num)
			colKey := fmt.Sprintf("col_%d_%d", j, num)
			sqKey := fmt.Sprintf("sq_%d_%d_%d", i/b.SubGridSize, j/b.SubGridSize, num)

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

func isOneToNine(s string) bool {
	if len(s) != 1 {
		return false
	}
	num, err := strconv.Atoi(s)
	return err == nil && num >= 1 && num <= 9
}
