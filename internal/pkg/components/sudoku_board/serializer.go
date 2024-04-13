package sudoku_board

import (
	"ads-cw/pkg/display"
	"strconv"
	"strings"
)

const (
	horizontalDivider = "-"
	verticalDivider   = "|"
	noValue           = "#"
	resetStyle        = "\033[0m"
	whiteBGBlack      = "\033[47m\033[30m" // White background, Black text
)

func SerializeBoard(board [][]Tile, pointer *display.Pointer) string {
	if len(board) == 0 || len(board[0]) == 0 {
		return ""
	}

	var builder strings.Builder
	rowDivider := createDivider(len(board[0])*4+1, horizontalDivider)

	builder.WriteString(rowDivider)
	for rowIndex, row := range board {
		x := -1
		if pointer != nil && pointer.Y == rowIndex {
			x = pointer.X
		}
		serializeRow(&builder, row, rowIndex > 0 && row[0].SubGrid != board[rowIndex-1][0].SubGrid, rowDivider, x, pointer.SelectedTileColour)
	}
	builder.WriteString(rowDivider)

	return builder.String()
}

func serializeRow(builder *strings.Builder, row []Tile, needsDivider bool, rowDivider string, pointerX int, selectionColour string) {
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
		serializeTile(builder, tile, columnIndex == pointerX, selectionColour)
		lastColumnSubGrid = tile.SubGrid
	}
	builder.WriteString(verticalDivider + "\n")
}

func serializeTile(builder *strings.Builder, tile Tile, selected bool, selectionColour string) {
	if selected {
		builder.WriteString(selectionColour)
	}

	if tile.Value == 0 {
		builder.WriteString(" " + noValue + " ")
	} else {
		builder.WriteString(" " + strconv.Itoa(tile.Value) + " ")
	}

	if selected {
		builder.WriteString(resetStyle)
	}
}

func createDivider(length int, dividerSymbol string) string {
	return strings.Repeat(dividerSymbol, length) + "\n"
}
