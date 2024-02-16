package sudoku_board

import (
	"strconv"
	"strings"
)

const (
	horizontalDivider = "-"
	verticalDivider   = "|"
	noValue           = "#"
)

type Serializer struct {
}

func NewSerializer() *Serializer {
	return &Serializer{}
}

func (s *Serializer) SerializeBoard(board [][]Tile) string {
	if len(board) == 0 || len(board[0]) == 0 {
		return ""
	}

	var builder strings.Builder
	rowDivider := s.createDivider(len(board[0])*4+1, horizontalDivider)

	builder.WriteString(rowDivider)
	for rowIndex, row := range board {
		s.serializeRow(&builder, row, rowIndex > 0 && row[0].SubGrid != board[rowIndex-1][0].SubGrid, rowDivider)
	}
	builder.WriteString(rowDivider)

	return builder.String()
}

func (s *Serializer) serializeRow(builder *strings.Builder, row []Tile, needsDivider bool, rowDivider string) {
	if needsDivider {
		builder.WriteString(rowDivider)
	}

	var lastColumnSubGrid = -1
	for columnIndex, tile := range row {
		if columnIndex == 0 || tile.SubGrid != lastColumnSubGrid {
			builder.WriteString(verticalDivider)
		} else {
			builder.WriteString(" ")
		}
		s.serializeTile(builder, tile)
		lastColumnSubGrid = tile.SubGrid
	}
	builder.WriteString(verticalDivider + "\n")
}

func (s *Serializer) serializeTile(builder *strings.Builder, tile Tile) {
	if tile.Value == 0 {
		builder.WriteString(" " + noValue + " ")
	} else {
		builder.WriteString(" " + strconv.Itoa(tile.Value) + " ")
	}
}

func (s *Serializer) createDivider(length int, dividerSymbol string) string {
	return strings.Repeat(dividerSymbol, length) + "\n"
}
