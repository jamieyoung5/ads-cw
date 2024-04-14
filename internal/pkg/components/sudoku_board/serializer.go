package sudoku_board

import (
	"ads-cw/pkg/display"
	"fmt"
	"strings"
)

const (
	verticalDividerSymbol     = "-"
	horizontalDividerSymbol   = "|"
	crossSectionDividerSymbol = "+"
	noValueSymbol             = "."
	resetStyle                = "\033[0m"          // anything after this will have no custom colouring applied
	whiteBGBlack              = "\033[47m\033[30m" // White background, Black text
)

func SerializeBoard(board [][]int, gridSize int, pointer *display.Pointer) string {
	var sb strings.Builder
	for y, row := range board {
		if y%gridSize == 0 && y != 0 {
			sb.WriteString(createVerticalDivider(gridSize))
		}
		for x, val := range row {
			if x%gridSize == 0 && x != 0 {
				sb.WriteString(horizontalDividerSymbol + " ")
			}

			func() {
				defer sb.WriteString(" ")
				if x == pointer.X && y == pointer.Y {
					defer sb.WriteString(resetStyle)
					sb.WriteString(pointer.SelectedTileColour)
				}

				if val == 0 {
					sb.WriteString(noValueSymbol)
				} else {
					sb.WriteString(fmt.Sprintf("%d", val))
				}
			}()
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func createVerticalDivider(gridSize int) (divider string) {

	squareDivider := createLine(gridSize * 2)
	divider = squareDivider

	for i := 0; i < gridSize-2; i++ {
		divider += crossSectionDividerSymbol + squareDivider + verticalDividerSymbol
	}

	return divider + crossSectionDividerSymbol + squareDivider + "\n"
}

func createLine(length int) (line string) {
	for i := 0; i < length; i++ {
		line += verticalDividerSymbol
	}

	return line
}
